// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: message.proto

package message

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

// TaskManagerClient is the client API for TaskManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskManagerClient interface {
	RegisterTask(ctx context.Context, in *RegisterTaskRequest, opts ...grpc.CallOption) (*RegisterTaskReply, error)
	GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskReply, error)
	TaskList(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListReply, error)
}

type taskManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskManagerClient(cc grpc.ClientConnInterface) TaskManagerClient {
	return &taskManagerClient{cc}
}

func (c *taskManagerClient) RegisterTask(ctx context.Context, in *RegisterTaskRequest, opts ...grpc.CallOption) (*RegisterTaskReply, error) {
	out := new(RegisterTaskReply)
	err := c.cc.Invoke(ctx, "/message.TaskManager/RegisterTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) GetTask(ctx context.Context, in *GetTaskRequest, opts ...grpc.CallOption) (*GetTaskReply, error) {
	out := new(GetTaskReply)
	err := c.cc.Invoke(ctx, "/message.TaskManager/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) TaskList(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListReply, error) {
	out := new(TaskListReply)
	err := c.cc.Invoke(ctx, "/message.TaskManager/TaskList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskManagerServer is the server API for TaskManager service.
// All implementations must embed UnimplementedTaskManagerServer
// for forward compatibility
type TaskManagerServer interface {
	RegisterTask(context.Context, *RegisterTaskRequest) (*RegisterTaskReply, error)
	GetTask(context.Context, *GetTaskRequest) (*GetTaskReply, error)
	TaskList(context.Context, *TaskListRequest) (*TaskListReply, error)
	mustEmbedUnimplementedTaskManagerServer()
}

// UnimplementedTaskManagerServer must be embedded to have forward compatible implementations.
type UnimplementedTaskManagerServer struct {
}

func (UnimplementedTaskManagerServer) RegisterTask(context.Context, *RegisterTaskRequest) (*RegisterTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterTask not implemented")
}
func (UnimplementedTaskManagerServer) GetTask(context.Context, *GetTaskRequest) (*GetTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTask not implemented")
}
func (UnimplementedTaskManagerServer) TaskList(context.Context, *TaskListRequest) (*TaskListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TaskList not implemented")
}
func (UnimplementedTaskManagerServer) mustEmbedUnimplementedTaskManagerServer() {}

// UnsafeTaskManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskManagerServer will
// result in compilation errors.
type UnsafeTaskManagerServer interface {
	mustEmbedUnimplementedTaskManagerServer()
}

func RegisterTaskManagerServer(s grpc.ServiceRegistrar, srv TaskManagerServer) {
	s.RegisterService(&TaskManager_ServiceDesc, srv)
}

func _TaskManager_RegisterTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).RegisterTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.TaskManager/RegisterTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).RegisterTask(ctx, req.(*RegisterTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.TaskManager/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).GetTask(ctx, req.(*GetTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_TaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).TaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.TaskManager/TaskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).TaskList(ctx, req.(*TaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskManager_ServiceDesc is the grpc.ServiceDesc for TaskManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.TaskManager",
	HandlerType: (*TaskManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterTask",
			Handler:    _TaskManager_RegisterTask_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _TaskManager_GetTask_Handler,
		},
		{
			MethodName: "TaskList",
			Handler:    _TaskManager_TaskList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
