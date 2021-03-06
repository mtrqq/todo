// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: api/tasks/rpc/tasks.proto

package rpc

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

// TasksAPIClient is the client API for TasksAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TasksAPIClient interface {
	New(ctx context.Context, in *NewTaskRequest, opts ...grpc.CallOption) (*ID, error)
	Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Task, error)
	Complete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error)
	List(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error)
}

type tasksAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTasksAPIClient(cc grpc.ClientConnInterface) TasksAPIClient {
	return &tasksAPIClient{cc}
}

func (c *tasksAPIClient) New(ctx context.Context, in *NewTaskRequest, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/rpc.TasksAPI/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksAPIClient) Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/rpc.TasksAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksAPIClient) Complete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/rpc.TasksAPI/Complete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksAPIClient) List(ctx context.Context, in *ListTasksRequest, opts ...grpc.CallOption) (*ListTasksResponse, error) {
	out := new(ListTasksResponse)
	err := c.cc.Invoke(ctx, "/rpc.TasksAPI/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksAPIServer is the server API for TasksAPI service.
// All implementations must embed UnimplementedTasksAPIServer
// for forward compatibility
type TasksAPIServer interface {
	New(context.Context, *NewTaskRequest) (*ID, error)
	Get(context.Context, *ID) (*Task, error)
	Complete(context.Context, *ID) (*ID, error)
	List(context.Context, *ListTasksRequest) (*ListTasksResponse, error)
	mustEmbedUnimplementedTasksAPIServer()
}

// UnimplementedTasksAPIServer must be embedded to have forward compatible implementations.
type UnimplementedTasksAPIServer struct {
}

func (UnimplementedTasksAPIServer) New(context.Context, *NewTaskRequest) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method New not implemented")
}
func (UnimplementedTasksAPIServer) Get(context.Context, *ID) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTasksAPIServer) Complete(context.Context, *ID) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Complete not implemented")
}
func (UnimplementedTasksAPIServer) List(context.Context, *ListTasksRequest) (*ListTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTasksAPIServer) mustEmbedUnimplementedTasksAPIServer() {}

// UnsafeTasksAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TasksAPIServer will
// result in compilation errors.
type UnsafeTasksAPIServer interface {
	mustEmbedUnimplementedTasksAPIServer()
}

func RegisterTasksAPIServer(s grpc.ServiceRegistrar, srv TasksAPIServer) {
	s.RegisterService(&TasksAPI_ServiceDesc, srv)
}

func _TasksAPI_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksAPIServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TasksAPI/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksAPIServer).New(ctx, req.(*NewTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TasksAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TasksAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksAPIServer).Get(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TasksAPI_Complete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksAPIServer).Complete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TasksAPI/Complete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksAPIServer).Complete(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TasksAPI_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksAPIServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.TasksAPI/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksAPIServer).List(ctx, req.(*ListTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TasksAPI_ServiceDesc is the grpc.ServiceDesc for TasksAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TasksAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.TasksAPI",
	HandlerType: (*TasksAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _TasksAPI_New_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TasksAPI_Get_Handler,
		},
		{
			MethodName: "Complete",
			Handler:    _TasksAPI_Complete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TasksAPI_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/tasks/rpc/tasks.proto",
}
