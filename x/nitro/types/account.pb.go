// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nitro/account.proto

package types

import (
	fmt "fmt"
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

type Account struct {
	Pubkey     string `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey"`
	Owner      string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner"`
	Lamports   uint64 `protobuf:"varint,3,opt,name=lamports,proto3" json:"lamports"`
	Slot       uint64 `protobuf:"varint,4,opt,name=slot,proto3" json:"slot"`
	Executable bool   `protobuf:"varint,5,opt,name=executable,proto3" json:"executable"`
	RentEpoch  uint64 `protobuf:"varint,6,opt,name=rent_epoch,json=rentEpoch,proto3" json:"rent_epoch"`
	Data       string `protobuf:"bytes,7,opt,name=data,proto3" json:"data"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_8baae6d592388adf, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *Account) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Account) GetLamports() uint64 {
	if m != nil {
		return m.Lamports
	}
	return 0
}

func (m *Account) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *Account) GetExecutable() bool {
	if m != nil {
		return m.Executable
	}
	return false
}

func (m *Account) GetRentEpoch() uint64 {
	if m != nil {
		return m.RentEpoch
	}
	return 0
}

func (m *Account) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "seiprotocol.seichain.nitro.Account")
}

func init() { proto.RegisterFile("nitro/account.proto", fileDescriptor_8baae6d592388adf) }

var fileDescriptor_8baae6d592388adf = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xbf, 0x4e, 0xfb, 0x30,
	0x10, 0xc7, 0xeb, 0xfe, 0xfa, 0xd7, 0xfa, 0x89, 0xc1, 0x30, 0x58, 0x15, 0x4a, 0xaa, 0x4e, 0x59,
	0x1a, 0x0f, 0x3c, 0x01, 0x91, 0x18, 0x58, 0x3d, 0xb2, 0x20, 0xc7, 0x58, 0xad, 0x45, 0x9a, 0x8b,
	0x62, 0x47, 0xb4, 0x6f, 0x81, 0xc4, 0x4b, 0x31, 0x76, 0x64, 0x8a, 0x50, 0xbb, 0xe5, 0x29, 0x50,
	0x2e, 0xa5, 0x74, 0xb9, 0xfb, 0xdc, 0xd7, 0x5f, 0xdf, 0xc9, 0x67, 0x7a, 0x9d, 0x5b, 0x5f, 0x82,
	0x50, 0x5a, 0x43, 0x95, 0xfb, 0xb8, 0x28, 0xc1, 0x03, 0x9b, 0x39, 0x63, 0x91, 0x34, 0x64, 0xb1,
	0x33, 0x56, 0xaf, 0x95, 0xcd, 0x63, 0x74, 0xce, 0x6e, 0x56, 0xb0, 0x02, 0x3c, 0x14, 0x2d, 0x75,
	0x37, 0x16, 0x1f, 0x7d, 0x3a, 0xbe, 0xef, 0x7a, 0xb0, 0x05, 0x1d, 0x15, 0x55, 0xfa, 0x6a, 0x76,
	0x9c, 0xcc, 0x49, 0x34, 0x4d, 0x68, 0x53, 0x87, 0x27, 0x45, 0x9e, 0x32, 0x0b, 0xe9, 0x10, 0xde,
	0x72, 0x53, 0xf2, 0x3e, 0x5a, 0xa6, 0x4d, 0x1d, 0x76, 0x82, 0xec, 0x12, 0x8b, 0xe8, 0x24, 0x53,
	0x9b, 0x02, 0x4a, 0xef, 0xf8, 0xbf, 0x39, 0x89, 0x06, 0xc9, 0xff, 0xa6, 0x0e, 0xcf, 0x9a, 0x3c,
	0x13, 0xbb, 0xa5, 0x03, 0x97, 0x81, 0xe7, 0x03, 0x74, 0x4d, 0x9a, 0x3a, 0xc4, 0x5a, 0x62, 0x64,
	0x31, 0xa5, 0x66, 0x6b, 0x74, 0xe5, 0x55, 0x9a, 0x19, 0x3e, 0x9c, 0x93, 0x68, 0x92, 0x5c, 0x35,
	0x75, 0x78, 0xa1, 0xca, 0x0b, 0x66, 0x4b, 0x4a, 0x4b, 0x93, 0xfb, 0x67, 0x53, 0x80, 0x5e, 0xf3,
	0x11, 0xf6, 0x44, 0xff, 0x9f, 0x2a, 0xa7, 0x2d, 0x3f, 0xb4, 0xd8, 0x0e, 0x7f, 0x51, 0x5e, 0xf1,
	0x31, 0x3e, 0x03, 0x87, 0xb7, 0xb5, 0xc4, 0x98, 0x3c, 0x7e, 0x1e, 0x02, 0xb2, 0x3f, 0x04, 0xe4,
	0xfb, 0x10, 0x90, 0xf7, 0x63, 0xd0, 0xdb, 0x1f, 0x83, 0xde, 0xd7, 0x31, 0xe8, 0x3d, 0x89, 0x95,
	0xf5, 0xeb, 0x2a, 0x8d, 0x35, 0x6c, 0x84, 0x33, 0x76, 0xf9, 0xbb, 0x6d, 0x2c, 0x70, 0xdd, 0x62,
	0x2b, 0xba, 0xaf, 0xf1, 0xbb, 0xc2, 0xb8, 0x74, 0x84, 0x8e, 0xbb, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x11, 0xcc, 0x58, 0xfb, 0xb0, 0x01, 0x00, 0x00,
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x3a
	}
	if m.RentEpoch != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.RentEpoch))
		i--
		dAtA[i] = 0x30
	}
	if m.Executable {
		i--
		if m.Executable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Slot != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.Slot))
		i--
		dAtA[i] = 0x20
	}
	if m.Lamports != 0 {
		i = encodeVarintAccount(dAtA, i, uint64(m.Lamports))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintAccount(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	if m.Lamports != 0 {
		n += 1 + sovAccount(uint64(m.Lamports))
	}
	if m.Slot != 0 {
		n += 1 + sovAccount(uint64(m.Slot))
	}
	if m.Executable {
		n += 2
	}
	if m.RentEpoch != 0 {
		n += 1 + sovAccount(uint64(m.RentEpoch))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovAccount(uint64(l))
	}
	return n
}

func sovAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccount(x uint64) (n int) {
	return sovAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccount
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lamports", wireType)
			}
			m.Lamports = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lamports |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Slot", wireType)
			}
			m.Slot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Slot |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Executable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
			m.Executable = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RentEpoch", wireType)
			}
			m.RentEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RentEpoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccount
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
				return ErrInvalidLengthAccount
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccount
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
func skipAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
					return 0, ErrIntOverflowAccount
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
				return 0, ErrInvalidLengthAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccount = fmt.Errorf("proto: unexpected end of group")
)