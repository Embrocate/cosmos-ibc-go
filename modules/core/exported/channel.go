package exported

// ChannelI defines the standard interface for a channel end.
type ChannelI interface {
	GetState() int32
	GetOrdering() int32
	GetCounterparty() CounterpartyChannelI
	GetConnectionHops() []string
	GetVersion() string
	ValidateBasic() error
}

// CounterpartyChannelI defines the standard interface for a channel end's
// counterparty.
type CounterpartyChannelI interface {
	GetPortID() string
	GetChannelID() string
	ValidateBasic() error
}

// PacketI defines the standard interface for IBC packets
type PacketI interface {
	GetSequence() uint64
	GetTimeoutHeight() Height
	GetTimeoutTimestamp() uint64
	GetSourcePort() string
	GetSourceChannel() string
	GetDestPort() string
	GetDestChannel() string
	GetData() []byte
	ValidateBasic() error
}

// Acknowledgement defines the interface used to return
// acknowledgements in the OnRecvPacket callback.
type Acknowledgement interface {
	// Success determines if the IBC application state should be persisted when handling `RecvPacket`.
	// During `OnRecvPacket` IBC application callback execution, all state changes are held in a cache store and committed if:
	// - the acknowledgement.Success() returns true
	// - a nil acknowledgement is returned (asynchronous acknowledgements)
	//
	// Note 1: IBC application callback events are always persisted so long as `RecvPacket` succeeds without error.
	//
	// Note 2: This method is independent of application level success/error which is encoded in the acknowledgement bytes in a protocol specific way.
	//
	// See https://github.com/cosmos/ibc-go/blob/v7.0.0/docs/ibc/apps.md for further explanations.
	Success() bool
	Acknowledgement() []byte
}
