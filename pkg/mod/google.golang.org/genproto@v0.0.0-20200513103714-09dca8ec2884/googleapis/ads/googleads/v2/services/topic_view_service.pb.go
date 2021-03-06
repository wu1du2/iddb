// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/topic_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
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

// Request message for [TopicViewService.GetTopicView][google.ads.googleads.v2.services.TopicViewService.GetTopicView].
type GetTopicViewRequest struct {
	// Required. The resource name of the topic view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopicViewRequest) Reset()         { *m = GetTopicViewRequest{} }
func (m *GetTopicViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopicViewRequest) ProtoMessage()    {}
func (*GetTopicViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d00f993f4e0946, []int{0}
}

func (m *GetTopicViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopicViewRequest.Unmarshal(m, b)
}
func (m *GetTopicViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopicViewRequest.Marshal(b, m, deterministic)
}
func (m *GetTopicViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopicViewRequest.Merge(m, src)
}
func (m *GetTopicViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopicViewRequest.Size(m)
}
func (m *GetTopicViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopicViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopicViewRequest proto.InternalMessageInfo

func (m *GetTopicViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetTopicViewRequest)(nil), "google.ads.googleads.v2.services.GetTopicViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/topic_view_service.proto", fileDescriptor_c6d00f993f4e0946)
}

var fileDescriptor_c6d00f993f4e0946 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xbf, 0xea, 0xd3, 0x40,
	0x1c, 0x27, 0x11, 0x04, 0xc3, 0x4f, 0x90, 0x88, 0x58, 0xa3, 0x60, 0x29, 0x1d, 0xa4, 0x94, 0x3b,
	0x88, 0x38, 0x78, 0xe2, 0x70, 0x45, 0xa8, 0x93, 0x16, 0x95, 0x0c, 0x12, 0x08, 0xd7, 0xe4, 0xdb,
	0x78, 0x90, 0xe4, 0x62, 0xee, 0x9a, 0x0e, 0xe2, 0xe2, 0x2b, 0xf8, 0x06, 0x8e, 0xbe, 0x81, 0xaf,
	0xd0, 0xd5, 0xcd, 0xc9, 0xc1, 0x49, 0xdf, 0x40, 0x1c, 0x24, 0xbd, 0x5c, 0x9a, 0x8a, 0xa1, 0xdb,
	0x87, 0xfb, 0xfc, 0xf9, 0xfe, 0x3b, 0xe7, 0x61, 0x2a, 0x44, 0x9a, 0x01, 0x66, 0x89, 0xc4, 0x1a,
	0x36, 0xa8, 0xf6, 0xb1, 0x84, 0xaa, 0xe6, 0x31, 0x48, 0xac, 0x44, 0xc9, 0xe3, 0xa8, 0xe6, 0xb0,
	0x8b, 0xda, 0x37, 0x54, 0x56, 0x42, 0x09, 0x77, 0xac, 0xf5, 0x88, 0x25, 0x12, 0x75, 0x56, 0x54,
	0xfb, 0xc8, 0x58, 0x3d, 0x7f, 0x28, 0xbc, 0x02, 0x29, 0xb6, 0xd5, 0x69, 0xba, 0x4e, 0xf5, 0xee,
	0x18, 0x4f, 0xc9, 0x31, 0x2b, 0x0a, 0xa1, 0x98, 0xe2, 0xa2, 0x90, 0x2d, 0x7b, 0xb3, 0xc7, 0xc6,
	0x19, 0x87, 0x42, 0xb5, 0xc4, 0xdd, 0x1e, 0xb1, 0xe1, 0x90, 0x25, 0xd1, 0x1a, 0xde, 0xb0, 0x9a,
	0x8b, 0xaa, 0x15, 0xdc, 0xea, 0x09, 0x4c, 0x79, 0x4d, 0x4d, 0x36, 0xce, 0xf5, 0x25, 0xa8, 0x57,
	0x4d, 0x27, 0x01, 0x87, 0xdd, 0x0b, 0x78, 0xbb, 0x05, 0xa9, 0xdc, 0xe7, 0xce, 0x55, 0x23, 0x8c,
	0x0a, 0x96, 0xc3, 0xc8, 0x1a, 0x5b, 0xf7, 0xae, 0x2c, 0x66, 0xdf, 0xa9, 0xfd, 0x9b, 0x4e, 0x9d,
	0xc9, 0x71, 0xe6, 0x16, 0x95, 0x5c, 0xa2, 0x58, 0xe4, 0xf8, 0x98, 0x74, 0x61, 0x02, 0x9e, 0xb1,
	0x1c, 0xfc, 0x5f, 0x96, 0x73, 0xad, 0xe3, 0x5e, 0xea, 0x25, 0xb9, 0x5f, 0x2c, 0xe7, 0xa2, 0x5f,
	0xdd, 0x7d, 0x80, 0xce, 0xed, 0x15, 0xfd, 0xa7, 0x5b, 0x6f, 0x3e, 0x68, 0xeb, 0x96, 0x8d, 0x3a,
	0xd3, 0xe4, 0xc9, 0x37, 0x7a, 0x3a, 0xdc, 0x87, 0xaf, 0x3f, 0x3e, 0xda, 0xc8, 0x9d, 0x37, 0xd7,
	0x79, 0x77, 0xc2, 0x3c, 0x8e, 0xb7, 0x52, 0x89, 0x1c, 0x2a, 0x89, 0x67, 0xfa, 0x5c, 0x4d, 0x82,
	0xc4, 0xb3, 0xf7, 0xde, 0xed, 0x3d, 0x1d, 0x0d, 0x6d, 0x61, 0xf1, 0xc7, 0x72, 0xa6, 0xb1, 0xc8,
	0xcf, 0x4e, 0xb3, 0xb8, 0xf1, 0xef, 0x4e, 0x56, 0xcd, 0x55, 0x56, 0xd6, 0xeb, 0xa7, 0xad, 0x35,
	0x15, 0x19, 0x2b, 0x52, 0x24, 0xaa, 0x14, 0xa7, 0x50, 0x1c, 0x6e, 0x86, 0x8f, 0xc5, 0x86, 0xbf,
	0xee, 0x23, 0x03, 0x3e, 0xd9, 0x97, 0x96, 0x94, 0x7e, 0xb6, 0xc7, 0x4b, 0x1d, 0x48, 0x13, 0x89,
	0x34, 0x6c, 0x50, 0xe0, 0xa3, 0xb6, 0xb0, 0xdc, 0x1b, 0x49, 0x48, 0x13, 0x19, 0x76, 0x92, 0x30,
	0xf0, 0x43, 0x23, 0xf9, 0x69, 0x4f, 0xf5, 0x3b, 0x21, 0x34, 0x91, 0x84, 0x74, 0x22, 0x42, 0x02,
	0x9f, 0x10, 0x23, 0x5b, 0x5f, 0x3e, 0xf4, 0x79, 0xff, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac,
	0xc6, 0x51, 0x0f, 0x61, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TopicViewServiceClient is the client API for TopicViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopicViewServiceClient interface {
	// Returns the requested topic view in full detail.
	GetTopicView(ctx context.Context, in *GetTopicViewRequest, opts ...grpc.CallOption) (*resources.TopicView, error)
}

type topicViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicViewServiceClient(cc grpc.ClientConnInterface) TopicViewServiceClient {
	return &topicViewServiceClient{cc}
}

func (c *topicViewServiceClient) GetTopicView(ctx context.Context, in *GetTopicViewRequest, opts ...grpc.CallOption) (*resources.TopicView, error) {
	out := new(resources.TopicView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.TopicViewService/GetTopicView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicViewServiceServer is the server API for TopicViewService service.
type TopicViewServiceServer interface {
	// Returns the requested topic view in full detail.
	GetTopicView(context.Context, *GetTopicViewRequest) (*resources.TopicView, error)
}

// UnimplementedTopicViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTopicViewServiceServer struct {
}

func (*UnimplementedTopicViewServiceServer) GetTopicView(ctx context.Context, req *GetTopicViewRequest) (*resources.TopicView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicView not implemented")
}

func RegisterTopicViewServiceServer(s *grpc.Server, srv TopicViewServiceServer) {
	s.RegisterService(&_TopicViewService_serviceDesc, srv)
}

func _TopicViewService_GetTopicView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicViewServiceServer).GetTopicView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.TopicViewService/GetTopicView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicViewServiceServer).GetTopicView(ctx, req.(*GetTopicViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopicViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.TopicViewService",
	HandlerType: (*TopicViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopicView",
			Handler:    _TopicViewService_GetTopicView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/topic_view_service.proto",
}
