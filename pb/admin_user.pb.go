// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: admin_user.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_admin_user_proto protoreflect.FileDescriptor

var file_admin_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x95,
	0x02, 0x0a, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x12, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2e,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x13, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x04, 0x5a, 0x02, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_admin_user_proto_goTypes = []interface{}{
	(*Empty)(nil),                  // 0: common.Empty
	(*CreateAdminUserRequest)(nil), // 1: entity.CreateAdminUserRequest
	(*UpdateAdminUserRequest)(nil), // 2: entity.UpdateAdminUserRequest
	(*AdminUserID)(nil),            // 3: entity.AdminUserID
	(*AdminUser)(nil),              // 4: entity.AdminUser
	(*AdminUserList)(nil),          // 5: entity.AdminUserList
}
var file_admin_user_proto_depIdxs = []int32{
	0, // 0: api.AdminUserService.GetMe:input_type -> common.Empty
	0, // 1: api.AdminUserService.GetAll:input_type -> common.Empty
	1, // 2: api.AdminUserService.Create:input_type -> entity.CreateAdminUserRequest
	2, // 3: api.AdminUserService.Update:input_type -> entity.UpdateAdminUserRequest
	3, // 4: api.AdminUserService.Delete:input_type -> entity.AdminUserID
	4, // 5: api.AdminUserService.GetMe:output_type -> entity.AdminUser
	5, // 6: api.AdminUserService.GetAll:output_type -> entity.AdminUserList
	4, // 7: api.AdminUserService.Create:output_type -> entity.AdminUser
	4, // 8: api.AdminUserService.Update:output_type -> entity.AdminUser
	0, // 9: api.AdminUserService.Delete:output_type -> common.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_user_proto_init() }
func file_admin_user_proto_init() {
	if File_admin_user_proto != nil {
		return
	}
	file_common_proto_init()
	file_entity_admin_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_user_proto_goTypes,
		DependencyIndexes: file_admin_user_proto_depIdxs,
	}.Build()
	File_admin_user_proto = out.File
	file_admin_user_proto_rawDesc = nil
	file_admin_user_proto_goTypes = nil
	file_admin_user_proto_depIdxs = nil
}
