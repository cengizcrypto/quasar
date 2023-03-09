// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qoracle/bandchain/bandchain.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// OracleScriptState defines the state of an oracle script which keeps track of oracle script request/response.
type OracleScriptState struct {
	ClientId              string     `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	CallData              *types.Any `protobuf:"bytes,2,opt,name=call_data,json=callData,proto3" json:"call_data,omitempty"`
	RequestPacketSequence uint64     `protobuf:"varint,3,opt,name=request_packet_sequence,json=requestPacketSequence,proto3" json:"request_packet_sequence,omitempty"`
	OracleRequestId       uint64     `protobuf:"varint,4,opt,name=oracle_request_id,json=oracleRequestId,proto3" json:"oracle_request_id,omitempty"`
	ResultPacketSequence  uint64     `protobuf:"varint,5,opt,name=result_packet_sequence,json=resultPacketSequence,proto3" json:"result_packet_sequence,omitempty"`
	Result                *types.Any `protobuf:"bytes,6,opt,name=result,proto3" json:"result,omitempty"`
	Failed                bool       `protobuf:"varint,7,opt,name=failed,proto3" json:"failed,omitempty"`
	StartedAtHeight       int64      `protobuf:"varint,8,opt,name=started_at_height,json=startedAtHeight,proto3" json:"started_at_height,omitempty"`
	UpdatedAtHeight       int64      `protobuf:"varint,9,opt,name=updated_at_height,json=updatedAtHeight,proto3" json:"updated_at_height,omitempty"`
}

func (m *OracleScriptState) Reset()         { *m = OracleScriptState{} }
func (m *OracleScriptState) String() string { return proto.CompactTextString(m) }
func (*OracleScriptState) ProtoMessage()    {}
func (*OracleScriptState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b19916fd11fc094c, []int{0}
}
func (m *OracleScriptState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleScriptState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleScriptState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleScriptState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleScriptState.Merge(m, src)
}
func (m *OracleScriptState) XXX_Size() int {
	return m.Size()
}
func (m *OracleScriptState) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleScriptState.DiscardUnknown(m)
}

var xxx_messageInfo_OracleScriptState proto.InternalMessageInfo

func (m *OracleScriptState) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *OracleScriptState) GetCallData() *types.Any {
	if m != nil {
		return m.CallData
	}
	return nil
}

func (m *OracleScriptState) GetRequestPacketSequence() uint64 {
	if m != nil {
		return m.RequestPacketSequence
	}
	return 0
}

func (m *OracleScriptState) GetOracleRequestId() uint64 {
	if m != nil {
		return m.OracleRequestId
	}
	return 0
}

func (m *OracleScriptState) GetResultPacketSequence() uint64 {
	if m != nil {
		return m.ResultPacketSequence
	}
	return 0
}

func (m *OracleScriptState) GetResult() *types.Any {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *OracleScriptState) GetFailed() bool {
	if m != nil {
		return m.Failed
	}
	return false
}

func (m *OracleScriptState) GetStartedAtHeight() int64 {
	if m != nil {
		return m.StartedAtHeight
	}
	return 0
}

func (m *OracleScriptState) GetUpdatedAtHeight() int64 {
	if m != nil {
		return m.UpdatedAtHeight
	}
	return 0
}

// CoinRatesCallData is the call data for coin rates script
type CoinRatesCallData struct {
	Symbols    []string `protobuf:"bytes,1,rep,name=symbols,proto3" json:"symbols,omitempty"`
	Multiplier uint64   `protobuf:"varint,2,opt,name=multiplier,proto3" json:"multiplier,omitempty"`
}

func (m *CoinRatesCallData) Reset()         { *m = CoinRatesCallData{} }
func (m *CoinRatesCallData) String() string { return proto.CompactTextString(m) }
func (*CoinRatesCallData) ProtoMessage()    {}
func (*CoinRatesCallData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b19916fd11fc094c, []int{1}
}
func (m *CoinRatesCallData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoinRatesCallData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoinRatesCallData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoinRatesCallData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinRatesCallData.Merge(m, src)
}
func (m *CoinRatesCallData) XXX_Size() int {
	return m.Size()
}
func (m *CoinRatesCallData) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinRatesCallData.DiscardUnknown(m)
}

var xxx_messageInfo_CoinRatesCallData proto.InternalMessageInfo

func (m *CoinRatesCallData) GetSymbols() []string {
	if m != nil {
		return m.Symbols
	}
	return nil
}

func (m *CoinRatesCallData) GetMultiplier() uint64 {
	if m != nil {
		return m.Multiplier
	}
	return 0
}

// CoinRatesResult is the result for coin rates script
type CoinRatesResult struct {
	Rates []uint64 `protobuf:"varint,1,rep,packed,name=rates,proto3" json:"rates,omitempty"`
}

func (m *CoinRatesResult) Reset()         { *m = CoinRatesResult{} }
func (m *CoinRatesResult) String() string { return proto.CompactTextString(m) }
func (*CoinRatesResult) ProtoMessage()    {}
func (*CoinRatesResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_b19916fd11fc094c, []int{2}
}
func (m *CoinRatesResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoinRatesResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoinRatesResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoinRatesResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinRatesResult.Merge(m, src)
}
func (m *CoinRatesResult) XXX_Size() int {
	return m.Size()
}
func (m *CoinRatesResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinRatesResult.DiscardUnknown(m)
}

var xxx_messageInfo_CoinRatesResult proto.InternalMessageInfo

func (m *CoinRatesResult) GetRates() []uint64 {
	if m != nil {
		return m.Rates
	}
	return nil
}

type OraclePrices struct {
	Prices          github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=prices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"prices"`
	UpdatedAtHeight int64                                       `protobuf:"varint,2,opt,name=updated_at_height,json=updatedAtHeight,proto3" json:"updated_at_height,omitempty"`
}

