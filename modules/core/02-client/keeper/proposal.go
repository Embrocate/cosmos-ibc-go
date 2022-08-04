package keeper

import (
	metrics "github.com/armon/go-metrics"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v5/modules/core/exported"
)

// ClientUpdateProposal will retrieve the subject and substitute client.
// A callback will occur to the subject client state with the client
// prefixed store being provided for both the subject and the substitute client.
// The IBC client implementations are responsible for validating the parameters of the
// subtitute (enusring they match the subject's parameters) as well as copying
// the necessary consensus states from the subtitute to the subject client
// store. The substitute must be Active and the subject must not be Active.
func (k Keeper) ClientUpdateProposal(ctx sdk.Context, p *types.ClientUpdateProposal) error {
	subjectClientState, found := k.GetClientState(ctx, p.SubjectClientId)
	if !found {
		return sdkerrors.Wrapf(types.ErrClientNotFound, "subject client with ID %s", p.SubjectClientId)
	}

	subjectClientStore := k.ClientStore(ctx, p.SubjectClientId)

	if status := subjectClientState.Status(ctx, subjectClientStore, k.cdc); status == exported.Active {
		return sdkerrors.Wrap(types.ErrInvalidUpdateClientProposal, "cannot update Active subject client")
	}

	substituteClientState, found := k.GetClientState(ctx, p.SubstituteClientId)
	if !found {
		return sdkerrors.Wrapf(types.ErrClientNotFound, "substitute client with ID %s", p.SubstituteClientId)
	}

	if subjectClientState.GetLatestHeight().GTE(substituteClientState.GetLatestHeight()) {
		return sdkerrors.Wrapf(types.ErrInvalidHeight, "subject client state latest height is greater or equal to substitute client state latest height (%s >= %s)", subjectClientState.GetLatestHeight(), substituteClientState.GetLatestHeight())
	}

	substituteClientStore := k.ClientStore(ctx, p.SubstituteClientId)

	if status := substituteClientState.Status(ctx, substituteClientStore, k.cdc); status != exported.Active {
		return sdkerrors.Wrapf(types.ErrClientStateNotFound, "substitute client is not Active, status is %s", status)
	}

	if err := subjectClientState.CheckSubstituteAndUpdateState(ctx, k.cdc, subjectClientStore, substituteClientStore, substituteClientState); err != nil {
		return err
	}

	k.Logger(ctx).Info("client updated after governance proposal passed", "client-id", p.SubjectClientId, "height")

	defer func() {
		telemetry.IncrCounterWithLabels(
			[]string{"ibc", "client", "update"},
			1,
			[]metrics.Label{
				telemetry.NewLabel(types.LabelClientType, substituteClientState.ClientType()),
				telemetry.NewLabel(types.LabelClientID, p.SubjectClientId),
				telemetry.NewLabel(types.LabelUpdateType, "proposal"),
			},
		)
	}()

	// emitting events in the keeper for proposal updates to clients
	EmitUpdateClientProposalEvent(ctx, p.SubjectClientId, substituteClientState.ClientType())

	return nil
}

// HandleUpgradeProposal sets the upgraded client state in the upgrade store. It clears
// an IBC client state and consensus state if a previous plan was set. Then  it
// will schedule an upgrade and finally set the upgraded client state in upgrade
// store.
func (k Keeper) HandleUpgradeProposal(ctx sdk.Context, p *types.UpgradeProposal) error {
	clientState, err := types.UnpackClientState(p.UpgradedClientState)
	if err != nil {
		return sdkerrors.Wrap(err, "could not unpack UpgradedClientState")
	}

	// zero out any custom fields before setting
	cs := clientState.ZeroCustomFields()
	bz, err := types.MarshalClientState(k.cdc, cs)
	if err != nil {
		return sdkerrors.Wrap(err, "could not marshal UpgradedClientState")
	}

	if err := k.upgradeKeeper.ScheduleUpgrade(ctx, p.Plan); err != nil {
		return err
	}

	// sets the new upgraded client in last height committed on this chain is at plan.Height,
	// since the chain will panic at plan.Height and new chain will resume at plan.Height
	if err = k.upgradeKeeper.SetUpgradedClient(ctx, p.Plan.Height, bz); err != nil {
		return err
	}

	// emitting an event for handling client upgrade proposal
	EmitUpgradeClientProposalEvent(ctx, p.Title, p.Plan.Height)

	return nil
}
