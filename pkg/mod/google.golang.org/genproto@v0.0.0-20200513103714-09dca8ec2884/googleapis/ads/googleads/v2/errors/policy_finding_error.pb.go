// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/policy_finding_error.proto

package errors

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

// Enum describing possible policy finding errors.
type PolicyFindingErrorEnum_PolicyFindingError int32

const (
	// Enum unspecified.
	PolicyFindingErrorEnum_UNSPECIFIED PolicyFindingErrorEnum_PolicyFindingError = 0
	// The received error code is not known in this version.
	PolicyFindingErrorEnum_UNKNOWN PolicyFindingErrorEnum_PolicyFindingError = 1
	// The resource has been disapproved since the policy summary includes
	// policy topics of type PROHIBITED.
	PolicyFindingErrorEnum_POLICY_FINDING PolicyFindingErrorEnum_PolicyFindingError = 2
	// The given policy topic does not exist.
	PolicyFindingErrorEnum_POLICY_TOPIC_NOT_FOUND PolicyFindingErrorEnum_PolicyFindingError = 3
)

var PolicyFindingErrorEnum_PolicyFindingError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "POLICY_FINDING",
	3: "POLICY_TOPIC_NOT_FOUND",
}

var PolicyFindingErrorEnum_PolicyFindingError_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"POLICY_FINDING":         2,
	"POLICY_TOPIC_NOT_FOUND": 3,
}

func (x PolicyFindingErrorEnum_PolicyFindingError) String() string {
	return proto.EnumName(PolicyFindingErrorEnum_PolicyFindingError_name, int32(x))
}

func (PolicyFindingErrorEnum_PolicyFindingError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4c9c6c15ab5b99b5, []int{0, 0}
}

// Container for enum describing possible policy finding errors.
type PolicyFindingErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyFindingErrorEnum) Reset()         { *m = PolicyFindingErrorEnum{} }
func (m *PolicyFindingErrorEnum) String() string { return proto.CompactTextString(m) }
func (*PolicyFindingErrorEnum) ProtoMessage()    {}
func (*PolicyFindingErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9c6c15ab5b99b5, []int{0}
}

func (m *PolicyFindingErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyFindingErrorEnum.Unmarshal(m, b)
}
func (m *PolicyFindingErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyFindingErrorEnum.Marshal(b, m, deterministic)
}
func (m *PolicyFindingErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyFindingErrorEnum.Merge(m, src)
}
func (m *PolicyFindingErrorEnum) XXX_Size() int {
	return xxx_messageInfo_PolicyFindingErrorEnum.Size(m)
}
func (m *PolicyFindingErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyFindingErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyFindingErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.PolicyFindingErrorEnum_PolicyFindingError", PolicyFindingErrorEnum_PolicyFindingError_name, PolicyFindingErrorEnum_PolicyFindingError_value)
	proto.RegisterType((*PolicyFindingErrorEnum)(nil), "google.ads.googleads.v2.errors.PolicyFindingErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/policy_finding_error.proto", fileDescriptor_4c9c6c15ab5b99b5)
}

var fileDescriptor_4c9c6c15ab5b99b5 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x5d, 0x07, 0x0a, 0x19, 0x68, 0xc9, 0x61, 0xc2, 0x90, 0x1d, 0xfa, 0x00, 0x29, 0xd4,
	0x93, 0xf1, 0xd4, 0xad, 0xed, 0x08, 0x4a, 0x1a, 0x70, 0x9b, 0x28, 0x85, 0xd2, 0x2d, 0x33, 0x04,
	0xb6, 0xa4, 0x34, 0x73, 0x20, 0xf8, 0x34, 0x1e, 0x7d, 0x14, 0x1f, 0xc5, 0xa3, 0x4f, 0x20, 0x6d,
	0xb6, 0x5e, 0x86, 0x9e, 0xfa, 0xf1, 0xef, 0xef, 0xfb, 0xf2, 0xfd, 0xff, 0xe0, 0x46, 0x68, 0x2d,
	0xd6, 0x2b, 0xbf, 0xe0, 0xc6, 0xb7, 0xb2, 0x56, 0xbb, 0xc0, 0x5f, 0x55, 0x95, 0xae, 0x8c, 0x5f,
	0xea, 0xb5, 0x5c, 0xbe, 0xe5, 0x2f, 0x52, 0x71, 0xa9, 0x44, 0xde, 0x4c, 0x51, 0x59, 0xe9, 0xad,
	0x86, 0x43, 0xcb, 0xa3, 0x82, 0x1b, 0xd4, 0x5a, 0xd1, 0x2e, 0x40, 0xd6, 0x3a, 0xb8, 0x3a, 0x44,
	0x97, 0xd2, 0x2f, 0x94, 0xd2, 0xdb, 0x62, 0x2b, 0xb5, 0x32, 0xd6, 0xed, 0xbd, 0x83, 0x3e, 0x6b,
	0xb2, 0x13, 0x1b, 0x1d, 0xd7, 0xa6, 0x58, 0xbd, 0x6e, 0xbc, 0x05, 0x80, 0xc7, 0x7f, 0xe0, 0x05,
	0xe8, 0xcd, 0xe8, 0x03, 0x8b, 0xc7, 0x24, 0x21, 0x71, 0xe4, 0x9e, 0xc0, 0x1e, 0x38, 0x9b, 0xd1,
	0x3b, 0x9a, 0x3e, 0x52, 0xb7, 0x03, 0x21, 0x38, 0x67, 0xe9, 0x3d, 0x19, 0x3f, 0xe5, 0x09, 0xa1,
	0x11, 0xa1, 0x13, 0xd7, 0x81, 0x03, 0xd0, 0xdf, 0xcf, 0xa6, 0x29, 0x23, 0xe3, 0x9c, 0xa6, 0xd3,
	0x3c, 0x49, 0x67, 0x34, 0x72, 0xbb, 0xa3, 0x9f, 0x0e, 0xf0, 0x96, 0x7a, 0x83, 0xfe, 0x5f, 0x61,
	0x74, 0x79, 0x5c, 0x84, 0xd5, 0xed, 0x59, 0xe7, 0x39, 0xda, 0x5b, 0x85, 0x5e, 0x17, 0x4a, 0x20,
	0x5d, 0x09, 0x5f, 0xac, 0x54, 0xb3, 0xdb, 0xe1, 0x90, 0xa5, 0x34, 0x7f, 0xdd, 0xf5, 0xd6, 0x7e,
	0x3e, 0x9c, 0xee, 0x24, 0x0c, 0x3f, 0x9d, 0xe1, 0xc4, 0x86, 0x85, 0xdc, 0x20, 0x2b, 0x6b, 0x35,
	0x0f, 0x50, 0xf3, 0xa4, 0xf9, 0x3a, 0x00, 0x59, 0xc8, 0x4d, 0xd6, 0x02, 0xd9, 0x3c, 0xc8, 0x2c,
	0xf0, 0xed, 0x78, 0x76, 0x8a, 0x71, 0xc8, 0x0d, 0xc6, 0x2d, 0x82, 0xf1, 0x3c, 0xc0, 0xd8, 0x42,
	0x8b, 0xd3, 0xa6, 0xdd, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x93, 0x9b, 0x52, 0xf4,
	0x01, 0x00, 0x00,
}
