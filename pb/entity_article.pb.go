// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: entity_article.proto

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

type ArticleID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ArticleID) Reset() {
	*x = ArticleID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleID) ProtoMessage() {}

func (x *ArticleID) ProtoReflect() protoreflect.Message {
	mi := &file_entity_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleID.ProtoReflect.Descriptor instead.
func (*ArticleID) Descriptor() ([]byte, []int) {
	return file_entity_article_proto_rawDescGZIP(), []int{0}
}

func (x *ArticleID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ArticleIDList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ArticleIDList) Reset() {
	*x = ArticleIDList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleIDList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleIDList) ProtoMessage() {}

func (x *ArticleIDList) ProtoReflect() protoreflect.Message {
	mi := &file_entity_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleIDList.ProtoReflect.Descriptor instead.
func (*ArticleIDList) Descriptor() ([]byte, []int) {
	return file_entity_article_proto_rawDescGZIP(), []int{1}
}

func (x *ArticleIDList) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type ArticleIDWithPager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // common.Pager pager = 2;
}

func (x *ArticleIDWithPager) Reset() {
	*x = ArticleIDWithPager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_article_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleIDWithPager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleIDWithPager) ProtoMessage() {}

func (x *ArticleIDWithPager) ProtoReflect() protoreflect.Message {
	mi := &file_entity_article_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleIDWithPager.ProtoReflect.Descriptor instead.
func (*ArticleIDWithPager) Descriptor() ([]byte, []int) {
	return file_entity_article_proto_rawDescGZIP(), []int{2}
}

func (x *ArticleIDWithPager) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserID       string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Title        string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description  string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt    string `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	ThumbnailURL string `protobuf:"bytes,6,opt,name=thumbnailURL,proto3" json:"thumbnailURL,omitempty"` // repeated entity.SearchHighlight highlights = 6;
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_article_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_entity_article_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_entity_article_proto_rawDescGZIP(), []int{3}
}

func (x *Article) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Article) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Article) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Article) GetThumbnailURL() string {
	if x != nil {
		return x.ThumbnailURL
	}
	return ""
}

type ArticleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Article `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ArticleList) Reset() {
	*x = ArticleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_article_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleList) ProtoMessage() {}

func (x *ArticleList) ProtoReflect() protoreflect.Message {
	mi := &file_entity_article_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleList.ProtoReflect.Descriptor instead.
func (*ArticleList) Descriptor() ([]byte, []int) {
	return file_entity_article_proto_rawDescGZIP(), []int{4}
}

func (x *ArticleList) GetItems() []*Article {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_entity_article_proto protoreflect.FileDescriptor

var file_entity_article_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x68, 0x69, 0x67, 0x68,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x09, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0d, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x44, 0x57, 0x69, 0x74, 0x68, 0x50, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xab, 0x01, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x52, 0x4c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x52, 0x4c, 0x22,
	0x34, 0x0a, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x04, 0x5a, 0x02, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_entity_article_proto_rawDescOnce sync.Once
	file_entity_article_proto_rawDescData = file_entity_article_proto_rawDesc
)

func file_entity_article_proto_rawDescGZIP() []byte {
	file_entity_article_proto_rawDescOnce.Do(func() {
		file_entity_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_article_proto_rawDescData)
	})
	return file_entity_article_proto_rawDescData
}

var file_entity_article_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_entity_article_proto_goTypes = []interface{}{
	(*ArticleID)(nil),          // 0: entity.ArticleID
	(*ArticleIDList)(nil),      // 1: entity.ArticleIDList
	(*ArticleIDWithPager)(nil), // 2: entity.ArticleIDWithPager
	(*Article)(nil),            // 3: entity.Article
	(*ArticleList)(nil),        // 4: entity.ArticleList
}
var file_entity_article_proto_depIdxs = []int32{
	3, // 0: entity.ArticleList.items:type_name -> entity.Article
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_entity_article_proto_init() }
func file_entity_article_proto_init() {
	if File_entity_article_proto != nil {
		return
	}
	file_common_proto_init()
	file_entity_search_highlight_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_entity_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleID); i {
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
		file_entity_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleIDList); i {
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
		file_entity_article_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleIDWithPager); i {
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
		file_entity_article_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_entity_article_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleList); i {
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
			RawDescriptor: file_entity_article_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_article_proto_goTypes,
		DependencyIndexes: file_entity_article_proto_depIdxs,
		MessageInfos:      file_entity_article_proto_msgTypes,
	}.Build()
	File_entity_article_proto = out.File
	file_entity_article_proto_rawDesc = nil
	file_entity_article_proto_goTypes = nil
	file_entity_article_proto_depIdxs = nil
}
