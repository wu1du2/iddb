// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/conversion_attribution_event_type.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The event type of conversions that are attributed to.
type ConversionAttributionEventTypeEnum_ConversionAttributionEventType int32

const (
	// Not specified.
	ConversionAttributionEventTypeEnum_UNSPECIFIED ConversionAttributionEventTypeEnum_ConversionAttributionEventType = 0
	// Represents value unknown in this version.
	ConversionAttributionEventTypeEnum_UNKNOWN ConversionAttributionEventTypeEnum_ConversionAttributionEventType = 1
	// The conversion is attributed to an impression.
	ConversionAttributionEventTypeEnum_IMPRESSION ConversionAttributionEventTypeEnum_ConversionAttributionEventType = 2
	// The conversion is attributed to an interaction.
	ConversionAttributionEventTypeEnum_INTERACTION ConversionAttributionEventTypeEnum_ConversionAttributionEventType = 3
)

var ConversionAttributionEventTypeEnum_ConversionAttributionEventType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "IMPRESSION",
	3: "INTERACTION",
}

var ConversionAttributionEventTypeEnum_ConversionAttributionEventType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"IMPRESSION":  2,
	"INTERACTION": 3,
}

func (x ConversionAttributionEventTypeEnum_ConversionAttributionEventType) String() string {
	return proto.EnumName(ConversionAttributionEventTypeEnum_ConversionAttributionEventType_name, int32(x))
}

func (ConversionAttributionEventTypeEnum_ConversionAttributionEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_658026e096ee40da, []int{0, 0}
}

// Container for enum indicating the event type the conversion is attributed to.
type ConversionAttributionEventTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionAttributionEventTypeEnum) Reset()         { *m = ConversionAttributionEventTypeEnum{} }
func (m *ConversionAttributionEventTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionAttributionEventTypeEnum) ProtoMessage()    {}
func (*ConversionAttributionEventTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_658026e096ee40da, []int{0}
}

func (m *ConversionAttributionEventTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionAttributionEventTypeEnum.Unmarshal(m, b)
}
func (m *ConversionAttributionEventTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionAttributionEventTypeEnum.Marshal(b, m, deterministic)
}
func (m *ConversionAttributionEventTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionAttributionEventTypeEnum.Merge(m, src)
}
func (m *ConversionAttributionEventTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionAttributionEventTypeEnum.Size(m)
}
func (m *ConversionAttributionEventTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionAttributionEventTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionAttributionEventTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.ConversionAttributionEventTypeEnum_ConversionAttributionEventType", ConversionAttributionEventTypeEnum_ConversionAttributionEventType_name, ConversionAttributionEventTypeEnum_ConversionAttributionEventType_value)
	proto.RegisterType((*ConversionAttributionEventTypeEnum)(nil), "google.ads.googleads.v3.enums.ConversionAttributionEventTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/conversion_attribution_event_type.proto", fileDescriptor_658026e096ee40da)
}

var fileDescriptor_658026e096ee40da = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4f, 0x4b, 0xc3, 0x30,
	0x1c, 0xb5, 0x1d, 0x28, 0x64, 0xa0, 0xa3, 0x47, 0x71, 0xc2, 0xe6, 0x3d, 0x3d, 0xf4, 0x16, 0x4f,
	0xd9, 0x8c, 0xa3, 0x88, 0x5d, 0xd9, 0x3f, 0x41, 0x0a, 0x23, 0x5b, 0x43, 0x28, 0x6c, 0x49, 0x69,
	0xb2, 0xc2, 0x3e, 0x80, 0x5f, 0xc4, 0xa3, 0x1f, 0xc5, 0x8f, 0xe2, 0x27, 0xf0, 0x28, 0x49, 0x5c,
	0x3d, 0xd9, 0x4b, 0x78, 0xe4, 0xf7, 0x7e, 0xef, 0xfd, 0xde, 0x03, 0x84, 0x4b, 0xc9, 0x77, 0x2c,
	0xa4, 0xb9, 0x0a, 0x1d, 0x34, 0xa8, 0x8e, 0x42, 0x26, 0x0e, 0x7b, 0x15, 0x6e, 0xa5, 0xa8, 0x59,
	0xa5, 0x0a, 0x29, 0xd6, 0x54, 0xeb, 0xaa, 0xd8, 0x1c, 0xb4, 0xc1, 0xac, 0x66, 0x42, 0xaf, 0xf5,
	0xb1, 0x64, 0xb0, 0xac, 0xa4, 0x96, 0x41, 0xdf, 0xed, 0x42, 0x9a, 0x2b, 0xd8, 0xc8, 0xc0, 0x3a,
	0x82, 0x56, 0xe6, 0xfa, 0xe6, 0xe4, 0x52, 0x16, 0x21, 0x15, 0x42, 0x6a, 0x6a, 0x64, 0x94, 0x5b,
	0x1e, 0xbe, 0x79, 0x60, 0x38, 0x6e, 0x8c, 0xf0, 0x9f, 0x0f, 0x31, 0x36, 0x8b, 0x63, 0xc9, 0x88,
	0x38, 0xec, 0x87, 0x6b, 0x70, 0xdb, 0xce, 0x0a, 0xae, 0x40, 0x77, 0x99, 0xcc, 0x53, 0x32, 0x8e,
	0x1f, 0x63, 0xf2, 0xd0, 0x3b, 0x0b, 0xba, 0xe0, 0x62, 0x99, 0x3c, 0x25, 0xd3, 0x97, 0xa4, 0xe7,
	0x05, 0x97, 0x00, 0xc4, 0xcf, 0xe9, 0x8c, 0xcc, 0xe7, 0xf1, 0x34, 0xe9, 0xf9, 0x86, 0x1d, 0x27,
	0x0b, 0x32, 0xc3, 0xe3, 0x85, 0xf9, 0xe8, 0x8c, 0xbe, 0x3d, 0x30, 0xd8, 0xca, 0x3d, 0x6c, 0xcd,
	0x32, 0xba, 0x6b, 0x3f, 0x22, 0x35, 0x91, 0x52, 0xef, 0x75, 0xf4, 0xab, 0xc2, 0xe5, 0x8e, 0x0a,
	0x0e, 0x65, 0xc5, 0x43, 0xce, 0x84, 0x0d, 0x7c, 0x2a, 0xba, 0x2c, 0xd4, 0x3f, 0xbd, 0xdf, 0xdb,
	0xf7, 0xdd, 0xef, 0x4c, 0x30, 0xfe, 0xf0, 0xfb, 0x13, 0x27, 0x85, 0x73, 0x05, 0x1d, 0x34, 0x68,
	0x15, 0x41, 0x53, 0x8b, 0xfa, 0x3c, 0xcd, 0x33, 0x9c, 0xab, 0xac, 0x99, 0x67, 0xab, 0x28, 0xb3,
	0xf3, 0x2f, 0x7f, 0xe0, 0x3e, 0x11, 0xc2, 0xb9, 0x42, 0xa8, 0x61, 0x20, 0xb4, 0x8a, 0x10, 0xb2,
	0x9c, 0xcd, 0xb9, 0x3d, 0x2c, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xa5, 0x4d, 0xf5, 0x0f,
	0x02, 0x00, 0x00,
}
