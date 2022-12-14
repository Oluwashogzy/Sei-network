// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/settlement.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type SettlementEntry struct {
	Account                string                                 `protobuf:"bytes,1,opt,name=account,proto3" json:"account"`
	PriceDenom             string                                 `protobuf:"bytes,2,opt,name=priceDenom,proto3" json:"price_denom"`
	AssetDenom             string                                 `protobuf:"bytes,3,opt,name=assetDenom,proto3" json:"asset_denom"`
	Quantity               github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=quantity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"quantity" yaml:"quantity"`
	ExecutionCostOrProceed github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=executionCostOrProceed,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"execution_cost_or_proceed" yaml:"execution_cost_or_proceed"`
	ExpectedCostOrProceed  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=expectedCostOrProceed,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"expected_cost_or_proceed" yaml:"expected_cost_or_proceed"`
	PositionDirection      string                                 `protobuf:"bytes,7,opt,name=positionDirection,proto3" json:"position_direction"`
	OrderType              string                                 `protobuf:"bytes,8,opt,name=orderType,proto3" json:"order_type"`
	OrderId                uint64                                 `protobuf:"varint,9,opt,name=orderId,proto3" json:"order_id"`
	Timestamp              uint64                                 `protobuf:"varint,10,opt,name=timestamp,proto3" json:"timestamp"`
	Height                 uint64                                 `protobuf:"varint,11,opt,name=height,proto3" json:"height"`
	SettlementId           uint64                                 `protobuf:"varint,12,opt,name=settlementId,proto3" json:"settlement_id"`
}

func (m *SettlementEntry) Reset()         { *m = SettlementEntry{} }
func (m *SettlementEntry) String() string { return proto.CompactTextString(m) }
func (*SettlementEntry) ProtoMessage()    {}
func (*SettlementEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_c24d83c09612bb1c, []int{0}
}
func (m *SettlementEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SettlementEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SettlementEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SettlementEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettlementEntry.Merge(m, src)
}
func (m *SettlementEntry) XXX_Size() int {
	return m.Size()
}
func (m *SettlementEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SettlementEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SettlementEntry proto.InternalMessageInfo

func (m *SettlementEntry) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *SettlementEntry) GetPriceDenom() string {
	if m != nil {
		return m.PriceDenom
	}
	return ""
}

func (m *SettlementEntry) GetAssetDenom() string {
	if m != nil {
		return m.AssetDenom
	}
	return ""
}

func (m *SettlementEntry) GetPositionDirection() string {
	if m != nil {
		return m.PositionDirection
	}
	return ""
}

func (m *SettlementEntry) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *SettlementEntry) GetOrderId() uint64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *SettlementEntry) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *SettlementEntry) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *SettlementEntry) GetSettlementId() uint64 {
	if m != nil {
		return m.SettlementId
	}
	return 0
}

type Settlements struct {
	Epoch   int64              `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch"`
	Entries []*SettlementEntry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries"`
}

func (m *Settlements) Reset()         { *m = Settlements{} }
func (m *Settlements) String() string { return proto.CompactTextString(m) }
func (*Settlements) ProtoMessage()    {}
func (*Settlements) Descriptor() ([]byte, []int) {
	return fileDescriptor_c24d83c09612bb1c, []int{1}
}
func (m *Settlements) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Settlements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Settlements.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Settlements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settlements.Merge(m, src)
}
func (m *Settlements) XXX_Size() int {
	return m.Size()
}
func (m *Settlements) XXX_DiscardUnknown() {
	xxx_messageInfo_Settlements.DiscardUnknown(m)
}

var xxx_messageInfo_Settlements proto.InternalMessageInfo

func (m *Settlements) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *Settlements) GetEntries() []*SettlementEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*SettlementEntry)(nil), "seiprotocol.seichain.dex.SettlementEntry")
	proto.RegisterType((*Settlements)(nil), "seiprotocol.seichain.dex.Settlements")
}

func init() { proto.RegisterFile("dex/settlement.proto", fileDescriptor_c24d83c09612bb1c) }

