// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/users/users.proto

package users

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 手机号
	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	// 手机验证码
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_users_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_users_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_proto_users_users_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *LoginReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 手机号
	Mobile string `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *GetCodeReq) Reset() {
	*x = GetCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_users_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCodeReq) ProtoMessage() {}

func (x *GetCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_users_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCodeReq.ProtoReflect.Descriptor instead.
func (*GetCodeReq) Descriptor() ([]byte, []int) {
	return file_proto_users_users_proto_rawDescGZIP(), []int{1}
}

func (x *GetCodeReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type Rep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Rep) Reset() {
	*x = Rep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_users_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rep) ProtoMessage() {}

func (x *Rep) ProtoReflect() protoreflect.Message {
	mi := &file_proto_users_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rep.ProtoReflect.Descriptor instead.
func (*Rep) Descriptor() ([]byte, []int) {
	return file_proto_users_users_proto_rawDescGZIP(), []int{2}
}

func (x *Rep) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Rep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_proto_users_users_proto protoreflect.FileDescriptor

var file_proto_users_users_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x12, 0x43, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xfa, 0x42, 0x13, 0x72, 0x11, 0x32, 0x0f, 0x28, 0x5e, 0x24,
	0x7c, 0x5e, 0x5b, 0x5c, 0x64, 0x5d, 0x7b, 0x31, 0x31, 0x7d, 0x24, 0x29, 0x92, 0x41, 0x12, 0x8a,
	0x01, 0x0f, 0x28, 0x5e, 0x24, 0x7c, 0x5e, 0x5b, 0x5c, 0x64, 0x5d, 0x7b, 0x31, 0x31, 0x7d, 0x24,
	0x29, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x51, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x43, 0x0a, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xfa, 0x42, 0x13,
	0x72, 0x11, 0x32, 0x0f, 0x28, 0x5e, 0x24, 0x7c, 0x5e, 0x5b, 0x5c, 0x64, 0x5d, 0x7b, 0x31, 0x31,
	0x7d, 0x24, 0x29, 0x92, 0x41, 0x12, 0x8a, 0x01, 0x0f, 0x28, 0x5e, 0x24, 0x7c, 0x5e, 0x5b, 0x5c,
	0x64, 0x5d, 0x7b, 0x31, 0x31, 0x7d, 0x24, 0x29, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x22, 0x35, 0x0a, 0x03, 0x52, 0x65, 0x70, 0x12, 0x1c, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xea, 0xde, 0x1f, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xa9, 0x01, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x4f, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x70, 0x22, 0x1b, 0x92, 0x41, 0x07, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x06, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a,
	0x01, 0x2a, 0x12, 0x4f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x22, 0x17, 0x92, 0x41, 0x07, 0x0a,
	0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x63,
	0x6f, 0x64, 0x65, 0x42, 0xe5, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x44, 0x6f, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x92, 0x41, 0xad, 0x01, 0x12, 0x44,
	0x0a, 0x17, 0x43, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x20, 0x77, 0x69,
	0x74, 0x68, 0x20, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x22, 0x21, 0x0a, 0x0c, 0xe6, 0xb1, 0x9f,
	0xe6, 0xb9, 0x96, 0xe5, 0xa4, 0xa7, 0xe7, 0x89, 0x9b, 0x1a, 0x11, 0x6e, 0x66, 0x61, 0x6e, 0x67,
	0x78, 0x75, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x06, 0x76, 0x31,
	0x2e, 0x30, 0x2e, 0x30, 0x22, 0x06, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x72, 0x39, 0x0a, 0x0b, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x12,
	0x2a, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6f, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_users_users_proto_rawDescOnce sync.Once
	file_proto_users_users_proto_rawDescData = file_proto_users_users_proto_rawDesc
)

func file_proto_users_users_proto_rawDescGZIP() []byte {
	file_proto_users_users_proto_rawDescOnce.Do(func() {
		file_proto_users_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_users_users_proto_rawDescData)
	})
	return file_proto_users_users_proto_rawDescData
}

var file_proto_users_users_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_users_users_proto_goTypes = []interface{}{
	(*LoginReq)(nil),   // 0: api.users.v1.LoginReq
	(*GetCodeReq)(nil), // 1: api.users.v1.GetCodeReq
	(*Rep)(nil),        // 2: api.users.v1.Rep
}
var file_proto_users_users_proto_depIdxs = []int32{
	0, // 0: api.users.v1.Users.Login:input_type -> api.users.v1.LoginReq
	1, // 1: api.users.v1.Users.GetCode:input_type -> api.users.v1.GetCodeReq
	2, // 2: api.users.v1.Users.Login:output_type -> api.users.v1.Rep
	2, // 3: api.users.v1.Users.GetCode:output_type -> api.users.v1.Rep
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_users_users_proto_init() }
func file_proto_users_users_proto_init() {
	if File_proto_users_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_users_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_proto_users_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCodeReq); i {
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
		file_proto_users_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rep); i {
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
			RawDescriptor: file_proto_users_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_users_users_proto_goTypes,
		DependencyIndexes: file_proto_users_users_proto_depIdxs,
		MessageInfos:      file_proto_users_users_proto_msgTypes,
	}.Build()
	File_proto_users_users_proto = out.File
	file_proto_users_users_proto_rawDesc = nil
	file_proto_users_users_proto_goTypes = nil
	file_proto_users_users_proto_depIdxs = nil
}