func (m *OraclePrices) Reset()         { *m = OraclePrices{} }
func (m *OraclePrices) String() string { return proto.CompactTextString(m) }
func (*OraclePrices) ProtoMessage()    {}
func (*OraclePrices) Descriptor() ([]byte, []int) {
	return fileDescriptor_b19916fd11fc094c, []int{3}
}
func (m *OraclePrices) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OraclePrices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OraclePrices.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OraclePrices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OraclePrices.Merge(m, src)
}
func (m *OraclePrices) XXX_Size() int {
	return m.Size()
}
func (m *OraclePrices) XXX_DiscardUnknown() {
	xxx_messageInfo_OraclePrices.DiscardUnknown(m)
}

var xxx_messageInfo_OraclePrices proto.InternalMessageInfo

func (m *OraclePrices) GetPrices() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Prices
	}
	return nil
}

func (m *OraclePrices) GetUpdatedAtHeight() int64 {
	if m != nil {
		return m.UpdatedAtHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*OracleScriptState)(nil), "quasarlabs.quasarnode.qoracle.bandchain.OracleScriptState")
	proto.RegisterType((*CoinRatesCallData)(nil), "quasarlabs.quasarnode.qoracle.bandchain.CoinRatesCallData")
	proto.RegisterType((*CoinRatesResult)(nil), "quasarlabs.quasarnode.qoracle.bandchain.CoinRatesResult")
	proto.RegisterType((*OraclePrices)(nil), "quasarlabs.quasarnode.qoracle.bandchain.OraclePrices")
}

func init() { proto.RegisterFile("qoracle/bandchain/bandchain.proto", fileDescriptor_b19916fd11fc094c) }

