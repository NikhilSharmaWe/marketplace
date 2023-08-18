// syntax = "proto3";

// option go_package = "github.com/NikhilSharmaWe/golangmicro/proto";

// service PriceFetcher {
// 	rpc FetchPrice(PriceRequest) returns (PriceResponse);
// }

// message PriceRequest {
// 	string ticker = 1;
// }

// message PriceResponse {
// 	string ticker = 1;
// 	float price = 2;
// }

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: proto/service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MarketplaceService_CreateShop_FullMethodName            = "/MarketplaceService/CreateShop"
	MarketplaceService_GetShopByID_FullMethodName           = "/MarketplaceService/GetShopByID"
	MarketplaceService_CreateProduct_FullMethodName         = "/MarketplaceService/CreateProduct"
	MarketplaceService_GetProductByID_FullMethodName        = "/MarketplaceService/GetProductByID"
	MarketplaceService_AddServiceableProduct_FullMethodName = "/MarketplaceService/AddServiceableProduct"
	MarketplaceService_CreateUser_FullMethodName            = "/MarketplaceService/CreateUser"
	MarketplaceService_GetUserByID_FullMethodName           = "/MarketplaceService/GetUserByID"
)

// MarketplaceServiceClient is the client API for MarketplaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketplaceServiceClient interface {
	// Shop-related methods
	CreateShop(ctx context.Context, in *CreateShopRequest, opts ...grpc.CallOption) (*Shop, error)
	GetShopByID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Shop, error)
	// Product-related methods
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*Product, error)
	GetProductByID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Product, error)
	// Serviceable products methods
	AddServiceableProduct(ctx context.Context, in *AddServiceableProductRequest, opts ...grpc.CallOption) (*ServiceableProduct, error)
	// User-related methods
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	GetUserByID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error)
}

type marketplaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketplaceServiceClient(cc grpc.ClientConnInterface) MarketplaceServiceClient {
	return &marketplaceServiceClient{cc}
}

func (c *marketplaceServiceClient) CreateShop(ctx context.Context, in *CreateShopRequest, opts ...grpc.CallOption) (*Shop, error) {
	out := new(Shop)
	err := c.cc.Invoke(ctx, MarketplaceService_CreateShop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceServiceClient) GetShopByID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Shop, error) {
	out := new(Shop)
	err := c.cc.Invoke(ctx, MarketplaceService_GetShopByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, MarketplaceService_CreateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceServiceClient) GetProductByID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, MarketplaceService_GetProductByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceServiceClient) AddServiceableProduct(ctx context.Context, in *AddServiceableProductRequest, opts ...grpc.CallOption) (*ServiceableProduct, error) {
	out := new(ServiceableProduct)
	err := c.cc.Invoke(ctx, MarketplaceService_AddServiceableProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, MarketplaceService_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceServiceClient) GetUserByID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, MarketplaceService_GetUserByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketplaceServiceServer is the server API for MarketplaceService service.
// All implementations must embed UnimplementedMarketplaceServiceServer
// for forward compatibility
type MarketplaceServiceServer interface {
	// Shop-related methods
	CreateShop(context.Context, *CreateShopRequest) (*Shop, error)
	GetShopByID(context.Context, *GetRequest) (*Shop, error)
	// Product-related methods
	CreateProduct(context.Context, *CreateProductRequest) (*Product, error)
	GetProductByID(context.Context, *GetRequest) (*Product, error)
	// Serviceable products methods
	AddServiceableProduct(context.Context, *AddServiceableProductRequest) (*ServiceableProduct, error)
	// User-related methods
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	GetUserByID(context.Context, *GetRequest) (*User, error)
	mustEmbedUnimplementedMarketplaceServiceServer()
}

// UnimplementedMarketplaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMarketplaceServiceServer struct {
}

func (UnimplementedMarketplaceServiceServer) CreateShop(context.Context, *CreateShopRequest) (*Shop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateShop not implemented")
}
func (UnimplementedMarketplaceServiceServer) GetShopByID(context.Context, *GetRequest) (*Shop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShopByID not implemented")
}
func (UnimplementedMarketplaceServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedMarketplaceServiceServer) GetProductByID(context.Context, *GetRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductByID not implemented")
}
func (UnimplementedMarketplaceServiceServer) AddServiceableProduct(context.Context, *AddServiceableProductRequest) (*ServiceableProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddServiceableProduct not implemented")
}
func (UnimplementedMarketplaceServiceServer) CreateUser(context.Context, *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMarketplaceServiceServer) GetUserByID(context.Context, *GetRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}
func (UnimplementedMarketplaceServiceServer) mustEmbedUnimplementedMarketplaceServiceServer() {}

// UnsafeMarketplaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketplaceServiceServer will
// result in compilation errors.
type UnsafeMarketplaceServiceServer interface {
	mustEmbedUnimplementedMarketplaceServiceServer()
}

func RegisterMarketplaceServiceServer(s grpc.ServiceRegistrar, srv MarketplaceServiceServer) {
	s.RegisterService(&MarketplaceService_ServiceDesc, srv)
}

func _MarketplaceService_CreateShop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServiceServer).CreateShop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarketplaceService_CreateShop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServiceServer).CreateShop(ctx, req.(*CreateShopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceService_GetShopByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServiceServer).GetShopByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarketplaceService_GetShopByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServiceServer).GetShopByID(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarketplaceService_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceService_GetProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServiceServer).GetProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarketplaceService_GetProductByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServiceServer).GetProductByID(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceService_AddServiceableProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddServiceableProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServiceServer).AddServiceableProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarketplaceService_AddServiceableProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServiceServer).AddServiceableProduct(ctx, req.(*AddServiceableProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarketplaceService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MarketplaceService_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServiceServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MarketplaceService_GetUserByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServiceServer).GetUserByID(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MarketplaceService_ServiceDesc is the grpc.ServiceDesc for MarketplaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarketplaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MarketplaceService",
	HandlerType: (*MarketplaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateShop",
			Handler:    _MarketplaceService_CreateShop_Handler,
		},
		{
			MethodName: "GetShopByID",
			Handler:    _MarketplaceService_GetShopByID_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _MarketplaceService_CreateProduct_Handler,
		},
		{
			MethodName: "GetProductByID",
			Handler:    _MarketplaceService_GetProductByID_Handler,
		},
		{
			MethodName: "AddServiceableProduct",
			Handler:    _MarketplaceService_AddServiceableProduct_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _MarketplaceService_CreateUser_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _MarketplaceService_GetUserByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
