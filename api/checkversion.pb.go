// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: checkversion.proto

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

type DeviceType int32

const (
	DeviceType_Android DeviceType = 0
	DeviceType_iOS     DeviceType = 1
	DeviceType_Windows DeviceType = 2
	DeviceType_OSX     DeviceType = 3
)

// Enum value maps for DeviceType.
var (
	DeviceType_name = map[int32]string{
		0: "Android",
		1: "iOS",
		2: "Windows",
		3: "OSX",
	}
	DeviceType_value = map[string]int32{
		"Android": 0,
		"iOS":     1,
		"Windows": 2,
		"OSX":     3,
	}
)

func (x DeviceType) Enum() *DeviceType {
	p := new(DeviceType)
	*p = x
	return p
}

func (x DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_checkversion_proto_enumTypes[0].Descriptor()
}

func (DeviceType) Type() protoreflect.EnumType {
	return &file_checkversion_proto_enumTypes[0]
}

func (x DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceType.Descriptor instead.
func (DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_checkversion_proto_rawDescGZIP(), []int{0}
}

//更新类型
type UpdateType int32

const (
	//普通--可以取消
	UpdateType_Normal UpdateType = 0
	//强制--不可取消
	UpdateType_Force UpdateType = 1
	//静默--不弹框直接下载
	UpdateType_Silence UpdateType = 2
)

// Enum value maps for UpdateType.
var (
	UpdateType_name = map[int32]string{
		0: "Normal",
		1: "Force",
		2: "Silence",
	}
	UpdateType_value = map[string]int32{
		"Normal":  0,
		"Force":   1,
		"Silence": 2,
	}
)

func (x UpdateType) Enum() *UpdateType {
	p := new(UpdateType)
	*p = x
	return p
}

func (x UpdateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateType) Descriptor() protoreflect.EnumDescriptor {
	return file_checkversion_proto_enumTypes[1].Descriptor()
}

func (UpdateType) Type() protoreflect.EnumType {
	return &file_checkversion_proto_enumTypes[1]
}

func (x UpdateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateType.Descriptor instead.
func (UpdateType) EnumDescriptor() ([]byte, []int) {
	return file_checkversion_proto_rawDescGZIP(), []int{1}
}

type CSCheckVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//本地版本号
	LocalVersion int32 `protobuf:"varint,1,opt,name=localVersion,proto3" json:"localVersion,omitempty"`
	//本地设备类型(Windows,Android,iOS,OSX)
	DeviceType DeviceType `protobuf:"varint,2,opt,name=deviceType,proto3,enum=DeviceType" json:"deviceType,omitempty"`
	//本地时间
	LocalTime int64 `protobuf:"varint,3,opt,name=localTime,proto3" json:"localTime,omitempty"`
}

func (x *CSCheckVersion) Reset() {
	*x = CSCheckVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkversion_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSCheckVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSCheckVersion) ProtoMessage() {}

func (x *CSCheckVersion) ProtoReflect() protoreflect.Message {
	mi := &file_checkversion_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSCheckVersion.ProtoReflect.Descriptor instead.
func (*CSCheckVersion) Descriptor() ([]byte, []int) {
	return file_checkversion_proto_rawDescGZIP(), []int{0}
}

func (x *CSCheckVersion) GetLocalVersion() int32 {
	if x != nil {
		return x.LocalVersion
	}
	return 0
}

func (x *CSCheckVersion) GetDeviceType() DeviceType {
	if x != nil {
		return x.DeviceType
	}
	return DeviceType_Android
}

func (x *CSCheckVersion) GetLocalTime() int64 {
	if x != nil {
		return x.LocalTime
	}
	return 0
}

//更新内容
type ContentData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//语言类型
	LanguageType int32 `protobuf:"varint,1,opt,name=languageType,proto3" json:"languageType,omitempty"`
	//标题
	UpdateTittle string `protobuf:"bytes,2,opt,name=updateTittle,proto3" json:"updateTittle,omitempty"`
	//内容
	UpdateString string `protobuf:"bytes,3,opt,name=updateString,proto3" json:"updateString,omitempty"`
}

func (x *ContentData) Reset() {
	*x = ContentData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkversion_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentData) ProtoMessage() {}

func (x *ContentData) ProtoReflect() protoreflect.Message {
	mi := &file_checkversion_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentData.ProtoReflect.Descriptor instead.
func (*ContentData) Descriptor() ([]byte, []int) {
	return file_checkversion_proto_rawDescGZIP(), []int{1}
}

func (x *ContentData) GetLanguageType() int32 {
	if x != nil {
		return x.LanguageType
	}
	return 0
}

func (x *ContentData) GetUpdateTittle() string {
	if x != nil {
		return x.UpdateTittle
	}
	return ""
}

func (x *ContentData) GetUpdateString() string {
	if x != nil {
		return x.UpdateString
	}
	return ""
}

type UpdateContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//更新的版本号
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	//更新类型
	UpdateType UpdateType `protobuf:"varint,2,opt,name=updateType,proto3,enum=UpdateType" json:"updateType,omitempty"`
	//更新内容
	Data []*ContentData `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	//更新发布时间
	Time int64 `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *UpdateContent) Reset() {
	*x = UpdateContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkversion_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateContent) ProtoMessage() {}

func (x *UpdateContent) ProtoReflect() protoreflect.Message {
	mi := &file_checkversion_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateContent.ProtoReflect.Descriptor instead.
func (*UpdateContent) Descriptor() ([]byte, []int) {
	return file_checkversion_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateContent) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *UpdateContent) GetUpdateType() UpdateType {
	if x != nil {
		return x.UpdateType
	}
	return UpdateType_Normal
}

func (x *UpdateContent) GetData() []*ContentData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UpdateContent) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type SCCheckVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//当前最新版本号
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	//当前更新的信息
	UpdateContent *UpdateContent `protobuf:"bytes,2,opt,name=updateContent,proto3" json:"updateContent,omitempty"`
}