var fileDescriptor_b19916fd11fc094c = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xbd, 0x6e, 0xdb, 0x3c,
	0x14, 0xb5, 0x62, 0xc7, 0xb1, 0x99, 0x0f, 0x08, 0x22, 0xf8, 0x4b, 0xd5, 0xb4, 0x50, 0x54, 0x2f,
	0x11, 0xfa, 0x43, 0xc2, 0x49, 0xd1, 0xa1, 0x5b, 0x7e, 0x86, 0x66, 0x28, 0x1a, 0xd0, 0x9d, 0xba,
	0x08, 0x14, 0xc5, 0xc8, 0x44, 0x64, 0x51, 0x16, 0xa9, 0xa2, 0x7e, 0x8b, 0xee, 0x9d, 0xbb, 0xf4,
	0x49, 0x32, 0x66, 0xec, 0xd4, 0x16, 0xf6, 0x8b, 0x14, 0x22, 0xa9, 0x38, 0x6d, 0x82, 0x4e, 0xba,
	0xf7, 0x9e, 0x73, 0x78, 0xae, 0x2e, 0x2f, 0xc1, 0x93, 0x99, 0x28, 0x09, 0xcd, 0x18, 0x8a, 0x49,
	0x9e, 0xd0, 0x09, 0xe1, 0xf9, 0x2a, 0x82, 0x45, 0x29, 0x94, 0x70, 0xf7, 0x67, 0x15, 0x91, 0xa4,
	0xcc, 0x48, 0x2c, 0xa1, 0x09, 0x73, 0x91, 0x30, 0x68, 0x85, 0xf0, 0x86, 0xbe, 0x3b, 0x48, 0x45,
	0x2a, 0xb4, 0x06, 0xd5, 0x91, 0x91, 0xef, 0x3e, 0x4c, 0x85, 0x48, 0x33, 0x86, 0x74, 0x16, 0x57,
	0x17, 0x88, 0xe4, 0x73, 0x0b, 0xf9, 0x54, 0xc8, 0xa9, 0x90, 0x28, 0x26, 0x92, 0xa1, 0x8f, 0xa3,
	0x98, 0x29, 0x32, 0x42, 0x54, 0x34, 0xce, 0xc3, 0x2f, 0x6d, 0xb0, 0xfd, 0x4e, 0xbb, 0x8c, 0x69,
	0xc9, 0x0b, 0x35, 0x56, 0x44, 0x31, 0xf7, 0x11, 0xe8, 0xd3, 0x8c, 0xb3, 0x5c, 0x45, 0x3c, 0xf1,
	0x9c, 0xc0, 0x09, 0xfb, 0xb8, 0x67, 0x0a, 0x67, 0x89, 0x3b, 0x02, 0x7d, 0x4a, 0xb2, 0x2c, 0x4a,
	0x88, 0x22, 0xde, 0x5a, 0xe0, 0x84, 0x9b, 0x07, 0x03, 0x68, 0x3a, 0x80, 0x4d, 0x07, 0xf0, 0x28,
	0x9f, 0xe3, 0x5e, 0x4d, 0x3b, 0x25, 0x8a, 0xb8, 0xaf, 0xc0, 0x83, 0x92, 0xcd, 0x2a, 0x26, 0x55,
	0x54, 0x10, 0x7a, 0xc9, 0x54, 0x24, 0xeb, 0x34, 0xa7, 0xcc, 0x6b, 0x07, 0x4e, 0xd8, 0xc1, 0xff,
	0x5b, 0xf8, 0x5c, 0xa3, 0x63, 0x0b, 0xba, 0x4f, 0xc1, 0xb6, 0x19, 0x41, 0xd4, 0xc8, 0x79, 0xe2,
	0x75, 0xb4, 0x62, 0xcb, 0x00, 0xd8, 0xd4, 0xcf, 0x12, 0xf7, 0x25, 0xd8, 0x29, 0x99, 0xac, 0xb2,
	0xbb, 0x16, 0xeb, 0x5a, 0x30, 0x30, 0xe8, 0x5f, 0x0e, 0xcf, 0x41, 0xd7, 0xd4, 0xbd, 0xee, 0x3f,
	0xfe, 0xc4, 0x72, 0xdc, 0x1d, 0xd0, 0xbd, 0x20, 0x3c, 0x63, 0x89, 0xb7, 0x11, 0x38, 0x61, 0x0f,
	0xdb, 0xac, 0xee, 0x53, 0x2a, 0x52, 0x2a, 0x96, 0x44, 0x44, 0x45, 0x13, 0xc6, 0xd3, 0x89, 0xf2,
	0x7a, 0x81, 0x13, 0xb6, 0xf1, 0x96, 0x05, 0x8e, 0xd4, 0x1b, 0x5d, 0xae, 0xb9, 0x55, 0x91, 0x90,
	0x3f, 0xb9, 0x7d, 0xc3, 0xb5, 0x40, 0xc3, 0x1d, 0xbe, 0x05, 0xdb, 0x27, 0x82, 0xe7, 0x98, 0x28,
	0x26, 0x4f, 0x9a, 0x61, 0x7a, 0x60, 0x43, 0xce, 0xa7, 0xb1, 0xc8, 0xa4, 0xe7, 0x04, 0xed, 0xb0,
	0x8f, 0x9b, 0xd4, 0xf5, 0x01, 0x98, 0x56, 0x99, 0xe2, 0x45, 0xc6, 0x59, 0xa9, 0xaf, 0xa6, 0x83,
	0x6f, 0x55, 0x86, 0xfb, 0x60, 0xeb, 0xe6, 0x38, 0x6c, 0xfe, 0x68, 0x00, 0xd6, 0xcb, 0x3a, 0xd5,
	0x47, 0x75, 0xb0, 0x49, 0x86, 0x5f, 0x1d, 0xf0, 0x9f, 0xd9, 0x8a, 0xf3, 0x92, 0x53, 0x26, 0x5d,
	0x0e, 0xba, 0x85, 0x8e, 0x34, 0x6f, 0xf3, 0xe0, 0x31, 0x34, 0x7b, 0x05, 0xeb, 0xbd, 0x82, 0x76,
	0xaf, 0xe0, 0x29, 0xa3, 0xf5, 0xf9, 0xc7, 0x87, 0x57, 0x3f, 0xf6, 0x5a, 0xdf, 0x7e, 0xee, 0x3d,
	0x4b, 0xb9, 0x9a, 0x54, 0x31, 0xa4, 0x62, 0x8a, 0xec, 0x1e, 0x9a, 0xcf, 0x0b, 0x99, 0x5c, 0x22,
	0x35, 0x2f, 0x98, 0x6c, 0x34, 0x12, 0x5b, 0x83, 0xfb, 0xe7, 0xb3, 0x76, 0xef, 0x7c, 0x8e, 0xdf,
	0x5f, 0x2d, 0x7c, 0xe7, 0x7a, 0xe1, 0x3b, 0xbf, 0x16, 0xbe, 0xf3, 0x79, 0xe9, 0xb7, 0xae, 0x97,
	0x7e, 0xeb, 0xfb, 0xd2, 0x6f, 0x7d, 0x78, 0x7d, 0xcb, 0x7a, 0xf5, 0xb8, 0xd0, 0xea, 0x71, 0xa1,
	0x4f, 0xe8, 0xee, 0xbb, 0xd4, 0x2d, 0xc5, 0x5d, 0x7d, 0xf7, 0x87, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xc7, 0xec, 0x79, 0x0b, 0xb9, 0x03, 0x00, 0x00,
}

