// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gvm/gvm/shard_block.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

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

type ShardBlock struct {
	Index      uint64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Parent     string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	LeftChild  string `protobuf:"bytes,3,opt,name=leftChild,proto3" json:"leftChild,omitempty"`
	RightChild string `protobuf:"bytes,4,opt,name=rightChild,proto3" json:"rightChild,omitempty"`
	Vdf        string `protobuf:"bytes,5,opt,name=vdf,proto3" json:"vdf,omitempty"`
}

func (m *ShardBlock) Reset()         { *m = ShardBlock{} }
func (m *ShardBlock) String() string { return proto.CompactTextString(m) }
func (*ShardBlock) ProtoMessage()    {}
func (*ShardBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d07d1b182f90bcd, []int{0}
}
func (m *ShardBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShardBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShardBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShardBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardBlock.Merge(m, src)
}
func (m *ShardBlock) XXX_Size() int {
	return m.Size()
}
func (m *ShardBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardBlock.DiscardUnknown(m)
}

var xxx_messageInfo_ShardBlock proto.InternalMessageInfo

func (m *ShardBlock) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ShardBlock) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ShardBlock) GetLeftChild() string {
	if m != nil {
		return m.LeftChild
	}
	return ""
}

func (m *ShardBlock) GetRightChild() string {
	if m != nil {
		return m.RightChild
	}
	return ""
}

func (m *ShardBlock) GetVdf() string {
	if m != nil {
		return m.Vdf
	}
	return ""
}

func init() {
	proto.RegisterType((*ShardBlock)(nil), "gvm.gvm.ShardBlock")
}

func init() { proto.RegisterFile("gvm/gvm/shard_block.proto", fileDescriptor_9d07d1b182f90bcd) }

var fileDescriptor_9d07d1b182f90bcd = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x2f, 0xcb, 0xd5,
	0x07, 0xe1, 0xe2, 0x8c, 0xc4, 0xa2, 0x94, 0xf8, 0xa4, 0x9c, 0xfc, 0xe4, 0x6c, 0xbd, 0x82, 0xa2,
	0xfc, 0x92, 0x7c, 0x21, 0xf6, 0xf4, 0xb2, 0x5c, 0xbd, 0xf4, 0xb2, 0x5c, 0xa5, 0x2e, 0x46, 0x2e,
	0xae, 0x60, 0x90, 0xb4, 0x13, 0x48, 0x56, 0x48, 0x84, 0x8b, 0x35, 0x33, 0x2f, 0x25, 0xb5, 0x42,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x25, 0x08, 0xc2, 0x11, 0x12, 0xe3, 0x62, 0x2b, 0x48, 0x2c, 0x4a,
	0xcd, 0x2b, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84, 0x64, 0xb8, 0x38, 0x73,
	0x52, 0xd3, 0x4a, 0x9c, 0x33, 0x32, 0x73, 0x52, 0x24, 0x98, 0xc1, 0x52, 0x08, 0x01, 0x21, 0x39,
	0x2e, 0xae, 0xa2, 0xcc, 0xf4, 0x0c, 0xa8, 0x34, 0x0b, 0x58, 0x1a, 0x49, 0x44, 0x48, 0x80, 0x8b,
	0xb9, 0x2c, 0x25, 0x4d, 0x82, 0x15, 0x2c, 0x01, 0x62, 0x3a, 0xd9, 0x9e, 0x78, 0x24, 0xc7, 0x78,
	0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7,
	0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x72, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e,
	0xae, 0x7e, 0x7a, 0x7e, 0x59, 0xae, 0x6e, 0x5e, 0x6a, 0x09, 0xd8, 0x6b, 0x15, 0x60, 0xb2, 0xa4,
	0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x37, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28,
	0xe3, 0x59, 0x04, 0xf8, 0x00, 0x00, 0x00,
}

func (m *ShardBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShardBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShardBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Vdf) > 0 {
		i -= len(m.Vdf)
		copy(dAtA[i:], m.Vdf)
		i = encodeVarintShardBlock(dAtA, i, uint64(len(m.Vdf)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.RightChild) > 0 {
		i -= len(m.RightChild)
		copy(dAtA[i:], m.RightChild)
		i = encodeVarintShardBlock(dAtA, i, uint64(len(m.RightChild)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.LeftChild) > 0 {
		i -= len(m.LeftChild)
		copy(dAtA[i:], m.LeftChild)
		i = encodeVarintShardBlock(dAtA, i, uint64(len(m.LeftChild)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Parent) > 0 {
		i -= len(m.Parent)
		copy(dAtA[i:], m.Parent)
		i = encodeVarintShardBlock(dAtA, i, uint64(len(m.Parent)))
		i--
		dAtA[i] = 0x12
	}
	if m.Index != 0 {
		i = encodeVarintShardBlock(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintShardBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovShardBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ShardBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovShardBlock(uint64(m.Index))
	}
	l = len(m.Parent)
	if l > 0 {
		n += 1 + l + sovShardBlock(uint64(l))
	}
	l = len(m.LeftChild)
	if l > 0 {
		n += 1 + l + sovShardBlock(uint64(l))
	}
	l = len(m.RightChild)
	if l > 0 {
		n += 1 + l + sovShardBlock(uint64(l))
	}
	l = len(m.Vdf)
	if l > 0 {
		n += 1 + l + sovShardBlock(uint64(l))
	}
	return n
}

func sovShardBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozShardBlock(x uint64) (n int) {
	return sovShardBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ShardBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShardBlock
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
			return fmt.Errorf("proto: ShardBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShardBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShardBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShardBlock
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
				return ErrInvalidLengthShardBlock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShardBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeftChild", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShardBlock
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
				return ErrInvalidLengthShardBlock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShardBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LeftChild = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RightChild", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShardBlock
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
				return ErrInvalidLengthShardBlock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShardBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RightChild = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vdf", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShardBlock
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
				return ErrInvalidLengthShardBlock
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShardBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vdf = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShardBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShardBlock
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
func skipShardBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShardBlock
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
					return 0, ErrIntOverflowShardBlock
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
					return 0, ErrIntOverflowShardBlock
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
				return 0, ErrInvalidLengthShardBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupShardBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthShardBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthShardBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShardBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupShardBlock = fmt.Errorf("proto: unexpected end of group")
)
