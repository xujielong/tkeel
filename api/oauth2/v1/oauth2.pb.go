//
//Copyright 2021 The tKeel Authors.
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//http://www.apache.org/licenses/LICENSE-2.0
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api/oauth2/v1/oauth2.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IssueOauth2TokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
}

func (x *IssueOauth2TokenRequest) Reset() {
	*x = IssueOauth2TokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_oauth2_v1_oauth2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueOauth2TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueOauth2TokenRequest) ProtoMessage() {}

func (x *IssueOauth2TokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_oauth2_v1_oauth2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueOauth2TokenRequest.ProtoReflect.Descriptor instead.
func (*IssueOauth2TokenRequest) Descriptor() ([]byte, []int) {
	return file_api_oauth2_v1_oauth2_proto_rawDescGZIP(), []int{0}
}

func (x *IssueOauth2TokenRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *IssueOauth2TokenRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

type IssueOauth2TokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	TokenType    string `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresIn    int32  `protobuf:"varint,4,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
}

func (x *IssueOauth2TokenResponse) Reset() {
	*x = IssueOauth2TokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_oauth2_v1_oauth2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueOauth2TokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueOauth2TokenResponse) ProtoMessage() {}

func (x *IssueOauth2TokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_oauth2_v1_oauth2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueOauth2TokenResponse.ProtoReflect.Descriptor instead.
func (*IssueOauth2TokenResponse) Descriptor() ([]byte, []int) {
	return file_api_oauth2_v1_oauth2_proto_rawDescGZIP(), []int{1}
}

func (x *IssueOauth2TokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *IssueOauth2TokenResponse) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *IssueOauth2TokenResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *IssueOauth2TokenResponse) GetExpiresIn() int32 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

type AddWhiteListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Secret   string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *AddWhiteListRequest) Reset() {
	*x = AddWhiteListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_oauth2_v1_oauth2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddWhiteListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddWhiteListRequest) ProtoMessage() {}

func (x *AddWhiteListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_oauth2_v1_oauth2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddWhiteListRequest.ProtoReflect.Descriptor instead.
func (*AddWhiteListRequest) Descriptor() ([]byte, []int) {
	return file_api_oauth2_v1_oauth2_proto_rawDescGZIP(), []int{2}
}

func (x *AddWhiteListRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AddWhiteListRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

var File_api_oauth2_v1_oauth2_proto protoreflect.FileDescriptor

var file_api_oauth2_v1_oauth2_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70,
	0x69, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x17, 0x49, 0x73, 0x73, 0x75, 0x65, 0x4f,
	0x61, 0x75, 0x74, 0x68, 0x32, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x22, 0xa0, 0x01, 0x0a, 0x18, 0x49, 0x73, 0x73, 0x75, 0x65, 0x4f, 0x61, 0x75,
	0x74, 0x68, 0x32, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x22, 0x4a, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x57, 0x68, 0x69,
	0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x32, 0xec, 0x01, 0x0a, 0x06, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x12, 0x77, 0x0a,
	0x10, 0x49, 0x73, 0x73, 0x75, 0x65, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x4f,
	0x61, 0x75, 0x74, 0x68, 0x32, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x07, 0x2f, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x32, 0x3a, 0x01, 0x2a, 0x12, 0x69, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x57, 0x68, 0x69,
	0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x32, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x32, 0x2f, 0x77, 0x68, 0x69, 0x74, 0x65, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x3a, 0x01,
	0x2a, 0x42, 0x3d, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_oauth2_v1_oauth2_proto_rawDescOnce sync.Once
	file_api_oauth2_v1_oauth2_proto_rawDescData = file_api_oauth2_v1_oauth2_proto_rawDesc
)

func file_api_oauth2_v1_oauth2_proto_rawDescGZIP() []byte {
	file_api_oauth2_v1_oauth2_proto_rawDescOnce.Do(func() {
		file_api_oauth2_v1_oauth2_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_oauth2_v1_oauth2_proto_rawDescData)
	})
	return file_api_oauth2_v1_oauth2_proto_rawDescData
}

var file_api_oauth2_v1_oauth2_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_oauth2_v1_oauth2_proto_goTypes = []interface{}{
	(*IssueOauth2TokenRequest)(nil),  // 0: api.oauth2.v1.IssueOauth2TokenRequest
	(*IssueOauth2TokenResponse)(nil), // 1: api.oauth2.v1.IssueOauth2TokenResponse
	(*AddWhiteListRequest)(nil),      // 2: api.oauth2.v1.AddWhiteListRequest
	(*emptypb.Empty)(nil),            // 3: google.protobuf.Empty
}
var file_api_oauth2_v1_oauth2_proto_depIdxs = []int32{
	0, // 0: api.oauth2.v1.Oauth2.IssueOauth2Token:input_type -> api.oauth2.v1.IssueOauth2TokenRequest
	2, // 1: api.oauth2.v1.Oauth2.AddWhiteList:input_type -> api.oauth2.v1.AddWhiteListRequest
	1, // 2: api.oauth2.v1.Oauth2.IssueOauth2Token:output_type -> api.oauth2.v1.IssueOauth2TokenResponse
	3, // 3: api.oauth2.v1.Oauth2.AddWhiteList:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_oauth2_v1_oauth2_proto_init() }
func file_api_oauth2_v1_oauth2_proto_init() {
	if File_api_oauth2_v1_oauth2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_oauth2_v1_oauth2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueOauth2TokenRequest); i {
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
		file_api_oauth2_v1_oauth2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueOauth2TokenResponse); i {
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
		file_api_oauth2_v1_oauth2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddWhiteListRequest); i {
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
			RawDescriptor: file_api_oauth2_v1_oauth2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_oauth2_v1_oauth2_proto_goTypes,
		DependencyIndexes: file_api_oauth2_v1_oauth2_proto_depIdxs,
		MessageInfos:      file_api_oauth2_v1_oauth2_proto_msgTypes,
	}.Build()
	File_api_oauth2_v1_oauth2_proto = out.File
	file_api_oauth2_v1_oauth2_proto_rawDesc = nil
	file_api_oauth2_v1_oauth2_proto_goTypes = nil
	file_api_oauth2_v1_oauth2_proto_depIdxs = nil
}
