// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: task/task.proto

package task

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

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskServiceClient interface {
	// 打招呼方法
	ListTask(ctx context.Context, in *ListTaskRequest, opts ...grpc.CallOption) (*ListTaskReply, error)
}

type taskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskServiceClient(cc grpc.ClientConnInterface) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) ListTask(ctx context.Context, in *ListTaskRequest, opts ...grpc.CallOption) (*ListTaskReply, error) {
	out := new(ListTaskReply)
	err := c.cc.Invoke(ctx, "/helloworld.TaskService/ListTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
// All implementations must embed UnimplementedTaskServiceServer
// for forward compatibility
type TaskServiceServer interface {
	// 打招呼方法
	ListTask(context.Context, *ListTaskRequest) (*ListTaskReply, error)
	mustEmbedUnimplementedTaskServiceServer()
}

// UnimplementedTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServiceServer struct {
}

func (UnimplementedTaskServiceServer) ListTask(context.Context, *ListTaskRequest) (*ListTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTask not implemented")
}
func (UnimplementedTaskServiceServer) mustEmbedUnimplementedTaskServiceServer() {}

// UnsafeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServiceServer will
// result in compilation errors.
type UnsafeTaskServiceServer interface {
	mustEmbedUnimplementedTaskServiceServer()
}

func RegisterTaskServiceServer(s grpc.ServiceRegistrar, srv TaskServiceServer) {
	s.RegisterService(&TaskService_ServiceDesc, srv)
}

func _TaskService_ListTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ListTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.TaskService/ListTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ListTask(ctx, req.(*ListTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskService_ServiceDesc is the grpc.ServiceDesc for TaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTask",
			Handler:    _TaskService_ListTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task/task.proto",
}
