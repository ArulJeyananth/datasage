// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: datasource.proto

package datasource

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

// DatasourceClient is the client API for Datasource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatasourceClient interface {
	AddDatasource(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	ListDatasource(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	DeleteDatasource(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	LogDatasource(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error)
	Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (*ScanResponse, error)
	ApplyPolicy(ctx context.Context, in *PolicyIdsRequest, opts ...grpc.CallOption) (*PolicyResponse, error)
}

type datasourceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatasourceClient(cc grpc.ClientConnInterface) DatasourceClient {
	return &datasourceClient{cc}
}

func (c *datasourceClient) AddDatasource(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/datasource.Datasource/AddDatasource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasourceClient) ListDatasource(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/datasource.Datasource/ListDatasource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasourceClient) DeleteDatasource(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/datasource.Datasource/DeleteDatasource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasourceClient) LogDatasource(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	out := new(LogResponse)
	err := c.cc.Invoke(ctx, "/datasource.Datasource/LogDatasource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasourceClient) Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (*ScanResponse, error) {
	out := new(ScanResponse)
	err := c.cc.Invoke(ctx, "/datasource.Datasource/Scan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasourceClient) ApplyPolicy(ctx context.Context, in *PolicyIdsRequest, opts ...grpc.CallOption) (*PolicyResponse, error) {
	out := new(PolicyResponse)
	err := c.cc.Invoke(ctx, "/datasource.Datasource/ApplyPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasourceServer is the server API for Datasource service.
// All implementations must embed UnimplementedDatasourceServer
// for forward compatibility
type DatasourceServer interface {
	AddDatasource(context.Context, *AddRequest) (*AddResponse, error)
	ListDatasource(context.Context, *ListRequest) (*ListResponse, error)
	DeleteDatasource(context.Context, *DeleteRequest) (*DeleteResponse, error)
	LogDatasource(context.Context, *LogRequest) (*LogResponse, error)
	Scan(context.Context, *ScanRequest) (*ScanResponse, error)
	ApplyPolicy(context.Context, *PolicyIdsRequest) (*PolicyResponse, error)
	mustEmbedUnimplementedDatasourceServer()
}

// UnimplementedDatasourceServer must be embedded to have forward compatible implementations.
type UnimplementedDatasourceServer struct {
}

func (UnimplementedDatasourceServer) AddDatasource(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDatasource not implemented")
}
func (UnimplementedDatasourceServer) ListDatasource(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDatasource not implemented")
}
func (UnimplementedDatasourceServer) DeleteDatasource(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDatasource not implemented")
}
func (UnimplementedDatasourceServer) LogDatasource(context.Context, *LogRequest) (*LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogDatasource not implemented")
}
func (UnimplementedDatasourceServer) Scan(context.Context, *ScanRequest) (*ScanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scan not implemented")
}
func (UnimplementedDatasourceServer) ApplyPolicy(context.Context, *PolicyIdsRequest) (*PolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyPolicy not implemented")
}
func (UnimplementedDatasourceServer) mustEmbedUnimplementedDatasourceServer() {}

// UnsafeDatasourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatasourceServer will
// result in compilation errors.
type UnsafeDatasourceServer interface {
	mustEmbedUnimplementedDatasourceServer()
}

func RegisterDatasourceServer(s grpc.ServiceRegistrar, srv DatasourceServer) {
	s.RegisterService(&Datasource_ServiceDesc, srv)
}

func _Datasource_AddDatasource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasourceServer).AddDatasource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datasource.Datasource/AddDatasource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasourceServer).AddDatasource(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datasource_ListDatasource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasourceServer).ListDatasource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datasource.Datasource/ListDatasource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasourceServer).ListDatasource(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datasource_DeleteDatasource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasourceServer).DeleteDatasource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datasource.Datasource/DeleteDatasource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasourceServer).DeleteDatasource(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datasource_LogDatasource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasourceServer).LogDatasource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datasource.Datasource/LogDatasource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasourceServer).LogDatasource(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datasource_Scan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasourceServer).Scan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datasource.Datasource/Scan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasourceServer).Scan(ctx, req.(*ScanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Datasource_ApplyPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasourceServer).ApplyPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/datasource.Datasource/ApplyPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasourceServer).ApplyPolicy(ctx, req.(*PolicyIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Datasource_ServiceDesc is the grpc.ServiceDesc for Datasource service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Datasource_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "datasource.Datasource",
	HandlerType: (*DatasourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDatasource",
			Handler:    _Datasource_AddDatasource_Handler,
		},
		{
			MethodName: "ListDatasource",
			Handler:    _Datasource_ListDatasource_Handler,
		},
		{
			MethodName: "DeleteDatasource",
			Handler:    _Datasource_DeleteDatasource_Handler,
		},
		{
			MethodName: "LogDatasource",
			Handler:    _Datasource_LogDatasource_Handler,
		},
		{
			MethodName: "Scan",
			Handler:    _Datasource_Scan_Handler,
		},
		{
			MethodName: "ApplyPolicy",
			Handler:    _Datasource_ApplyPolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datasource.proto",
}