var fileDescriptor_c24d83c09612bb1c = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x3f, 0x6f, 0xdb, 0x3c,
	0x10, 0xc6, 0xad, 0xfc, 0xb1, 0x63, 0x3a, 0x79, 0x83, 0x10, 0x79, 0x03, 0xb6, 0x83, 0x68, 0x08,
	0x68, 0x90, 0xa2, 0x8d, 0x04, 0xb4, 0xe8, 0xd2, 0xd1, 0x75, 0x51, 0x64, 0x28, 0x1a, 0xb0, 0x9d,
	0xba, 0x08, 0x0a, 0x75, 0xb0, 0x89, 0x46, 0xa2, 0x2a, 0xd2, 0x80, 0xbd, 0xf5, 0x23, 0xf4, 0x3b,
	0x74, 0xe8, 0x57, 0xc9, 0x98, 0xb1, 0xe8, 0x40, 0x14, 0xc9, 0xa6, 0x31, 0x5b, 0xb7, 0x42, 0x94,
	0x65, 0x25, 0xad, 0x3d, 0x64, 0xe2, 0xf1, 0xb9, 0xdf, 0xdd, 0x3d, 0x86, 0x79, 0x42, 0xfb, 0x31,
	0x4c, 0x03, 0x05, 0x5a, 0x9f, 0x43, 0x02, 0xa9, 0xf6, 0xb3, 0x5c, 0x6a, 0x89, 0x89, 0x02, 0x61,
	0x23, 0x2e, 0xcf, 0x7d, 0x05, 0x82, 0x8f, 0x23, 0x91, 0xfa, 0x31, 0x4c, 0x1f, 0xee, 0x8f, 0xe4,
	0x48, 0xda, 0x54, 0x50, 0x46, 0x15, 0xef, 0xfd, 0x6e, 0xa3, 0xdd, 0xf7, 0x8b, 0x26, 0xaf, 0x53,
	0x9d, 0xcf, 0xf0, 0x23, 0xd4, 0x89, 0x38, 0x97, 0x93, 0x54, 0x13, 0xa7, 0xef, 0x1c, 0x75, 0x07,
	0xbd, 0xc2, 0xd0, 0x5a, 0x62, 0x75, 0x80, 0x03, 0x84, 0xb2, 0x5c, 0x70, 0x18, 0x42, 0x2a, 0x13,
	0xb2, 0x66, 0xc9, 0xdd, 0xc2, 0xd0, 0x9e, 0x55, 0xc3, 0xb8, 0x94, 0xd9, 0x2d, 0xa4, 0x2c, 0x88,
	0x94, 0x02, 0x5d, 0x15, 0xac, 0x37, 0x05, 0x56, 0xad, 0x0b, 0x1a, 0x04, 0x0b, 0xb4, 0xf5, 0x79,
	0x12, 0xa5, 0x5a, 0xe8, 0x19, 0xd9, 0xb0, 0xf8, 0xdb, 0x0b, 0x43, 0x5b, 0x3f, 0x0d, 0x3d, 0x1c,
	0x09, 0x3d, 0x9e, 0x9c, 0xf9, 0x5c, 0x26, 0x01, 0x97, 0x2a, 0x91, 0x6a, 0x7e, 0x1c, 0xab, 0xf8,
	0x53, 0xa0, 0x67, 0x19, 0x28, 0x7f, 0x08, 0xbc, 0x30, 0x74, 0xd1, 0xe1, 0xc6, 0xd0, 0xdd, 0x59,
	0x94, 0x9c, 0xbf, 0xf4, 0x6a, 0xc5, 0x63, 0x8b, 0x24, 0xfe, 0xee, 0xa0, 0x03, 0x98, 0x02, 0x9f,
	0x68, 0x21, 0xd3, 0x57, 0x52, 0xe9, 0x77, 0xf9, 0x69, 0x2e, 0x39, 0x40, 0x4c, 0x36, 0xed, 0x64,
	0x79, 0xef, 0xc9, 0x0f, 0x16, 0xfd, 0x42, 0x2e, 0x95, 0x0e, 0x65, 0x1e, 0x66, 0x55, 0xcb, 0x1b,
	0x43, 0xfb, 0x95, 0x95, 0x95, 0x88, 0xc7, 0x56, 0xd8, 0xc1, 0xdf, 0x1c, 0xf4, 0x3f, 0x4c, 0x33,
	0xe0, 0x1a, 0xe2, 0xbb, 0x46, 0xdb, 0xd6, 0x68, 0x72, 0x6f, 0xa3, 0xa4, 0x6e, 0xb7, 0xc4, 0x27,
	0xad, 0x7d, 0x2e, 0x27, 0x3c, 0xb6, 0xdc, 0x0b, 0x1e, 0xa2, 0xbd, 0x4c, 0x2a, 0x51, 0xda, 0x1f,
	0x8a, 0x1c, 0x78, 0x19, 0x90, 0x8e, 0x35, 0x78, 0x50, 0x18, 0x8a, 0xeb, 0x64, 0x18, 0xd7, 0x59,
	0xf6, 0x6f, 0x01, 0x7e, 0x8a, 0xba, 0x32, 0x8f, 0x21, 0xff, 0x30, 0xcb, 0x80, 0x6c, 0xd9, 0xea,
	0xff, 0x0a, 0x43, 0x91, 0x15, 0xc3, 0xf2, 0x37, 0xb0, 0x06, 0xc0, 0x87, 0xa8, 0x63, 0x2f, 0x27,
	0x31, 0xe9, 0xf6, 0x9d, 0xa3, 0x8d, 0xc1, 0x76, 0xf9, 0xff, 0x57, 0xac, 0x88, 0x59, 0x9d, 0xc4,
	0x4f, 0x50, 0x57, 0x8b, 0x04, 0x94, 0x8e, 0x92, 0x8c, 0x20, 0x4b, 0xee, 0x14, 0x86, 0x36, 0x22,
	0x6b, 0x42, 0xec, 0xa1, 0xf6, 0x18, 0xc4, 0x68, 0xac, 0x49, 0xcf, 0x92, 0xa8, 0x30, 0x74, 0xae,
	0xb0, 0xf9, 0x89, 0x5f, 0xa0, 0xed, 0x66, 0x11, 0x4f, 0x62, 0xb2, 0x6d, 0xc9, 0xbd, 0xc2, 0xd0,
	0x9d, 0x46, 0x2f, 0x2d, 0xdc, 0xc1, 0xbc, 0x2f, 0x0e, 0xea, 0x35, 0xbb, 0xa7, 0x30, 0x45, 0x9b,
	0x90, 0x49, 0x3e, 0xb6, 0x5b, 0xb7, 0x3e, 0xe8, 0x16, 0x86, 0x56, 0x02, 0xab, 0x0e, 0x7c, 0x8a,
	0x3a, 0x90, 0xea, 0x5c, 0x80, 0x22, 0x6b, 0xfd, 0xf5, 0xa3, 0xde, 0xb3, 0xc7, 0xfe, 0xaa, 0x75,
	0xf7, 0xff, 0x5a, 0xea, 0x6a, 0x87, 0xe7, 0xd5, 0xac, 0x0e, 0x06, 0x6f, 0x2e, 0xae, 0x5c, 0xe7,
	0xf2, 0xca, 0x75, 0x7e, 0x5d, 0xb9, 0xce, 0xd7, 0x6b, 0xb7, 0x75, 0x79, 0xed, 0xb6, 0x7e, 0x5c,
	0xbb, 0xad, 0x8f, 0xc7, 0xb7, 0x9e, 0x8f, 0x02, 0x71, 0x5c, 0x4f, 0xb1, 0x17, 0x3b, 0x26, 0x98,
	0x06, 0xe5, 0x27, 0xc8, 0xbe, 0xa4, 0xb3, 0xb6, 0xcd, 0x3f, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x07, 0xa3, 0x0d, 0xb2, 0x96, 0x04, 0x00, 0x00,
}

