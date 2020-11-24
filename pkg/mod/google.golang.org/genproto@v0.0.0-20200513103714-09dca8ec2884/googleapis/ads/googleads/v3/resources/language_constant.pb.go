// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/language_constant.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A language.
type LanguageConstant struct {
	// Output only. The resource name of the language constant.
	// Language constant resource names have the form:
	//
	// `languageConstants/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the language constant.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The language code, e.g. "en_US", "en_AU", "es", "fr", etc.
	Code *wrappers.StringValue `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	// Output only. The full name of the language in English, e.g., "English (US)", "Spanish",
	// etc.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Whether the language is targetable.
	Targetable           *wrappers.BoolValue `protobuf:"bytes,5,opt,name=targetable,proto3" json:"targetable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LanguageConstant) Reset()         { *m = LanguageConstant{} }
func (m *LanguageConstant) String() string { return proto.CompactTextString(m) }
func (*LanguageConstant) ProtoMessage()    {}
func (*LanguageConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d50094a07423762, []int{0}
}

func (m *LanguageConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LanguageConstant.Unmarshal(m, b)
}
func (m *LanguageConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LanguageConstant.Marshal(b, m, deterministic)
}
func (m *LanguageConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LanguageConstant.Merge(m, src)
}
func (m *LanguageConstant) XXX_Size() int {
	return xxx_messageInfo_LanguageConstant.Size(m)
}
func (m *LanguageConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_LanguageConstant.DiscardUnknown(m)
}

var xxx_messageInfo_LanguageConstant proto.InternalMessageInfo

func (m *LanguageConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *LanguageConstant) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *LanguageConstant) GetCode() *wrappers.StringValue {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *LanguageConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *LanguageConstant) GetTargetable() *wrappers.BoolValue {
	if m != nil {
		return m.Targetable
	}
	return nil
}

func init() {
	proto.RegisterType((*LanguageConstant)(nil), "google.ads.googleads.v3.resources.LanguageConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/language_constant.proto", fileDescriptor_1d50094a07423762)
}

