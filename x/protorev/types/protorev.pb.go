// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/protorev/v1beta1/protorev.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// TokenPairArbRoutes tracks all of the hot routes for a given pair of tokens
type TokenPairArbRoutes struct {
	// Stores all of the possible hot paths for a given pair of tokens
	ArbRoutes []*Route `protobuf:"bytes,1,rep,name=arb_routes,json=arbRoutes,proto3" json:"arb_routes,omitempty"`
	// Token denomination of the first asset
	TokenIn string `protobuf:"bytes,2,opt,name=token_in,json=tokenIn,proto3" json:"token_in,omitempty"`
	// Token denomination of the second asset
	TokenOut string `protobuf:"bytes,3,opt,name=token_out,json=tokenOut,proto3" json:"token_out,omitempty"`
}

func (m *TokenPairArbRoutes) Reset()         { *m = TokenPairArbRoutes{} }
func (m *TokenPairArbRoutes) String() string { return proto.CompactTextString(m) }
func (*TokenPairArbRoutes) ProtoMessage()    {}
func (*TokenPairArbRoutes) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9f2391fd9fec01, []int{0}
}
func (m *TokenPairArbRoutes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenPairArbRoutes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenPairArbRoutes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenPairArbRoutes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenPairArbRoutes.Merge(m, src)
}
func (m *TokenPairArbRoutes) XXX_Size() int {
	return m.Size()
}
func (m *TokenPairArbRoutes) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenPairArbRoutes.DiscardUnknown(m)
}

var xxx_messageInfo_TokenPairArbRoutes proto.InternalMessageInfo

func (m *TokenPairArbRoutes) GetArbRoutes() []*Route {
	if m != nil {
		return m.ArbRoutes
	}
	return nil
}

func (m *TokenPairArbRoutes) GetTokenIn() string {
	if m != nil {
		return m.TokenIn
	}
	return ""
}

func (m *TokenPairArbRoutes) GetTokenOut() string {
	if m != nil {
		return m.TokenOut
	}
	return ""
}

