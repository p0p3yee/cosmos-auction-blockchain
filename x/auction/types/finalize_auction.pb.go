// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auction/auction/finalize_auction.proto

package types

import (
	fmt "fmt"
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

type FinalizeAuction struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id         uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	AuctionId  uint64 `protobuf:"varint,3,opt,name=auctionId,proto3" json:"auctionId,omitempty"`
	BidId      uint64 `protobuf:"varint,4,opt,name=bidId,proto3" json:"bidId,omitempty"`
	FinalPrice uint64 `protobuf:"varint,5,opt,name=finalPrice,proto3" json:"finalPrice,omitempty"`
}

func (m *FinalizeAuction) Reset()         { *m = FinalizeAuction{} }
func (m *FinalizeAuction) String() string { return proto.CompactTextString(m) }
func (*FinalizeAuction) ProtoMessage()    {}
func (*FinalizeAuction) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a04c5586b41bb2a, []int{0}
}
func (m *FinalizeAuction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FinalizeAuction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FinalizeAuction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FinalizeAuction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinalizeAuction.Merge(m, src)
}
func (m *FinalizeAuction) XXX_Size() int {
	return m.Size()
}
func (m *FinalizeAuction) XXX_DiscardUnknown() {
	xxx_messageInfo_FinalizeAuction.DiscardUnknown(m)
}

var xxx_messageInfo_FinalizeAuction proto.InternalMessageInfo

func (m *FinalizeAuction) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *FinalizeAuction) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FinalizeAuction) GetAuctionId() uint64 {
	if m != nil {
		return m.AuctionId
	}
	return 0
}

func (m *FinalizeAuction) GetBidId() uint64 {
	if m != nil {
		return m.BidId
	}
	return 0
}

func (m *FinalizeAuction) GetFinalPrice() uint64 {
	if m != nil {
		return m.FinalPrice
	}
	return 0
}

func init() {
	proto.RegisterType((*FinalizeAuction)(nil), "auction.auction.FinalizeAuction")
}

func init() {
	proto.RegisterFile("auction/auction/finalize_auction.proto", fileDescriptor_7a04c5586b41bb2a)
}

var fileDescriptor_7a04c5586b41bb2a = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4b, 0x2c, 0x4d, 0x2e,
	0xc9, 0xcc, 0xcf, 0xd3, 0x87, 0xd1, 0x69, 0x99, 0x79, 0x89, 0x39, 0x99, 0x55, 0xa9, 0xf1, 0x50,
	0x01, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x7e, 0x18, 0x17, 0x4a, 0x2b, 0xf5, 0x33, 0x72,
	0xf1, 0xbb, 0x41, 0xd5, 0x3a, 0x42, 0xc4, 0x84, 0x24, 0xb8, 0xd8, 0x93, 0x8b, 0x52, 0x13, 0x4b,
	0xf2, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x3e, 0x2e, 0xa6, 0xcc,
	0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x96, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x19, 0x2e, 0x4e, 0xa8,
	0x41, 0x9e, 0x29, 0x12, 0xcc, 0x60, 0x61, 0x84, 0x80, 0x90, 0x08, 0x17, 0x6b, 0x52, 0x66, 0x8a,
	0x67, 0x8a, 0x04, 0x0b, 0x58, 0x06, 0xc2, 0x11, 0x92, 0xe3, 0xe2, 0x02, 0x3b, 0x2e, 0xa0, 0x28,
	0x33, 0x39, 0x55, 0x82, 0x15, 0x2c, 0x85, 0x24, 0xe2, 0x64, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3,
	0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xe2, 0x30, 0xcf, 0x55, 0xc0, 0xbd, 0x59, 0x52, 0x59, 0x90, 0x5a,
	0x9c, 0xc4, 0x06, 0xf6, 0x9c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xba, 0x6a, 0xa2, 0x32, 0x06,
	0x01, 0x00, 0x00,
}

func (m *FinalizeAuction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FinalizeAuction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FinalizeAuction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FinalPrice != 0 {
		i = encodeVarintFinalizeAuction(dAtA, i, uint64(m.FinalPrice))
		i--
		dAtA[i] = 0x28
	}
	if m.BidId != 0 {
		i = encodeVarintFinalizeAuction(dAtA, i, uint64(m.BidId))
		i--
		dAtA[i] = 0x20
	}
	if m.AuctionId != 0 {
		i = encodeVarintFinalizeAuction(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x18
	}
	if m.Id != 0 {
		i = encodeVarintFinalizeAuction(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintFinalizeAuction(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFinalizeAuction(dAtA []byte, offset int, v uint64) int {
	offset -= sovFinalizeAuction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FinalizeAuction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovFinalizeAuction(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovFinalizeAuction(uint64(m.Id))
	}
	if m.AuctionId != 0 {
		n += 1 + sovFinalizeAuction(uint64(m.AuctionId))
	}
	if m.BidId != 0 {
		n += 1 + sovFinalizeAuction(uint64(m.BidId))
	}
	if m.FinalPrice != 0 {
		n += 1 + sovFinalizeAuction(uint64(m.FinalPrice))
	}
	return n
}

func sovFinalizeAuction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFinalizeAuction(x uint64) (n int) {
	return sovFinalizeAuction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FinalizeAuction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFinalizeAuction
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
			return fmt.Errorf("proto: FinalizeAuction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FinalizeAuction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinalizeAuction
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
				return ErrInvalidLengthFinalizeAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFinalizeAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinalizeAuction
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinalizeAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidId", wireType)
			}
			m.BidId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinalizeAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BidId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalPrice", wireType)
			}
			m.FinalPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFinalizeAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinalPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFinalizeAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFinalizeAuction
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
func skipFinalizeAuction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFinalizeAuction
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
					return 0, ErrIntOverflowFinalizeAuction
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
					return 0, ErrIntOverflowFinalizeAuction
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
				return 0, ErrInvalidLengthFinalizeAuction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFinalizeAuction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFinalizeAuction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFinalizeAuction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFinalizeAuction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFinalizeAuction = fmt.Errorf("proto: unexpected end of group")
)