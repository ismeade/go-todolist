// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: task/task.proto

package task

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 定义请求的message
type ListTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *ListTaskRequest) Reset() {
	*x = ListTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTaskRequest) ProtoMessage() {}

func (x *ListTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTaskRequest.ProtoReflect.Descriptor instead.
func (*ListTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{0}
}

func (x *ListTaskRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 定义响应的message
type ListTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count     int32        `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	TaskReply []*TaskReply `protobuf:"bytes,2,rep,name=taskReply,proto3" json:"taskReply,omitempty"`
}

func (x *ListTaskReply) Reset() {
	*x = ListTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTaskReply) ProtoMessage() {}

func (x *ListTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTaskReply.ProtoReflect.Descriptor instead.
func (*ListTaskReply) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{1}
}

func (x *ListTaskReply) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListTaskReply) GetTaskReply() []*TaskReply {
	if x != nil {
		return x.TaskReply
	}
	return nil
}

type TaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId   int32  `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	TaskName string `protobuf:"bytes,2,opt,name=taskName,proto3" json:"taskName,omitempty"`
}

func (x *TaskReply) Reset() {
	*x = TaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskReply) ProtoMessage() {}

func (x *TaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_task_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskReply.ProtoReflect.Descriptor instead.
func (*TaskReply) Descriptor() ([]byte, []int) {
	return file_task_task_proto_rawDescGZIP(), []int{2}
}

func (x *TaskReply) GetTaskId() int32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *TaskReply) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

var File_task_task_proto protoreflect.FileDescriptor

var file_task_task_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x29, 0x0a,
	0x0f, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x33, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x3f, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x61, 0x73,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x53, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x1b, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x6f,
	0x2d, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_task_task_proto_rawDescOnce sync.Once
	file_task_task_proto_rawDescData = file_task_task_proto_rawDesc
)

func file_task_task_proto_rawDescGZIP() []byte {
	file_task_task_proto_rawDescOnce.Do(func() {
		file_task_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_task_proto_rawDescData)
	})
	return file_task_task_proto_rawDescData
}

var file_task_task_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_task_task_proto_goTypes = []interface{}{
	(*ListTaskRequest)(nil), // 0: helloworld.ListTaskRequest
	(*ListTaskReply)(nil),   // 1: helloworld.ListTaskReply
	(*TaskReply)(nil),       // 2: helloworld.TaskReply
}
var file_task_task_proto_depIdxs = []int32{
	2, // 0: helloworld.ListTaskReply.taskReply:type_name -> helloworld.TaskReply
	0, // 1: helloworld.TaskService.ListTask:input_type -> helloworld.ListTaskRequest
	1, // 2: helloworld.TaskService.ListTask:output_type -> helloworld.ListTaskReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_task_task_proto_init() }
func file_task_task_proto_init() {
	if File_task_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_task_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTaskRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTaskReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_task_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_task_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_task_proto_goTypes,
		DependencyIndexes: file_task_task_proto_depIdxs,
		MessageInfos:      file_task_task_proto_msgTypes,
	}.Build()
	File_task_task_proto = out.File
	file_task_task_proto_rawDesc = nil
	file_task_task_proto_goTypes = nil
	file_task_task_proto_depIdxs = nil
}