// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

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

type TransactionInput struct {
	Pos                  *Idx       `protobuf:"bytes,1,opt,name=pos,proto3" json:"pos,omitempty"`
	Sk                   *SecretKey `protobuf:"bytes,2,opt,name=sk,proto3" json:"sk,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TransactionInput) Reset()         { *m = TransactionInput{} }
func (m *TransactionInput) String() string { return proto.CompactTextString(m) }
func (*TransactionInput) ProtoMessage()    {}
func (*TransactionInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}

func (m *TransactionInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionInput.Unmarshal(m, b)
}
func (m *TransactionInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionInput.Marshal(b, m, deterministic)
}
func (m *TransactionInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionInput.Merge(m, src)
}
func (m *TransactionInput) XXX_Size() int {
	return xxx_messageInfo_TransactionInput.Size(m)
}
func (m *TransactionInput) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionInput.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionInput proto.InternalMessageInfo

func (m *TransactionInput) GetPos() *Idx {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *TransactionInput) GetSk() *SecretKey {
	if m != nil {
		return m.Sk
	}
	return nil
}

type TransactionOutput struct {
	NoteType             NoteType   `protobuf:"varint,1,opt,name=note_type,json=noteType,proto3,enum=phoenix.NoteType" json:"note_type,omitempty"`
	Pk                   *PublicKey `protobuf:"bytes,2,opt,name=pk,proto3" json:"pk,omitempty"`
	Value                uint64     `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TransactionOutput) Reset()         { *m = TransactionOutput{} }
func (m *TransactionOutput) String() string { return proto.CompactTextString(m) }
func (*TransactionOutput) ProtoMessage()    {}
func (*TransactionOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{1}
}

func (m *TransactionOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionOutput.Unmarshal(m, b)
}
func (m *TransactionOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionOutput.Marshal(b, m, deterministic)
}
func (m *TransactionOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionOutput.Merge(m, src)
}
func (m *TransactionOutput) XXX_Size() int {
	return xxx_messageInfo_TransactionOutput.Size(m)
}
func (m *TransactionOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionOutput.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionOutput proto.InternalMessageInfo

func (m *TransactionOutput) GetNoteType() NoteType {
	if m != nil {
		return m.NoteType
	}
	return NoteType_TRANSPARENT
}

func (m *TransactionOutput) GetPk() *PublicKey {
	if m != nil {
		return m.Pk
	}
	return nil
}

func (m *TransactionOutput) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Transaction struct {
	Inputs               []*TransactionInput  `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs              []*TransactionOutput `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{2}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetInputs() []*TransactionInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Transaction) GetOutputs() []*TransactionOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionInput)(nil), "phoenix.TransactionInput")
	proto.RegisterType((*TransactionOutput)(nil), "phoenix.TransactionOutput")
	proto.RegisterType((*Transaction)(nil), "phoenix.Transaction")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_2cc4e03d2c28c490) }

var fileDescriptor_2cc4e03d2c28c490 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4f, 0x84, 0x30,
	0x10, 0x85, 0x03, 0xe8, 0xae, 0x0e, 0xc6, 0x48, 0xe3, 0x01, 0x39, 0x98, 0x0d, 0x27, 0x4e, 0x24,
	0xa2, 0x7f, 0x62, 0x63, 0xa2, 0xa6, 0x6e, 0xbc, 0x1a, 0x16, 0x27, 0x91, 0xb0, 0x69, 0x27, 0x74,
	0xd8, 0x2c, 0x07, 0xff, 0xbb, 0x69, 0x41, 0x24, 0xba, 0xb7, 0x69, 0xdf, 0xbc, 0xf7, 0xf5, 0x15,
	0x22, 0x6e, 0x4b, 0x65, 0xca, 0x8a, 0x6b, 0xad, 0x72, 0x6a, 0x35, 0x6b, 0xb1, 0xa4, 0x4f, 0x8d,
	0xaa, 0x3e, 0x24, 0xd0, 0x60, 0x6f, 0x86, 0xcb, 0x04, 0x94, 0x66, 0x1c, 0xe6, 0xf4, 0x0d, 0xae,
	0x36, 0xbf, 0xae, 0xb5, 0xa2, 0x8e, 0xc5, 0x2d, 0x04, 0xa4, 0x4d, 0xec, 0xad, 0xbc, 0x2c, 0x2c,
	0x2e, 0xf2, 0x31, 0x22, 0x5f, 0x7f, 0x1c, 0xa4, 0x15, 0x44, 0x0a, 0xbe, 0x69, 0x62, 0xdf, 0xc9,
	0x62, 0x92, 0x5f, 0xb1, 0x6a, 0x91, 0x1f, 0xb1, 0x97, 0xbe, 0x69, 0xd2, 0x2f, 0x88, 0x66, 0xb9,
	0xcf, 0x1d, 0xdb, 0xe0, 0x1c, 0xce, 0x2d, 0xfa, 0x9d, 0x7b, 0x42, 0x17, 0x7f, 0x59, 0x44, 0x93,
	0xff, 0x49, 0x33, 0x6e, 0x7a, 0x42, 0x79, 0xa6, 0xc6, 0xc9, 0x82, 0xe8, 0x3f, 0xe8, 0xa5, 0xdb,
	0xee, 0xea, 0xca, 0x81, 0xa8, 0x11, 0xd7, 0x70, 0xba, 0x2f, 0x77, 0x1d, 0xc6, 0xc1, 0xca, 0xcb,
	0x4e, 0xe4, 0x70, 0x48, 0xf7, 0x10, 0xce, 0xf0, 0xe2, 0x0e, 0x16, 0xb5, 0xad, 0x66, 0x4b, 0x05,
	0x59, 0x58, 0xdc, 0x4c, 0x61, 0x7f, 0xcb, 0xcb, 0x71, 0x51, 0x3c, 0xc0, 0x52, 0xbb, 0x57, 0x9b,
	0xd8, 0x77, 0x9e, 0xe4, 0x98, 0x67, 0x28, 0x26, 0x7f, 0x56, 0xb7, 0x0b, 0xf7, 0xab, 0xf7, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x31, 0x43, 0x8a, 0x8b, 0x01, 0x00, 0x00,
}
