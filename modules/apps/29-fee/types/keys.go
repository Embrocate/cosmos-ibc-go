package types

import "fmt"

const (
	// ModuleName defines the 29-fee name
	ModuleName = "feeibc"

	// StoreKey is the store key string for IBC transfer
	StoreKey = ModuleName

	// PortKey is the port id that is wrapped by fee middleware
	PortKey = "feetransfer"

	// RouterKey is the message route for IBC transfer
	RouterKey = ModuleName

	// QuerierRoute is the querier route for IBC transfer
	QuerierRoute = ModuleName

	Version = "fee29-1"

	// FeeEnabledPrefix is the key prefix for storing fee enabled flag
	FeeEnabledKeyPrefix = "feeEnabled"

	// RelayerAddressKeyPrefix is the key prefix for relayer address mapping
	RelayerAddressKeyPrefix = "relayerAddress"

	// RelayerAddressKeyPrefix is the key prefix for relayer address mapping
	FeeInEscrowPrefix = "feeInEscrow"
)

// FeeEnabledKey returns the key that stores a flag to determine if fee logic should
// be enabled for the given port and channel identifiers.
func FeeEnabledKey(portID, channelID string) []byte {
	return []byte(fmt.Sprintf("%s/%s/%s", FeeEnabledKeyPrefix, portID, channelID))
}

// KeyRelayerAddress returns the key for relayer address -> counteryparty address mapping
func KeyRelayerAddress(address string) []byte {
	return []byte(fmt.Sprintf("%s/%s", RelayerAddressKeyPrefix, address))
}

// KeyFeeInEscrow returns the key for escrowed fees
func KeyFeeInEscrow(account, channelId, sequenceId string) []byte {
	return []byte(fmt.Sprintf("%s/%s/%s/packet/%s", FeeInEscrowPrefix, account, channelId, sequenceId))
}
