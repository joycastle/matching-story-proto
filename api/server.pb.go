// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: server.proto

package api

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

type S2CBackCode int32

const (
	S2CBackCode_DEFAULT S2CBackCode = 0
	//服务端处理错误
	S2CBackCode_SERVER_ERROR S2CBackCode = 1000
	//帐号被封
	S2CBackCode_ACCOUNT_LOCK S2CBackCode = 1001
	//帐号被删除
	S2CBackCode_ACCOUNT_DEL S2CBackCode = 1002
	//版本校验失败
	S2CBackCode_VERSION_CODE_ERROR S2CBackCode = 1003
	//校验码错误
	S2CBackCode_ENCRYPT_CODE_ERROR S2CBackCode = 1004
	//登录校验错误
	S2CBackCode_LOGIN_CHECK_ERROR S2CBackCode = 1005
	//保存数据金币校验错误
	S2CBackCode_SAVE_SILVER_CHECK_ERROR S2CBackCode = 1016
	//保存数据爱心校验错误
	S2CBackCode_SAVE_ENERGY_CHECK_ERROR S2CBackCode = 1007
	//保存数据失败
	S2CBackCode_SAVE_DATA_ERROR   S2CBackCode = 1008
	S2CBackCode_SERVER_INFO_ERROR S2CBackCode = 1050
	//好友已满
	S2CBackCode_FRIEND_HAS_FULL S2CBackCode = 1051
	//充值失败
	S2CBackCode_RECHARGE_FAILED S2CBackCode = 1052
	//没有订单信息
	S2CBackCode_NO_RECHARGE_ID S2CBackCode = 1053
	//订单已完成
	S2CBackCode_RECHARGE_HAS_DONE S2CBackCode = 1054
	//绑定账号失败
	S2CBackCode_ACCOUNT_BIND_ERROR S2CBackCode = 1055
	//成功
	S2CBackCode_SUCCESS S2CBackCode = 2000
)

// Enum value maps for S2CBackCode.
var (
	S2CBackCode_name = map[int32]string{
		0:    "DEFAULT",
		1000: "SERVER_ERROR",
		1001: "ACCOUNT_LOCK",
		1002: "ACCOUNT_DEL",
		1003: "VERSION_CODE_ERROR",
		1004: "ENCRYPT_CODE_ERROR",
		1005: "LOGIN_CHECK_ERROR",
		1016: "SAVE_SILVER_CHECK_ERROR",
		1007: "SAVE_ENERGY_CHECK_ERROR",
		1008: "SAVE_DATA_ERROR",
		1050: "SERVER_INFO_ERROR",
		1051: "FRIEND_HAS_FULL",
		1052: "RECHARGE_FAILED",
		1053: "NO_RECHARGE_ID",
		1054: "RECHARGE_HAS_DONE",
		1055: "ACCOUNT_BIND_ERROR",
		2000: "SUCCESS",
	}
	S2CBackCode_value = map[string]int32{
		"DEFAULT":                 0,
		"SERVER_ERROR":            1000,
		"ACCOUNT_LOCK":            1001,
		"ACCOUNT_DEL":             1002,
		"VERSION_CODE_ERROR":      1003,
		"ENCRYPT_CODE_ERROR":      1004,
		"LOGIN_CHECK_ERROR":       1005,
		"SAVE_SILVER_CHECK_ERROR": 1016,
		"SAVE_ENERGY_CHECK_ERROR": 1007,
		"SAVE_DATA_ERROR":         1008,
		"SERVER_INFO_ERROR":       1050,
		"FRIEND_HAS_FULL":         1051,
		"RECHARGE_FAILED":         1052,
		"NO_RECHARGE_ID":          1053,
		"RECHARGE_HAS_DONE":       1054,
		"ACCOUNT_BIND_ERROR":      1055,
		"SUCCESS":                 2000,
	}
)

func (x S2CBackCode) Enum() *S2CBackCode {
	p := new(S2CBackCode)
	*p = x
	return p
}

func (x S2CBackCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (S2CBackCode) Descriptor() protoreflect.EnumDescriptor {
	return file_server_proto_enumTypes[0].Descriptor()
}

func (S2CBackCode) Type() protoreflect.EnumType {
	return &file_server_proto_enumTypes[0]
}

func (x S2CBackCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use S2CBackCode.Descriptor instead.
func (S2CBackCode) EnumDescriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

type SCErrorMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SCErrorMsg) Reset() {
	*x = SCErrorMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SCErrorMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCErrorMsg) ProtoMessage() {}

func (x *SCErrorMsg) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCErrorMsg.ProtoReflect.Descriptor instead.
func (*SCErrorMsg) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0c,
	0x0a, 0x0a, 0x53, 0x43, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x2a, 0x86, 0x03, 0x0a,
	0x0b, 0x53, 0x32, 0x43, 0x42, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0c, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xe8, 0x07, 0x12, 0x11, 0x0a, 0x0c,
	0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x10, 0xe9, 0x07, 0x12,
	0x10, 0x0a, 0x0b, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x10, 0xea,
	0x07, 0x12, 0x17, 0x0a, 0x12, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xeb, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x45, 0x4e,
	0x43, 0x52, 0x59, 0x50, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0xec, 0x07, 0x12, 0x16, 0x0a, 0x11, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x43, 0x48, 0x45,
	0x43, 0x4b, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xed, 0x07, 0x12, 0x1c, 0x0a, 0x17, 0x53,
	0x41, 0x56, 0x45, 0x5f, 0x53, 0x49, 0x4c, 0x56, 0x45, 0x52, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xf8, 0x07, 0x12, 0x1c, 0x0a, 0x17, 0x53, 0x41, 0x56,
	0x45, 0x5f, 0x45, 0x4e, 0x45, 0x52, 0x47, 0x59, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0xef, 0x07, 0x12, 0x14, 0x0a, 0x0f, 0x53, 0x41, 0x56, 0x45, 0x5f,
	0x44, 0x41, 0x54, 0x41, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xf0, 0x07, 0x12, 0x16, 0x0a,
	0x11, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x9a, 0x08, 0x12, 0x14, 0x0a, 0x0f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x5f,
	0x48, 0x41, 0x53, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x9b, 0x08, 0x12, 0x14, 0x0a, 0x0f, 0x52,
	0x45, 0x43, 0x48, 0x41, 0x52, 0x47, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x9c,
	0x08, 0x12, 0x13, 0x0a, 0x0e, 0x4e, 0x4f, 0x5f, 0x52, 0x45, 0x43, 0x48, 0x41, 0x52, 0x47, 0x45,
	0x5f, 0x49, 0x44, 0x10, 0x9d, 0x08, 0x12, 0x16, 0x0a, 0x11, 0x52, 0x45, 0x43, 0x48, 0x41, 0x52,
	0x47, 0x45, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x9e, 0x08, 0x12, 0x17,
	0x0a, 0x12, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x42, 0x49, 0x4e, 0x44, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x9f, 0x08, 0x12, 0x0c, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45,
	0x53, 0x53, 0x10, 0xd0, 0x0f, 0x42, 0x17, 0x48, 0x01, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x61, 0x70,
	0x69, 0xaa, 0x02, 0x0a, 0x4d, 0x4d, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_server_proto_goTypes = []interface{}{
	(S2CBackCode)(0),   // 0: S2CBackCode
	(*SCErrorMsg)(nil), // 1: SCErrorMsg
}
var file_server_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SCErrorMsg); i {
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
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		EnumInfos:         file_server_proto_enumTypes,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}
