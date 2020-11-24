// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/payments_account_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
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

// Request message for fetching all accessible payments accounts.
type ListPaymentsAccountsRequest struct {
	// Required. The ID of the customer to apply the PaymentsAccount list operation to.
	CustomerId           string   `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPaymentsAccountsRequest) Reset()         { *m = ListPaymentsAccountsRequest{} }
func (m *ListPaymentsAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*ListPaymentsAccountsRequest) ProtoMessage()    {}
func (*ListPaymentsAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68acb38ba4f095b7, []int{0}
}

func (m *ListPaymentsAccountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPaymentsAccountsRequest.Unmarshal(m, b)
}
func (m *ListPaymentsAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPaymentsAccountsRequest.Marshal(b, m, deterministic)
}
func (m *ListPaymentsAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPaymentsAccountsRequest.Merge(m, src)
}
func (m *ListPaymentsAccountsRequest) XXX_Size() int {
	return xxx_messageInfo_ListPaymentsAccountsRequest.Size(m)
}
func (m *ListPaymentsAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPaymentsAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPaymentsAccountsRequest proto.InternalMessageInfo

func (m *ListPaymentsAccountsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

// Response message for [PaymentsAccountService.ListPaymentsAccounts][google.ads.googleads.v1.services.PaymentsAccountService.ListPaymentsAccounts].
type ListPaymentsAccountsResponse struct {
	// The list of accessible payments accounts.
	PaymentsAccounts     []*resources.PaymentsAccount `protobuf:"bytes,1,rep,name=payments_accounts,json=paymentsAccounts,proto3" json:"payments_accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ListPaymentsAccountsResponse) Reset()         { *m = ListPaymentsAccountsResponse{} }
func (m *ListPaymentsAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*ListPaymentsAccountsResponse) ProtoMessage()    {}
func (*ListPaymentsAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68acb38ba4f095b7, []int{1}
}

func (m *ListPaymentsAccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPaymentsAccountsResponse.Unmarshal(m, b)
}
func (m *ListPaymentsAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPaymentsAccountsResponse.Marshal(b, m, deterministic)
}
func (m *ListPaymentsAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPaymentsAccountsResponse.Merge(m, src)
}
func (m *ListPaymentsAccountsResponse) XXX_Size() int {
	return xxx_messageInfo_ListPaymentsAccountsResponse.Size(m)
}
func (m *ListPaymentsAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPaymentsAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPaymentsAccountsResponse proto.InternalMessageInfo

func (m *ListPaymentsAccountsResponse) GetPaymentsAccounts() []*resources.PaymentsAccount {
	if m != nil {
		return m.PaymentsAccounts
	}
	return nil
}

func init() {
	proto.RegisterType((*ListPaymentsAccountsRequest)(nil), "google.ads.googleads.v1.services.ListPaymentsAccountsRequest")
	proto.RegisterType((*ListPaymentsAccountsResponse)(nil), "google.ads.googleads.v1.services.ListPaymentsAccountsResponse")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/payments_account_service.proto", fileDescriptor_68acb38ba4f095b7)
}

var fileDescriptor_68acb38ba4f095b7 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0x26, 0x29, 0x08, 0xce, 0x5e, 0x34, 0x88, 0x96, 0x76, 0xc1, 0x52, 0x7a, 0x58, 0x3c, 0xcc,
	0x98, 0x7a, 0x91, 0x91, 0x55, 0xa6, 0x0a, 0xab, 0xe0, 0xa1, 0x54, 0xe8, 0x41, 0x0a, 0x61, 0x36,
	0x33, 0xc6, 0x81, 0x76, 0x26, 0xe6, 0x4d, 0x0b, 0x22, 0x22, 0xec, 0x5f, 0xd8, 0x7f, 0xe0, 0xd1,
	0x9f, 0xb2, 0x57, 0x6f, 0x7a, 0xf1, 0xe0, 0xc9, 0x5f, 0x21, 0xd9, 0xc9, 0x64, 0x43, 0x48, 0x5d,
	0xf0, 0xf6, 0xf1, 0xde, 0xf7, 0xbe, 0xef, 0xcd, 0x97, 0x17, 0xf4, 0x2c, 0x33, 0x26, 0x5b, 0x4b,
	0xc2, 0x05, 0x10, 0x07, 0x4b, 0xb4, 0x8b, 0x09, 0xc8, 0x62, 0xa7, 0x52, 0x09, 0x24, 0xe7, 0x1f,
	0x37, 0x52, 0x5b, 0x48, 0x78, 0x9a, 0x9a, 0xad, 0xb6, 0x49, 0xd5, 0xc1, 0x79, 0x61, 0xac, 0x89,
	0x46, 0x6e, 0x0a, 0x73, 0x01, 0xb8, 0x16, 0xc0, 0xbb, 0x18, 0x7b, 0x81, 0xc1, 0xe3, 0x7d, 0x16,
	0x85, 0x04, 0xb3, 0x2d, 0xba, 0x3c, 0x9c, 0xf6, 0xe0, 0xd0, 0x4f, 0xe6, 0x8a, 0x70, 0xad, 0x8d,
	0xe5, 0x56, 0x19, 0x0d, 0x55, 0xf7, 0x5e, 0xa3, 0x9b, 0xae, 0x95, 0xac, 0xc7, 0xee, 0x37, 0x1a,
	0xef, 0x94, 0x5c, 0x8b, 0xe4, 0x54, 0xbe, 0xe7, 0x3b, 0x65, 0x0a, 0x47, 0x18, 0x3f, 0x47, 0xc3,
	0xd7, 0x0a, 0xec, 0xbc, 0x72, 0x65, 0xce, 0x14, 0x16, 0xf2, 0xc3, 0x56, 0x82, 0x8d, 0x26, 0xe8,
	0x20, 0xdd, 0x82, 0x35, 0x1b, 0x59, 0x24, 0x4a, 0xf4, 0x83, 0x51, 0x70, 0x74, 0x73, 0xd6, 0xfb,
	0xc5, 0xc2, 0x05, 0xf2, 0xf5, 0x57, 0x62, 0xfc, 0x05, 0x1d, 0x76, 0x8b, 0x40, 0x6e, 0x34, 0xc8,
	0x28, 0x41, 0xb7, 0xdb, 0xcf, 0x82, 0x7e, 0x30, 0xea, 0x1d, 0x1d, 0x4c, 0xa7, 0x78, 0x5f, 0x68,
	0x75, 0x24, 0xb8, 0xa5, 0xbb, 0xb8, 0x95, 0xb7, 0x8c, 0xa6, 0xe7, 0x21, 0xba, 0xdb, 0x62, 0xbd,
	0x71, 0x99, 0x47, 0x3f, 0x03, 0x74, 0xa7, 0x6b, 0xb9, 0xe8, 0x18, 0x5f, 0xf7, 0xb9, 0xf0, 0x3f,
	0x92, 0x19, 0x3c, 0xfd, 0xdf, 0x71, 0x97, 0xc9, 0xf8, 0xc5, 0x0f, 0xd6, 0x8c, 0xf6, 0xec, 0xfb,
	0xef, 0xf3, 0xf0, 0x61, 0x84, 0xcb, 0x63, 0xf0, 0x65, 0x20, 0x9f, 0x1a, 0x8c, 0xe3, 0x07, 0x9f,
	0x49, 0xfb, 0xe1, 0x83, 0xe1, 0x05, 0xeb, 0x5f, 0x99, 0x57, 0x28, 0x57, 0x80, 0x53, 0xb3, 0x99,
	0x9d, 0x85, 0x68, 0x92, 0x9a, 0xcd, 0xb5, 0x8b, 0xce, 0x86, 0xdd, 0xd9, 0xcd, 0xcb, 0x0b, 0x99,
	0x07, 0x6f, 0x5f, 0x56, 0x02, 0x99, 0x59, 0x73, 0x9d, 0x61, 0x53, 0x64, 0x24, 0x93, 0xfa, 0xf2,
	0x7e, 0xc8, 0x95, 0xe5, 0xfe, 0xff, 0xe6, 0x89, 0x07, 0x5f, 0xc3, 0xde, 0x09, 0x63, 0xdf, 0xc2,
	0xd1, 0x89, 0x13, 0x64, 0x02, 0xb0, 0x83, 0x25, 0x5a, 0xc6, 0xb8, 0x32, 0x86, 0x0b, 0x4f, 0x59,
	0x31, 0x01, 0xab, 0x9a, 0xb2, 0x5a, 0xc6, 0x2b, 0x4f, 0xf9, 0x13, 0x4e, 0x5c, 0x9d, 0x52, 0x26,
	0x80, 0xd2, 0x9a, 0x44, 0xe9, 0x32, 0xa6, 0xd4, 0xd3, 0x4e, 0x6f, 0x5c, 0xee, 0xf9, 0xe8, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x47, 0x71, 0xe6, 0xde, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PaymentsAccountServiceClient is the client API for PaymentsAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaymentsAccountServiceClient interface {
	// Returns all payments accounts associated with all managers
	// between the login customer ID and specified serving customer in the
	// hierarchy, inclusive.
	ListPaymentsAccounts(ctx context.Context, in *ListPaymentsAccountsRequest, opts ...grpc.CallOption) (*ListPaymentsAccountsResponse, error)
}

type paymentsAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentsAccountServiceClient(cc grpc.ClientConnInterface) PaymentsAccountServiceClient {
	return &paymentsAccountServiceClient{cc}
}

func (c *paymentsAccountServiceClient) ListPaymentsAccounts(ctx context.Context, in *ListPaymentsAccountsRequest, opts ...grpc.CallOption) (*ListPaymentsAccountsResponse, error) {
	out := new(ListPaymentsAccountsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.PaymentsAccountService/ListPaymentsAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentsAccountServiceServer is the server API for PaymentsAccountService service.
type PaymentsAccountServiceServer interface {
	// Returns all payments accounts associated with all managers
	// between the login customer ID and specified serving customer in the
	// hierarchy, inclusive.
	ListPaymentsAccounts(context.Context, *ListPaymentsAccountsRequest) (*ListPaymentsAccountsResponse, error)
}

// UnimplementedPaymentsAccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPaymentsAccountServiceServer struct {
}

func (*UnimplementedPaymentsAccountServiceServer) ListPaymentsAccounts(ctx context.Context, req *ListPaymentsAccountsRequest) (*ListPaymentsAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPaymentsAccounts not implemented")
}

func RegisterPaymentsAccountServiceServer(s *grpc.Server, srv PaymentsAccountServiceServer) {
	s.RegisterService(&_PaymentsAccountService_serviceDesc, srv)
}

func _PaymentsAccountService_ListPaymentsAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPaymentsAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentsAccountServiceServer).ListPaymentsAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.PaymentsAccountService/ListPaymentsAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentsAccountServiceServer).ListPaymentsAccounts(ctx, req.(*ListPaymentsAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentsAccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.PaymentsAccountService",
	HandlerType: (*PaymentsAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPaymentsAccounts",
			Handler:    _PaymentsAccountService_ListPaymentsAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/payments_account_service.proto",
}