func (m *SettlementEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SettlementEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SettlementEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SettlementId != 0 {
		i = encodeVarintSettlement(dAtA, i, uint64(m.SettlementId))
		i--
		dAtA[i] = 0x60
	}
	if m.Height != 0 {
		i = encodeVarintSettlement(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x58
	}
	if m.Timestamp != 0 {
		i = encodeVarintSettlement(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x50
	}
	if m.OrderId != 0 {
		i = encodeVarintSettlement(dAtA, i, uint64(m.OrderId))
		i--
		dAtA[i] = 0x48
	}
	if len(m.OrderType) > 0 {
		i -= len(m.OrderType)
		copy(dAtA[i:], m.OrderType)
		i = encodeVarintSettlement(dAtA, i, uint64(len(m.OrderType)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.PositionDirection) > 0 {
		i -= len(m.PositionDirection)
		copy(dAtA[i:], m.PositionDirection)
		i = encodeVarintSettlement(dAtA, i, uint64(len(m.PositionDirection)))
		i--
		dAtA[i] = 0x3a
	}
	{
		size := m.ExpectedCostOrProceed.Size()
		i -= size
		if _, err := m.ExpectedCostOrProceed.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSettlement(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.ExecutionCostOrProceed.Size()
		i -= size
		if _, err := m.ExecutionCostOrProceed.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSettlement(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Quantity.Size()
		i -= size
		if _, err := m.Quantity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSettlement(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.AssetDenom) > 0 {
		i -= len(m.AssetDenom)
		copy(dAtA[i:], m.AssetDenom)
		i = encodeVarintSettlement(dAtA, i, uint64(len(m.AssetDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PriceDenom) > 0 {
		i -= len(m.PriceDenom)
		copy(dAtA[i:], m.PriceDenom)
		i = encodeVarintSettlement(dAtA, i, uint64(len(m.PriceDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintSettlement(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Settlements) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Settlements) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Settlements) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSettlement(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Epoch != 0 {
		i = encodeVarintSettlement(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSettlement(dAtA []byte, offset int, v uint64) int {
	offset -= sovSettlement(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SettlementEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovSettlement(uint64(l))
	}
	l = len(m.PriceDenom)
	if l > 0 {
		n += 1 + l + sovSettlement(uint64(l))
	}
	l = len(m.AssetDenom)
	if l > 0 {
		n += 1 + l + sovSettlement(uint64(l))
	}
	l = m.Quantity.Size()
	n += 1 + l + sovSettlement(uint64(l))
	l = m.ExecutionCostOrProceed.Size()
	n += 1 + l + sovSettlement(uint64(l))
	l = m.ExpectedCostOrProceed.Size()
	n += 1 + l + sovSettlement(uint64(l))
	l = len(m.PositionDirection)
	if l > 0 {
		n += 1 + l + sovSettlement(uint64(l))
	}
	l = len(m.OrderType)
	if l > 0 {
		n += 1 + l + sovSettlement(uint64(l))
	}
	if m.OrderId != 0 {
		n += 1 + sovSettlement(uint64(m.OrderId))
	}
	if m.Timestamp != 0 {
		n += 1 + sovSettlement(uint64(m.Timestamp))
	}
	if m.Height != 0 {
		n += 1 + sovSettlement(uint64(m.Height))
	}
	if m.SettlementId != 0 {
		n += 1 + sovSettlement(uint64(m.SettlementId))
	}
	return n
}

func (m *Settlements) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Epoch != 0 {
		n += 1 + sovSettlement(uint64(m.Epoch))
	}
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovSettlement(uint64(l))
		}
	}
	return n
}

func sovSettlement(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSettlement(x uint64) (n int) {
	return sovSettlement(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SettlementEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSettlement
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
			return fmt.Errorf("proto: SettlementEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SettlementEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PriceDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AssetDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quantity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecutionCostOrProceed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExecutionCostOrProceed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpectedCostOrProceed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExpectedCostOrProceed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PositionDirection", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PositionDirection = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			m.OrderId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SettlementId", wireType)
			}
			m.SettlementId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SettlementId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSettlement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSettlement
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
func (m *Settlements) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSettlement
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
			return fmt.Errorf("proto: Settlements: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Settlements: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSettlement
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
				return ErrInvalidLengthSettlement
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSettlement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, &SettlementEntry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSettlement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSettlement
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
func skipSettlement(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSettlement
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
					return 0, ErrIntOverflowSettlement
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
					return 0, ErrIntOverflowSettlement
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
				return 0, ErrInvalidLengthSettlement
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSettlement
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSettlement
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSettlement        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSettlement          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSettlement = fmt.Errorf("proto: unexpected end of group")
)
