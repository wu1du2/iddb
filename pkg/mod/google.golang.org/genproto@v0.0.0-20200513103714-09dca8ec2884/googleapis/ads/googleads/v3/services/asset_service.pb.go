// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/asset_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v3/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [AssetService.GetAsset][google.ads.googleads.v3.services.AssetService.GetAsset]
type GetAssetRequest struct {
	// Required. The resource name of the asset to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAssetRequest) Reset()         { *m = GetAssetRequest{} }
func (m *GetAssetRequest) String() string { return proto.CompactTextString(m) }
func (*GetAssetRequest) ProtoMessage()    {}
func (*GetAssetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a43928a39eb915, []int{0}
}

func (m *GetAssetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAssetRequest.Unmarshal(m, b)
}
func (m *GetAssetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAssetRequest.Marshal(b, m, deterministic)
}
func (m *GetAssetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAssetRequest.Merge(m, src)
}
func (m *GetAssetRequest) XXX_Size() int {
	return xxx_messageInfo_GetAssetRequest.Size(m)
}
func (m *GetAssetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAssetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAssetRequest proto.InternalMessageInfo

func (m *GetAssetRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AssetService.MutateAssets][google.ads.googleads.v3.services.AssetService.MutateAssets]
type MutateAssetsRequest struct {
	// Required. The ID of the customer whose assets are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual assets.
	Operations           []*AssetOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MutateAssetsRequest) Reset()         { *m = MutateAssetsRequest{} }
func (m *MutateAssetsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAssetsRequest) ProtoMessage()    {}
func (*MutateAssetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a43928a39eb915, []int{1}
}

func (m *MutateAssetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAssetsRequest.Unmarshal(m, b)
}
func (m *MutateAssetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAssetsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAssetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAssetsRequest.Merge(m, src)
}
func (m *MutateAssetsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAssetsRequest.Size(m)
}
func (m *MutateAssetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAssetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAssetsRequest proto.InternalMessageInfo

func (m *MutateAssetsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAssetsRequest) GetOperations() []*AssetOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// A single operation to create an asset. Supported asset types are
// YoutubeVideoAsset, MediaBundleAsset, ImageAsset, and LeadFormAsset. TextAsset
// should be created with Ad inline.
type AssetOperation struct {
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AssetOperation_Create
	Operation            isAssetOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AssetOperation) Reset()         { *m = AssetOperation{} }
func (m *AssetOperation) String() string { return proto.CompactTextString(m) }
func (*AssetOperation) ProtoMessage()    {}
func (*AssetOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a43928a39eb915, []int{2}
}

func (m *AssetOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetOperation.Unmarshal(m, b)
}
func (m *AssetOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetOperation.Marshal(b, m, deterministic)
}
func (m *AssetOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetOperation.Merge(m, src)
}
func (m *AssetOperation) XXX_Size() int {
	return xxx_messageInfo_AssetOperation.Size(m)
}
func (m *AssetOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AssetOperation proto.InternalMessageInfo

type isAssetOperation_Operation interface {
	isAssetOperation_Operation()
}

type AssetOperation_Create struct {
	Create *resources.Asset `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

func (*AssetOperation_Create) isAssetOperation_Operation() {}

func (m *AssetOperation) GetOperation() isAssetOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AssetOperation) GetCreate() *resources.Asset {
	if x, ok := m.GetOperation().(*AssetOperation_Create); ok {
		return x.Create
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AssetOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AssetOperation_Create)(nil),
	}
}

// Response message for an asset mutate.
type MutateAssetsResponse struct {
	// All results for the mutate.
	Results              []*MutateAssetResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MutateAssetsResponse) Reset()         { *m = MutateAssetsResponse{} }
func (m *MutateAssetsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAssetsResponse) ProtoMessage()    {}
func (*MutateAssetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a43928a39eb915, []int{3}
}

func (m *MutateAssetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAssetsResponse.Unmarshal(m, b)
}
func (m *MutateAssetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAssetsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAssetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAssetsResponse.Merge(m, src)
}
func (m *MutateAssetsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAssetsResponse.Size(m)
}
func (m *MutateAssetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAssetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAssetsResponse proto.InternalMessageInfo

func (m *MutateAssetsResponse) GetResults() []*MutateAssetResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the asset mutate.
type MutateAssetResult struct {
	// The resource name returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAssetResult) Reset()         { *m = MutateAssetResult{} }
func (m *MutateAssetResult) String() string { return proto.CompactTextString(m) }
func (*MutateAssetResult) ProtoMessage()    {}
func (*MutateAssetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a43928a39eb915, []int{4}
}

func (m *MutateAssetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAssetResult.Unmarshal(m, b)
}
func (m *MutateAssetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAssetResult.Marshal(b, m, deterministic)
}
func (m *MutateAssetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAssetResult.Merge(m, src)
}
func (m *MutateAssetResult) XXX_Size() int {
	return xxx_messageInfo_MutateAssetResult.Size(m)
}
func (m *MutateAssetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAssetResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAssetResult proto.InternalMessageInfo

func (m *MutateAssetResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAssetRequest)(nil), "google.ads.googleads.v3.services.GetAssetRequest")
	proto.RegisterType((*MutateAssetsRequest)(nil), "google.ads.googleads.v3.services.MutateAssetsRequest")
	proto.RegisterType((*AssetOperation)(nil), "google.ads.googleads.v3.services.AssetOperation")
	proto.RegisterType((*MutateAssetsResponse)(nil), "google.ads.googleads.v3.services.MutateAssetsResponse")
	proto.RegisterType((*MutateAssetResult)(nil), "google.ads.googleads.v3.services.MutateAssetResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/asset_service.proto", fileDescriptor_09a43928a39eb915)
}

var fileDescriptor_09a43928a39eb915 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x8e, 0x54, 0xe8, 0x26, 0x80, 0x6a, 0x10, 0x84, 0x80, 0x20, 0x32, 0x11, 0x8a, 0x02,
	0xac, 0x4b, 0x0c, 0x08, 0x19, 0x55, 0x62, 0x73, 0x49, 0x11, 0x2a, 0xad, 0x82, 0xc8, 0x01, 0x45,
	0x44, 0x5b, 0x7b, 0x31, 0x96, 0x62, 0x6f, 0xf0, 0x6c, 0x72, 0xa9, 0x7a, 0xe1, 0x11, 0xe0, 0x0d,
	0x38, 0x01, 0x8f, 0xd2, 0x03, 0x17, 0x6e, 0x3d, 0xf5, 0xc0, 0x89, 0x27, 0x40, 0x9c, 0x90, 0xbd,
	0x5e, 0xc7, 0x01, 0xa2, 0xb4, 0xb7, 0x89, 0xe7, 0xfb, 0x99, 0xd9, 0x99, 0x09, 0xba, 0xef, 0x73,
	0xee, 0x8f, 0x98, 0x45, 0x3d, 0xb0, 0x64, 0x98, 0x44, 0x53, 0xdb, 0x02, 0x16, 0x4f, 0x03, 0x97,
	0x81, 0x45, 0x01, 0x98, 0x18, 0x66, 0x3f, 0xf1, 0x38, 0xe6, 0x82, 0x1b, 0x75, 0x09, 0xc5, 0xd4,
	0x03, 0x9c, 0xb3, 0xf0, 0xd4, 0xc6, 0x8a, 0x55, 0xbb, 0xbb, 0x48, 0x37, 0x66, 0xc0, 0x27, 0x71,
	0x2e, 0x2c, 0x05, 0x6b, 0xd7, 0x14, 0x7c, 0x1c, 0x58, 0x34, 0x8a, 0xb8, 0xa0, 0x22, 0xe0, 0x11,
	0x64, 0xd9, 0xcb, 0x85, 0xac, 0x3b, 0x0a, 0x58, 0xa4, 0x68, 0x37, 0x0a, 0x89, 0x37, 0x01, 0x1b,
	0x79, 0xc3, 0x5d, 0xf6, 0x96, 0x4e, 0x03, 0x1e, 0x67, 0x80, 0x2b, 0x05, 0x80, 0x72, 0x96, 0x29,
	0xf3, 0x35, 0x3a, 0xdf, 0x65, 0x82, 0x24, 0x45, 0xf4, 0xd8, 0xbb, 0x09, 0x03, 0x61, 0x3c, 0x43,
	0x67, 0x15, 0x68, 0x18, 0xd1, 0x90, 0x55, 0xb5, 0xba, 0xd6, 0x5c, 0xed, 0xdc, 0x3a, 0x22, 0xfa,
	0x6f, 0x52, 0x47, 0xd7, 0x67, 0xad, 0x66, 0xd1, 0x38, 0x00, 0xec, 0xf2, 0xd0, 0x92, 0x2a, 0x15,
	0x45, 0x7e, 0x4e, 0x43, 0x66, 0x7e, 0xd0, 0xd0, 0x85, 0xad, 0x89, 0xa0, 0x82, 0xa5, 0x59, 0x50,
	0x26, 0x0d, 0x54, 0x76, 0x27, 0x20, 0x78, 0xc8, 0xe2, 0x61, 0xe0, 0x65, 0x16, 0xa5, 0x23, 0xa2,
	0xf7, 0x90, 0xfa, 0xfe, 0xd4, 0x33, 0x5e, 0x22, 0xc4, 0xc7, 0x2c, 0x96, 0xcf, 0x50, 0xd5, 0xeb,
	0xa5, 0x66, 0xb9, 0xbd, 0x8e, 0x97, 0x3d, 0x3b, 0x4e, 0xad, 0xb6, 0x15, 0x31, 0x93, 0x9d, 0x09,
	0x99, 0x14, 0x9d, 0x9b, 0x87, 0x18, 0x1d, 0xb4, 0xe2, 0xc6, 0x8c, 0x0a, 0xd9, 0x6c, 0xb9, 0xdd,
	0x5c, 0x68, 0x92, 0x4f, 0x4e, 0xba, 0x6c, 0x9e, 0xea, 0x65, 0xcc, 0x4e, 0x19, 0xad, 0xe6, 0x1e,
	0x26, 0x43, 0x17, 0xe7, 0xdb, 0x86, 0x31, 0x8f, 0x80, 0x19, 0x5b, 0xe8, 0x74, 0xcc, 0x60, 0x32,
	0x12, 0xaa, 0x1d, 0x7b, 0x79, 0x3b, 0x05, 0xa1, 0x5e, 0xca, 0xed, 0x29, 0x0d, 0xf3, 0x11, 0x5a,
	0xfb, 0x27, 0x6b, 0xdc, 0xfc, 0xef, 0x00, 0xe7, 0x07, 0xd3, 0xfe, 0x5c, 0x42, 0x95, 0x94, 0xf4,
	0x42, 0xda, 0x18, 0x5f, 0x34, 0x74, 0x46, 0xad, 0x82, 0x71, 0x6f, 0x79, 0x55, 0x7f, 0xad, 0x4d,
	0xed, 0xd8, 0x4f, 0x66, 0x3e, 0x39, 0x24, 0xf3, 0x05, 0xbe, 0xff, 0xfe, 0xe3, 0xa3, 0xde, 0x32,
	0x9a, 0xc9, 0x65, 0xec, 0xcd, 0x65, 0x36, 0xd4, 0x32, 0x80, 0xd5, 0x92, 0xa7, 0x02, 0x56, 0x6b,
	0xdf, 0xf8, 0xa6, 0xa1, 0x4a, 0xf1, 0x79, 0x8d, 0x07, 0x27, 0x7a, 0x45, 0xb5, 0x85, 0xb5, 0x87,
	0x27, 0xa5, 0xc9, 0x29, 0x9a, 0xdb, 0x87, 0xe4, 0x52, 0x61, 0x7d, 0xef, 0xcc, 0x76, 0x2b, 0x6d,
	0x65, 0xdd, 0xbc, 0x9d, 0xb4, 0x32, 0xab, 0x7d, 0xaf, 0x00, 0xde, 0x68, 0xed, 0x67, 0x9d, 0x38,
	0x61, 0xaa, 0xed, 0x68, 0xad, 0xda, 0xd5, 0x03, 0x52, 0x5d, 0x74, 0x57, 0x9d, 0x5f, 0x1a, 0x6a,
	0xb8, 0x3c, 0x5c, 0x5a, 0x6b, 0x67, 0xad, 0x38, 0xd0, 0x9d, 0xe4, 0xbe, 0x77, 0xb4, 0x57, 0x9b,
	0x19, 0xcd, 0xe7, 0x23, 0x1a, 0xf9, 0x98, 0xc7, 0xbe, 0xe5, 0xb3, 0x28, 0xbd, 0x7e, 0x6b, 0x66,
	0xb4, 0xf8, 0xaf, 0xef, 0xb1, 0x0a, 0x3e, 0xe9, 0xa5, 0x2e, 0x21, 0x5f, 0xf5, 0x7a, 0x57, 0x0a,
	0x12, 0x0f, 0xb0, 0x0c, 0x93, 0xa8, 0x6f, 0xe3, 0xcc, 0x18, 0x0e, 0x14, 0x64, 0x40, 0x3c, 0x18,
	0xe4, 0x90, 0x41, 0xdf, 0x1e, 0x28, 0xc8, 0x4f, 0xbd, 0x21, 0xbf, 0x3b, 0x0e, 0xf1, 0xc0, 0x71,
	0x72, 0x90, 0xe3, 0xf4, 0x6d, 0xc7, 0x51, 0xb0, 0xdd, 0x95, 0xb4, 0x4e, 0xfb, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x36, 0xb8, 0xa8, 0x1f, 0xa1, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AssetServiceClient is the client API for AssetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AssetServiceClient interface {
	// Returns the requested asset in full detail.
	GetAsset(ctx context.Context, in *GetAssetRequest, opts ...grpc.CallOption) (*resources.Asset, error)
	// Creates assets. Operation statuses are returned.
	MutateAssets(ctx context.Context, in *MutateAssetsRequest, opts ...grpc.CallOption) (*MutateAssetsResponse, error)
}

type assetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetServiceClient(cc grpc.ClientConnInterface) AssetServiceClient {
	return &assetServiceClient{cc}
}

func (c *assetServiceClient) GetAsset(ctx context.Context, in *GetAssetRequest, opts ...grpc.CallOption) (*resources.Asset, error) {
	out := new(resources.Asset)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.AssetService/GetAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetServiceClient) MutateAssets(ctx context.Context, in *MutateAssetsRequest, opts ...grpc.CallOption) (*MutateAssetsResponse, error) {
	out := new(MutateAssetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.AssetService/MutateAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServiceServer is the server API for AssetService service.
type AssetServiceServer interface {
	// Returns the requested asset in full detail.
	GetAsset(context.Context, *GetAssetRequest) (*resources.Asset, error)
	// Creates assets. Operation statuses are returned.
	MutateAssets(context.Context, *MutateAssetsRequest) (*MutateAssetsResponse, error)
}

// UnimplementedAssetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAssetServiceServer struct {
}

func (*UnimplementedAssetServiceServer) GetAsset(ctx context.Context, req *GetAssetRequest) (*resources.Asset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAsset not implemented")
}
func (*UnimplementedAssetServiceServer) MutateAssets(ctx context.Context, req *MutateAssetsRequest) (*MutateAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAssets not implemented")
}

func RegisterAssetServiceServer(s *grpc.Server, srv AssetServiceServer) {
	s.RegisterService(&_AssetService_serviceDesc, srv)
}

func _AssetService_GetAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).GetAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.AssetService/GetAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).GetAsset(ctx, req.(*GetAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetService_MutateAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServiceServer).MutateAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.AssetService/MutateAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServiceServer).MutateAssets(ctx, req.(*MutateAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AssetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.AssetService",
	HandlerType: (*AssetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAsset",
			Handler:    _AssetService_GetAsset_Handler,
		},
		{
			MethodName: "MutateAssets",
			Handler:    _AssetService_MutateAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/asset_service.proto",
}