// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: pkg/api/chat.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          string                 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageRequest) Reset() {
	*x = MessageRequest{}
	mi := &file_pkg_api_chat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageRequest) ProtoMessage() {}

func (x *MessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_chat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageRequest.ProtoReflect.Descriptor instead.
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return file_pkg_api_chat_proto_rawDescGZIP(), []int{0}
}

func (x *MessageRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *MessageRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type MessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          string                 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Timestamp     string                 `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	mi := &file_pkg_api_chat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_chat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_pkg_api_chat_proto_rawDescGZIP(), []int{1}
}

func (x *MessageResponse) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *MessageResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *MessageResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type ConnectRequst struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          string                 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectRequst) Reset() {
	*x = ConnectRequst{}
	mi := &file_pkg_api_chat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectRequst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequst) ProtoMessage() {}

func (x *ConnectRequst) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_api_chat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequst.ProtoReflect.Descriptor instead.
func (*ConnectRequst) Descriptor() ([]byte, []int) {
	return file_pkg_api_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectRequst) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

var File_pkg_api_chat_proto protoreflect.FileDescriptor

const file_pkg_api_chat_proto_rawDesc = "" +
	"\n" +
	"\x12pkg/api/chat.proto\x12\x03api\"8\n" +
	"\x0eMessageRequest\x12\x12\n" +
	"\x04user\x18\x01 \x01(\tR\x04user\x12\x12\n" +
	"\x04text\x18\x02 \x01(\tR\x04text\"W\n" +
	"\x0fMessageResponse\x12\x12\n" +
	"\x04user\x18\x01 \x01(\tR\x04user\x12\x12\n" +
	"\x04text\x18\x02 \x01(\tR\x04text\x12\x1c\n" +
	"\ttimestamp\x18\x03 \x01(\tR\ttimestamp\"#\n" +
	"\rConnectRequst\x12\x12\n" +
	"\x04user\x18\x01 \x01(\tR\x04user2\x7f\n" +
	"\x04Chat\x128\n" +
	"\vSendMessage\x12\x13.api.MessageRequest\x1a\x14.api.MessageResponse\x12=\n" +
	"\x0eRecieveMessage\x12\x13.api.MessageRequest\x1a\x14.api.MessageResponse0\x01B+Z)github.com/DKowalski25/gRpas_Chat/pkg/apib\x06proto3"

var (
	file_pkg_api_chat_proto_rawDescOnce sync.Once
	file_pkg_api_chat_proto_rawDescData []byte
)

func file_pkg_api_chat_proto_rawDescGZIP() []byte {
	file_pkg_api_chat_proto_rawDescOnce.Do(func() {
		file_pkg_api_chat_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_api_chat_proto_rawDesc), len(file_pkg_api_chat_proto_rawDesc)))
	})
	return file_pkg_api_chat_proto_rawDescData
}

var file_pkg_api_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_api_chat_proto_goTypes = []any{
	(*MessageRequest)(nil),  // 0: api.MessageRequest
	(*MessageResponse)(nil), // 1: api.MessageResponse
	(*ConnectRequst)(nil),   // 2: api.ConnectRequst
}
var file_pkg_api_chat_proto_depIdxs = []int32{
	0, // 0: api.Chat.SendMessage:input_type -> api.MessageRequest
	0, // 1: api.Chat.RecieveMessage:input_type -> api.MessageRequest
	1, // 2: api.Chat.SendMessage:output_type -> api.MessageResponse
	1, // 3: api.Chat.RecieveMessage:output_type -> api.MessageResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_api_chat_proto_init() }
func file_pkg_api_chat_proto_init() {
	if File_pkg_api_chat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_api_chat_proto_rawDesc), len(file_pkg_api_chat_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_api_chat_proto_goTypes,
		DependencyIndexes: file_pkg_api_chat_proto_depIdxs,
		MessageInfos:      file_pkg_api_chat_proto_msgTypes,
	}.Build()
	File_pkg_api_chat_proto = out.File
	file_pkg_api_chat_proto_goTypes = nil
	file_pkg_api_chat_proto_depIdxs = nil
}
