// Code generated by protoc-gen-go. DO NOT EDIT.
// source: artifacts.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "www.velocidex.com/golang/velociraptor/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Artifacts struct {
	Names                []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifacts) Reset()         { *m = Artifacts{} }
func (m *Artifacts) String() string { return proto.CompactTextString(m) }
func (*Artifacts) ProtoMessage()    {}
func (*Artifacts) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_382238d2a2b379e1, []int{0}
}
func (m *Artifacts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifacts.Unmarshal(m, b)
}
func (m *Artifacts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifacts.Marshal(b, m, deterministic)
}
func (dst *Artifacts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifacts.Merge(dst, src)
}
func (m *Artifacts) XXX_Size() int {
	return xxx_messageInfo_Artifacts.Size(m)
}
func (m *Artifacts) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifacts.DiscardUnknown(m)
}

var xxx_messageInfo_Artifacts proto.InternalMessageInfo

func (m *Artifacts) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type ArtifactCollectorArgs struct {
	Artifacts            *Artifacts `protobuf:"bytes,1,opt,name=artifacts,proto3" json:"artifacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ArtifactCollectorArgs) Reset()         { *m = ArtifactCollectorArgs{} }
func (m *ArtifactCollectorArgs) String() string { return proto.CompactTextString(m) }
func (*ArtifactCollectorArgs) ProtoMessage()    {}
func (*ArtifactCollectorArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_artifacts_382238d2a2b379e1, []int{1}
}
func (m *ArtifactCollectorArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArtifactCollectorArgs.Unmarshal(m, b)
}
func (m *ArtifactCollectorArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArtifactCollectorArgs.Marshal(b, m, deterministic)
}
func (dst *ArtifactCollectorArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArtifactCollectorArgs.Merge(dst, src)
}
func (m *ArtifactCollectorArgs) XXX_Size() int {
	return xxx_messageInfo_ArtifactCollectorArgs.Size(m)
}
func (m *ArtifactCollectorArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_ArtifactCollectorArgs.DiscardUnknown(m)
}

var xxx_messageInfo_ArtifactCollectorArgs proto.InternalMessageInfo

func (m *ArtifactCollectorArgs) GetArtifacts() *Artifacts {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

func init() {
	proto.RegisterType((*Artifacts)(nil), "proto.Artifacts")
	proto.RegisterType((*ArtifactCollectorArgs)(nil), "proto.ArtifactCollectorArgs")
}

func init() { proto.RegisterFile("artifacts.proto", fileDescriptor_artifacts_382238d2a2b379e1) }

var fileDescriptor_artifacts_382238d2a2b379e1 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2c, 0x2a, 0xc9,
	0x4c, 0x4b, 0x4c, 0x2e, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52,
	0x56, 0xe5, 0xe5, 0xe5, 0x7a, 0x65, 0xa9, 0x39, 0xf9, 0xc9, 0x99, 0x29, 0xa9, 0x15, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0xfa, 0x10, 0xc1, 0xa2, 0xc4, 0x82, 0x92,
	0xfc, 0x22, 0x7d, 0xb0, 0x62, 0xfd, 0xe2, 0xd4, 0xdc, 0xc4, 0xbc, 0x92, 0xcc, 0x64, 0x88, 0x11,
	0x4a, 0xce, 0x5c, 0x9c, 0x8e, 0x30, 0x53, 0x85, 0xcc, 0xb8, 0x58, 0xf3, 0x12, 0x73, 0x53, 0x8b,
	0x25, 0x18, 0x15, 0x98, 0x35, 0x38, 0x9d, 0x14, 0x1e, 0xfd, 0x79, 0x7c, 0x84, 0x51, 0x4a, 0x48,
	0x22, 0x24, 0x23, 0x55, 0x01, 0x6e, 0xb7, 0x42, 0x49, 0xbe, 0x42, 0x4e, 0x62, 0x69, 0x5e, 0x72,
	0x86, 0x5e, 0x10, 0x44, 0xb9, 0x52, 0x1d, 0x97, 0x28, 0xcc, 0x10, 0xe7, 0xfc, 0x9c, 0x9c, 0xd4,
	0xe4, 0x92, 0xfc, 0x22, 0xc7, 0xa2, 0xf4, 0x62, 0xa1, 0x54, 0x2e, 0x4e, 0xb8, 0x3e, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0x6e, 0x23, 0x01, 0x88, 0xc5, 0x7a, 0x70, 0x5b, 0x9d, 0x4c, 0xc1, 0xd6, 0xe8,
	0xe3, 0xb6, 0x46, 0x49, 0x18, 0xae, 0x58, 0x21, 0x24, 0x5f, 0xc1, 0x07, 0x2c, 0x1a, 0x84, 0x30,
	0x39, 0x89, 0x0d, 0x6c, 0xa4, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xd4, 0xe6, 0x14, 0x21,
	0x01, 0x00, 0x00,
}