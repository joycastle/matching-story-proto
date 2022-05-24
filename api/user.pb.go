// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: user.proto

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

// 玩家基本信息
type UserBaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`              // ID
	Level          int32  `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`        // 等级
	Name           string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`           // 姓名
	HeadIcon       string `protobuf:"bytes,4,opt,name=headIcon,proto3" json:"headIcon,omitempty"`   // 头像
	Country        string `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`     // 国家
	GuildID        int64  `protobuf:"varint,6,opt,name=guildID,proto3" json:"guildID,omitempty"`    // 公会ID
	Badge          string `protobuf:"bytes,7,opt,name=Badge,proto3" json:"Badge,omitempty"`         //　公会徽章
	GuildName      string `protobuf:"bytes,8,opt,name=guildName,proto3" json:"guildName,omitempty"` // 公会名字
	Like           int32  `protobuf:"varint,9,opt,name=like,proto3" json:"like,omitempty"`          // 玩家点赞
	Channel        int32  `protobuf:"varint,10,opt,name=channel,proto3" json:"channel,omitempty"`   //渠道
	Star           int32  `protobuf:"varint,11,opt,name=star,proto3" json:"star,omitempty"`         // 星星 满级玩法，星星个数，可以理解为level2
	StarActivityID int64  `protobuf:"varint,12,opt,name=starActivityID,proto3" json:"starActivityID,omitempty"`
}

func (x *UserBaseInfo) Reset() {
	*x = UserBaseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBaseInfo) ProtoMessage() {}

func (x *UserBaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBaseInfo.ProtoReflect.Descriptor instead.
func (*UserBaseInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserBaseInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserBaseInfo) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *UserBaseInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserBaseInfo) GetHeadIcon() string {
	if x != nil {
		return x.HeadIcon
	}
	return ""
}

func (x *UserBaseInfo) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *UserBaseInfo) GetGuildID() int64 {
	if x != nil {
		return x.GuildID
	}
	return 0
}

func (x *UserBaseInfo) GetBadge() string {
	if x != nil {
		return x.Badge
	}
	return ""
}

func (x *UserBaseInfo) GetGuildName() string {
	if x != nil {
		return x.GuildName
	}
	return ""
}

func (x *UserBaseInfo) GetLike() int32 {
	if x != nil {
		return x.Like
	}
	return 0
}

func (x *UserBaseInfo) GetChannel() int32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

func (x *UserBaseInfo) GetStar() int32 {
	if x != nil {
		return x.Star
	}
	return 0
}

func (x *UserBaseInfo) GetStarActivityID() int64 {
	if x != nil {
		return x.StarActivityID
	}
	return 0
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x02, 0x0a,
	0x0c, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x49,
	0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64, 0x49,
	0x63, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x67, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x61, 0x64, 0x67, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x42, 0x61, 0x64, 0x67, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x69, 0x6b, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x61,
	0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x74, 0x61, 0x72, 0x12, 0x26, 0x0a,
	0x0e, 0x73, 0x74, 0x61, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x49, 0x44, 0x42, 0x17, 0x48, 0x01, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x61, 0x70,
	0x69, 0xaa, 0x02, 0x0a, 0x4d, 0x4d, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_proto_goTypes = []interface{}{
	(*UserBaseInfo)(nil), // 0: UserBaseInfo
}
var file_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBaseInfo); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
