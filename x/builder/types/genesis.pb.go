// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pob/builder/v1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the genesis state of the x/builder module.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_287f1bdff5ccfc33, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// Params defines the parameters of the x/builder module.
type Params struct {
	// max_bundle_size is the maximum number of transactions that can be bundled
	// in a single bundle.
	MaxBundleSize uint32 `protobuf:"varint,1,opt,name=max_bundle_size,json=maxBundleSize,proto3" json:"max_bundle_size,omitempty"`
	// escrow_account_address is the address of the account that will receive a
	// portion of the bid proceeds.
	EscrowAccountAddress []byte `protobuf:"bytes,2,opt,name=escrow_account_address,json=escrowAccountAddress,proto3" json:"escrow_account_address,omitempty"`
	// reserve_fee specifies the bid floor for the auction.
	ReserveFee types.Coin `protobuf:"bytes,3,opt,name=reserve_fee,json=reserveFee,proto3" json:"reserve_fee"`
	// min_bid_increment specifies the minimum amount that the next bid must be
	// greater than the previous bid.
	MinBidIncrement types.Coin `protobuf:"bytes,4,opt,name=min_bid_increment,json=minBidIncrement,proto3" json:"min_bid_increment"`
	// front_running_protection specifies whether front running and sandwich
	// attack protection is enabled.
	FrontRunningProtection bool `protobuf:"varint,5,opt,name=front_running_protection,json=frontRunningProtection,proto3" json:"front_running_protection,omitempty"`
	// proposer_fee defines the portion of the winning bid that goes to the block
	// proposer that proposed the block.
	ProposerFee github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=proposer_fee,json=proposerFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"proposer_fee"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_287f1bdff5ccfc33, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxBundleSize() uint32 {
	if m != nil {
		return m.MaxBundleSize
	}
	return 0
}

func (m *Params) GetEscrowAccountAddress() []byte {
	if m != nil {
		return m.EscrowAccountAddress
	}
	return nil
}

func (m *Params) GetReserveFee() types.Coin {
	if m != nil {
		return m.ReserveFee
	}
	return types.Coin{}
}

func (m *Params) GetMinBidIncrement() types.Coin {
	if m != nil {
		return m.MinBidIncrement
	}
	return types.Coin{}
}

func (m *Params) GetFrontRunningProtection() bool {
	if m != nil {
		return m.FrontRunningProtection
	}
	return false
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "pob.builder.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "pob.builder.v1.Params")
}

func init() { proto.RegisterFile("pob/builder/v1/genesis.proto", fileDescriptor_287f1bdff5ccfc33) }

var fileDescriptor_287f1bdff5ccfc33 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6e, 0xd4, 0x3c,
	0x14, 0x85, 0x27, 0x7f, 0xfb, 0x8f, 0xa8, 0x67, 0x4a, 0xd5, 0xa8, 0x1a, 0x85, 0x82, 0xd2, 0xa8,
	0x8b, 0x32, 0xaa, 0x54, 0x5b, 0x03, 0x2c, 0x10, 0xbb, 0x09, 0xa5, 0x88, 0xdd, 0x90, 0xee, 0xd8,
	0x44, 0x4e, 0x72, 0x1b, 0xac, 0xd6, 0x76, 0x64, 0x3b, 0x61, 0xe8, 0x23, 0xb0, 0xe2, 0x31, 0x58,
	0xa1, 0x3e, 0x46, 0x97, 0x5d, 0x22, 0x16, 0x15, 0x9a, 0x59, 0xf4, 0x35, 0x50, 0xec, 0xb4, 0x94,
	0x25, 0x9b, 0x24, 0xf2, 0x77, 0x7d, 0xce, 0x3d, 0xb9, 0x17, 0x3d, 0xa9, 0x64, 0x46, 0xb2, 0x9a,
	0x9d, 0x15, 0xa0, 0x48, 0x33, 0x21, 0x25, 0x08, 0xd0, 0x4c, 0xe3, 0x4a, 0x49, 0x23, 0xfd, 0x87,
	0x95, 0xcc, 0x70, 0x47, 0x71, 0x33, 0xd9, 0xde, 0x2a, 0x65, 0x29, 0x2d, 0x22, 0xed, 0x97, 0xab,
	0xda, 0x0e, 0x73, 0xa9, 0xb9, 0xd4, 0x24, 0xa3, 0x1a, 0x48, 0x33, 0xc9, 0xc0, 0xd0, 0x09, 0xc9,
	0x25, 0x13, 0x1d, 0xdf, 0xa4, 0x9c, 0x09, 0x49, 0xec, 0xd3, 0x1d, 0xed, 0x1e, 0xa2, 0xe1, 0x5b,
	0xe7, 0x74, 0x6c, 0xa8, 0x01, 0xff, 0x05, 0xea, 0x57, 0x54, 0x51, 0xae, 0x03, 0x2f, 0xf2, 0xc6,
	0x83, 0x67, 0x23, 0xfc, 0xb7, 0x33, 0x9e, 0x59, 0x1a, 0xaf, 0x5e, 0x5e, 0xef, 0xf4, 0x92, 0xae,
	0x76, 0xf7, 0xfb, 0x0a, 0xea, 0x3b, 0xe0, 0xef, 0xa1, 0x0d, 0x4e, 0xe7, 0x69, 0x56, 0x8b, 0xe2,
	0x0c, 0x52, 0xcd, 0xce, 0xc1, 0x2a, 0xad, 0x27, 0xeb, 0x9c, 0xce, 0x63, 0x7b, 0x7a, 0xcc, 0xce,
	0x5b, 0xa3, 0x11, 0xe8, 0x5c, 0xc9, 0x4f, 0x29, 0xcd, 0x73, 0x59, 0x0b, 0x93, 0xd2, 0xa2, 0x50,
	0xa0, 0x75, 0xf0, 0x5f, 0xe4, 0x8d, 0x87, 0xc9, 0x96, 0xa3, 0x53, 0x07, 0xa7, 0x8e, 0xf9, 0x6f,
	0xd0, 0x40, 0x81, 0x06, 0xd5, 0x40, 0x7a, 0x02, 0x10, 0xac, 0xd8, 0x1e, 0x1f, 0x61, 0x97, 0x1b,
	0xb7, 0xb9, 0x71, 0x97, 0x1b, 0xbf, 0x96, 0x4c, 0xc4, 0x6b, 0x6d, 0x9b, 0xdf, 0x6e, 0x2e, 0xf6,
	0xbd, 0x04, 0x75, 0x17, 0x8f, 0x00, 0xfc, 0x19, 0xda, 0xe4, 0x4c, 0xa4, 0x19, 0x2b, 0x52, 0x26,
	0x72, 0x05, 0x1c, 0x84, 0x09, 0x56, 0xff, 0x41, 0x6c, 0x83, 0x33, 0x11, 0xb3, 0xe2, 0xdd, 0xed,
	0x65, 0xff, 0x25, 0x0a, 0x4e, 0x94, 0x14, 0x26, 0x55, 0xb5, 0x10, 0x4c, 0x94, 0x69, 0xfb, 0x7b,
	0x21, 0x37, 0x4c, 0x8a, 0xe0, 0xff, 0xc8, 0x1b, 0x3f, 0x48, 0x46, 0x96, 0x27, 0x0e, 0xcf, 0xee,
	0xa8, 0xff, 0x1e, 0x0d, 0x2b, 0x25, 0x2b, 0xa9, 0x41, 0xd9, 0x4c, 0xfd, 0xc8, 0x1b, 0xaf, 0xc5,
	0xb8, 0xf5, 0xfa, 0x79, 0xbd, 0xb3, 0x57, 0x32, 0xf3, 0xb1, 0xce, 0x70, 0x2e, 0x39, 0xe9, 0xa6,
	0xeb, 0x5e, 0x07, 0xba, 0x38, 0x25, 0xe6, 0x73, 0x05, 0x1a, 0x1f, 0x42, 0x9e, 0x0c, 0x6e, 0x35,
	0x8e, 0x00, 0x5e, 0x45, 0x5f, 0x6e, 0x2e, 0xf6, 0x1f, 0xdf, 0xab, 0x9b, 0xdf, 0x6d, 0x56, 0x37,
	0xbe, 0xe9, 0xe5, 0x22, 0xf4, 0xae, 0x16, 0xa1, 0xf7, 0x6b, 0x11, 0x7a, 0x5f, 0x97, 0x61, 0xef,
	0x6a, 0x19, 0xf6, 0x7e, 0x2c, 0xc3, 0xde, 0x87, 0xa7, 0xf7, 0x0c, 0xf5, 0x29, 0xab, 0x0e, 0x38,
	0x34, 0xa4, 0xdd, 0xcd, 0x3f, 0x1a, 0xd6, 0x35, 0xeb, 0xdb, 0x05, 0x7a, 0xfe, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x3d, 0xb6, 0x59, 0xc6, 0xb9, 0x02, 0x00, 0x00,
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
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
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

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ProposerFee.Size()
		i -= size
		if _, err := m.ProposerFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.FrontRunningProtection {
		i--
		if m.FrontRunningProtection {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.MinBidIncrement.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.ReserveFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.EscrowAccountAddress) > 0 {
		i -= len(m.EscrowAccountAddress)
		copy(dAtA[i:], m.EscrowAccountAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.EscrowAccountAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.MaxBundleSize != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxBundleSize))
		i--
		dAtA[i] = 0x8
	}
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxBundleSize != 0 {
		n += 1 + sovGenesis(uint64(m.MaxBundleSize))
	}
	l = len(m.EscrowAccountAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.ReserveFee.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.MinBidIncrement.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.FrontRunningProtection {
		n += 2
	}
	l = m.ProposerFee.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBundleSize", wireType)
			}
			m.MaxBundleSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxBundleSize |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EscrowAccountAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EscrowAccountAddress = append(m.EscrowAccountAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.EscrowAccountAddress == nil {
				m.EscrowAccountAddress = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReserveFee", wireType)
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
			if err := m.ReserveFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBidIncrement", wireType)
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
			if err := m.MinBidIncrement.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FrontRunningProtection", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.FrontRunningProtection = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ProposerFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
