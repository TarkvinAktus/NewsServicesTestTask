// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package messages

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

type News struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *News) Reset()         { *m = News{} }
func (m *News) String() string { return proto.CompactTextString(m) }
func (*News) ProtoMessage()    {}
func (*News) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *News) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_News.Unmarshal(m, b)
}
func (m *News) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_News.Marshal(b, m, deterministic)
}
func (m *News) XXX_Merge(src proto.Message) {
	xxx_messageInfo_News.Merge(m, src)
}
func (m *News) XXX_Size() int {
	return xxx_messageInfo_News.Size(m)
}
func (m *News) XXX_DiscardUnknown() {
	xxx_messageInfo_News.DiscardUnknown(m)
}

var xxx_messageInfo_News proto.InternalMessageInfo

func (m *News) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *News) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *News) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*News)(nil), "News")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2, 0xe3, 0x62, 0xf1, 0x4b,
	0x2d, 0x2f, 0x16, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62,
	0xca, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0x8b, 0x80, 0xd9,
	0x20, 0xb1, 0x94, 0xc4, 0x92, 0x44, 0x09, 0x66, 0x88, 0x18, 0x88, 0x9d, 0xc4, 0x06, 0x36, 0xc6,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x93, 0x08, 0xc0, 0xfe, 0x58, 0x00, 0x00, 0x00,
}
