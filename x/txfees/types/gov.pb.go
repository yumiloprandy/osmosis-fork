// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/txfees/v1beta1/gov.proto

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

// UpdateFeeTokenProposal is a gov Content type for adding a new whitelisted fee
// token. It must specify a denom along with gamm pool ID to use as a spot price
// calculator. It can be used to add a new denom to the whitelist It can also be
// used to update the Pool to associate with the denom. If Pool ID is set to 0,
// it will remove the denom from the whitelisted set.
type UpdateFeeTokenProposal struct {
	Title       string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	Feetoken    FeeToken `protobuf:"bytes,3,opt,name=feetoken,proto3" json:"feetoken" yaml:"fee_token"`
}

func (m *UpdateFeeTokenProposal) Reset()      { *m = UpdateFeeTokenProposal{} }
func (*UpdateFeeTokenProposal) ProtoMessage() {}
func (*UpdateFeeTokenProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c4a51bafc82863d, []int{0}
}
func (m *UpdateFeeTokenProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateFeeTokenProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateFeeTokenProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateFeeTokenProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFeeTokenProposal.Merge(m, src)
}
func (m *UpdateFeeTokenProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateFeeTokenProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFeeTokenProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFeeTokenProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpdateFeeTokenProposal)(nil), "osmosis.txfees.v1beta1.UpdateFeeTokenProposal")
}

func init() { proto.RegisterFile("osmosis/txfees/v1beta1/gov.proto", fileDescriptor_2c4a51bafc82863d) }

var fileDescriptor_2c4a51bafc82863d = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0x2f, 0xce, 0xcd,
	0x2f, 0xce, 0x2c, 0xd6, 0x2f, 0xa9, 0x48, 0x4b, 0x4d, 0x2d, 0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d,
	0x49, 0x34, 0xd4, 0x4f, 0xcf, 0x2f, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x83, 0xaa,
	0xd0, 0x83, 0xa8, 0xd0, 0x83, 0xaa, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd1, 0x07,
	0xb1, 0x20, 0xaa, 0xa5, 0x54, 0x71, 0x98, 0x97, 0x96, 0x9a, 0x5a, 0x92, 0x9f, 0x9d, 0x9a, 0x07,
	0x51, 0xa6, 0xf4, 0x80, 0x91, 0x4b, 0x2c, 0xb4, 0x20, 0x25, 0xb1, 0x24, 0xd5, 0x2d, 0x35, 0x35,
	0x04, 0x24, 0x11, 0x50, 0x94, 0x5f, 0x90, 0x5f, 0x9c, 0x98, 0x23, 0xa4, 0xc6, 0xc5, 0x5a, 0x92,
	0x59, 0x92, 0x93, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0x24, 0xf0, 0xe9, 0x9e, 0x3c, 0x4f,
	0x65, 0x62, 0x6e, 0x8e, 0x95, 0x12, 0x58, 0x58, 0x29, 0x08, 0x22, 0x2d, 0x64, 0xc1, 0xc5, 0x9d,
	0x92, 0x5a, 0x9c, 0x5c, 0x94, 0x59, 0x50, 0x92, 0x99, 0x9f, 0x27, 0xc1, 0x04, 0x56, 0x2d, 0xf6,
	0xe9, 0x9e, 0xbc, 0x10, 0x44, 0x35, 0x92, 0xa4, 0x52, 0x10, 0xb2, 0x52, 0xa1, 0x30, 0x2e, 0x0e,
	0x98, 0x73, 0x24, 0x98, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x14, 0xf4, 0xb0, 0x7b, 0x52, 0x0f, 0xe6,
	0x3a, 0x27, 0x89, 0x13, 0xf7, 0xe4, 0x19, 0x3e, 0xdd, 0x93, 0x17, 0x80, 0x18, 0x9e, 0x96, 0x9a,
	0x1a, 0x0f, 0x36, 0x40, 0x29, 0x08, 0x6e, 0x96, 0x15, 0x4f, 0xc7, 0x02, 0x79, 0x86, 0x19, 0x0b,
	0xe4, 0x19, 0x5e, 0x2c, 0x90, 0x67, 0x74, 0xf2, 0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63,
	0x39, 0x86, 0x28, 0xc3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xa8,
	0xbd, 0xba, 0x39, 0x89, 0x49, 0xc5, 0x30, 0x8e, 0x7e, 0x99, 0xb9, 0x7e, 0x05, 0x2c, 0x00, 0x4b,
	0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xc1, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x40,
	0x95, 0x3f, 0xf6, 0xaf, 0x01, 0x00, 0x00,
}

func (this *UpdateFeeTokenProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateFeeTokenProposal)
	if !ok {
		that2, ok := that.(UpdateFeeTokenProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if !this.Feetoken.Equal(&that1.Feetoken) {
		return false
	}
	return true
}
func (m *UpdateFeeTokenProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateFeeTokenProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateFeeTokenProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Feetoken.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UpdateFeeTokenProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = m.Feetoken.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UpdateFeeTokenProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: UpdateFeeTokenProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateFeeTokenProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feetoken", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Feetoken.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
