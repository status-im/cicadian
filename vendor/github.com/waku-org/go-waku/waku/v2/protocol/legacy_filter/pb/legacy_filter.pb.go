// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: legacy_filter.proto

// 12/WAKU2-FILTER rfc: https://rfc.vac.dev/spec/12/
// Protocol identifier: /vac/waku/filter/2.0.0-beta1

package pb

import (
	pb "github.com/waku-org/go-waku/waku/v2/protocol/pb"
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

type FilterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscribe      bool                           `protobuf:"varint,1,opt,name=subscribe,proto3" json:"subscribe,omitempty"`
	Topic          string                         `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	ContentFilters []*FilterRequest_ContentFilter `protobuf:"bytes,3,rep,name=content_filters,json=contentFilters,proto3" json:"content_filters,omitempty"`
}

func (x *FilterRequest) Reset() {
	*x = FilterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_legacy_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterRequest) ProtoMessage() {}

func (x *FilterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_legacy_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterRequest.ProtoReflect.Descriptor instead.
func (*FilterRequest) Descriptor() ([]byte, []int) {
	return file_legacy_filter_proto_rawDescGZIP(), []int{0}
}

func (x *FilterRequest) GetSubscribe() bool {
	if x != nil {
		return x.Subscribe
	}
	return false
}

func (x *FilterRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *FilterRequest) GetContentFilters() []*FilterRequest_ContentFilter {
	if x != nil {
		return x.ContentFilters
	}
	return nil
}

type MessagePush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*pb.WakuMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *MessagePush) Reset() {
	*x = MessagePush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_legacy_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessagePush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagePush) ProtoMessage() {}

func (x *MessagePush) ProtoReflect() protoreflect.Message {
	mi := &file_legacy_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagePush.ProtoReflect.Descriptor instead.
func (*MessagePush) Descriptor() ([]byte, []int) {
	return file_legacy_filter_proto_rawDescGZIP(), []int{1}
}

func (x *MessagePush) GetMessages() []*pb.WakuMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

type FilterRpc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string         `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Request   *FilterRequest `protobuf:"bytes,2,opt,name=request,proto3,oneof" json:"request,omitempty"`
	Push      *MessagePush   `protobuf:"bytes,3,opt,name=push,proto3,oneof" json:"push,omitempty"`
}

func (x *FilterRpc) Reset() {
	*x = FilterRpc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_legacy_filter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterRpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterRpc) ProtoMessage() {}

func (x *FilterRpc) ProtoReflect() protoreflect.Message {
	mi := &file_legacy_filter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterRpc.ProtoReflect.Descriptor instead.
func (*FilterRpc) Descriptor() ([]byte, []int) {
	return file_legacy_filter_proto_rawDescGZIP(), []int{2}
}

func (x *FilterRpc) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *FilterRpc) GetRequest() *FilterRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *FilterRpc) GetPush() *MessagePush {
	if x != nil {
		return x.Push
	}
	return nil
}

type FilterRequest_ContentFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentTopic string `protobuf:"bytes,1,opt,name=content_topic,json=contentTopic,proto3" json:"content_topic,omitempty"`
}

func (x *FilterRequest_ContentFilter) Reset() {
	*x = FilterRequest_ContentFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_legacy_filter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterRequest_ContentFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterRequest_ContentFilter) ProtoMessage() {}

func (x *FilterRequest_ContentFilter) ProtoReflect() protoreflect.Message {
	mi := &file_legacy_filter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterRequest_ContentFilter.ProtoReflect.Descriptor instead.
func (*FilterRequest_ContentFilter) Descriptor() ([]byte, []int) {
	return file_legacy_filter_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FilterRequest_ContentFilter) GetContentTopic() string {
	if x != nil {
		return x.ContentTopic
	}
	return ""
}

var File_legacy_filter_proto protoreflect.FileDescriptor

var file_legacy_filter_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x77, 0x61, 0x6b, 0x75, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1d, 0x77, 0x61, 0x6b, 0x75,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x0d, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x59, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x77, 0x61, 0x6b, 0x75, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x34, 0x0a, 0x0d, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x22, 0x47, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x75, 0x73, 0x68, 0x12,
	0x38, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x61, 0x6b, 0x75, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x61, 0x6b, 0x75, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x09, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x70, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x77, 0x61, 0x6b, 0x75, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x04, 0x70, 0x75, 0x73,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x77, 0x61, 0x6b, 0x75, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x75, 0x73, 0x68, 0x48, 0x01, 0x52, 0x04, 0x70, 0x75, 0x73,
	0x68, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_legacy_filter_proto_rawDescOnce sync.Once
	file_legacy_filter_proto_rawDescData = file_legacy_filter_proto_rawDesc
)

func file_legacy_filter_proto_rawDescGZIP() []byte {
	file_legacy_filter_proto_rawDescOnce.Do(func() {
		file_legacy_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_legacy_filter_proto_rawDescData)
	})
	return file_legacy_filter_proto_rawDescData
}

var file_legacy_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_legacy_filter_proto_goTypes = []interface{}{
	(*FilterRequest)(nil),               // 0: waku.filter.v2beta1.FilterRequest
	(*MessagePush)(nil),                 // 1: waku.filter.v2beta1.MessagePush
	(*FilterRpc)(nil),                   // 2: waku.filter.v2beta1.FilterRpc
	(*FilterRequest_ContentFilter)(nil), // 3: waku.filter.v2beta1.FilterRequest.ContentFilter
	(*pb.WakuMessage)(nil),              // 4: waku.message.v1.WakuMessage
}
var file_legacy_filter_proto_depIdxs = []int32{
	3, // 0: waku.filter.v2beta1.FilterRequest.content_filters:type_name -> waku.filter.v2beta1.FilterRequest.ContentFilter
	4, // 1: waku.filter.v2beta1.MessagePush.messages:type_name -> waku.message.v1.WakuMessage
	0, // 2: waku.filter.v2beta1.FilterRpc.request:type_name -> waku.filter.v2beta1.FilterRequest
	1, // 3: waku.filter.v2beta1.FilterRpc.push:type_name -> waku.filter.v2beta1.MessagePush
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_legacy_filter_proto_init() }
func file_legacy_filter_proto_init() {
	if File_legacy_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_legacy_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterRequest); i {
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
		file_legacy_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessagePush); i {
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
		file_legacy_filter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterRpc); i {
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
		file_legacy_filter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterRequest_ContentFilter); i {
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
	file_legacy_filter_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_legacy_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_legacy_filter_proto_goTypes,
		DependencyIndexes: file_legacy_filter_proto_depIdxs,
		MessageInfos:      file_legacy_filter_proto_msgTypes,
	}.Build()
	File_legacy_filter_proto = out.File
	file_legacy_filter_proto_rawDesc = nil
	file_legacy_filter_proto_goTypes = nil
	file_legacy_filter_proto_depIdxs = nil
}
