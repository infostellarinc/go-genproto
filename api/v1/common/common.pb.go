// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stellarstation/api/v1/common/common.proto

package common

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

// A set of bits - this is not necessarily aligned to bytes (8-bits).
type Bits struct {
	// The number of bits in this message. Only the most-significant `length` bits of `payload` will
	// be used.
	Length uint32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	// The actual bits in this message. Only the most-significant `length` bits of `payload` will
	// be used.
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bits) Reset()         { *m = Bits{} }
func (m *Bits) String() string { return proto.CompactTextString(m) }
func (*Bits) ProtoMessage()    {}
func (*Bits) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fe638f514c4aab9, []int{0}
}

func (m *Bits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bits.Unmarshal(m, b)
}
func (m *Bits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bits.Marshal(b, m, deterministic)
}
func (m *Bits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bits.Merge(m, src)
}
func (m *Bits) XXX_Size() int {
	return xxx_messageInfo_Bits.Size(m)
}
func (m *Bits) XXX_DiscardUnknown() {
	xxx_messageInfo_Bits.DiscardUnknown(m)
}

var xxx_messageInfo_Bits proto.InternalMessageInfo

func (m *Bits) GetLength() uint32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Bits) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// An az/el angle.
type Angle struct {
	// The azimuth.
	Azimuth float64 `protobuf:"fixed64,1,opt,name=azimuth,proto3" json:"azimuth,omitempty"`
	// The elevation.
	Elevation            float64  `protobuf:"fixed64,2,opt,name=elevation,proto3" json:"elevation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Angle) Reset()         { *m = Angle{} }
func (m *Angle) String() string { return proto.CompactTextString(m) }
func (*Angle) ProtoMessage()    {}
func (*Angle) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fe638f514c4aab9, []int{1}
}

func (m *Angle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Angle.Unmarshal(m, b)
}
func (m *Angle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Angle.Marshal(b, m, deterministic)
}
func (m *Angle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Angle.Merge(m, src)
}
func (m *Angle) XXX_Size() int {
	return xxx_messageInfo_Angle.Size(m)
}
func (m *Angle) XXX_DiscardUnknown() {
	xxx_messageInfo_Angle.DiscardUnknown(m)
}

var xxx_messageInfo_Angle proto.InternalMessageInfo

func (m *Angle) GetAzimuth() float64 {
	if m != nil {
		return m.Azimuth
	}
	return 0
}

func (m *Angle) GetElevation() float64 {
	if m != nil {
		return m.Elevation
	}
	return 0
}

func init() {
	proto.RegisterType((*Bits)(nil), "stellarstation.api.v1.common.Bits")
	proto.RegisterType((*Angle)(nil), "stellarstation.api.v1.common.Angle")
}

func init() {
	proto.RegisterFile("stellarstation/api/v1/common/common.proto", fileDescriptor_4fe638f514c4aab9)
}

var fileDescriptor_4fe638f514c4aab9 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0x2e, 0x49, 0xcd,
	0xc9, 0x49, 0x2c, 0x2a, 0x2e, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f,
	0x33, 0xd4, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x52, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0x32, 0xa8, 0x4a, 0xf5, 0x12, 0x0b, 0x32, 0xf5, 0xca, 0x0c, 0xf5, 0x20, 0x6a, 0x94, 0x2c,
	0xb8, 0x58, 0x9c, 0x32, 0x4b, 0x8a, 0x85, 0xc4, 0xb8, 0xd8, 0x72, 0x52, 0xf3, 0xd2, 0x4b, 0x32,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x83, 0xa0, 0x3c, 0x21, 0x09, 0x2e, 0xf6, 0x82, 0xc4, 0xca,
	0x9c, 0xfc, 0xc4, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x18, 0x57, 0xc9, 0x9e, 0x8b,
	0xd5, 0x31, 0x2f, 0x3d, 0x27, 0x15, 0xa4, 0x24, 0xb1, 0x2a, 0x33, 0xb7, 0x14, 0xaa, 0x97, 0x31,
	0x08, 0xc6, 0x15, 0x92, 0xe1, 0xe2, 0x4c, 0xcd, 0x49, 0x2d, 0x03, 0xdb, 0x0b, 0xd6, 0xce, 0x18,
	0x84, 0x10, 0x70, 0xca, 0xe1, 0x52, 0x48, 0xce, 0xcf, 0xd5, 0xc3, 0xe7, 0x3c, 0x27, 0x6e, 0x67,
	0x30, 0x1d, 0x00, 0xf2, 0x49, 0x00, 0x63, 0x94, 0x65, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x12, 0x48,
	0x56, 0x3f, 0x33, 0x2f, 0x2d, 0x1f, 0xaa, 0x37, 0x33, 0x2f, 0x59, 0x3f, 0x3d, 0x5f, 0x17, 0x5f,
	0x98, 0x24, 0xb1, 0x81, 0x43, 0xc3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x58, 0x96, 0x9b,
	0x3a, 0x01, 0x00, 0x00,
}
