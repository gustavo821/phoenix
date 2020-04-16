// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consensus.proto

package rusk

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Call during block generation
type DistributeRequest struct {
	TotalReward           uint64       `protobuf:"fixed64,1,opt,name=total_reward,json=totalReward,proto3" json:"total_reward,omitempty"`
	ProvisionersAddresses []byte       `protobuf:"bytes,2,opt,name=provisioners_addresses,json=provisionersAddresses,proto3" json:"provisioners_addresses,omitempty"`
	BgPk                  *PublicKey   `protobuf:"bytes,3,opt,name=bg_pk,json=bgPk,proto3" json:"bg_pk,omitempty"`
	Tx                    *Transaction `protobuf:"bytes,4,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}     `json:"-"`
	XXX_unrecognized      []byte       `json:"-"`
	XXX_sizecache         int32        `json:"-"`
}

func (m *DistributeRequest) Reset()         { *m = DistributeRequest{} }
func (m *DistributeRequest) String() string { return proto.CompactTextString(m) }
func (*DistributeRequest) ProtoMessage()    {}
func (*DistributeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{0}
}

func (m *DistributeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributeRequest.Unmarshal(m, b)
}
func (m *DistributeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributeRequest.Marshal(b, m, deterministic)
}
func (m *DistributeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributeRequest.Merge(m, src)
}
func (m *DistributeRequest) XXX_Size() int {
	return xxx_messageInfo_DistributeRequest.Size(m)
}
func (m *DistributeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DistributeRequest proto.InternalMessageInfo

func (m *DistributeRequest) GetTotalReward() uint64 {
	if m != nil {
		return m.TotalReward
	}
	return 0
}

func (m *DistributeRequest) GetProvisionersAddresses() []byte {
	if m != nil {
		return m.ProvisionersAddresses
	}
	return nil
}

func (m *DistributeRequest) GetBgPk() *PublicKey {
	if m != nil {
		return m.BgPk
	}
	return nil
}

func (m *DistributeRequest) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

// Call from consensus
type WithdrawRequest struct {
	BlsKey               []byte       `protobuf:"bytes,1,opt,name=bls_key,json=blsKey,proto3" json:"bls_key,omitempty"`
	Sig                  []byte       `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	Msg                  []byte       `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Tx                   *Transaction `protobuf:"bytes,4,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *WithdrawRequest) Reset()         { *m = WithdrawRequest{} }
func (m *WithdrawRequest) String() string { return proto.CompactTextString(m) }
func (*WithdrawRequest) ProtoMessage()    {}
func (*WithdrawRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{1}
}

func (m *WithdrawRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawRequest.Unmarshal(m, b)
}
func (m *WithdrawRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawRequest.Marshal(b, m, deterministic)
}
func (m *WithdrawRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawRequest.Merge(m, src)
}
func (m *WithdrawRequest) XXX_Size() int {
	return xxx_messageInfo_WithdrawRequest.Size(m)
}
func (m *WithdrawRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawRequest proto.InternalMessageInfo

func (m *WithdrawRequest) GetBlsKey() []byte {
	if m != nil {
		return m.BlsKey
	}
	return nil
}

func (m *WithdrawRequest) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *WithdrawRequest) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *WithdrawRequest) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

