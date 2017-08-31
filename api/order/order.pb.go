// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order/order.proto

/*
Package order is a generated protocol buffer package.

It is generated from these files:
	order/order.proto

It has these top-level messages:
	Order
	CreateOrderRequest
	GetOrderRequest
	ListOrdersRequest
	ListOrdersResponse
	DeleteOrderRequest
*/
package order

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/mwitkow/go-proto-validators"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Order struct {
	Id    string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	User  string   `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Items []string `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Order) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

type CreateOrderRequest struct {
	User  string   `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Items []string `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (m *CreateOrderRequest) Reset()                    { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()               {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateOrderRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *CreateOrderRequest) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

type GetOrderRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetOrderRequest) Reset()                    { *m = GetOrderRequest{} }
func (m *GetOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOrderRequest) ProtoMessage()               {}
func (*GetOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetOrderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListOrdersRequest struct {
	// user is optional
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *ListOrdersRequest) Reset()                    { *m = ListOrdersRequest{} }
func (m *ListOrdersRequest) String() string            { return proto.CompactTextString(m) }
func (*ListOrdersRequest) ProtoMessage()               {}
func (*ListOrdersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListOrdersRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type ListOrdersResponse struct {
	Orders []*Order `protobuf:"bytes,1,rep,name=orders" json:"orders,omitempty"`
}

func (m *ListOrdersResponse) Reset()                    { *m = ListOrdersResponse{} }
func (m *ListOrdersResponse) String() string            { return proto.CompactTextString(m) }
func (*ListOrdersResponse) ProtoMessage()               {}
func (*ListOrdersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListOrdersResponse) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type DeleteOrderRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteOrderRequest) Reset()                    { *m = DeleteOrderRequest{} }
func (m *DeleteOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteOrderRequest) ProtoMessage()               {}
func (*DeleteOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteOrderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Order)(nil), "order.Order")
	proto.RegisterType((*CreateOrderRequest)(nil), "order.CreateOrderRequest")
	proto.RegisterType((*GetOrderRequest)(nil), "order.GetOrderRequest")
	proto.RegisterType((*ListOrdersRequest)(nil), "order.ListOrdersRequest")
	proto.RegisterType((*ListOrdersResponse)(nil), "order.ListOrdersResponse")
	proto.RegisterType((*DeleteOrderRequest)(nil), "order.DeleteOrderRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for OrderService service

type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error)
	ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error)
	DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*Order, error)
	UpdateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
}

type orderServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/order.OrderService/CreateOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/order.OrderService/GetOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error) {
	out := new(ListOrdersResponse)
	err := grpc.Invoke(ctx, "/order.OrderService/ListOrders", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/order.OrderService/DeleteOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/order.OrderService/UpdateOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderService service

type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*Order, error)
	GetOrder(context.Context, *GetOrderRequest) (*Order, error)
	ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error)
	DeleteOrder(context.Context, *DeleteOrderRequest) (*Order, error)
	UpdateOrder(context.Context, *Order) (*Order, error)
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/ListOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListOrders(ctx, req.(*ListOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/DeleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).DeleteOrder(ctx, req.(*DeleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _OrderService_ListOrders_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _OrderService_DeleteOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _OrderService_UpdateOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/order.proto",
}

func init() { proto.RegisterFile("order/order.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x8a, 0xda, 0x40,
	0x14, 0xc6, 0x49, 0x52, 0xa5, 0x9e, 0x48, 0xc5, 0x43, 0xff, 0xc4, 0xd0, 0x0b, 0x3b, 0x08, 0x15,
	0x41, 0x87, 0x5a, 0xe8, 0x85, 0x17, 0xa5, 0xa5, 0x85, 0x52, 0x68, 0x11, 0x2c, 0x3e, 0x40, 0x34,
	0x43, 0x3a, 0x54, 0x33, 0x69, 0x66, 0xd4, 0x8b, 0x65, 0x6f, 0xf6, 0x11, 0x76, 0x1f, 0x6d, 0x5f,
	0x61, 0x1f, 0x64, 0xd9, 0x99, 0x6c, 0xcc, 0x3f, 0xf6, 0x26, 0xcc, 0xe4, 0x9c, 0xef, 0x3b, 0xdf,
	0xf9, 0x25, 0xd0, 0x17, 0x69, 0xc8, 0x52, 0xaa, 0x9f, 0xb3, 0x24, 0x15, 0x4a, 0x60, 0x4b, 0x5f,
	0xfc, 0xb7, 0x91, 0x10, 0xd1, 0x8e, 0xd1, 0x20, 0xe1, 0x34, 0x88, 0x63, 0xa1, 0x02, 0xc5, 0x45,
	0x2c, 0x4d, 0x93, 0xff, 0x29, 0xe2, 0xea, 0xef, 0x61, 0x33, 0xdb, 0x8a, 0x3d, 0xdd, 0x9f, 0xb8,
	0xfa, 0x27, 0x4e, 0x34, 0x12, 0x53, 0x5d, 0x9c, 0x1e, 0x83, 0x1d, 0x0f, 0x03, 0x25, 0x52, 0x49,
	0xf3, 0xa3, 0xd1, 0x91, 0xaf, 0xd0, 0x5a, 0x3e, 0xd8, 0xe3, 0x0b, 0xb0, 0x79, 0xe8, 0x59, 0x43,
	0x6b, 0xdc, 0x59, 0xd9, 0x3c, 0x44, 0x84, 0x67, 0x07, 0xc9, 0x52, 0xcf, 0xd6, 0x6f, 0xf4, 0x19,
	0x5f, 0x42, 0x8b, 0x2b, 0xb6, 0x97, 0x9e, 0x33, 0x74, 0xc6, 0x9d, 0x95, 0xb9, 0x90, 0xcf, 0x80,
	0xdf, 0x52, 0x16, 0x28, 0xa6, 0x8d, 0x56, 0xec, 0xff, 0x81, 0x49, 0x95, 0xeb, 0xad, 0x26, 0xbd,
	0x5d, 0xd4, 0xbf, 0x83, 0xde, 0x0f, 0xa6, 0x4a, 0xe2, 0x4a, 0x18, 0xf2, 0x1e, 0xfa, 0xbf, 0xb8,
	0x34, 0x3d, 0xf2, 0x89, 0x09, 0x64, 0x01, 0x58, 0x6c, 0x94, 0x89, 0x88, 0x25, 0xc3, 0x11, 0xb4,
	0x35, 0x43, 0xe9, 0x59, 0x43, 0x67, 0xec, 0xce, 0xbb, 0x33, 0xc3, 0xd7, 0xcc, 0xcc, 0x6a, 0x64,
	0x04, 0xf8, 0x9d, 0xed, 0x58, 0x65, 0x8f, 0x4a, 0x94, 0xf9, 0xb5, 0x03, 0x5d, 0xdd, 0xf0, 0x87,
	0xa5, 0x47, 0xbe, 0x65, 0xf8, 0x1b, 0xdc, 0xc2, 0xfa, 0x38, 0xc8, 0xbc, 0xeb, 0x48, 0xfc, 0xd2,
	0x58, 0xf2, 0xea, 0xea, 0xf6, 0xee, 0xc6, 0xee, 0x11, 0xa0, 0xc7, 0x0f, 0xe6, 0x7b, 0xcb, 0x85,
	0x35, 0xc1, 0x9f, 0xf0, 0xfc, 0x91, 0x06, 0xbe, 0xce, 0x04, 0x15, 0x3c, 0x15, 0xa3, 0x37, 0xda,
	0xa8, 0x8f, 0xbd, 0xb3, 0x11, 0xbd, 0xe0, 0xe1, 0x25, 0xae, 0x01, 0xce, 0x30, 0xd0, 0xcb, 0x44,
	0x35, 0x90, 0xfe, 0xa0, 0xa1, 0x62, 0xc8, 0x11, 0xd4, 0xde, 0x5d, 0x2c, 0x84, 0xc4, 0x25, 0xb8,
	0x05, 0x4e, 0xf9, 0xc2, 0x75, 0x76, 0xcd, 0x39, 0x27, 0xb5, 0x9c, 0x5f, 0xc0, 0x5d, 0x27, 0x61,
	0x4e, 0xb0, 0xa4, 0x6a, 0x86, 0xe6, 0x97, 0xa1, 0x6d, 0xda, 0xfa, 0x67, 0xfe, 0x78, 0x1f, 0x00,
	0x00, 0xff, 0xff, 0xc7, 0x41, 0xd8, 0x1d, 0x3e, 0x03, 0x00, 0x00,
}