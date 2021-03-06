// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: code.proto

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

type PANGU_CODE int32

const (
	PANGU_CODE_PANGU_CODE_UNSPECIFIED                      PANGU_CODE = 0    // 无效code
	PANGU_CODE_PANGU_CODE_OK                               PANGU_CODE = 2000 // 成功
	PANGU_CODE_PANGU_CODE_NOT_LOGGED_IN                    PANGU_CODE = 4010 // 未登录
	PANGU_CODE_PANGU_CODE_UNAUTHORIZED                     PANGU_CODE = 4011 // 未授权
	PANGU_CODE_PANGU_CODE_PARAMETER_ILLEGAL                PANGU_CODE = 4030 // 参数不合法
	PANGU_CODE_PANGU_CODE_UNTRUSTED_USER                   PANGU_CODE = 4031 // 非法的用户Id
	PANGU_CODE_PANGU_CODE_ROUTING_NOT_EXIST                PANGU_CODE = 4040 // 路由不存在
	PANGU_CODE_PANGU_CODE_REQUEST_TIMEOUT                  PANGU_CODE = 4080 // 服务器等待客户端发送的请求超时
	PANGU_CODE_PANGU_CODE_SERVER_ERR                       PANGU_CODE = 5000 // 系统错误
	PANGU_CODE_PANGU_CODE_HTTP_UNPACK_ERR                  PANGU_CODE = 6000
	PANGU_CODE_PANGU_CODE_HTTP_PACK_ERR                    PANGU_CODE = 6001
	PANGU_CODE_PANGU_CODE_REDIS_ERR                        PANGU_CODE = 6002
	PANGU_CODE_PANGU_CODE_NO_DATA                          PANGU_CODE = 6101 // 没有数据
	PANGU_CODE_PANGU_CODE_MODEL_ADD_ERR                    PANGU_CODE = 6102 // 添加错误
	PANGU_CODE_PANGU_CODE_MODEL_DEL_ERR                    PANGU_CODE = 6103 // 删除错误
	PANGU_CODE_PANGU_CODE_MODEL_STORE_ERR                  PANGU_CODE = 6104 // 存储错误
	PANGU_CODE_PANGU_CODE_PLAYER_IN_GUILD                  PANGU_CODE = 7000 // 玩家在公会中
	PANGU_CODE_PANGU_CODE_GUILD_NOT_EXIST                  PANGU_CODE = 7001 // 公会不存在
	PANGU_CODE_PANGU_CODE_PLAYER_NO_GUILD                  PANGU_CODE = 7002 // 玩家不在公会中
	PANGU_CODE_PANGU_CODE_PLAYER_NOT_EXIST                 PANGU_CODE = 7003 // 用户不存在
	PANGU_CODE_PANGU_CODE_GUILD_LIMIT_LEVEL                PANGU_CODE = 7004 // 限制等级不能大于会长等级
	PANGU_CODE_PANGU_CODE_GUILD_MAX_MEMBER_NUM             PANGU_CODE = 7005 // 公会人数已满
	PANGU_CODE_PANGU_CODE_PLAYER_NOT_LEADER                PANGU_CODE = 7006 // 不是会长
	PANGU_CODE_PANGU_CODE_GUILD_NOT_KICK_SELF              PANGU_CODE = 7007 // 不能删除自己
	PANGU_CODE_PANGU_CODE_PLAYER_IS_CHEAT_USER             PANGU_CODE = 7008 // 作弊玩家
	PANGU_CODE_PANGU_CODE_REQUEST_INTERVAL_TOO_SHORT       PANGU_CODE = 7100 // 请求资源间隔过短(时间间隔可配置->app.guild.helpInterval)
	PANGU_CODE_PANGU_CODE_HELP_IS_SATISFIED                PANGU_CODE = 7101 // 该请求帮助已被满足，无法帮助(该体力请求已完成)
	PANGU_CODE_PANGU_CODE_HELP_IS_CONSUMED                 PANGU_CODE = 7105 // 该资源已被消费(体力已经被领取，无法再次领取)
	PANGU_CODE_PANGU_CODE_ACTIVITY_ID_ERROR                PANGU_CODE = 8000 // 活动ID错误
	PANGU_CODE_PANGU_CODE_ACTIVITY_PLAYER_NO_GROUP         PANGU_CODE = 8001 // 玩家不在组中
	PANGU_CODE_PANGU_CODE_ACTIVITY_NO_OPEN                 PANGU_CODE = 8002 // 活动没开启
	PANGU_CODE_PANGU_CODE_ACTIVITY_RECEIVE_AWARD           PANGU_CODE = 8003 // 活动已领取奖励
	PANGU_CODE_PANGU_CODE_PANGU_CODE_JOIN_ACTIVITY_TIMEOUT PANGU_CODE = 8004 // 超时不能加入活动
	PANGU_CODE_PANGU_CODE_ACTIVITY_PLAYER_IN_GROUP         PANGU_CODE = 8005 // 玩家不在组中
	PANGU_CODE_PANGU_CODE_ACTIVITY_STOP_ENTER              PANGU_CODE = 8006 // 停止进入
	PANGU_CODE_PANGU_CODE_ACTIVITY_REWARD_NOT_START        PANGU_CODE = 8007 // 未开始领奖
	PANGU_CODE_PANGU_CODE_PIGEON_GRPC_ERR                  PANGU_CODE = 10000
	PANGU_CODE_PANGU_CODE_PIGEON_MISS_IGNORE_MYSELF        PANGU_CODE = 10001 // ignoreMyself缺未提供userId
	PANGU_CODE_PANGU_CODE_PIGEON_GRPC_SEND_ERR             PANGU_CODE = 10002
	PANGU_CODE_PANGU_CODE_MADEL_RUSH                       PANGU_CODE = 11031 // 服务被删除
	PANGU_CODE_PANGU_CODE_MADEL_RUSH_JOIN_END              PANGU_CODE = 11032 // 加入已关闭
)

