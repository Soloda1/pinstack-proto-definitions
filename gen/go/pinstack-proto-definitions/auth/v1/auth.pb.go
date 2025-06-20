// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: auth/auth.proto

package authv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

// Token pair message
type TokenPair struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TokenPair) Reset() {
	*x = TokenPair{}
	mi := &file_auth_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenPair) ProtoMessage() {}

func (x *TokenPair) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenPair.ProtoReflect.Descriptor instead.
func (*TokenPair) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *TokenPair) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *TokenPair) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

// Request messages
type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	FullName      *string                `protobuf:"bytes,4,opt,name=full_name,json=fullName,proto3,oneof" json:"full_name,omitempty"`
	Bio           *string                `protobuf:"bytes,5,opt,name=bio,proto3,oneof" json:"bio,omitempty"`
	AvatarUrl     *string                `protobuf:"bytes,6,opt,name=avatar_url,json=avatarUrl,proto3,oneof" json:"avatar_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_auth_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

func (x *RegisterRequest) GetBio() string {
	if x != nil && x.Bio != nil {
		return *x.Bio
	}
	return ""
}

func (x *RegisterRequest) GetAvatarUrl() string {
	if x != nil && x.AvatarUrl != nil {
		return *x.AvatarUrl
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Login         string                 `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"` // can be username or email
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_auth_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RefreshRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefreshToken  string                 `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshRequest) Reset() {
	*x = RefreshRequest{}
	mi := &file_auth_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshRequest) ProtoMessage() {}

func (x *RefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshRequest.ProtoReflect.Descriptor instead.
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *RefreshRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type LogoutRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefreshToken  string                 `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	mi := &file_auth_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *LogoutRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type UpdatePasswordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OldPassword   string                 `protobuf:"bytes,2,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword   string                 `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePasswordRequest) Reset() {
	*x = UpdatePasswordRequest{}
	mi := &file_auth_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePasswordRequest) ProtoMessage() {}

func (x *UpdatePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePasswordRequest.ProtoReflect.Descriptor instead.
func (*UpdatePasswordRequest) Descriptor() ([]byte, []int) {
	return file_auth_auth_proto_rawDescGZIP(), []int{5}
}

func (x *UpdatePasswordRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePasswordRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *UpdatePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

var File_auth_auth_proto protoreflect.FileDescriptor

const file_auth_auth_proto_rawDesc = "" +
	"\n" +
	"\x0fauth/auth.proto\x12\aauth.v1\x1a\x1bgoogle/protobuf/empty.proto\"S\n" +
	"\tTokenPair\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12#\n" +
	"\rrefresh_token\x18\x02 \x01(\tR\frefreshToken\"\xe1\x01\n" +
	"\x0fRegisterRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\x12 \n" +
	"\tfull_name\x18\x04 \x01(\tH\x00R\bfullName\x88\x01\x01\x12\x15\n" +
	"\x03bio\x18\x05 \x01(\tH\x01R\x03bio\x88\x01\x01\x12\"\n" +
	"\n" +
	"avatar_url\x18\x06 \x01(\tH\x02R\tavatarUrl\x88\x01\x01B\f\n" +
	"\n" +
	"_full_nameB\x06\n" +
	"\x04_bioB\r\n" +
	"\v_avatar_url\"@\n" +
	"\fLoginRequest\x12\x14\n" +
	"\x05login\x18\x01 \x01(\tR\x05login\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"5\n" +
	"\x0eRefreshRequest\x12#\n" +
	"\rrefresh_token\x18\x01 \x01(\tR\frefreshToken\"4\n" +
	"\rLogoutRequest\x12#\n" +
	"\rrefresh_token\x18\x01 \x01(\tR\frefreshToken\"m\n" +
	"\x15UpdatePasswordRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12!\n" +
	"\fold_password\x18\x02 \x01(\tR\voldPassword\x12!\n" +
	"\fnew_password\x18\x03 \x01(\tR\vnewPassword2\xc1\x02\n" +
	"\vAuthService\x12:\n" +
	"\bRegister\x12\x18.auth.v1.RegisterRequest\x1a\x12.auth.v1.TokenPair\"\x00\x124\n" +
	"\x05Login\x12\x15.auth.v1.LoginRequest\x1a\x12.auth.v1.TokenPair\"\x00\x128\n" +
	"\aRefresh\x12\x17.auth.v1.RefreshRequest\x1a\x12.auth.v1.TokenPair\"\x00\x12:\n" +
	"\x06Logout\x12\x16.auth.v1.LogoutRequest\x1a\x16.google.protobuf.Empty\"\x00\x12J\n" +
	"\x0eUpdatePassword\x12\x1e.auth.v1.UpdatePasswordRequest\x1a\x16.google.protobuf.Empty\"\x00B+Z)pinstack-proto-definitions/auth/v1;authv1b\x06proto3"

var (
	file_auth_auth_proto_rawDescOnce sync.Once
	file_auth_auth_proto_rawDescData []byte
)

func file_auth_auth_proto_rawDescGZIP() []byte {
	file_auth_auth_proto_rawDescOnce.Do(func() {
		file_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_auth_auth_proto_rawDesc), len(file_auth_auth_proto_rawDesc)))
	})
	return file_auth_auth_proto_rawDescData
}

var file_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_auth_auth_proto_goTypes = []any{
	(*TokenPair)(nil),             // 0: auth.v1.TokenPair
	(*RegisterRequest)(nil),       // 1: auth.v1.RegisterRequest
	(*LoginRequest)(nil),          // 2: auth.v1.LoginRequest
	(*RefreshRequest)(nil),        // 3: auth.v1.RefreshRequest
	(*LogoutRequest)(nil),         // 4: auth.v1.LogoutRequest
	(*UpdatePasswordRequest)(nil), // 5: auth.v1.UpdatePasswordRequest
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_auth_auth_proto_depIdxs = []int32{
	1, // 0: auth.v1.AuthService.Register:input_type -> auth.v1.RegisterRequest
	2, // 1: auth.v1.AuthService.Login:input_type -> auth.v1.LoginRequest
	3, // 2: auth.v1.AuthService.Refresh:input_type -> auth.v1.RefreshRequest
	4, // 3: auth.v1.AuthService.Logout:input_type -> auth.v1.LogoutRequest
	5, // 4: auth.v1.AuthService.UpdatePassword:input_type -> auth.v1.UpdatePasswordRequest
	0, // 5: auth.v1.AuthService.Register:output_type -> auth.v1.TokenPair
	0, // 6: auth.v1.AuthService.Login:output_type -> auth.v1.TokenPair
	0, // 7: auth.v1.AuthService.Refresh:output_type -> auth.v1.TokenPair
	6, // 8: auth.v1.AuthService.Logout:output_type -> google.protobuf.Empty
	6, // 9: auth.v1.AuthService.UpdatePassword:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_auth_proto_init() }
func file_auth_auth_proto_init() {
	if File_auth_auth_proto != nil {
		return
	}
	file_auth_auth_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_auth_auth_proto_rawDesc), len(file_auth_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_auth_proto_goTypes,
		DependencyIndexes: file_auth_auth_proto_depIdxs,
		MessageInfos:      file_auth_auth_proto_msgTypes,
	}.Build()
	File_auth_auth_proto = out.File
	file_auth_auth_proto_goTypes = nil
	file_auth_auth_proto_depIdxs = nil
}
