// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: gvm/shard/query.proto

package shard

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
	Query_Params_FullMethodName          = "/gvm.shard.Query/Params"
	Query_Shard_FullMethodName           = "/gvm.shard.Query/Shard"
	Query_ShardAll_FullMethodName        = "/gvm.shard.Query/ShardAll"
	Query_ToParent_FullMethodName        = "/gvm.shard.Query/ToParent"
	Query_ToParentAll_FullMethodName     = "/gvm.shard.Query/ToParentAll"
	Query_ToLeftChild_FullMethodName     = "/gvm.shard.Query/ToLeftChild"
	Query_ToLeftChildAll_FullMethodName  = "/gvm.shard.Query/ToLeftChildAll"
	Query_ToRightChild_FullMethodName    = "/gvm.shard.Query/ToRightChild"
	Query_ToRightChildAll_FullMethodName = "/gvm.shard.Query/ToRightChildAll"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Shard items.
	Shard(ctx context.Context, in *QueryGetShardRequest, opts ...grpc.CallOption) (*QueryGetShardResponse, error)
	ShardAll(ctx context.Context, in *QueryAllShardRequest, opts ...grpc.CallOption) (*QueryAllShardResponse, error)
	// Queries a list of ToParent items.
	ToParent(ctx context.Context, in *QueryGetToParentRequest, opts ...grpc.CallOption) (*QueryGetToParentResponse, error)
	ToParentAll(ctx context.Context, in *QueryAllToParentRequest, opts ...grpc.CallOption) (*QueryAllToParentResponse, error)
	// Queries a list of ToLeftChild items.
	ToLeftChild(ctx context.Context, in *QueryGetToLeftChildRequest, opts ...grpc.CallOption) (*QueryGetToLeftChildResponse, error)
	ToLeftChildAll(ctx context.Context, in *QueryAllToLeftChildRequest, opts ...grpc.CallOption) (*QueryAllToLeftChildResponse, error)
	// Queries a list of ToRightChild items.
	ToRightChild(ctx context.Context, in *QueryGetToRightChildRequest, opts ...grpc.CallOption) (*QueryGetToRightChildResponse, error)
	ToRightChildAll(ctx context.Context, in *QueryAllToRightChildRequest, opts ...grpc.CallOption) (*QueryAllToRightChildResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Shard(ctx context.Context, in *QueryGetShardRequest, opts ...grpc.CallOption) (*QueryGetShardResponse, error) {
	out := new(QueryGetShardResponse)
	err := c.cc.Invoke(ctx, Query_Shard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ShardAll(ctx context.Context, in *QueryAllShardRequest, opts ...grpc.CallOption) (*QueryAllShardResponse, error) {
	out := new(QueryAllShardResponse)
	err := c.cc.Invoke(ctx, Query_ShardAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ToParent(ctx context.Context, in *QueryGetToParentRequest, opts ...grpc.CallOption) (*QueryGetToParentResponse, error) {
	out := new(QueryGetToParentResponse)
	err := c.cc.Invoke(ctx, Query_ToParent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ToParentAll(ctx context.Context, in *QueryAllToParentRequest, opts ...grpc.CallOption) (*QueryAllToParentResponse, error) {
	out := new(QueryAllToParentResponse)
	err := c.cc.Invoke(ctx, Query_ToParentAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ToLeftChild(ctx context.Context, in *QueryGetToLeftChildRequest, opts ...grpc.CallOption) (*QueryGetToLeftChildResponse, error) {
	out := new(QueryGetToLeftChildResponse)
	err := c.cc.Invoke(ctx, Query_ToLeftChild_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ToLeftChildAll(ctx context.Context, in *QueryAllToLeftChildRequest, opts ...grpc.CallOption) (*QueryAllToLeftChildResponse, error) {
	out := new(QueryAllToLeftChildResponse)
	err := c.cc.Invoke(ctx, Query_ToLeftChildAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ToRightChild(ctx context.Context, in *QueryGetToRightChildRequest, opts ...grpc.CallOption) (*QueryGetToRightChildResponse, error) {
	out := new(QueryGetToRightChildResponse)
	err := c.cc.Invoke(ctx, Query_ToRightChild_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ToRightChildAll(ctx context.Context, in *QueryAllToRightChildRequest, opts ...grpc.CallOption) (*QueryAllToRightChildResponse, error) {
	out := new(QueryAllToRightChildResponse)
	err := c.cc.Invoke(ctx, Query_ToRightChildAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Shard items.
	Shard(context.Context, *QueryGetShardRequest) (*QueryGetShardResponse, error)
	ShardAll(context.Context, *QueryAllShardRequest) (*QueryAllShardResponse, error)
	// Queries a list of ToParent items.
	ToParent(context.Context, *QueryGetToParentRequest) (*QueryGetToParentResponse, error)
	ToParentAll(context.Context, *QueryAllToParentRequest) (*QueryAllToParentResponse, error)
	// Queries a list of ToLeftChild items.
	ToLeftChild(context.Context, *QueryGetToLeftChildRequest) (*QueryGetToLeftChildResponse, error)
	ToLeftChildAll(context.Context, *QueryAllToLeftChildRequest) (*QueryAllToLeftChildResponse, error)
	// Queries a list of ToRightChild items.
	ToRightChild(context.Context, *QueryGetToRightChildRequest) (*QueryGetToRightChildResponse, error)
	ToRightChildAll(context.Context, *QueryAllToRightChildRequest) (*QueryAllToRightChildResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Shard(context.Context, *QueryGetShardRequest) (*QueryGetShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shard not implemented")
}
func (UnimplementedQueryServer) ShardAll(context.Context, *QueryAllShardRequest) (*QueryAllShardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShardAll not implemented")
}
func (UnimplementedQueryServer) ToParent(context.Context, *QueryGetToParentRequest) (*QueryGetToParentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToParent not implemented")
}
func (UnimplementedQueryServer) ToParentAll(context.Context, *QueryAllToParentRequest) (*QueryAllToParentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToParentAll not implemented")
}
func (UnimplementedQueryServer) ToLeftChild(context.Context, *QueryGetToLeftChildRequest) (*QueryGetToLeftChildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToLeftChild not implemented")
}
func (UnimplementedQueryServer) ToLeftChildAll(context.Context, *QueryAllToLeftChildRequest) (*QueryAllToLeftChildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToLeftChildAll not implemented")
}
func (UnimplementedQueryServer) ToRightChild(context.Context, *QueryGetToRightChildRequest) (*QueryGetToRightChildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToRightChild not implemented")
}
func (UnimplementedQueryServer) ToRightChildAll(context.Context, *QueryAllToRightChildRequest) (*QueryAllToRightChildResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToRightChildAll not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Shard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Shard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Shard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Shard(ctx, req.(*QueryGetShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ShardAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllShardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ShardAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ShardAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ShardAll(ctx, req.(*QueryAllShardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ToParent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetToParentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ToParent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ToParent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ToParent(ctx, req.(*QueryGetToParentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ToParentAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllToParentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ToParentAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ToParentAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ToParentAll(ctx, req.(*QueryAllToParentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ToLeftChild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetToLeftChildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ToLeftChild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ToLeftChild_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ToLeftChild(ctx, req.(*QueryGetToLeftChildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ToLeftChildAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllToLeftChildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ToLeftChildAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ToLeftChildAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ToLeftChildAll(ctx, req.(*QueryAllToLeftChildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ToRightChild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetToRightChildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ToRightChild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ToRightChild_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ToRightChild(ctx, req.(*QueryGetToRightChildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ToRightChildAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllToRightChildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ToRightChildAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ToRightChildAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ToRightChildAll(ctx, req.(*QueryAllToRightChildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gvm.shard.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Shard",
			Handler:    _Query_Shard_Handler,
		},
		{
			MethodName: "ShardAll",
			Handler:    _Query_ShardAll_Handler,
		},
		{
			MethodName: "ToParent",
			Handler:    _Query_ToParent_Handler,
		},
		{
			MethodName: "ToParentAll",
			Handler:    _Query_ToParentAll_Handler,
		},
		{
			MethodName: "ToLeftChild",
			Handler:    _Query_ToLeftChild_Handler,
		},
		{
			MethodName: "ToLeftChildAll",
			Handler:    _Query_ToLeftChildAll_Handler,
		},
		{
			MethodName: "ToRightChild",
			Handler:    _Query_ToRightChild_Handler,
		},
		{
			MethodName: "ToRightChildAll",
			Handler:    _Query_ToRightChildAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gvm/shard/query.proto",
}
