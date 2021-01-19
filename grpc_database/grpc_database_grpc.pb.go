// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpcdb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GRPCDatabaseClient is the client API for GRPCDatabase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GRPCDatabaseClient interface {
	CreateDatabase(ctx context.Context, in *Database, opts ...grpc.CallOption) (*DatabaseResult, error)
	CreateTable(ctx context.Context, in *TableRequest, opts ...grpc.CallOption) (*TableResponse, error)
	AlterTable(ctx context.Context, in *TableRequest, opts ...grpc.CallOption) (*TableResponse, error)
	ExecuteSelect(ctx context.Context, in *SelectQuery, opts ...grpc.CallOption) (*SelectQueryResult, error)
	ExecuteInsert(ctx context.Context, in *InsertQueryRequest, opts ...grpc.CallOption) (*InsertQueryResponse, error)
	ExecuteUpdate(ctx context.Context, in *UpdateQuery, opts ...grpc.CallOption) (*UpdateQueryResult, error)
}

type gRPCDatabaseClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCDatabaseClient(cc grpc.ClientConnInterface) GRPCDatabaseClient {
	return &gRPCDatabaseClient{cc}
}

func (c *gRPCDatabaseClient) CreateDatabase(ctx context.Context, in *Database, opts ...grpc.CallOption) (*DatabaseResult, error) {
	out := new(DatabaseResult)
	err := c.cc.Invoke(ctx, "/grpc_database.GRPCDatabase/CreateDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCDatabaseClient) CreateTable(ctx context.Context, in *TableRequest, opts ...grpc.CallOption) (*TableResponse, error) {
	out := new(TableResponse)
	err := c.cc.Invoke(ctx, "/grpc_database.GRPCDatabase/CreateTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCDatabaseClient) AlterTable(ctx context.Context, in *TableRequest, opts ...grpc.CallOption) (*TableResponse, error) {
	out := new(TableResponse)
	err := c.cc.Invoke(ctx, "/grpc_database.GRPCDatabase/AlterTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCDatabaseClient) ExecuteSelect(ctx context.Context, in *SelectQuery, opts ...grpc.CallOption) (*SelectQueryResult, error) {
	out := new(SelectQueryResult)
	err := c.cc.Invoke(ctx, "/grpc_database.GRPCDatabase/ExecuteSelect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCDatabaseClient) ExecuteInsert(ctx context.Context, in *InsertQueryRequest, opts ...grpc.CallOption) (*InsertQueryResponse, error) {
	out := new(InsertQueryResponse)
	err := c.cc.Invoke(ctx, "/grpc_database.GRPCDatabase/ExecuteInsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCDatabaseClient) ExecuteUpdate(ctx context.Context, in *UpdateQuery, opts ...grpc.CallOption) (*UpdateQueryResult, error) {
	out := new(UpdateQueryResult)
	err := c.cc.Invoke(ctx, "/grpc_database.GRPCDatabase/ExecuteUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCDatabaseServer is the server API for GRPCDatabase service.
// All implementations must embed UnimplementedGRPCDatabaseServer
// for forward compatibility
type GRPCDatabaseServer interface {
	CreateDatabase(context.Context, *Database) (*DatabaseResult, error)
	CreateTable(context.Context, *TableRequest) (*TableResponse, error)
	AlterTable(context.Context, *TableRequest) (*TableResponse, error)
	ExecuteSelect(context.Context, *SelectQuery) (*SelectQueryResult, error)
	ExecuteInsert(context.Context, *InsertQueryRequest) (*InsertQueryResponse, error)
	ExecuteUpdate(context.Context, *UpdateQuery) (*UpdateQueryResult, error)
	mustEmbedUnimplementedGRPCDatabaseServer()
}

// UnimplementedGRPCDatabaseServer must be embedded to have forward compatible implementations.
type UnimplementedGRPCDatabaseServer struct {
}

func (UnimplementedGRPCDatabaseServer) CreateDatabase(context.Context, *Database) (*DatabaseResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDatabase not implemented")
}
func (UnimplementedGRPCDatabaseServer) CreateTable(context.Context, *TableRequest) (*TableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTable not implemented")
}
func (UnimplementedGRPCDatabaseServer) AlterTable(context.Context, *TableRequest) (*TableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlterTable not implemented")
}
func (UnimplementedGRPCDatabaseServer) ExecuteSelect(context.Context, *SelectQuery) (*SelectQueryResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteSelect not implemented")
}
func (UnimplementedGRPCDatabaseServer) ExecuteInsert(context.Context, *InsertQueryRequest) (*InsertQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteInsert not implemented")
}
func (UnimplementedGRPCDatabaseServer) ExecuteUpdate(context.Context, *UpdateQuery) (*UpdateQueryResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteUpdate not implemented")
}
func (UnimplementedGRPCDatabaseServer) mustEmbedUnimplementedGRPCDatabaseServer() {}

// UnsafeGRPCDatabaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GRPCDatabaseServer will
// result in compilation errors.
type UnsafeGRPCDatabaseServer interface {
	mustEmbedUnimplementedGRPCDatabaseServer()
}

func RegisterGRPCDatabaseServer(s grpc.ServiceRegistrar, srv GRPCDatabaseServer) {
	s.RegisterService(&_GRPCDatabase_serviceDesc, srv)
}

func _GRPCDatabase_CreateDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Database)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCDatabaseServer).CreateDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_database.GRPCDatabase/CreateDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCDatabaseServer).CreateDatabase(ctx, req.(*Database))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCDatabase_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCDatabaseServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_database.GRPCDatabase/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCDatabaseServer).CreateTable(ctx, req.(*TableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCDatabase_AlterTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCDatabaseServer).AlterTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_database.GRPCDatabase/AlterTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCDatabaseServer).AlterTable(ctx, req.(*TableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCDatabase_ExecuteSelect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCDatabaseServer).ExecuteSelect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_database.GRPCDatabase/ExecuteSelect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCDatabaseServer).ExecuteSelect(ctx, req.(*SelectQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCDatabase_ExecuteInsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCDatabaseServer).ExecuteInsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_database.GRPCDatabase/ExecuteInsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCDatabaseServer).ExecuteInsert(ctx, req.(*InsertQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCDatabase_ExecuteUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCDatabaseServer).ExecuteUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_database.GRPCDatabase/ExecuteUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCDatabaseServer).ExecuteUpdate(ctx, req.(*UpdateQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _GRPCDatabase_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_database.GRPCDatabase",
	HandlerType: (*GRPCDatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDatabase",
			Handler:    _GRPCDatabase_CreateDatabase_Handler,
		},
		{
			MethodName: "CreateTable",
			Handler:    _GRPCDatabase_CreateTable_Handler,
		},
		{
			MethodName: "AlterTable",
			Handler:    _GRPCDatabase_AlterTable_Handler,
		},
		{
			MethodName: "ExecuteSelect",
			Handler:    _GRPCDatabase_ExecuteSelect_Handler,
		},
		{
			MethodName: "ExecuteInsert",
			Handler:    _GRPCDatabase_ExecuteInsert_Handler,
		},
		{
			MethodName: "ExecuteUpdate",
			Handler:    _GRPCDatabase_ExecuteUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_database.proto",
}
