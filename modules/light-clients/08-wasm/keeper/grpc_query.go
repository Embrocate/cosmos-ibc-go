package keeper

import (
	"context"
	"encoding/hex"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"

	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/internal/ibcwasm"
	"github.com/cosmos/ibc-go/modules/light-clients/08-wasm/types"
)

var _ types.QueryServer = (*Keeper)(nil)

// Code implements the Query/Code gRPC method
func (k Keeper) Code(goCtx context.Context, req *types.QueryCodeRequest) (*types.QueryCodeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	checksum, err := hex.DecodeString(req.Checksum)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid code hash")
	}

	// Only return code hashes we previously stored, not arbitrary code hashes that might be stored via e.g Wasmd.
	if !types.HasCodeHash(sdk.UnwrapSDKContext(goCtx), checksum) {
		return nil, status.Error(codes.NotFound, errorsmod.Wrap(types.ErrWasmCodeHashNotFound, req.Checksum).Error())
	}

	code, err := k.wasmVM.GetCode(checksum)
	if err != nil {
		return nil, status.Error(codes.NotFound, errorsmod.Wrap(types.ErrWasmCodeHashNotFound, req.Checksum).Error())
	}

	return &types.QueryCodeResponse{
		Data: code,
	}, nil
}

// Checksums implements the Query/Checksums gRPC method. It returns a list of hex encoded checksums stored.
func (Keeper) Checksums(goCtx context.Context, req *types.QueryChecksumsRequest) (*types.QueryChecksumsResponse, error) {
	checksums, pageRes, err := sdkquery.CollectionPaginate(
		goCtx,
		ibcwasm.CodeHashes,
		req.Pagination,
		func(key []byte, value collections.NoValue) (string, error) {
			return hex.EncodeToString(key), nil
		})
	if err != nil {
		return nil, err
	}

	return &types.QueryChecksumsResponse{
		Checksums:  checksums,
		Pagination: pageRes,
	}, nil
}
