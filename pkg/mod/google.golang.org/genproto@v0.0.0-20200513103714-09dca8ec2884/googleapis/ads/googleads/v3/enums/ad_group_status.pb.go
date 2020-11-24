// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/ad_group_status.proto

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

// The possible statuses of an ad group.
type AdGroupStatusEnum_AdGroupStatus int32

const (
	// The status has not been specified.
	AdGroupStatusEnum_UNSPECIFIED AdGroupStatusEnum_AdGroupStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupStatusEnum_UNKNOWN AdGroupStatusEnum_AdGroupStatus = 1
	// The ad group is enabled.
	AdGroupStatusEnum_ENABLED AdGroupStatusEnum_AdGroupStatus = 2
	// The ad group is paused.
	AdGroupStatusEnum_PAUSED AdGroupStatusEnum_AdGroupStatus = 3
	// The ad group is removed.
	AdGroupStatusEnum_REMOVED AdGroupStatusEnum_AdGroupStatus = 4
)

var AdGroupStatusEnum_AdGroupStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "PAUSED",
	4: "REMOVED",
}

var AdGroupStatusEnum_AdGroupStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"PAUSED":      3,
	"REMOVED":     4,
}

func (x AdGroupStatusEnum_AdGroupStatus) String() string {
	return proto.EnumName(AdGroupStatusEnum_AdGroupStatus_name, int32(x))
}

func (AdGroupStatusEnum_AdGroupStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ec520dd54d7ad64c, []int{0, 0}
}

// Container for enum describing possible statuses of an ad group.
type AdGroupStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupStatusEnum) Reset()         { *m = AdGroupStatusEnum{} }
func (m *AdGroupStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupStatusEnum) ProtoMessage()    {}
func (*AdGroupStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec520dd54d7ad64c, []int{0}
}

func (m *AdGroupStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupStatusEnum.Unmarshal(m, b)
}
func (m *AdGroupStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupStatusEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupStatusEnum.Merge(m, src)
}
func (m *AdGroupStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupStatusEnum.Size(m)
}
func (m *AdGroupStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.AdGroupStatusEnum_AdGroupStatus", AdGroupStatusEnum_AdGroupStatus_name, AdGroupStatusEnum_AdGroupStatus_value)
	proto.RegisterType((*AdGroupStatusEnum)(nil), "google.ads.googleads.v3.enums.AdGroupStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/ad_group_status.proto", fileDescriptor_ec520dd54d7ad64c)
}

var fileDescriptor_ec520dd54d7ad64c = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x18, 0x77, 0x9d, 0x4c, 0xc8, 0x10, 0x6b, 0x8f, 0xe2, 0x0e, 0xdb, 0x03, 0x24, 0x87, 0xdc, 0xe2,
	0x29, 0xb5, 0x71, 0x0c, 0xb5, 0x2b, 0x96, 0x55, 0x90, 0xc2, 0x88, 0x66, 0xc4, 0xc1, 0x9a, 0x94,
	0xa6, 0xdd, 0x03, 0x79, 0xf4, 0x51, 0x7c, 0x11, 0xc1, 0xa7, 0x90, 0xa4, 0x6b, 0x61, 0x07, 0xbd,
	0x94, 0x5f, 0xbf, 0xdf, 0x9f, 0xfc, 0xbe, 0x0f, 0x60, 0xa9, 0xb5, 0xdc, 0x6d, 0x10, 0x17, 0x06,
	0xb5, 0xd0, 0xa2, 0x3d, 0x46, 0x1b, 0xd5, 0x14, 0x06, 0x71, 0xb1, 0x96, 0x95, 0x6e, 0xca, 0xb5,
	0xa9, 0x79, 0xdd, 0x18, 0x58, 0x56, 0xba, 0xd6, 0xc1, 0xa4, 0x55, 0x42, 0x2e, 0x0c, 0xec, 0x4d,
	0x70, 0x8f, 0xa1, 0x33, 0x5d, 0x5d, 0x77, 0x99, 0xe5, 0x16, 0x71, 0xa5, 0x74, 0xcd, 0xeb, 0xad,
	0x56, 0x07, 0xf3, 0xec, 0x1d, 0x5c, 0x52, 0x31, 0xb7, 0xa1, 0xa9, 0xcb, 0x64, 0xaa, 0x29, 0x66,
	0x29, 0x38, 0x3f, 0x1a, 0x06, 0x17, 0x60, 0xbc, 0x8a, 0xd3, 0x84, 0xdd, 0x2e, 0xee, 0x16, 0x2c,
	0xf2, 0x4f, 0x82, 0x31, 0x38, 0x5b, 0xc5, 0xf7, 0xf1, 0xf2, 0x39, 0xf6, 0x07, 0xf6, 0x87, 0xc5,
	0x34, 0x7c, 0x60, 0x91, 0xef, 0x05, 0x00, 0x8c, 0x12, 0xba, 0x4a, 0x59, 0xe4, 0x0f, 0x2d, 0xf1,
	0xc4, 0x1e, 0x97, 0x19, 0x8b, 0xfc, 0xd3, 0xf0, 0x7b, 0x00, 0xa6, 0x6f, 0xba, 0x80, 0xff, 0xb6,
	0x0d, 0x83, 0xa3, 0x87, 0x13, 0xdb, 0x31, 0x19, 0xbc, 0x84, 0x07, 0x93, 0xd4, 0x3b, 0xae, 0x24,
	0xd4, 0x95, 0x44, 0x72, 0xa3, 0xdc, 0x06, 0xdd, 0x9d, 0xca, 0xad, 0xf9, 0xe3, 0x6c, 0x37, 0xee,
	0xfb, 0xe1, 0x0d, 0xe7, 0x94, 0x7e, 0x7a, 0x93, 0x79, 0x1b, 0x45, 0x85, 0x81, 0x2d, 0xb4, 0x28,
	0xc3, 0xd0, 0x6e, 0x6e, 0xbe, 0x3a, 0x3e, 0xa7, 0xc2, 0xe4, 0x3d, 0x9f, 0x67, 0x38, 0x77, 0xfc,
	0x8f, 0x37, 0x6d, 0x87, 0x84, 0x50, 0x61, 0x08, 0xe9, 0x15, 0x84, 0x64, 0x98, 0x10, 0xa7, 0x79,
	0x1d, 0xb9, 0x62, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x15, 0x6f, 0x49, 0xce, 0x01, 0x00,
	0x00,
}