var fileDescriptor_1d50094a07423762 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x8a, 0xd4, 0x30,
	0x1c, 0xc6, 0x99, 0x76, 0x15, 0x8c, 0x0a, 0x52, 0x10, 0xc6, 0x71, 0xd1, 0x5d, 0x61, 0x60, 0x45,
	0x48, 0xd4, 0xaa, 0x60, 0x3c, 0x48, 0xea, 0x61, 0x51, 0x44, 0x96, 0x11, 0x7b, 0x90, 0x81, 0x21,
	0xd3, 0x64, 0x63, 0xa1, 0x93, 0x7f, 0x49, 0xd2, 0xf1, 0x20, 0x5e, 0x7c, 0x94, 0x3d, 0xfa, 0x28,
	0x3e, 0xc5, 0x9e, 0xf7, 0x11, 0x3c, 0x49, 0xdb, 0xb4, 0x53, 0x66, 0x40, 0xdd, 0xdb, 0xd7, 0xfe,
	0xbf, 0xdf, 0x97, 0x2f, 0xe1, 0x8f, 0x5e, 0x2a, 0x00, 0x55, 0x48, 0xc2, 0x85, 0x25, 0xad, 0xac,
	0xd5, 0x3a, 0x26, 0x46, 0x5a, 0xa8, 0x4c, 0x26, 0x2d, 0x29, 0xb8, 0x56, 0x15, 0x57, 0x72, 0x91,
	0x81, 0xb6, 0x8e, 0x6b, 0x87, 0x4b, 0x03, 0x0e, 0xa2, 0xc3, 0xd6, 0x8f, 0xb9, 0xb0, 0xb8, 0x47,
	0xf1, 0x3a, 0xc6, 0x3d, 0x3a, 0xb9, 0xdf, 0xa5, 0x97, 0x39, 0x39, 0xcd, 0x65, 0x21, 0x16, 0x4b,
	0xf9, 0x85, 0xaf, 0x73, 0x30, 0x6d, 0xc6, 0xe4, 0xce, 0xc0, 0xd0, 0x61, 0x7e, 0x74, 0xcf, 0x8f,
	0x9a, 0xaf, 0x65, 0x75, 0x4a, 0xbe, 0x1a, 0x5e, 0x96, 0xd2, 0x58, 0x3f, 0xdf, 0x1f, 0xa0, 0x5c,
	0x6b, 0x70, 0xdc, 0xe5, 0xa0, 0xfd, 0xf4, 0xc1, 0x59, 0x88, 0x6e, 0xbd, 0xf7, 0xc5, 0xdf, 0xf8,
	0xde, 0x51, 0x8a, 0x6e, 0x76, 0x87, 0x2c, 0x34, 0x5f, 0xc9, 0xf1, 0xe8, 0x60, 0x74, 0x74, 0x2d,
	0x79, 0x72, 0xce, 0xc2, 0xdf, 0xec, 0x11, 0x7a, 0xb8, 0xb9, 0x85, 0x57, 0x65, 0x6e, 0x71, 0x06,
	0x2b, 0xb2, 0x9d, 0x34, 0xbb, 0xd1, 0xe5, 0x7c, 0xe0, 0x2b, 0x19, 0x3d, 0x46, 0x41, 0x2e, 0xc6,
	0xc1, 0xc1, 0xe8, 0xe8, 0xfa, 0xd3, 0xbb, 0x9e, 0xc5, 0x5d, 0x6f, 0xfc, 0x56, 0xbb, 0x17, 0xcf,
	0x52, 0x5e, 0x54, 0x32, 0x09, 0xcf, 0x59, 0x38, 0x0b, 0x72, 0x11, 0x3d, 0x47, 0x7b, 0x19, 0x08,
	0x39, 0x0e, 0x1b, 0x66, 0x7f, 0x87, 0xf9, 0xe8, 0x4c, 0xae, 0xd5, 0x00, 0x6a, 0xec, 0x35, 0xd6,
	0xf4, 0xde, 0xfb, 0x6f, 0xac, 0xb6, 0x47, 0xaf, 0x11, 0x72, 0xdc, 0x28, 0xe9, 0xf8, 0xb2, 0x90,
	0xe3, 0x2b, 0x0d, 0x3c, 0xd9, 0x81, 0x13, 0x80, 0x62, 0x80, 0x0e, 0x10, 0xfa, 0xe9, 0x82, 0xcd,
	0x2e, 0xf1, 0x3c, 0xd1, 0xb4, 0xd8, 0xfa, 0x63, 0xc9, 0xb7, 0x9d, 0x35, 0xfa, 0x9e, 0xfc, 0x08,
	0xd0, 0x34, 0x83, 0x15, 0xfe, 0xe7, 0x22, 0x25, 0xb7, 0xb7, 0x8f, 0x38, 0xa9, 0x5b, 0x9f, 0x8c,
	0x3e, 0xbf, 0xf3, 0xac, 0x82, 0x3a, 0x1f, 0x83, 0x51, 0x44, 0x49, 0xdd, 0xdc, 0x89, 0x6c, 0x5a,
	0xfe, 0x65, 0xbd, 0x5f, 0xf5, 0xea, 0x2c, 0x08, 0x8f, 0x19, 0xfb, 0x19, 0x1c, 0x1e, 0xb7, 0x91,
	0x4c, 0x58, 0xdc, 0xca, 0x5a, 0xa5, 0x31, 0x9e, 0x75, 0xce, 0x5f, 0x9d, 0x67, 0xce, 0x84, 0x9d,
	0xf7, 0x9e, 0x79, 0x1a, 0xcf, 0x7b, 0xcf, 0x45, 0x30, 0x6d, 0x07, 0x94, 0x32, 0x61, 0x29, 0xed,
	0x5d, 0x94, 0xa6, 0x31, 0xa5, 0xbd, 0x6f, 0x79, 0xb5, 0x29, 0x1b, 0xff, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x95, 0xa7, 0xea, 0x01, 0x8a, 0x03, 0x00, 0x00,
}
