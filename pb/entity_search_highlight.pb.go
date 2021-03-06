// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: entity_search_highlight.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SearchHighlight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val       string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
	MatchWord string `protobuf:"bytes,3,opt,name=matchWord,proto3" json:"matchWord,omitempty"`
	LinkId    string `protobuf:"bytes,4,opt,name=link_id,json=linkId,proto3" json:"link_id,omitempty"`
}

func (x *SearchHighlight) Reset() {
	*x = SearchHighlight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_search_highlight_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchHighlight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchHighlight) ProtoMessage() {}

func (x *SearchHighlight) ProtoReflect() protoreflect.Message {
	mi := &file_entity_search_highlight_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchHighlight.ProtoReflect.Descriptor instead.
func (*SearchHighlight) Descriptor() ([]byte, []int) {
	return file_entity_search_highlight_proto_rawDescGZIP(), []int{0}
}

func (x *SearchHighlight) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SearchHighlight) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

func (x *SearchHighlight) GetMatchWord() string {
	if x != nil {
		return x.MatchWord
	}
	return ""
}

func (x *SearchHighlight) GetLinkId() string {
	if x != nil {
		return x.LinkId
	}
	return ""
}

var File_entity_search_highlight_proto protoreflect.FileDescriptor

var file_entity_search_highlight_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f,
	0x68, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x6c, 0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x48, 0x69, 0x67, 0x68, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x12, 0x1c,
	0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c,
	0x69, 0x6e, 0x6b, 0x49, 0x64, 0x42, 0x04, 0x5a, 0x02, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_entity_search_highlight_proto_rawDescOnce sync.Once
	file_entity_search_highlight_proto_rawDescData = file_entity_search_highlight_proto_rawDesc
)

func file_entity_search_highlight_proto_rawDescGZIP() []byte {
	file_entity_search_highlight_proto_rawDescOnce.Do(func() {
		file_entity_search_highlight_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_search_highlight_proto_rawDescData)
	})
	return file_entity_search_highlight_proto_rawDescData
}

var file_entity_search_highlight_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_entity_search_highlight_proto_goTypes = []interface{}{
	(*SearchHighlight)(nil), // 0: entity.SearchHighlight
}
var file_entity_search_highlight_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_entity_search_highlight_proto_init() }
func file_entity_search_highlight_proto_init() {
	if File_entity_search_highlight_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entity_search_highlight_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchHighlight); i {
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
			RawDescriptor: file_entity_search_highlight_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_search_highlight_proto_goTypes,
		DependencyIndexes: file_entity_search_highlight_proto_depIdxs,
		MessageInfos:      file_entity_search_highlight_proto_msgTypes,
	}.Build()
	File_entity_search_highlight_proto = out.File
	file_entity_search_highlight_proto_rawDesc = nil
	file_entity_search_highlight_proto_goTypes = nil
	file_entity_search_highlight_proto_depIdxs = nil
}