// Call from consensus
type StakeRequest struct {
	BlsKey               []byte       `protobuf:"bytes,1,opt,name=bls_key,json=blsKey,proto3" json:"bls_key,omitempty"`
	Value                uint64       `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	Locktime             uint64       `protobuf:"fixed64,4,opt,name=locktime,proto3" json:"locktime,omitempty"`
	Tx                   *Transaction `protobuf:"bytes,5,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StakeRequest) Reset()         { *m = StakeRequest{} }
func (m *StakeRequest) String() string { return proto.CompactTextString(m) }
func (*StakeRequest) ProtoMessage()    {}
func (*StakeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{2}
}

func (m *StakeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakeRequest.Unmarshal(m, b)
}
func (m *StakeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakeRequest.Marshal(b, m, deterministic)
}
func (m *StakeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeRequest.Merge(m, src)
}
func (m *StakeRequest) XXX_Size() int {
	return xxx_messageInfo_StakeRequest.Size(m)
}
func (m *StakeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StakeRequest proto.InternalMessageInfo

func (m *StakeRequest) GetBlsKey() []byte {
	if m != nil {
		return m.BlsKey
	}
	return nil
}

func (m *StakeRequest) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *StakeRequest) GetLocktime() uint64 {
	if m != nil {
		return m.Locktime
	}
	return 0
}

func (m *StakeRequest) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

// Call from block generator
type BidRequest struct {
	M                    []byte       `protobuf:"bytes,1,opt,name=m,proto3" json:"m,omitempty"`
	Value                uint64       `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	Locktime             uint64       `protobuf:"fixed64,3,opt,name=locktime,proto3" json:"locktime,omitempty"`
	Tx                   *Transaction `protobuf:"bytes,4,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BidRequest) Reset()         { *m = BidRequest{} }
func (m *BidRequest) String() string { return proto.CompactTextString(m) }
func (*BidRequest) ProtoMessage()    {}
func (*BidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{3}
}

func (m *BidRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidRequest.Unmarshal(m, b)
}
func (m *BidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidRequest.Marshal(b, m, deterministic)
}
func (m *BidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidRequest.Merge(m, src)
}
func (m *BidRequest) XXX_Size() int {
	return xxx_messageInfo_BidRequest.Size(m)
}
func (m *BidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BidRequest proto.InternalMessageInfo

func (m *BidRequest) GetM() []byte {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *BidRequest) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *BidRequest) GetLocktime() uint64 {
	if m != nil {
		return m.Locktime
	}
	return 0
}

func (m *BidRequest) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

// Call from consensus
type SlashRequest struct {
	BlsKey               []byte       `protobuf:"bytes,1,opt,name=bls_key,json=blsKey,proto3" json:"bls_key,omitempty"`
	Step                 uint32       `protobuf:"varint,2,opt,name=step,proto3" json:"step,omitempty"`
	Height               uint64       `protobuf:"fixed64,3,opt,name=height,proto3" json:"height,omitempty"`
	FirstMsg             []byte       `protobuf:"bytes,4,opt,name=first_msg,json=firstMsg,proto3" json:"first_msg,omitempty"`
	FirstSig             []byte       `protobuf:"bytes,5,opt,name=first_sig,json=firstSig,proto3" json:"first_sig,omitempty"`
	SecondMsg            []byte       `protobuf:"bytes,6,opt,name=second_msg,json=secondMsg,proto3" json:"second_msg,omitempty"`
	SecondSig            []byte       `protobuf:"bytes,7,opt,name=second_sig,json=secondSig,proto3" json:"second_sig,omitempty"`
	Tx                   *Transaction `protobuf:"bytes,8,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SlashRequest) Reset()         { *m = SlashRequest{} }
func (m *SlashRequest) String() string { return proto.CompactTextString(m) }
func (*SlashRequest) ProtoMessage()    {}
func (*SlashRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{4}
}

func (m *SlashRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlashRequest.Unmarshal(m, b)
}
func (m *SlashRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlashRequest.Marshal(b, m, deterministic)
}
func (m *SlashRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlashRequest.Merge(m, src)
}
func (m *SlashRequest) XXX_Size() int {
	return xxx_messageInfo_SlashRequest.Size(m)
}
func (m *SlashRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SlashRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SlashRequest proto.InternalMessageInfo

func (m *SlashRequest) GetBlsKey() []byte {
	if m != nil {
		return m.BlsKey
	}
	return nil
}

func (m *SlashRequest) GetStep() uint32 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *SlashRequest) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *SlashRequest) GetFirstMsg() []byte {
	if m != nil {
		return m.FirstMsg
	}
	return nil
}

func (m *SlashRequest) GetFirstSig() []byte {
	if m != nil {
		return m.FirstSig
	}
	return nil
}

func (m *SlashRequest) GetSecondMsg() []byte {
	if m != nil {
		return m.SecondMsg
	}
	return nil
}

func (m *SlashRequest) GetSecondSig() []byte {
	if m != nil {
		return m.SecondSig
	}
	return nil
}

func (m *SlashRequest) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func init() {
	proto.RegisterType((*DistributeRequest)(nil), "rusk.DistributeRequest")
	proto.RegisterType((*WithdrawRequest)(nil), "rusk.WithdrawRequest")
	proto.RegisterType((*StakeRequest)(nil), "rusk.StakeRequest")
	proto.RegisterType((*BidRequest)(nil), "rusk.BidRequest")
	proto.RegisterType((*SlashRequest)(nil), "rusk.SlashRequest")
}

func init() {
	proto.RegisterFile("consensus.proto", fileDescriptor_56f0f2c53b3de771)
}

var fileDescriptor_56f0f2c53b3de771 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0xc4, 0x71, 0xd3, 0xa9, 0x51, 0xc8, 0x0a, 0x4a, 0x14, 0x84, 0x94, 0x56, 0x1c,
	0x72, 0xca, 0x01, 0xc4, 0x03, 0x80, 0xb8, 0x55, 0x48, 0xd5, 0x06, 0x89, 0xa3, 0xb5, 0xb6, 0x07,
	0x67, 0x65, 0xc7, 0x9b, 0xee, 0x8c, 0xdb, 0x46, 0xe2, 0xb9, 0x78, 0x2c, 0x9e, 0x01, 0xed, 0xda,
	0xc1, 0x39, 0x40, 0xca, 0x6d, 0x66, 0xfe, 0x99, 0xfd, 0xfe, 0x3f, 0x91, 0x61, 0x92, 0x99, 0x9a,
	0xb0, 0xa6, 0x86, 0x56, 0x3b, 0x6b, 0xd8, 0x88, 0xd0, 0x36, 0x54, 0xce, 0xa7, 0x6c, 0x55, 0x4d,
	0x2a, 0x63, 0x6d, 0xea, 0x56, 0x98, 0x43, 0x89, 0xfb, 0x6e, 0xe9, 0xfa, 0x67, 0x00, 0xd3, 0xcf,
	0x9a, 0xd8, 0xea, 0xb4, 0x61, 0x94, 0x78, 0xd7, 0x20, 0xb1, 0xb8, 0x82, 0x98, 0x0d, 0xab, 0x2a,
	0xb1, 0xf8, 0xa0, 0x6c, 0x3e, 0x0b, 0x16, 0xc1, 0x32, 0x92, 0x17, 0x7e, 0x26, 0xfd, 0x48, 0x7c,
	0x80, 0xcb, 0x9d, 0x35, 0xf7, 0x9a, 0xb4, 0xa9, 0xd1, 0x52, 0xa2, 0xf2, 0xdc, 0x22, 0x11, 0xd2,
	0x6c, 0xb0, 0x08, 0x96, 0xb1, 0x7c, 0x79, 0xac, 0x7e, 0x3c, 0x88, 0xe2, 0x2d, 0x8c, 0xd2, 0x22,
	0xd9, 0x95, 0xb3, 0xe1, 0x22, 0x58, 0x5e, 0xbc, 0x9b, 0xac, 0x9c, 0xc9, 0xd5, 0x6d, 0x93, 0x56,
	0x3a, 0xbb, 0xc1, 0xbd, 0x0c, 0xd3, 0xe2, 0xb6, 0x14, 0x57, 0x30, 0xe0, 0xc7, 0x59, 0xe8, 0x57,
	0xa6, 0xed, 0xca, 0xd7, 0x3e, 0x86, 0x1c, 0xf0, 0xe3, 0xf5, 0x1d, 0x4c, 0xbe, 0x69, 0xde, 0xe4,
	0x56, 0x3d, 0x1c, 0x5c, 0xbf, 0x82, 0xb3, 0xb4, 0xa2, 0xa4, 0xc4, 0xbd, 0x37, 0x1c, 0xcb, 0x28,
	0xad, 0xe8, 0x06, 0xf7, 0xe2, 0x39, 0x0c, 0x49, 0x17, 0x9d, 0x31, 0x57, 0xba, 0xc9, 0x96, 0x0a,
	0x6f, 0x22, 0x96, 0xae, 0xfc, 0x1f, 0xe4, 0x0f, 0x88, 0xd7, 0xac, 0x4a, 0x7c, 0x92, 0xf7, 0x02,
	0x46, 0xf7, 0xaa, 0x6a, 0xd0, 0xbf, 0x1f, 0xc9, 0xb6, 0x11, 0x73, 0x18, 0x57, 0x26, 0x2b, 0x59,
	0x6f, 0xd1, 0x73, 0x22, 0xf9, 0xa7, 0xef, 0xe8, 0xa3, 0x53, 0x74, 0x03, 0xf0, 0x49, 0xe7, 0x07,
	0x76, 0x0c, 0xc1, 0xb6, 0xa3, 0x06, 0xdb, 0x1e, 0x38, 0xf8, 0x17, 0x70, 0xf8, 0x57, 0xe0, 0xc9,
	0xb8, 0xbf, 0x02, 0x88, 0xd7, 0x95, 0xa2, 0xcd, 0x93, 0x79, 0x05, 0x84, 0xc4, 0xb8, 0xf3, 0xf4,
	0x67, 0xd2, 0xd7, 0xe2, 0x12, 0xa2, 0x0d, 0xea, 0x62, 0xc3, 0x1d, 0xba, 0xeb, 0xc4, 0x6b, 0x38,
	0xff, 0xae, 0x2d, 0x71, 0xe2, 0x7e, 0xff, 0xd0, 0x3f, 0x33, 0xf6, 0x83, 0x2f, 0x54, 0xf4, 0xa2,
	0xfb, 0xbb, 0x46, 0x47, 0xe2, 0x5a, 0x17, 0xe2, 0x0d, 0x00, 0x61, 0x66, 0xea, 0xdc, 0x9f, 0x46,
	0x5e, 0x3d, 0x6f, 0x27, 0xee, 0xb6, 0x97, 0xdd, 0xf1, 0xd9, 0xb1, 0xec, 0xae, 0xdb, 0xc0, 0xe3,
	0x13, 0x81, 0xd3, 0xc8, 0x7f, 0x12, 0xef, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x81, 0x56,
	0x72, 0x4a, 0x03, 0x00, 0x00,
}