// Route is a hot route for a given pair of tokens
type Route struct {
	// The pool IDs that are travered in the directed cyclic graph (traversed left
	// -> right)
	Trades []*Trade `protobuf:"bytes,1,rep,name=trades,proto3" json:"trades,omitempty"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9f2391fd9fec01, []int{1}
}
func (m *Route) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Route.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return m.Size()
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetTrades() []*Trade {
	if m != nil {
		return m.Trades
	}
	return nil
}

// Trade is a single trade in a route
type Trade struct {
	// The pool IDs that are travered in the directed cyclic graph (traversed left
	// -> right)
	Pool uint64 `protobuf:"varint,1,opt,name=pool,proto3" json:"pool,omitempty"`
	// The denom of token A that is traded
	TokenIn string `protobuf:"bytes,2,opt,name=token_in,json=tokenIn,proto3" json:"token_in,omitempty"`
	// The denom of token B that is traded
	TokenOut string `protobuf:"bytes,3,opt,name=token_out,json=tokenOut,proto3" json:"token_out,omitempty"`
}

func (m *Trade) Reset()         { *m = Trade{} }
func (m *Trade) String() string { return proto.CompactTextString(m) }
func (*Trade) ProtoMessage()    {}
func (*Trade) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9f2391fd9fec01, []int{2}
}
func (m *Trade) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Trade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Trade.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Trade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trade.Merge(m, src)
}
func (m *Trade) XXX_Size() int {
	return m.Size()
}
func (m *Trade) XXX_DiscardUnknown() {
	xxx_messageInfo_Trade.DiscardUnknown(m)
}

var xxx_messageInfo_Trade proto.InternalMessageInfo

func (m *Trade) GetPool() uint64 {
	if m != nil {
		return m.Pool
	}
	return 0
}

func (m *Trade) GetTokenIn() string {
	if m != nil {
		return m.TokenIn
	}
	return ""
}

func (m *Trade) GetTokenOut() string {
	if m != nil {
		return m.TokenOut
	}
	return ""
}

// PoolStatistics contains the number of trades the module has executed after a
// swap on a given pool and the profits from the trades
type PoolStatistics struct {
	// profits is the total profit from all trades on this pool
	Profits []*types.Coin `protobuf:"bytes,1,rep,name=profits,proto3" json:"profits,omitempty"`
	// number_of_trades is the number of trades the module has executed
	NumberOfTrades github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=number_of_trades,json=numberOfTrades,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"number_of_trades"`
	// pool_id is the id of the pool
	PoolId uint64 `protobuf:"varint,3,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
}

func (m *PoolStatistics) Reset()         { *m = PoolStatistics{} }
func (m *PoolStatistics) String() string { return proto.CompactTextString(m) }
func (*PoolStatistics) ProtoMessage()    {}
func (*PoolStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e9f2391fd9fec01, []int{3}
}
func (m *PoolStatistics) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolStatistics.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolStatistics.Merge(m, src)
}
func (m *PoolStatistics) XXX_Size() int {
	return m.Size()
}
func (m *PoolStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_PoolStatistics proto.InternalMessageInfo

func (m *PoolStatistics) GetProfits() []*types.Coin {
	if m != nil {
		return m.Profits
	}
	return nil
}

func (m *PoolStatistics) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func init() {
	proto.RegisterType((*TokenPairArbRoutes)(nil), "osmosis.protorev.v1beta1.TokenPairArbRoutes")
	proto.RegisterType((*Route)(nil), "osmosis.protorev.v1beta1.Route")
	proto.RegisterType((*Trade)(nil), "osmosis.protorev.v1beta1.Trade")
	proto.RegisterType((*PoolStatistics)(nil), "osmosis.protorev.v1beta1.PoolStatistics")
}

func init() {
	proto.RegisterFile("osmosis/protorev/v1beta1/protorev.proto", fileDescriptor_1e9f2391fd9fec01)
}

var fileDescriptor_1e9f2391fd9fec01 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xcf, 0x8a, 0xd3, 0x40,
	0x18, 0xcf, 0xb8, 0xd9, 0xd6, 0x8e, 0xb0, 0xc8, 0x20, 0x98, 0xae, 0x90, 0x94, 0x3d, 0x68, 0x2f,
	0x3b, 0xa1, 0xee, 0x82, 0xe0, 0x41, 0x70, 0x05, 0xa1, 0x17, 0x77, 0x19, 0xf7, 0xa0, 0x5e, 0xc2,
	0x4c, 0x92, 0xd6, 0x61, 0xdb, 0x7c, 0x61, 0x66, 0x52, 0xf4, 0x2d, 0xc4, 0x27, 0xf0, 0x31, 0x7c,
	0x84, 0x3d, 0xee, 0x51, 0x3c, 0x14, 0x69, 0x2f, 0x3e, 0x86, 0x64, 0x66, 0x1a, 0xbd, 0x08, 0x82,
	0xa7, 0x7c, 0xbf, 0x3f, 0xf3, 0x7d, 0xbf, 0x6f, 0x26, 0xf8, 0x11, 0xe8, 0x25, 0x68, 0xa9, 0xd3,
	0x5a, 0x81, 0x01, 0x55, 0xae, 0xd2, 0xd5, 0x44, 0x94, 0x86, 0x4f, 0x3a, 0x82, 0xda, 0x82, 0x44,
	0xde, 0x48, 0x3b, 0xde, 0x1b, 0x0f, 0x87, 0xb9, 0x95, 0x32, 0x2b, 0xa4, 0x0e, 0x38, 0xd7, 0xe1,
	0xbd, 0x39, 0xcc, 0xc1, 0xf1, 0x6d, 0xe5, 0xd9, 0xd8, 0x79, 0x52, 0xc1, 0x75, 0xd9, 0x8d, 0xcb,
	0x41, 0x56, 0x4e, 0x3f, 0xfa, 0x8c, 0x30, 0xb9, 0x84, 0xab, 0xb2, 0xba, 0xe0, 0x52, 0x3d, 0x57,
	0x82, 0x41, 0x63, 0x4a, 0x4d, 0x9e, 0x61, 0xcc, 0x95, 0xc8, 0x94, 0x45, 0x11, 0x1a, 0xed, 0x8d,
	0xef, 0x3c, 0x4e, 0xe8, 0xdf, 0x62, 0x51, 0x7b, 0x8a, 0x0d, 0x78, 0x77, 0x7e, 0x88, 0x6f, 0x9b,
	0xb6, 0x6b, 0x26, 0xab, 0xe8, 0xd6, 0x08, 0x8d, 0x07, 0xac, 0x6f, 0xf1, 0xb4, 0x22, 0x0f, 0xf0,
	0xc0, 0x49, 0xd0, 0x98, 0x68, 0xcf, 0x6a, 0xce, 0x7b, 0xde, 0x98, 0xa7, 0xe1, 0xcf, 0x2f, 0x09,
	0x3a, 0x7a, 0x89, 0xf7, 0x6d, 0x1f, 0xf2, 0x04, 0xf7, 0x8c, 0xe2, 0xc5, 0xbf, 0x44, 0xb8, 0x6c,
	0x7d, 0xcc, 0xdb, 0x7d, 0x9f, 0xb7, 0x78, 0xdf, 0xd2, 0x84, 0xe0, 0xb0, 0x06, 0x58, 0x44, 0x68,
	0x84, 0xc6, 0x21, 0xb3, 0xf5, 0x7f, 0x46, 0xfc, 0x8a, 0xf0, 0xc1, 0x05, 0xc0, 0xe2, 0xb5, 0xe1,
	0x46, 0x6a, 0x23, 0x73, 0x4d, 0x4e, 0x70, 0xbf, 0x56, 0x30, 0x93, 0x66, 0x97, 0x76, 0x48, 0xfd,
	0x03, 0xb5, 0x97, 0xdf, 0x05, 0x7d, 0x01, 0xb2, 0x62, 0x3b, 0x27, 0x79, 0x83, 0xef, 0x56, 0xcd,
	0x52, 0x94, 0x2a, 0x83, 0x59, 0xe6, 0x77, 0xb5, 0x69, 0xce, 0xe8, 0xf5, 0x3a, 0x09, 0xbe, 0xaf,
	0x93, 0x87, 0x73, 0x69, 0xde, 0x37, 0x82, 0xe6, 0xb0, 0xf4, 0x0f, 0xee, 0x3f, 0xc7, 0xba, 0xb8,
	0x4a, 0xcd, 0xc7, 0xba, 0xd4, 0x74, 0x5a, 0x19, 0x76, 0xe0, 0xfa, 0x9c, 0xcf, 0xec, 0xca, 0x9a,
	0xdc, 0xc7, 0xfd, 0x76, 0xcf, 0x4c, 0x16, 0x76, 0x85, 0x90, 0xf5, 0x5a, 0x38, 0x2d, 0xce, 0x5e,
	0x5d, 0x6f, 0x62, 0x74, 0xb3, 0x89, 0xd1, 0x8f, 0x4d, 0x8c, 0x3e, 0x6d, 0xe3, 0xe0, 0x66, 0x1b,
	0x07, 0xdf, 0xb6, 0x71, 0xf0, 0xee, 0xf4, 0x8f, 0x51, 0xfe, 0xa2, 0x8f, 0x17, 0x5c, 0xe8, 0x1d,
	0x48, 0x57, 0x93, 0xd3, 0xf4, 0xc3, 0xef, 0xdf, 0xd7, 0x0e, 0x17, 0x3d, 0x8b, 0x4f, 0x7e, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xda, 0x8e, 0xa1, 0x76, 0xdf, 0x02, 0x00, 0x00,
}

func (this *TokenPairArbRoutes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenPairArbRoutes)
	if !ok {
		that2, ok := that.(TokenPairArbRoutes)
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
	if len(this.ArbRoutes) != len(that1.ArbRoutes) {
		return false
	}
	for i := range this.ArbRoutes {
		if !this.ArbRoutes[i].Equal(that1.ArbRoutes[i]) {
			return false
		}
	}
	if this.TokenIn != that1.TokenIn {
		return false
	}
	if this.TokenOut != that1.TokenOut {
		return false
	}
	return true
}
func (this *Route) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route)
	if !ok {
		that2, ok := that.(Route)
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
	if len(this.Trades) != len(that1.Trades) {
		return false
	}
	for i := range this.Trades {
		if !this.Trades[i].Equal(that1.Trades[i]) {
			return false
		}
	}
	return true
}
func (this *Trade) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Trade)
	if !ok {
		that2, ok := that.(Trade)
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
	if this.Pool != that1.Pool {
		return false
	}
	if this.TokenIn != that1.TokenIn {
		return false
	}
	if this.TokenOut != that1.TokenOut {
		return false
	}
	return true
}
func (m *TokenPairArbRoutes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenPairArbRoutes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenPairArbRoutes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenOut) > 0 {
		i -= len(m.TokenOut)
		copy(dAtA[i:], m.TokenOut)
		i = encodeVarintProtorev(dAtA, i, uint64(len(m.TokenOut)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenIn) > 0 {
		i -= len(m.TokenIn)
		copy(dAtA[i:], m.TokenIn)
		i = encodeVarintProtorev(dAtA, i, uint64(len(m.TokenIn)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ArbRoutes) > 0 {
		for iNdEx := len(m.ArbRoutes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ArbRoutes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProtorev(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Route) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Route) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Route) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Trades) > 0 {
		for iNdEx := len(m.Trades) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Trades[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProtorev(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Trade) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Trade) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Trade) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenOut) > 0 {
		i -= len(m.TokenOut)
		copy(dAtA[i:], m.TokenOut)
		i = encodeVarintProtorev(dAtA, i, uint64(len(m.TokenOut)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenIn) > 0 {
		i -= len(m.TokenIn)
		copy(dAtA[i:], m.TokenIn)
		i = encodeVarintProtorev(dAtA, i, uint64(len(m.TokenIn)))
		i--
		dAtA[i] = 0x12
	}
	if m.Pool != 0 {
		i = encodeVarintProtorev(dAtA, i, uint64(m.Pool))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PoolStatistics) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolStatistics) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolStatistics) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolId != 0 {
		i = encodeVarintProtorev(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.NumberOfTrades.Size()
		i -= size
		if _, err := m.NumberOfTrades.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProtorev(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Profits) > 0 {
		for iNdEx := len(m.Profits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Profits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProtorev(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintProtorev(dAtA []byte, offset int, v uint64) int {
	offset -= sovProtorev(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenPairArbRoutes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ArbRoutes) > 0 {
		for _, e := range m.ArbRoutes {
			l = e.Size()
			n += 1 + l + sovProtorev(uint64(l))
		}
	}
	l = len(m.TokenIn)
	if l > 0 {
		n += 1 + l + sovProtorev(uint64(l))
	}
	l = len(m.TokenOut)
	if l > 0 {
		n += 1 + l + sovProtorev(uint64(l))
	}
	return n
}

func (m *Route) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Trades) > 0 {
		for _, e := range m.Trades {
			l = e.Size()
			n += 1 + l + sovProtorev(uint64(l))
		}
	}
	return n
}

func (m *Trade) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pool != 0 {
		n += 1 + sovProtorev(uint64(m.Pool))
	}
	l = len(m.TokenIn)
	if l > 0 {
		n += 1 + l + sovProtorev(uint64(l))
	}
	l = len(m.TokenOut)
	if l > 0 {
		n += 1 + l + sovProtorev(uint64(l))
	}
	return n
}

func (m *PoolStatistics) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Profits) > 0 {
		for _, e := range m.Profits {
			l = e.Size()
			n += 1 + l + sovProtorev(uint64(l))
		}
	}
	l = m.NumberOfTrades.Size()
	n += 1 + l + sovProtorev(uint64(l))
	if m.PoolId != 0 {
		n += 1 + sovProtorev(uint64(m.PoolId))
	}
	return n
}

func sovProtorev(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProtorev(x uint64) (n int) {
	return sovProtorev(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenPairArbRoutes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtorev
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
			return fmt.Errorf("proto: TokenPairArbRoutes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenPairArbRoutes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArbRoutes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ArbRoutes = append(m.ArbRoutes, &Route{})
			if err := m.ArbRoutes[len(m.ArbRoutes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenIn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenOut = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtorev(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtorev
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
func (m *Route) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtorev
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
			return fmt.Errorf("proto: Route: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Route: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trades", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Trades = append(m.Trades, &Trade{})
			if err := m.Trades[len(m.Trades)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtorev(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtorev
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
func (m *Trade) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtorev
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
			return fmt.Errorf("proto: Trade: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Trade: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
			}
			m.Pool = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pool |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenIn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenOut = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtorev(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtorev
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
func (m *PoolStatistics) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtorev
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
			return fmt.Errorf("proto: PoolStatistics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolStatistics: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Profits = append(m.Profits, &types.Coin{})
			if err := m.Profits[len(m.Profits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumberOfTrades", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
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
				return ErrInvalidLengthProtorev
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProtorev
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NumberOfTrades.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtorev
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProtorev(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProtorev
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
func skipProtorev(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtorev
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
					return 0, ErrIntOverflowProtorev
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
					return 0, ErrIntOverflowProtorev
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
				return 0, ErrInvalidLengthProtorev
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProtorev
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProtorev
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProtorev        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtorev          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProtorev = fmt.Errorf("proto: unexpected end of group")
)