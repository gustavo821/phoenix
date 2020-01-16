// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keys.proto

package phoenix

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

type SecretKey struct {
	A                    *Scalar  `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    *Scalar  `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecretKey) Reset()         { *m = SecretKey{} }
func (m *SecretKey) String() string { return proto.CompactTextString(m) }
func (*SecretKey) ProtoMessage()    {}
func (*SecretKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9084e97af2346a26, []int{0}
}

func (m *SecretKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecretKey.Unmarshal(m, b)
}
func (m *SecretKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecretKey.Marshal(b, m, deterministic)
}
func (m *SecretKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretKey.Merge(m, src)
}
func (m *SecretKey) XXX_Size() int {
	return xxx_messageInfo_SecretKey.Size(m)
}
func (m *SecretKey) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretKey.DiscardUnknown(m)
}

var xxx_messageInfo_SecretKey proto.InternalMessageInfo

func (m *SecretKey) GetA() *Scalar {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *SecretKey) GetB() *Scalar {
	if m != nil {
		return m.B
	}
	return nil
}

type ViewKey struct {
	A                    *Scalar          `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	BG                   *CompressedPoint `protobuf:"bytes,2,opt,name=b_g,json=bG,proto3" json:"b_g,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ViewKey) Reset()         { *m = ViewKey{} }
func (m *ViewKey) String() string { return proto.CompactTextString(m) }
func (*ViewKey) ProtoMessage()    {}
func (*ViewKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9084e97af2346a26, []int{1}
}

func (m *ViewKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewKey.Unmarshal(m, b)
}
func (m *ViewKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewKey.Marshal(b, m, deterministic)
}
func (m *ViewKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewKey.Merge(m, src)
}
func (m *ViewKey) XXX_Size() int {
	return xxx_messageInfo_ViewKey.Size(m)
}
func (m *ViewKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewKey.DiscardUnknown(m)
}

var xxx_messageInfo_ViewKey proto.InternalMessageInfo

func (m *ViewKey) GetA() *Scalar {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *ViewKey) GetBG() *CompressedPoint {
	if m != nil {
		return m.BG
	}
	return nil
}

type PublicKey struct {
	AG                   *CompressedPoint `protobuf:"bytes,1,opt,name=a_g,json=aG,proto3" json:"a_g,omitempty"`
	BG                   *CompressedPoint `protobuf:"bytes,2,opt,name=b_g,json=bG,proto3" json:"b_g,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PublicKey) Reset()         { *m = PublicKey{} }
func (m *PublicKey) String() string { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()    {}
func (*PublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9084e97af2346a26, []int{2}
}

func (m *PublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublicKey.Unmarshal(m, b)
}
func (m *PublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublicKey.Marshal(b, m, deterministic)
}
func (m *PublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublicKey.Merge(m, src)
}
func (m *PublicKey) XXX_Size() int {
	return xxx_messageInfo_PublicKey.Size(m)
}
func (m *PublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_PublicKey proto.InternalMessageInfo

func (m *PublicKey) GetAG() *CompressedPoint {
	if m != nil {
		return m.AG
	}
	return nil
}

func (m *PublicKey) GetBG() *CompressedPoint {
	if m != nil {
		return m.BG
	}
	return nil
}

func init() {
	proto.RegisterType((*SecretKey)(nil), "phoenix.SecretKey")
	proto.RegisterType((*ViewKey)(nil), "phoenix.ViewKey")
	proto.RegisterType((*PublicKey)(nil), "phoenix.PublicKey")
}

func init() { proto.RegisterFile("keys.proto", fileDescriptor_9084e97af2346a26) }

var fileDescriptor_9084e97af2346a26 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4e, 0xad, 0x2c,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0xc8, 0xc8, 0x4f, 0xcd, 0xcb, 0xac, 0x90,
	0xe2, 0x4e, 0xcb, 0x4c, 0xcd, 0x49, 0x81, 0x88, 0x2a, 0x79, 0x72, 0x71, 0x06, 0xa7, 0x26, 0x17,
	0xa5, 0x96, 0x78, 0xa7, 0x56, 0x0a, 0xc9, 0x72, 0x31, 0x26, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x1b, 0xf1, 0xeb, 0x41, 0x95, 0xeb, 0x05, 0x27, 0x27, 0xe6, 0x24, 0x16, 0x05, 0x31, 0x26, 0x82,
	0xa4, 0x93, 0x24, 0x98, 0x70, 0x48, 0x27, 0x29, 0x05, 0x73, 0xb1, 0x87, 0x65, 0xa6, 0x96, 0x13,
	0x61, 0x90, 0x26, 0x17, 0x73, 0x52, 0x7c, 0x3a, 0xd4, 0x28, 0x09, 0xb8, 0x02, 0xe7, 0xfc, 0xdc,
	0x82, 0xa2, 0xd4, 0xe2, 0xe2, 0xd4, 0x94, 0x80, 0xfc, 0xcc, 0xbc, 0x92, 0x20, 0xa6, 0x24, 0x77,
	0xa5, 0x44, 0x2e, 0xce, 0x80, 0xd2, 0xa4, 0x9c, 0xcc, 0x64, 0x90, 0xb1, 0x9a, 0x5c, 0xcc, 0x89,
	0xf1, 0xe9, 0x50, 0x83, 0xf1, 0xe8, 0x4b, 0x74, 0x27, 0xc1, 0x8a, 0x24, 0x36, 0x70, 0x48, 0x18,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xfb, 0x20, 0x91, 0x2d, 0x01, 0x00, 0x00,
}
