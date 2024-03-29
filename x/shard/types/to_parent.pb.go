// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gvm/shard/to_parent.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type ToParent struct {
	Id     uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender string     `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Amount types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
	Info   string     `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
}

func (m *ToParent) Reset()         { *m = ToParent{} }
func (m *ToParent) String() string { return proto.CompactTextString(m) }
func (*ToParent) ProtoMessage()    {}
func (*ToParent) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d2bbb1ff98fa342, []int{0}
}
func (m *ToParent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ToParent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ToParent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ToParent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToParent.Merge(m, src)
}
func (m *ToParent) XXX_Size() int {
	return m.Size()
}
func (m *ToParent) XXX_DiscardUnknown() {
	xxx_messageInfo_ToParent.DiscardUnknown(m)
}

var xxx_messageInfo_ToParent proto.InternalMessageInfo

func (m *ToParent) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ToParent) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *ToParent) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *ToParent) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto.RegisterType((*ToParent)(nil), "gvm.shard.ToParent")
}

func init() { proto.RegisterFile("gvm/shard/to_parent.proto", fileDescriptor_2d2bbb1ff98fa342) }

var fileDescriptor_2d2bbb1ff98fa342 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0xcf, 0x4a, 0xc4, 0x30,
	0x18, 0xc4, 0x9b, 0x5a, 0x8a, 0x1b, 0xc1, 0x43, 0x10, 0xe9, 0xee, 0x21, 0x16, 0x41, 0xe8, 0xc5,
	0x84, 0xd5, 0x83, 0x47, 0x61, 0x7d, 0x01, 0x29, 0x9e, 0xbc, 0x48, 0xff, 0xc4, 0x6c, 0x0e, 0xc9,
	0x57, 0x9a, 0x6c, 0xd1, 0x93, 0xaf, 0xe0, 0x63, 0xed, 0x71, 0x8f, 0x9e, 0x44, 0xda, 0x17, 0x91,
	0xa6, 0xdd, 0xdb, 0x24, 0x33, 0xcc, 0x8f, 0x6f, 0xf0, 0x52, 0x76, 0x9a, 0xdb, 0x6d, 0xd1, 0xd6,
	0xdc, 0xc1, 0x5b, 0x53, 0xb4, 0xc2, 0x38, 0xd6, 0xb4, 0xe0, 0x80, 0x2c, 0x64, 0xa7, 0x99, 0xb7,
	0x56, 0x17, 0x12, 0x24, 0xf8, 0x5f, 0x3e, 0xaa, 0x29, 0xb0, 0xa2, 0x15, 0x58, 0x0d, 0x96, 0x97,
	0x85, 0x15, 0xbc, 0x5b, 0x97, 0xc2, 0x15, 0x6b, 0x5e, 0x81, 0x32, 0x93, 0x7f, 0xfd, 0x85, 0x4f,
	0x5f, 0xe0, 0xd9, 0x57, 0x92, 0x73, 0x1c, 0xaa, 0x3a, 0x41, 0x29, 0xca, 0xa2, 0x3c, 0x54, 0x35,
	0xb9, 0xc4, 0xb1, 0x15, 0xa6, 0x16, 0x6d, 0x12, 0xa6, 0x28, 0x5b, 0xe4, 0xf3, 0x8b, 0x3c, 0xe0,
	0xb8, 0xd0, 0xb0, 0x33, 0x2e, 0x39, 0x49, 0x51, 0x76, 0x76, 0xb7, 0x64, 0x13, 0x84, 0x8d, 0x10,
	0x36, 0x43, 0xd8, 0x13, 0x28, 0xb3, 0x89, 0xf6, 0xbf, 0x57, 0x41, 0x3e, 0xc7, 0x09, 0xc1, 0x91,
	0x32, 0xef, 0x90, 0x44, 0xbe, 0xce, 0xeb, 0xcd, 0xe3, 0xbe, 0xa7, 0xe8, 0xd0, 0x53, 0xf4, 0xd7,
	0x53, 0xf4, 0x3d, 0xd0, 0xe0, 0x30, 0xd0, 0xe0, 0x67, 0xa0, 0xc1, 0xeb, 0x8d, 0x54, 0x6e, 0xbb,
	0x2b, 0x59, 0x05, 0x9a, 0x4b, 0xe8, 0xf4, 0xad, 0x11, 0x8e, 0x8f, 0x53, 0x7c, 0x1c, 0xc7, 0xf8,
	0x6c, 0x84, 0x2d, 0x63, 0x7f, 0xc8, 0xfd, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x1b, 0xa3,
	0x04, 0x26, 0x01, 0x00, 0x00,
}

func (m *ToParent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ToParent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ToParent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Info) > 0 {
		i -= len(m.Info)
		copy(dAtA[i:], m.Info)
		i = encodeVarintToParent(dAtA, i, uint64(len(m.Info)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintToParent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintToParent(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintToParent(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintToParent(dAtA []byte, offset int, v uint64) int {
	offset -= sovToParent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ToParent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovToParent(uint64(m.Id))
	}
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovToParent(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovToParent(uint64(l))
	l = len(m.Info)
	if l > 0 {
		n += 1 + l + sovToParent(uint64(l))
	}
	return n
}

func sovToParent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozToParent(x uint64) (n int) {
	return sovToParent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ToParent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToParent
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
			return fmt.Errorf("proto: ToParent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ToParent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToParent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToParent
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
				return ErrInvalidLengthToParent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToParent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToParent
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
				return ErrInvalidLengthToParent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthToParent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToParent
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
				return ErrInvalidLengthToParent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToParent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Info = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToParent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToParent
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
func skipToParent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowToParent
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
					return 0, ErrIntOverflowToParent
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
					return 0, ErrIntOverflowToParent
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
				return 0, ErrInvalidLengthToParent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupToParent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthToParent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthToParent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowToParent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupToParent = fmt.Errorf("proto: unexpected end of group")
)
