// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: v1_rpc_delete_snippet_topic.proto

package pb

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

type DeleteSnippetTopicRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SnippetTopicId string `protobuf:"bytes,1,opt,name=snippet_topic_id,json=snippetTopicId,proto3" json:"snippet_topic_id,omitempty"`
}

func (x *DeleteSnippetTopicRequest) Reset() {
	*x = DeleteSnippetTopicRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rpc_delete_snippet_topic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSnippetTopicRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSnippetTopicRequest) ProtoMessage() {}

func (x *DeleteSnippetTopicRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rpc_delete_snippet_topic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSnippetTopicRequest.ProtoReflect.Descriptor instead.
func (*DeleteSnippetTopicRequest) Descriptor() ([]byte, []int) {
	return file_v1_rpc_delete_snippet_topic_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteSnippetTopicRequest) GetSnippetTopicId() string {
	if x != nil {
		return x.SnippetTopicId
	}
	return ""
}

var File_v1_rpc_delete_snippet_topic_proto protoreflect.FileDescriptor

var file_v1_rpc_delete_snippet_topic_proto_rawDesc = []byte{
	0x0a, 0x21, 0x76, 0x31, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f,
	0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x19, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x6e, 0x69, 0x70, 0x70,
	0x65, 0x74, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x49,
	0x64, 0x42, 0x0f, 0x5a, 0x0d, 0x73, 0x77, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_rpc_delete_snippet_topic_proto_rawDescOnce sync.Once
	file_v1_rpc_delete_snippet_topic_proto_rawDescData = file_v1_rpc_delete_snippet_topic_proto_rawDesc
)

func file_v1_rpc_delete_snippet_topic_proto_rawDescGZIP() []byte {
	file_v1_rpc_delete_snippet_topic_proto_rawDescOnce.Do(func() {
		file_v1_rpc_delete_snippet_topic_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_rpc_delete_snippet_topic_proto_rawDescData)
	})
	return file_v1_rpc_delete_snippet_topic_proto_rawDescData
}

var file_v1_rpc_delete_snippet_topic_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v1_rpc_delete_snippet_topic_proto_goTypes = []interface{}{
	(*DeleteSnippetTopicRequest)(nil), // 0: proto.DeleteSnippetTopicRequest
}
var file_v1_rpc_delete_snippet_topic_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_rpc_delete_snippet_topic_proto_init() }
func file_v1_rpc_delete_snippet_topic_proto_init() {
	if File_v1_rpc_delete_snippet_topic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_rpc_delete_snippet_topic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSnippetTopicRequest); i {
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
			RawDescriptor: file_v1_rpc_delete_snippet_topic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_rpc_delete_snippet_topic_proto_goTypes,
		DependencyIndexes: file_v1_rpc_delete_snippet_topic_proto_depIdxs,
		MessageInfos:      file_v1_rpc_delete_snippet_topic_proto_msgTypes,
	}.Build()
	File_v1_rpc_delete_snippet_topic_proto = out.File
	file_v1_rpc_delete_snippet_topic_proto_rawDesc = nil
	file_v1_rpc_delete_snippet_topic_proto_goTypes = nil
	file_v1_rpc_delete_snippet_topic_proto_depIdxs = nil
}
