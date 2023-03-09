// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qoracle/pool.proto

package types

import (
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Pool defines the generalized structure of a liquidity pool coming from any source chain to qoracle.
type Pool struct {
	Id        string                                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Assets    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=assets,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"assets" yaml:"token"`
	TVL       github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,3,opt,name=tvl,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tvl"`
	APY       github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,4,opt,name=apy,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"apy"`
	Raw       *types1.Any                              `protobuf:"bytes,5,opt,name=raw,proto3" json:"raw,omitempty"`
	UpdatedAt time.Time                                `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3,stdtime" json:"updated_at"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_e83c544ee23a2441, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Pool) GetAssets() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *Pool) GetRaw() *types1.Any {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *Pool) GetUpdatedAt() time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Pool)(nil), "quasarlabs.quasarnode.qoracle.Pool")
}

func init() { proto.RegisterFile("qoracle/pool.proto", fileDescriptor_e83c544ee23a2441) }

var fileDescriptor_e83c544ee23a2441 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x8e, 0xd3, 0x40,
	0x10, 0x86, 0xbd, 0xf1, 0x11, 0x71, 0x7b, 0x27, 0x0a, 0xeb, 0x0a, 0x5f, 0x24, 0x6c, 0x2b, 0xc5,
	0xc9, 0x0d, 0xbb, 0xdc, 0x21, 0x1a, 0x2a, 0xe2, 0x43, 0x42, 0x42, 0x14, 0x27, 0x2b, 0x42, 0x82,
	0x06, 0xad, 0xed, 0xc5, 0x58, 0xb1, 0x3d, 0x8e, 0x77, 0x1d, 0xf0, 0x2b, 0x50, 0xe5, 0x39, 0x78,
	0x92, 0x94, 0x29, 0x11, 0x85, 0x83, 0x9c, 0x37, 0xe0, 0x09, 0xd0, 0xda, 0x8e, 0x82, 0xa0, 0x41,
	0x54, 0x3b, 0xab, 0xf9, 0xe7, 0x9b, 0xd9, 0xd9, 0x1f, 0x1b, 0x4b, 0x28, 0x59, 0x98, 0x72, 0x5a,
	0x00, 0xa4, 0xa4, 0x28, 0x41, 0x82, 0xf1, 0x70, 0x59, 0x31, 0xc1, 0xca, 0x94, 0x05, 0x82, 0xf4,
	0x61, 0x0e, 0x11, 0x27, 0x83, 0x72, 0x72, 0x11, 0x43, 0x0c, 0x9d, 0x92, 0xaa, 0xa8, 0x2f, 0x9a,
	0x5c, 0xc6, 0x00, 0xb1, 0xe2, 0xa8, 0x5b, 0x50, 0x7d, 0xa0, 0x2c, 0xaf, 0x87, 0x94, 0xfd, 0x67,
	0x4a, 0x26, 0x19, 0x17, 0x92, 0x65, 0xc5, 0x20, 0xb0, 0x42, 0x10, 0x19, 0x08, 0x1a, 0x30, 0xc1,
	0xe9, 0xea, 0x3a, 0xe0, 0x92, 0x5d, 0xd3, 0x10, 0x92, 0xbc, 0xcf, 0x4f, 0xbf, 0xe8, 0xf8, 0xe4,
	0x0e, 0x20, 0x35, 0x1e, 0xe0, 0x51, 0x12, 0x99, 0xc8, 0x41, 0xee, 0xa9, 0x3f, 0x4a, 0x22, 0x43,
	0xe0, 0x31, 0x13, 0x82, 0x4b, 0x61, 0x8e, 0x1c, 0xdd, 0x3d, 0xbb, 0xb9, 0x24, 0x3d, 0x89, 0x28,
	0x12, 0x19, 0x48, 0xe4, 0x16, 0x92, 0xdc, 0x7b, 0xbe, 0x69, 0x6c, 0xed, 0x67, 0x63, 0x9f, 0xd7,
	0x2c, 0x4b, 0x9f, 0x4d, 0x25, 0x2c, 0x78, 0x3e, 0xfd, 0xba, 0xb3, 0xdd, 0x38, 0x91, 0x1f, 0xab,
	0x80, 0x84, 0x90, 0xd1, 0x61, 0x8c, 0xfe, 0x78, 0x24, 0xa2, 0x05, 0x95, 0x75, 0xc1, 0x45, 0x07,
	0x10, 0xfe, 0xd0, 0xca, 0x78, 0x89, 0x75, 0xb9, 0x4a, 0x4d, 0xdd, 0x41, 0xee, 0xb9, 0xf7, 0x54,
	0x61, 0xbf, 0x37, 0xf6, 0xd5, 0x3f, 0x60, 0x5e, 0xf0, 0xb0, 0x6d, 0x6c, 0x7d, 0xfe, 0xe6, 0xb5,
	0xaf, 0x08, 0x0a, 0xc4, 0x8a, 0xda, 0x3c, 0xf9, 0x5f, 0xd0, 0xec, 0xee, 0xad, 0xaf, 0x08, 0xc6,
	0x15, 0xd6, 0x4b, 0xf6, 0xc9, 0xbc, 0xe7, 0x20, 0xf7, 0xec, 0xe6, 0x82, 0xf4, 0xeb, 0x26, 0x87,
	0x75, 0x93, 0x59, 0x5e, 0xfb, 0x4a, 0x60, 0xdc, 0x62, 0x5c, 0x15, 0x11, 0x93, 0x3c, 0x7a, 0xcf,
	0xa4, 0x39, 0xee, 0xe4, 0x93, 0xbf, 0xe4, 0xf3, 0xc3, 0xef, 0x78, 0xf7, 0xd5, 0x4c, 0xeb, 0x9d,
	0x8d, 0xfc, 0xd3, 0xa1, 0x6e, 0x26, 0xbd, 0x57, 0x9b, 0xd6, 0x42, 0xdb, 0xd6, 0x42, 0x3f, 0x5a,
	0x0b, 0xad, 0xf7, 0x96, 0xb6, 0xdd, 0x5b, 0xda, 0xb7, 0xbd, 0xa5, 0xbd, 0x7b, 0xfc, 0xdb, 0xe8,
	0x47, 0x0b, 0xd1, 0xa3, 0x85, 0xe8, 0x67, 0x7a, 0xb0, 0x5b, 0xf7, 0x90, 0x60, 0xdc, 0x35, 0x7d,
	0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xac, 0x8c, 0xb7, 0xc4, 0x86, 0x02, 0x00, 0x00,
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.UpdatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPool(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	if m.Raw != nil {
		{
			size, err := m.Raw.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	{
		size := m.APY.Size()
		i -= size
		if _, err := m.APY.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.TVL.Size()
		i -= size
		if _, err := m.TVL.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Assets) > 0 {
		for iNdEx := len(m.Assets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Assets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	if len(m.Assets) > 0 {
		for _, e := range m.Assets {
			l = e.Size()
			n += 1 + l + sovPool(uint64(l))
		}
	}
	l = m.TVL.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.APY.Size()
	n += 1 + l + sovPool(uint64(l))
	if m.Raw != nil {
		l = m.Raw.Size()
		n += 1 + l + sovPool(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt)
	n += 1 + l + sovPool(uint64(l))
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Assets = append(m.Assets, types.Coin{})
			if err := m.Assets[len(m.Assets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TVL", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TVL.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field APY", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.APY.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Raw == nil {
				m.Raw = &types1.Any{}
			}
			if err := m.Raw.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)