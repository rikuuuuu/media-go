// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: entity_admin_user.proto

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

type AdminUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *AdminUser) Reset() {
	*x = AdminUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_admin_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminUser) ProtoMessage() {}

func (x *AdminUser) ProtoReflect() protoreflect.Message {
	mi := &file_entity_admin_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminUser.ProtoReflect.Descriptor instead.
func (*AdminUser) Descriptor() ([]byte, []int) {
	return file_entity_admin_user_proto_rawDescGZIP(), []int{0}
}

func (x *AdminUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AdminUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdminUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AdminUserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AdminUserID) Reset() {
	*x = AdminUserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_admin_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminUserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminUserID) ProtoMessage() {}

func (x *AdminUserID) ProtoReflect() protoreflect.Message {
	mi := &file_entity_admin_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminUserID.ProtoReflect.Descriptor instead.
func (*AdminUserID) Descriptor() ([]byte, []int) {
	return file_entity_admin_user_proto_rawDescGZIP(), []int{1}
}

func (x *AdminUserID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateAdminUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateAdminUserRequest) Reset() {
	*x = CreateAdminUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_admin_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdminUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdminUserRequest) ProtoMessage() {}

func (x *CreateAdminUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entity_admin_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdminUserRequest.ProtoReflect.Descriptor instead.
func (*CreateAdminUserRequest) Descriptor() ([]byte, []int) {
	return file_entity_admin_user_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAdminUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateAdminUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UpdateAdminUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateAdminUserRequest) Reset() {
	*x = UpdateAdminUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_admin_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdminUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminUserRequest) ProtoMessage() {}

func (x *UpdateAdminUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entity_admin_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdminUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateAdminUserRequest) Descriptor() ([]byte, []int) {
	return file_entity_admin_user_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAdminUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateAdminUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AccessToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AccessToken) Reset() {
	*x = AccessToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_admin_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessToken) ProtoMessage() {}

func (x *AccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_entity_admin_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessToken.ProtoReflect.Descriptor instead.
func (*AccessToken) Descriptor() ([]byte, []int) {
	return file_entity_admin_user_proto_rawDescGZIP(), []int{4}
}

func (x *AccessToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AdminUserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*AdminUser `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *AdminUserList) Reset() {
	*x = AdminUserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_admin_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminUserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminUserList) ProtoMessage() {}

func (x *AdminUserList) ProtoReflect() protoreflect.Message {
	mi := &file_entity_admin_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminUserList.ProtoReflect.Descriptor instead.
func (*AdminUserList) Descriptor() ([]byte, []int) {
	return file_entity_admin_user_proto_rawDescGZIP(), []int{5}
}

func (x *AdminUserList) GetItems() []*AdminUser {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_entity_admin_user_proto protoreflect.FileDescriptor

var file_entity_admin_user_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x22, 0x45, 0x0a, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x1d, 0x0a, 0x0b, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x3c, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x23, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x38, 0x0a, 0x0d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x42, 0x04, 0x5a, 0x02, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_admin_user_proto_rawDescOnce sync.Once
	file_entity_admin_user_proto_rawDescData = file_entity_admin_user_proto_rawDesc
)

func file_entity_admin_user_proto_rawDescGZIP() []byte {
	file_entity_admin_user_proto_rawDescOnce.Do(func() {
		file_entity_admin_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_admin_user_proto_rawDescData)
	})
	return file_entity_admin_user_proto_rawDescData
}

var file_entity_admin_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_entity_admin_user_proto_goTypes = []interface{}{
	(*AdminUser)(nil),              // 0: entity.AdminUser
	(*AdminUserID)(nil),            // 1: entity.AdminUserID
	(*CreateAdminUserRequest)(nil), // 2: entity.CreateAdminUserRequest
	(*UpdateAdminUserRequest)(nil), // 3: entity.UpdateAdminUserRequest
	(*AccessToken)(nil),            // 4: entity.AccessToken
	(*AdminUserList)(nil),          // 5: entity.AdminUserList
}
var file_entity_admin_user_proto_depIdxs = []int32{
	0, // 0: entity.AdminUserList.items:type_name -> entity.AdminUser
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_entity_admin_user_proto_init() }
func file_entity_admin_user_proto_init() {
	if File_entity_admin_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entity_admin_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminUser); i {
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
		file_entity_admin_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminUserID); i {
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
		file_entity_admin_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAdminUserRequest); i {
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
		file_entity_admin_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAdminUserRequest); i {
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
		file_entity_admin_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessToken); i {
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
		file_entity_admin_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminUserList); i {
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
			RawDescriptor: file_entity_admin_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_admin_user_proto_goTypes,
		DependencyIndexes: file_entity_admin_user_proto_depIdxs,
		MessageInfos:      file_entity_admin_user_proto_msgTypes,
	}.Build()
	File_entity_admin_user_proto = out.File
	file_entity_admin_user_proto_rawDesc = nil
	file_entity_admin_user_proto_goTypes = nil
	file_entity_admin_user_proto_depIdxs = nil
}
