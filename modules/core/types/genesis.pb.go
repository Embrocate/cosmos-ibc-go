// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/types/v1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	types1 "github.com/cosmos/ibc-go/v5/modules/core/03-connection/types"
	types2 "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the ibc module's genesis state.
type GenesisState struct {
	// ICS002 - Clients genesis state
	ClientGenesis types.GenesisState `protobuf:"bytes,1,opt,name=client_genesis,json=clientGenesis,proto3" json:"client_genesis" yaml:"client_genesis"`
	// ICS003 - Connections genesis state
	ConnectionGenesis types1.GenesisState `protobuf:"bytes,2,opt,name=connection_genesis,json=connectionGenesis,proto3" json:"connection_genesis" yaml:"connection_genesis"`
	// ICS004 - Channel genesis state
	ChannelGenesis types2.GenesisState `protobuf:"bytes,3,opt,name=channel_genesis,json=channelGenesis,proto3" json:"channel_genesis" yaml:"channel_genesis"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a49c5663e6fc59, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetClientGenesis() types.GenesisState {
	if m != nil {
		return m.ClientGenesis
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetConnectionGenesis() types1.GenesisState {
	if m != nil {
		return m.ConnectionGenesis
	}
	return types1.GenesisState{}
}

func (m *GenesisState) GetChannelGenesis() types2.GenesisState {
	if m != nil {
		return m.ChannelGenesis
	}
	return types2.GenesisState{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ibc.core.types.v1.GenesisState")
}

func init() { proto.RegisterFile("ibc/core/types/v1/genesis.proto", fileDescriptor_b9a49c5663e6fc59) }

var fileDescriptor_b9a49c5663e6fc59 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0x93, 0x5e, 0xe9, 0x0e, 0x01, 0x8a, 0x1a, 0x01, 0x82, 0x4a, 0xb8, 0x6d, 0xd4, 0x81,
	0x05, 0x5b, 0x05, 0x26, 0xc6, 0x2e, 0x30, 0x87, 0x8d, 0x05, 0x25, 0xc6, 0xa4, 0x46, 0x89, 0x4f,
	0x55, 0xbb, 0x91, 0xfa, 0x16, 0x3c, 0x56, 0xc7, 0x8e, 0x88, 0xa1, 0x42, 0xc9, 0x1b, 0xf0, 0x04,
	0xa8, 0xb1, 0x49, 0x52, 0x79, 0x8b, 0xfe, 0xf3, 0x9d, 0xff, 0x3b, 0x4a, 0xe2, 0x0d, 0x78, 0x4c,
	0x09, 0x85, 0x05, 0x23, 0x6a, 0x35, 0x67, 0x92, 0xe4, 0x13, 0x92, 0x30, 0xc1, 0x24, 0x97, 0x78,
	0xbe, 0x00, 0x05, 0x7e, 0x8f, 0xc7, 0x14, 0xef, 0x00, 0x5c, 0x01, 0x38, 0x9f, 0xf4, 0x4f, 0x12,
	0x48, 0xa0, 0x9a, 0x92, 0xdd, 0x93, 0x06, 0xfb, 0xc3, 0xba, 0x89, 0xa6, 0x9c, 0x09, 0x65, 0x55,
	0xf5, 0xc7, 0x0d, 0x01, 0x42, 0x30, 0xaa, 0x38, 0x08, 0x9b, 0x1a, 0x35, 0xd4, 0x2c, 0x12, 0x82,
	0xa5, 0x16, 0x12, 0x7c, 0x75, 0xbc, 0xc3, 0x07, 0x9d, 0x3c, 0xa9, 0x48, 0x31, 0xff, 0xcd, 0xeb,
	0x6a, 0xe9, 0x8b, 0x01, 0xcf, 0xdd, 0xa1, 0x7b, 0x75, 0x70, 0x33, 0xc4, 0xf5, 0xf5, 0x7a, 0x8e,
	0xf3, 0x09, 0x6e, 0x6f, 0x4e, 0x2f, 0xd7, 0xdb, 0x81, 0xf3, 0xb3, 0x1d, 0x9c, 0xae, 0xa2, 0x2c,
	0xbd, 0x0f, 0xf6, 0x5b, 0x82, 0xf0, 0x48, 0x07, 0x66, 0xc5, 0xcf, 0x3d, 0xbf, 0x39, 0xbd, 0x76,
	0x75, 0x2a, 0xd7, 0xb8, 0xe5, 0xaa, 0x19, 0xcb, 0x37, 0x32, 0xbe, 0x0b, 0xe3, 0xb3, 0xda, 0x82,
	0xb0, 0xd7, 0x84, 0x7f, 0xde, 0x77, 0xef, 0xd8, 0xbc, 0x8c, 0x5a, 0xfa, 0xaf, 0x92, 0x8e, 0x5a,
	0x52, 0x0d, 0x58, 0x46, 0x64, 0x8c, 0x67, 0xc6, 0xb8, 0xdf, 0x13, 0x84, 0x5d, 0x93, 0x98, 0xa5,
	0xe9, 0xe3, 0xba, 0x40, 0xee, 0xa6, 0x40, 0xee, 0x77, 0x81, 0xdc, 0x8f, 0x12, 0x39, 0x9b, 0x12,
	0x39, 0x9f, 0x25, 0x72, 0x9e, 0x71, 0xc2, 0xd5, 0x6c, 0x19, 0x63, 0x0a, 0x19, 0xa1, 0x20, 0x33,
	0x90, 0x84, 0xc7, 0xf4, 0x3a, 0x01, 0x92, 0xdf, 0x91, 0x0c, 0x5e, 0x97, 0x29, 0x93, 0xad, 0x7f,
	0x29, 0xfe, 0x5f, 0x7d, 0xad, 0xdb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xbe, 0x54, 0x19,
	0x64, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ChannelGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.ConnectionGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ClientGenesis.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ClientGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.ConnectionGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.ChannelGenesis.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClientGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ConnectionGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelGenesis", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChannelGenesis.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