func (x *SCCheckVersion) Reset() {
	*x = SCCheckVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkversion_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SCCheckVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCCheckVersion) ProtoMessage() {}

func (x *SCCheckVersion) ProtoReflect() protoreflect.Message {
	mi := &file_checkversion_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCCheckVersion.ProtoReflect.Descriptor instead.
func (*SCCheckVersion) Descriptor() ([]byte, []int) {
	return file_checkversion_proto_rawDescGZIP(), []int{3}
}

func (x *SCCheckVersion) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *SCCheckVersion) GetUpdateContent() *UpdateContent {
	if x != nil {
		return x.UpdateContent
	}
	return nil
}

//请求ABTest
type CSABTestConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//当前版本号
	LocalVersion int32 `protobuf:"varint,1,opt,name=localVersion,proto3" json:"localVersion,omitempty"`
	//设备ID
	DeviceID string `protobuf:"bytes,2,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	//请求的AB测试类型列表
	AbType []int32 `protobuf:"varint,3,rep,packed,name=abType,proto3" json:"abType,omitempty"`
}

func (x *CSABTestConfig) Reset() {
	*x = CSABTestConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkversion_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSABTestConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSABTestConfig) ProtoMessage() {}

func (x *CSABTestConfig) ProtoReflect() protoreflect.Message {
	mi := &file_checkversion_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSABTestConfig.ProtoReflect.Descriptor instead.
func (*CSABTestConfig) Descriptor() ([]byte, []int) {
	return file_checkversion_proto_rawDescGZIP(), []int{4}
}

func (x *CSABTestConfig) GetLocalVersion() int32 {
	if x != nil {
		return x.LocalVersion
	}
	return 0
}

func (x *CSABTestConfig) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *CSABTestConfig) GetAbType() []int32 {
	if x != nil {
		return x.AbType
	}
	return nil
}

//ABTest组
type SCABTestConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//Server端回复的ABTest组
	AbConfig []*ABConfig `protobuf:"bytes,1,rep,name=abConfig,proto3" json:"abConfig,omitempty"`
}

func (x *SCABTestConfig) Reset() {
	*x = SCABTestConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkversion_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SCABTestConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SCABTestConfig) ProtoMessage() {}

func (x *SCABTestConfig) ProtoReflect() protoreflect.Message {
	mi := &file_checkversion_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SCABTestConfig.ProtoReflect.Descriptor instead.
func (*SCABTestConfig) Descriptor() ([]byte, []int) {
	return file_checkversion_proto_rawDescGZIP(), []int{5}
}

func (x *SCABTestConfig) GetAbConfig() []*ABConfig {
	if x != nil {
		return x.AbConfig
	}
	return nil
}

var File_checkversion_proto protoreflect.FileDescriptor

var file_checkversion_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x0e, 0x43, 0x53, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x79, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x74, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x8c,
	0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x60, 0x0a,
	0x0e, 0x53, 0x43, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x0d, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x68, 0x0a, 0x0e, 0x43, 0x53, 0x41, 0x42, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x62, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x06, 0x61, 0x62, 0x54, 0x79, 0x70, 0x65, 0x22, 0x37, 0x0a, 0x0e, 0x53, 0x43, 0x41,
	0x42, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x08, 0x61,
	0x62, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x41, 0x42, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x61, 0x62, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2a, 0x38, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x69, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x73, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x53, 0x58, 0x10, 0x03, 0x2a, 0x30, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x6f,
	0x72, 0x6d, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x10, 0x02, 0x42, 0x17,
	0x48, 0x01, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x61, 0x70, 0x69, 0xaa, 0x02, 0x0a, 0x4d, 0x4d, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_checkversion_proto_rawDescOnce sync.Once
	file_checkversion_proto_rawDescData = file_checkversion_proto_rawDesc
)

func file_checkversion_proto_rawDescGZIP() []byte {
	file_checkversion_proto_rawDescOnce.Do(func() {
		file_checkversion_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkversion_proto_rawDescData)
	})
	return file_checkversion_proto_rawDescData
}

var file_checkversion_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_checkversion_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_checkversion_proto_goTypes = []interface{}{
	(DeviceType)(0),        // 0: DeviceType
	(UpdateType)(0),        // 1: UpdateType
	(*CSCheckVersion)(nil), // 2: CSCheckVersion
	(*ContentData)(nil),    // 3: ContentData
	(*UpdateContent)(nil),  // 4: UpdateContent
	(*SCCheckVersion)(nil), // 5: SCCheckVersion
	(*CSABTestConfig)(nil), // 6: CSABTestConfig
	(*SCABTestConfig)(nil), // 7: SCABTestConfig
	(*ABConfig)(nil),       // 8: ABConfig
}
var file_checkversion_proto_depIdxs = []int32{
	0, // 0: CSCheckVersion.deviceType:type_name -> DeviceType
	1, // 1: UpdateContent.updateType:type_name -> UpdateType
	3, // 2: UpdateContent.data:type_name -> ContentData
	4, // 3: SCCheckVersion.updateContent:type_name -> UpdateContent
	8, // 4: SCABTestConfig.abConfig:type_name -> ABConfig
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_checkversion_proto_init() }
func file_checkversion_proto_init() {
	if File_checkversion_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_checkversion_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSCheckVersion); i {
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
		file_checkversion_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentData); i {
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
		file_checkversion_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateContent); i {
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
		file_checkversion_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SCCheckVersion); i {
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
		file_checkversion_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSABTestConfig); i {
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
		file_checkversion_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SCABTestConfig); i {
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
			RawDescriptor: file_checkversion_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_checkversion_proto_goTypes,
		DependencyIndexes: file_checkversion_proto_depIdxs,
		EnumInfos:         file_checkversion_proto_enumTypes,
		MessageInfos:      file_checkversion_proto_msgTypes,
	}.Build()
	File_checkversion_proto = out.File
	file_checkversion_proto_rawDesc = nil
	file_checkversion_proto_goTypes = nil
	file_checkversion_proto_depIdxs = nil
}
