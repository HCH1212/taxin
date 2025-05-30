// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: api/user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Password      string                 `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	Like          []string               `protobuf:"bytes,2,rep,name=like,proto3" json:"like,omitempty"`
	Username      string                 `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	mi := &file_api_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterReq) GetLike() []string {
	if x != nil {
		return x.Like
	}
	return nil
}

func (x *RegisterReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type RegisterResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResp) Reset() {
	*x = RegisterResp{}
	mi := &file_api_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResp) ProtoMessage() {}

func (x *RegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResp.ProtoReflect.Descriptor instead.
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	mi := &file_api_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[2]
	if x != nil {
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
	return file_api_user_proto_rawDescGZIP(), []int{2}
}

func (x *LoginReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	mi := &file_api_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type UserInfoReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoReq) Reset() {
	*x = UserInfoReq{}
	mi := &file_api_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoReq) ProtoMessage() {}

func (x *UserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoReq.ProtoReflect.Descriptor instead.
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{4}
}

type UserInfoResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Like          []string               `protobuf:"bytes,2,rep,name=like,proto3" json:"like,omitempty"`
	LikeEmbedding []float32              `protobuf:"fixed32,3,rep,packed,name=like_embedding,json=likeEmbedding,proto3" json:"like_embedding,omitempty"`
	CreateAt      string                 `protobuf:"bytes,4,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	UpdateAt      string                 `protobuf:"bytes,5,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	Username      string                 `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfoResp) Reset() {
	*x = UserInfoResp{}
	mi := &file_api_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResp) ProtoMessage() {}

func (x *UserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResp.ProtoReflect.Descriptor instead.
func (*UserInfoResp) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{5}
}

func (x *UserInfoResp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserInfoResp) GetLike() []string {
	if x != nil {
		return x.Like
	}
	return nil
}

func (x *UserInfoResp) GetLikeEmbedding() []float32 {
	if x != nil {
		return x.LikeEmbedding
	}
	return nil
}

func (x *UserInfoResp) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

func (x *UserInfoResp) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

func (x *UserInfoResp) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_api_user_proto protoreflect.FileDescriptor

const file_api_user_proto_rawDesc = "" +
	"\n" +
	"\x0eapi/user.proto\x12\x04user\"Y\n" +
	"\vRegisterReq\x12\x1a\n" +
	"\bpassword\x18\x01 \x01(\tR\bpassword\x12\x12\n" +
	"\x04like\x18\x02 \x03(\tR\x04like\x12\x1a\n" +
	"\busername\x18\x03 \x01(\tR\busername\"'\n" +
	"\fRegisterResp\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"?\n" +
	"\bLoginReq\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\".\n" +
	"\tLoginResp\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\"\r\n" +
	"\vUserInfoReq\"\xb8\x01\n" +
	"\fUserInfoResp\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x12\n" +
	"\x04like\x18\x02 \x03(\tR\x04like\x12%\n" +
	"\x0elike_embedding\x18\x03 \x03(\x02R\rlikeEmbedding\x12\x1b\n" +
	"\tcreate_at\x18\x04 \x01(\tR\bcreateAt\x12\x1b\n" +
	"\tupdate_at\x18\x05 \x01(\tR\bupdateAt\x12\x1a\n" +
	"\busername\x18\x06 \x01(\tR\busername2\xa0\x01\n" +
	"\vUserService\x121\n" +
	"\bRegister\x12\x11.user.RegisterReq\x1a\x12.user.RegisterResp\x12(\n" +
	"\x05Login\x12\x0e.user.LoginReq\x1a\x0f.user.LoginResp\x124\n" +
	"\vGetUserInfo\x12\x11.user.UserInfoReq\x1a\x12.user.UserInfoRespB\aZ\x05/userb\x06proto3"

var (
	file_api_user_proto_rawDescOnce sync.Once
	file_api_user_proto_rawDescData []byte
)

func file_api_user_proto_rawDescGZIP() []byte {
	file_api_user_proto_rawDescOnce.Do(func() {
		file_api_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_user_proto_rawDesc), len(file_api_user_proto_rawDesc)))
	})
	return file_api_user_proto_rawDescData
}

var file_api_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_user_proto_goTypes = []any{
	(*RegisterReq)(nil),  // 0: user.RegisterReq
	(*RegisterResp)(nil), // 1: user.RegisterResp
	(*LoginReq)(nil),     // 2: user.LoginReq
	(*LoginResp)(nil),    // 3: user.LoginResp
	(*UserInfoReq)(nil),  // 4: user.UserInfoReq
	(*UserInfoResp)(nil), // 5: user.UserInfoResp
}
var file_api_user_proto_depIdxs = []int32{
	0, // 0: user.UserService.Register:input_type -> user.RegisterReq
	2, // 1: user.UserService.Login:input_type -> user.LoginReq
	4, // 2: user.UserService.GetUserInfo:input_type -> user.UserInfoReq
	1, // 3: user.UserService.Register:output_type -> user.RegisterResp
	3, // 4: user.UserService.Login:output_type -> user.LoginResp
	5, // 5: user.UserService.GetUserInfo:output_type -> user.UserInfoResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_proto_init() }
func file_api_user_proto_init() {
	if File_api_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_user_proto_rawDesc), len(file_api_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_user_proto_goTypes,
		DependencyIndexes: file_api_user_proto_depIdxs,
		MessageInfos:      file_api_user_proto_msgTypes,
	}.Build()
	File_api_user_proto = out.File
	file_api_user_proto_goTypes = nil
	file_api_user_proto_depIdxs = nil
}
