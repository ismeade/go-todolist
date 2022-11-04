package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"go-todolist/api/protobuf-spec/task"
)

type TaskService struct {
	task.UnimplementedTaskServiceServer
}

func NewServer() *TaskService {
	return &TaskService{}
}

//
func (s *TaskService) ListTask(ctx context.Context, in *task.ListTaskRequest) (*task.ListTaskReply, error) {
	return &task.ListTaskReply{Count: 10}, nil
}

func main() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()                             // 创建gRPC服务器
	task.RegisterTaskServiceServer(s, &TaskService{}) // 在gRPC服务端注册服务
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