func (m *OracleScriptState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleScriptState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleScriptState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdatedAtHeight != 0 {
		i = encodeVarintBandchain(dAtA, i, uint64(m.UpdatedAtHeight))
		i--
		dAtA[i] = 0x48
	}
	if m.StartedAtHeight != 0 {
		i = encodeVarintBandchain(dAtA, i, uint64(m.StartedAtHeight))
		i--
		dAtA[i] = 0x40
	}
	if m.Failed {
		i--
		if m.Failed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.Result != nil {
		{
			size, err := m.Result.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBandchain(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.ResultPacketSequence != 0 {
		i = encodeVarintBandchain(dAtA, i, uint64(m.ResultPacketSequence))
		i--
		dAtA[i] = 0x28
	}
	if m.OracleRequestId != 0 {
		i = encodeVarintBandchain(dAtA, i, uint64(m.OracleRequestId))
		i--
		dAtA[i] = 0x20
	}
	if m.RequestPacketSequence != 0 {
		i = encodeVarintBandchain(dAtA, i, uint64(m.RequestPacketSequence))
		i--
		dAtA[i] = 0x18
	}
	if m.CallData != nil {
		{
			size, err := m.CallData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBandchain(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintBandchain(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CoinRatesCallData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoinRatesCallData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoinRatesCallData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Multiplier != 0 {
		i = encodeVarintBandchain(dAtA, i, uint64(m.Multiplier))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Symbols) > 0 {
		for iNdEx := len(m.Symbols) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Symbols[iNdEx])
			copy(dAtA[i:], m.Symbols[iNdEx])
			i = encodeVarintBandchain(dAtA, i, uint64(len(m.Symbols[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CoinRatesResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoinRatesResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoinRatesResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rates) > 0 {
		dAtA4 := make([]byte, len(m.Rates)*10)
		var j3 int
		for _, num := range m.Rates {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintBandchain(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OraclePrices) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OraclePrices) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OraclePrices) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UpdatedAtHeight != 0 {
		i = encodeVarintBandchain(dAtA, i, uint64(m.UpdatedAtHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Prices) > 0 {
		for iNdEx := len(m.Prices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Prices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBandchain(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintBandchain(dAtA []byte, offset int, v uint64) int {
	offset -= sovBandchain(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OracleScriptState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovBandchain(uint64(l))
	}
	if m.CallData != nil {
		l = m.CallData.Size()
		n += 1 + l + sovBandchain(uint64(l))
	}
	if m.RequestPacketSequence != 0 {
		n += 1 + sovBandchain(uint64(m.RequestPacketSequence))
	}
	if m.OracleRequestId != 0 {
		n += 1 + sovBandchain(uint64(m.OracleRequestId))
	}
	if m.ResultPacketSequence != 0 {
		n += 1 + sovBandchain(uint64(m.ResultPacketSequence))
	}
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovBandchain(uint64(l))
	}
	if m.Failed {
		n += 2
	}
	if m.StartedAtHeight != 0 {
		n += 1 + sovBandchain(uint64(m.StartedAtHeight))
	}
	if m.UpdatedAtHeight != 0 {
		n += 1 + sovBandchain(uint64(m.UpdatedAtHeight))
	}
	return n
}

func (m *CoinRatesCallData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Symbols) > 0 {
		for _, s := range m.Symbols {
			l = len(s)
			n += 1 + l + sovBandchain(uint64(l))
		}
	}
	if m.Multiplier != 0 {
		n += 1 + sovBandchain(uint64(m.Multiplier))
	}
	return n
}

func (m *CoinRatesResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rates) > 0 {
		l = 0
		for _, e := range m.Rates {
			l += sovBandchain(uint64(e))
		}
		n += 1 + sovBandchain(uint64(l)) + l
	}
	return n
}

func (m *OraclePrices) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Prices) > 0 {
		for _, e := range m.Prices {
			l = e.Size()
			n += 1 + l + sovBandchain(uint64(l))
		}
	}
	if m.UpdatedAtHeight != 0 {
		n += 1 + sovBandchain(uint64(m.UpdatedAtHeight))
	}
	return n
}

func sovBandchain(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBandchain(x uint64) (n int) {
	return sovBandchain(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OracleScriptState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBandchain
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
			return fmt.Errorf("proto: OracleScriptState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleScriptState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandchain
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
				return ErrInvalidLengthBandchain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBandchain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandchain
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
				return ErrInvalidLengthBandchain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBandchain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CallData == nil {
				m.CallData = &types.Any{}
			}
			if err := m.CallData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestPacketSequence", wireType)
			}
			m.RequestPacketSequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestPacketSequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OracleRequestId", wireType)
			}
			m.OracleRequestId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OracleRequestId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultPacketSequence", wireType)
			}
			m.ResultPacketSequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResultPacketSequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandchain
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
				return ErrInvalidLengthBandchain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBandchain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Result == nil {
				m.Result = &types.Any{}
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandchain
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
			m.Failed = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartedAtHeight", wireType)
			}
			m.StartedAtHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartedAtHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAtHeight", wireType)
			}
			m.UpdatedAtHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAtHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBandchain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBandchain
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
func (m *CoinRatesCallData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBandchain
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
			return fmt.Errorf("proto: CoinRatesCallData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoinRatesCallData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbols", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandchain
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
				return ErrInvalidLengthBandchain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBandchain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbols = append(m.Symbols, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
			}
			m.Multiplier = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Multiplier |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBandchain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBandchain
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
func (m *CoinRatesResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBandchain
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
			return fmt.Errorf("proto: CoinRatesResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoinRatesResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBandchain
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Rates = append(m.Rates, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBandchain
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthBandchain
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthBandchain
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Rates) == 0 {
					m.Rates = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBandchain
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Rates = append(m.Rates, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Rates", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBandchain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBandchain
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
func (m *OraclePrices) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBandchain
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
			return fmt.Errorf("proto: OraclePrices: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OraclePrices: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandchain
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
				return ErrInvalidLengthBandchain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBandchain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prices = append(m.Prices, types1.DecCoin{})
			if err := m.Prices[len(m.Prices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAtHeight", wireType)
			}
			m.UpdatedAtHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBandchain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdatedAtHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBandchain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBandchain
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
func skipBandchain(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBandchain
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
					return 0, ErrIntOverflowBandchain
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
					return 0, ErrIntOverflowBandchain
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
				return 0, ErrInvalidLengthBandchain
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBandchain
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBandchain
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBandchain        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBandchain          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBandchain = fmt.Errorf("proto: unexpected end of group")
)