// Enum value maps for PANGU_CODE.
var (
	PANGU_CODE_name = map[int32]string{
		0:     "PANGU_CODE_UNSPECIFIED",
		2000:  "PANGU_CODE_OK",
		4010:  "PANGU_CODE_NOT_LOGGED_IN",
		4011:  "PANGU_CODE_UNAUTHORIZED",
		4030:  "PANGU_CODE_PARAMETER_ILLEGAL",
		4031:  "PANGU_CODE_UNTRUSTED_USER",
		4040:  "PANGU_CODE_ROUTING_NOT_EXIST",
		4080:  "PANGU_CODE_REQUEST_TIMEOUT",
		5000:  "PANGU_CODE_SERVER_ERR",
		6000:  "PANGU_CODE_HTTP_UNPACK_ERR",
		6001:  "PANGU_CODE_HTTP_PACK_ERR",
		6002:  "PANGU_CODE_REDIS_ERR",
		6101:  "PANGU_CODE_NO_DATA",
		6102:  "PANGU_CODE_MODEL_ADD_ERR",
		6103:  "PANGU_CODE_MODEL_DEL_ERR",
		6104:  "PANGU_CODE_MODEL_STORE_ERR",
		7000:  "PANGU_CODE_PLAYER_IN_GUILD",
		7001:  "PANGU_CODE_GUILD_NOT_EXIST",
		7002:  "PANGU_CODE_PLAYER_NO_GUILD",
		7003:  "PANGU_CODE_PLAYER_NOT_EXIST",
		7004:  "PANGU_CODE_GUILD_LIMIT_LEVEL",
		7005:  "PANGU_CODE_GUILD_MAX_MEMBER_NUM",
		7006:  "PANGU_CODE_PLAYER_NOT_LEADER",
		7007:  "PANGU_CODE_GUILD_NOT_KICK_SELF",
		7008:  "PANGU_CODE_PLAYER_IS_CHEAT_USER",
		7100:  "PANGU_CODE_REQUEST_INTERVAL_TOO_SHORT",
		7101:  "PANGU_CODE_HELP_IS_SATISFIED",
		7105:  "PANGU_CODE_HELP_IS_CONSUMED",
		8000:  "PANGU_CODE_ACTIVITY_ID_ERROR",
		8001:  "PANGU_CODE_ACTIVITY_PLAYER_NO_GROUP",
		8002:  "PANGU_CODE_ACTIVITY_NO_OPEN",
		8003:  "PANGU_CODE_ACTIVITY_RECEIVE_AWARD",
		8004:  "PANGU_CODE_PANGU_CODE_JOIN_ACTIVITY_TIMEOUT",
		8005:  "PANGU_CODE_ACTIVITY_PLAYER_IN_GROUP",
		8006:  "PANGU_CODE_ACTIVITY_STOP_ENTER",
		8007:  "PANGU_CODE_ACTIVITY_REWARD_NOT_START",
		10000: "PANGU_CODE_PIGEON_GRPC_ERR",
		10001: "PANGU_CODE_PIGEON_MISS_IGNORE_MYSELF",
		10002: "PANGU_CODE_PIGEON_GRPC_SEND_ERR",
		11031: "PANGU_CODE_MADEL_RUSH",
		11032: "PANGU_CODE_MADEL_RUSH_JOIN_END",
	}
	PANGU_CODE_value = map[string]int32{
		"PANGU_CODE_UNSPECIFIED":                      0,
		"PANGU_CODE_OK":                               2000,
		"PANGU_CODE_NOT_LOGGED_IN":                    4010,
		"PANGU_CODE_UNAUTHORIZED":                     4011,
		"PANGU_CODE_PARAMETER_ILLEGAL":                4030,
		"PANGU_CODE_UNTRUSTED_USER":                   4031,
		"PANGU_CODE_ROUTING_NOT_EXIST":                4040,
		"PANGU_CODE_REQUEST_TIMEOUT":                  4080,
		"PANGU_CODE_SERVER_ERR":                       5000,
		"PANGU_CODE_HTTP_UNPACK_ERR":                  6000,
		"PANGU_CODE_HTTP_PACK_ERR":                    6001,
		"PANGU_CODE_REDIS_ERR":                        6002,
		"PANGU_CODE_NO_DATA":                          6101,
		"PANGU_CODE_MODEL_ADD_ERR":                    6102,
		"PANGU_CODE_MODEL_DEL_ERR":                    6103,
		"PANGU_CODE_MODEL_STORE_ERR":                  6104,
		"PANGU_CODE_PLAYER_IN_GUILD":                  7000,
		"PANGU_CODE_GUILD_NOT_EXIST":                  7001,
		"PANGU_CODE_PLAYER_NO_GUILD":                  7002,
		"PANGU_CODE_PLAYER_NOT_EXIST":                 7003,
		"PANGU_CODE_GUILD_LIMIT_LEVEL":                7004,
		"PANGU_CODE_GUILD_MAX_MEMBER_NUM":             7005,
		"PANGU_CODE_PLAYER_NOT_LEADER":                7006,
		"PANGU_CODE_GUILD_NOT_KICK_SELF":              7007,
		"PANGU_CODE_PLAYER_IS_CHEAT_USER":             7008,
		"PANGU_CODE_REQUEST_INTERVAL_TOO_SHORT":       7100,
		"PANGU_CODE_HELP_IS_SATISFIED":                7101,
		"PANGU_CODE_HELP_IS_CONSUMED":                 7105,
		"PANGU_CODE_ACTIVITY_ID_ERROR":                8000,
		"PANGU_CODE_ACTIVITY_PLAYER_NO_GROUP":         8001,
		"PANGU_CODE_ACTIVITY_NO_OPEN":                 8002,
		"PANGU_CODE_ACTIVITY_RECEIVE_AWARD":           8003,
		"PANGU_CODE_PANGU_CODE_JOIN_ACTIVITY_TIMEOUT": 8004,
		"PANGU_CODE_ACTIVITY_PLAYER_IN_GROUP":         8005,
		"PANGU_CODE_ACTIVITY_STOP_ENTER":              8006,
		"PANGU_CODE_ACTIVITY_REWARD_NOT_START":        8007,
		"PANGU_CODE_PIGEON_GRPC_ERR":                  10000,
		"PANGU_CODE_PIGEON_MISS_IGNORE_MYSELF":        10001,
		"PANGU_CODE_PIGEON_GRPC_SEND_ERR":             10002,
		"PANGU_CODE_MADEL_RUSH":                       11031,
		"PANGU_CODE_MADEL_RUSH_JOIN_END":              11032,
	}
)

