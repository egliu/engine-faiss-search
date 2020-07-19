// Code generated by protoc-gen-go. DO NOT EDIT.
// source: faisssearch/v1/engine-faiss-search-service.proto

package v1pb

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

// Hello is a hello.
type Hello struct {
	Hello                int64    `protobuf:"varint,1,opt,name=hello,proto3" json:"hello,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_12a81b90d6740d5b, []int{0}
}

func (m *Hello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hello.Unmarshal(m, b)
}
func (m *Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hello.Marshal(b, m, deterministic)
}
func (m *Hello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hello.Merge(m, src)
}
func (m *Hello) XXX_Size() int {
	return xxx_messageInfo_Hello.Size(m)
}
func (m *Hello) XXX_DiscardUnknown() {
	xxx_messageInfo_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_Hello proto.InternalMessageInfo

func (m *Hello) GetHello() int64 {
	if m != nil {
		return m.Hello
	}
	return 0
}

func init() {
	proto.RegisterType((*Hello)(nil), "faisssearch.v1.Hello")
}

func init() {
	proto.RegisterFile("faisssearch/v1/engine-faiss-search-service.proto", fileDescriptor_12a81b90d6740d5b)
}

var fileDescriptor_12a81b90d6740d5b = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0x4b, 0xcc, 0x2c,
	0x2e, 0x2e, 0x4e, 0x4d, 0x2c, 0x4a, 0xce, 0xd0, 0x2f, 0x33, 0xd4, 0x4f, 0xcd, 0x4b, 0xcf, 0xcc,
	0x4b, 0xd5, 0x05, 0x8b, 0xea, 0x42, 0x84, 0x75, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0xf5,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xf8, 0x90, 0x74, 0xe8, 0x95, 0x19, 0x2a, 0xc9, 0x72, 0xb1,
	0x7a, 0xa4, 0xe6, 0xe4, 0xe4, 0x0b, 0x89, 0x70, 0xb1, 0x66, 0x80, 0x18, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xcc, 0x41, 0x10, 0x8e, 0x93, 0x35, 0x97, 0x50, 0x72, 0x7e, 0xae, 0x1e, 0xaa, 0x26, 0x27,
	0x59, 0x57, 0xb0, 0x3d, 0x6e, 0x20, 0xd1, 0x60, 0xb0, 0x68, 0x30, 0xc4, 0x92, 0x00, 0x90, 0x1d,
	0x01, 0x8c, 0x51, 0x2c, 0x65, 0x86, 0x05, 0x49, 0x49, 0x6c, 0x60, 0x2b, 0x8d, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x74, 0x5e, 0x50, 0x81, 0xa6, 0x00, 0x00, 0x00,
}