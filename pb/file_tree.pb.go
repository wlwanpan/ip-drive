// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file_tree.proto

package pb

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FSNode_Type int32

const (
	FSNode_FILE FSNode_Type = 0
	FSNode_DIR  FSNode_Type = 1
)

var FSNode_Type_name = map[int32]string{
	0: "FILE",
	1: "DIR",
}

var FSNode_Type_value = map[string]int32{
	"FILE": 0,
	"DIR":  1,
}

func (x FSNode_Type) String() string {
	return proto.EnumName(FSNode_Type_name, int32(x))
}

func (FSNode_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_718d290bcea536a3, []int{0, 0}
}

type FSNode struct {
	Id                   []byte    `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Path                 string    `protobuf:"bytes,2,opt,name=Path,proto3" json:"Path,omitempty"`
	Links                []*FSNode `protobuf:"bytes,4,rep,name=Links,proto3" json:"Links,omitempty"`
	Source               string    `protobuf:"bytes,5,opt,name=Source,proto3" json:"Source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FSNode) Reset()         { *m = FSNode{} }
func (m *FSNode) String() string { return proto.CompactTextString(m) }
func (*FSNode) ProtoMessage()    {}
func (*FSNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_718d290bcea536a3, []int{0}
}

func (m *FSNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FSNode.Unmarshal(m, b)
}
func (m *FSNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FSNode.Marshal(b, m, deterministic)
}
func (m *FSNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FSNode.Merge(m, src)
}
func (m *FSNode) XXX_Size() int {
	return xxx_messageInfo_FSNode.Size(m)
}
func (m *FSNode) XXX_DiscardUnknown() {
	xxx_messageInfo_FSNode.DiscardUnknown(m)
}

var xxx_messageInfo_FSNode proto.InternalMessageInfo

func (m *FSNode) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *FSNode) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *FSNode) GetLinks() []*FSNode {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *FSNode) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type FSTree struct {
	Owner                string   `protobuf:"bytes,1,opt,name=Owner,proto3" json:"Owner,omitempty"`
	Head                 *FSNode  `protobuf:"bytes,2,opt,name=Head,proto3" json:"Head,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FSTree) Reset()         { *m = FSTree{} }
func (m *FSTree) String() string { return proto.CompactTextString(m) }
func (*FSTree) ProtoMessage()    {}
func (*FSTree) Descriptor() ([]byte, []int) {
	return fileDescriptor_718d290bcea536a3, []int{1}
}

func (m *FSTree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FSTree.Unmarshal(m, b)
}
func (m *FSTree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FSTree.Marshal(b, m, deterministic)
}
func (m *FSTree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FSTree.Merge(m, src)
}
func (m *FSTree) XXX_Size() int {
	return xxx_messageInfo_FSTree.Size(m)
}
func (m *FSTree) XXX_DiscardUnknown() {
	xxx_messageInfo_FSTree.DiscardUnknown(m)
}

var xxx_messageInfo_FSTree proto.InternalMessageInfo

func (m *FSTree) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *FSTree) GetHead() *FSNode {
	if m != nil {
		return m.Head
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.FSNode_Type", FSNode_Type_name, FSNode_Type_value)
	proto.RegisterType((*FSNode)(nil), "pb.FSNode")
	proto.RegisterType((*FSTree)(nil), "pb.FSTree")
}

func init() { proto.RegisterFile("file_tree.proto", fileDescriptor_718d290bcea536a3) }

var fileDescriptor_718d290bcea536a3 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xcb, 0xcc, 0x49,
	0x8d, 0x2f, 0x29, 0x4a, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52,
	0x6a, 0x64, 0xe4, 0x62, 0x73, 0x0b, 0xf6, 0xcb, 0x4f, 0x49, 0x15, 0xe2, 0xe3, 0x62, 0xf2, 0x4c,
	0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x62, 0xf2, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0x09, 0x48,
	0x2c, 0xc9, 0x90, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x14, 0xb8, 0x58, 0x7d,
	0x32, 0xf3, 0xb2, 0x8b, 0x25, 0x58, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0xb8, 0xf4, 0x0a, 0x92, 0xf4,
	0x20, 0xda, 0x83, 0x20, 0x12, 0x42, 0x62, 0x5c, 0x6c, 0xc1, 0xf9, 0xa5, 0x45, 0xc9, 0xa9, 0x12,
	0xac, 0x60, 0x7d, 0x50, 0x9e, 0x92, 0x24, 0x17, 0x4b, 0x48, 0x65, 0x41, 0xaa, 0x10, 0x07, 0x17,
	0x8b, 0x9b, 0xa7, 0x8f, 0xab, 0x00, 0x83, 0x10, 0x3b, 0x17, 0xb3, 0x8b, 0x67, 0x90, 0x00, 0xa3,
	0x92, 0x1d, 0xc8, 0x09, 0x21, 0x45, 0xa9, 0xa9, 0x42, 0x22, 0x5c, 0xac, 0xfe, 0xe5, 0x79, 0xa9,
	0x45, 0x60, 0x57, 0x70, 0x06, 0x41, 0x38, 0x42, 0x72, 0x5c, 0x2c, 0x1e, 0xa9, 0x89, 0x29, 0x60,
	0x87, 0xa0, 0xda, 0x09, 0x16, 0x4f, 0x62, 0x03, 0x7b, 0xc7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x53, 0xf5, 0x63, 0xf0, 0xe1, 0x00, 0x00, 0x00,
}