func (x PANGU_CODE) Enum() *PANGU_CODE {
	p := new(PANGU_CODE)
	*p = x
	return p
}

func (x PANGU_CODE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PANGU_CODE) Descriptor() protoreflect.EnumDescriptor {
	return file_code_proto_enumTypes[0].Descriptor()
}

func (PANGU_CODE) Type() protoreflect.EnumType {
	return &file_code_proto_enumTypes[0]
}

func (x PANGU_CODE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PANGU_CODE.Descriptor instead.
func (PANGU_CODE) EnumDescriptor() ([]byte, []int) {
	return file_code_proto_rawDescGZIP(), []int{0}
}

var File_code_proto protoreflect.FileDescriptor

var file_code_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x92, 0x0b, 0x0a,
	0x0a, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x12, 0x1a, 0x0a, 0x16, 0x50,
	0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0d, 0x50, 0x41, 0x4e, 0x47, 0x55,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4f, 0x4b, 0x10, 0xd0, 0x0f, 0x12, 0x1d, 0x0a, 0x18, 0x50,
	0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x4f,
	0x47, 0x47, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x10, 0xaa, 0x1f, 0x12, 0x1c, 0x0a, 0x17, 0x50, 0x41,
	0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f,
	0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0xab, 0x1f, 0x12, 0x21, 0x0a, 0x1c, 0x50, 0x41, 0x4e, 0x47,
	0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x45, 0x54, 0x45, 0x52,
	0x5f, 0x49, 0x4c, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x10, 0xbe, 0x1f, 0x12, 0x1e, 0x0a, 0x19, 0x50,
	0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x54, 0x52, 0x55, 0x53,
	0x54, 0x45, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0xbf, 0x1f, 0x12, 0x21, 0x0a, 0x1c, 0x50,
	0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x49, 0x4e,
	0x47, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0xc8, 0x1f, 0x12, 0x1f,
	0x0a, 0x1a, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0xf0, 0x1f, 0x12,
	0x1a, 0x0a, 0x15, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x88, 0x27, 0x12, 0x1f, 0x0a, 0x1a, 0x50,
	0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x55,
	0x4e, 0x50, 0x41, 0x43, 0x4b, 0x5f, 0x45, 0x52, 0x52, 0x10, 0xf0, 0x2e, 0x12, 0x1d, 0x0a, 0x18,
	0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x5f,
	0x50, 0x41, 0x43, 0x4b, 0x5f, 0x45, 0x52, 0x52, 0x10, 0xf1, 0x2e, 0x12, 0x19, 0x0a, 0x14, 0x50,
	0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45, 0x44, 0x49, 0x53, 0x5f,
	0x45, 0x52, 0x52, 0x10, 0xf2, 0x2e, 0x12, 0x17, 0x0a, 0x12, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0xd5, 0x2f, 0x12,
	0x1d, 0x0a, 0x18, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4d, 0x4f,
	0x44, 0x45, 0x4c, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x10, 0xd6, 0x2f, 0x12, 0x1d,
	0x0a, 0x18, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4d, 0x4f, 0x44,
	0x45, 0x4c, 0x5f, 0x44, 0x45, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x10, 0xd7, 0x2f, 0x12, 0x1f, 0x0a,
	0x1a, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45,
	0x4c, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x10, 0xd8, 0x2f, 0x12, 0x1f,
	0x0a, 0x1a, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x4c, 0x41,
	0x59, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x5f, 0x47, 0x55, 0x49, 0x4c, 0x44, 0x10, 0xd8, 0x36, 0x12,
	0x1f, 0x0a, 0x1a, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x47, 0x55,
	0x49, 0x4c, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0xd9, 0x36,
	0x12, 0x1f, 0x0a, 0x1a, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x50,
	0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x5f, 0x47, 0x55, 0x49, 0x4c, 0x44, 0x10, 0xda,
	0x36, 0x12, 0x20, 0x0a, 0x1b, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54,
	0x10, 0xdb, 0x36, 0x12, 0x21, 0x0a, 0x1c, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x47, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x10, 0xdc, 0x36, 0x12, 0x24, 0x0a, 0x1f, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x47, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x4d,
	0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x4e, 0x55, 0x4d, 0x10, 0xdd, 0x36, 0x12, 0x21, 0x0a, 0x1c,
	0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45,
	0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0xde, 0x36, 0x12,
	0x23, 0x0a, 0x1e, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x47, 0x55,
	0x49, 0x4c, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4b, 0x49, 0x43, 0x4b, 0x5f, 0x53, 0x45, 0x4c,
	0x46, 0x10, 0xdf, 0x36, 0x12, 0x24, 0x0a, 0x1f, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x49, 0x53, 0x5f, 0x43, 0x48, 0x45,
	0x41, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0xe0, 0x36, 0x12, 0x2a, 0x0a, 0x25, 0x50, 0x41,
	0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x56, 0x41, 0x4c, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x53, 0x48,
	0x4f, 0x52, 0x54, 0x10, 0xbc, 0x37, 0x12, 0x21, 0x0a, 0x1c, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x48, 0x45, 0x4c, 0x50, 0x5f, 0x49, 0x53, 0x5f, 0x53, 0x41, 0x54,
	0x49, 0x53, 0x46, 0x49, 0x45, 0x44, 0x10, 0xbd, 0x37, 0x12, 0x20, 0x0a, 0x1b, 0x50, 0x41, 0x4e,
	0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x48, 0x45, 0x4c, 0x50, 0x5f, 0x49, 0x53, 0x5f,
	0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x44, 0x10, 0xc1, 0x37, 0x12, 0x21, 0x0a, 0x1c, 0x50,
	0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49,
	0x54, 0x59, 0x5f, 0x49, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xc0, 0x3e, 0x12, 0x28,
	0x0a, 0x23, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x5f,
	0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0xc1, 0x3e, 0x12, 0x20, 0x0a, 0x1b, 0x50, 0x41, 0x4e, 0x47,
	0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f,
	0x4e, 0x4f, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0xc2, 0x3e, 0x12, 0x26, 0x0a, 0x21, 0x50, 0x41,
	0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54,
	0x59, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x49, 0x56, 0x45, 0x5f, 0x41, 0x57, 0x41, 0x52, 0x44, 0x10,
	0xc3, 0x3e, 0x12, 0x30, 0x0a, 0x2b, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4a, 0x4f, 0x49, 0x4e,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55,
	0x54, 0x10, 0xc4, 0x3e, 0x12, 0x28, 0x0a, 0x23, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x50, 0x4c, 0x41, 0x59,
	0x45, 0x52, 0x5f, 0x49, 0x4e, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0xc5, 0x3e, 0x12, 0x23,
	0x0a, 0x1e, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x5f, 0x45, 0x4e, 0x54, 0x45, 0x52,
	0x10, 0xc6, 0x3e, 0x12, 0x29, 0x0a, 0x24, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x57, 0x41, 0x52,
	0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0xc7, 0x3e, 0x12, 0x1f,
	0x0a, 0x1a, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x49, 0x47,
	0x45, 0x4f, 0x4e, 0x5f, 0x47, 0x52, 0x50, 0x43, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x90, 0x4e, 0x12,
	0x29, 0x0a, 0x24, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x49,
	0x47, 0x45, 0x4f, 0x4e, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x5f, 0x49, 0x47, 0x4e, 0x4f, 0x52, 0x45,
	0x5f, 0x4d, 0x59, 0x53, 0x45, 0x4c, 0x46, 0x10, 0x91, 0x4e, 0x12, 0x24, 0x0a, 0x1f, 0x50, 0x41,
	0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x49, 0x47, 0x45, 0x4f, 0x4e, 0x5f,
	0x47, 0x52, 0x50, 0x43, 0x5f, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x92, 0x4e,
	0x12, 0x1a, 0x0a, 0x15, 0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4d,
	0x41, 0x44, 0x45, 0x4c, 0x5f, 0x52, 0x55, 0x53, 0x48, 0x10, 0x97, 0x56, 0x12, 0x23, 0x0a, 0x1e,
	0x50, 0x41, 0x4e, 0x47, 0x55, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4d, 0x41, 0x44, 0x45, 0x4c,
	0x5f, 0x52, 0x55, 0x53, 0x48, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x5f, 0x45, 0x4e, 0x44, 0x10, 0x98,
	0x56, 0x42, 0x17, 0x48, 0x01, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x61, 0x70, 0x69, 0xaa, 0x02, 0x0a,
	0x4d, 0x4d, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_code_proto_rawDescOnce sync.Once
	file_code_proto_rawDescData = file_code_proto_rawDesc
)

func file_code_proto_rawDescGZIP() []byte {
	file_code_proto_rawDescOnce.Do(func() {
		file_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_code_proto_rawDescData)
	})
	return file_code_proto_rawDescData
}

var file_code_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_code_proto_goTypes = []interface{}{
	(PANGU_CODE)(0), // 0: PANGU_CODE
}
var file_code_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_code_proto_init() }
func file_code_proto_init() {
	if File_code_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_code_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_code_proto_goTypes,
		DependencyIndexes: file_code_proto_depIdxs,
		EnumInfos:         file_code_proto_enumTypes,
	}.Build()
	File_code_proto = out.File
	file_code_proto_rawDesc = nil
	file_code_proto_goTypes = nil
	file_code_proto_depIdxs = nil
}
