// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: engine/auth/authorized_profile.proto

package auth

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

type PartyIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartyCode string `protobuf:"bytes,1,opt,name=party_code,json=partyCode,proto3" json:"party_code,omitempty"`
	// base64 encoded version of the DER encoded public key
	PublicKey string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *PartyIdentity) Reset() {
	*x = PartyIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_auth_authorized_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyIdentity) ProtoMessage() {}

func (x *PartyIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_engine_auth_authorized_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyIdentity.ProtoReflect.Descriptor instead.
func (*PartyIdentity) Descriptor() ([]byte, []int) {
	return file_engine_auth_authorized_profile_proto_rawDescGZIP(), []int{0}
}

func (x *PartyIdentity) GetPartyCode() string {
	if x != nil {
		return x.PartyCode
	}
	return ""
}

func (x *PartyIdentity) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type DriverIdentity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// base64 encoded version of the DER encoded public key
	PublicKey string `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *DriverIdentity) Reset() {
	*x = DriverIdentity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_auth_authorized_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverIdentity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverIdentity) ProtoMessage() {}

func (x *DriverIdentity) ProtoReflect() protoreflect.Message {
	mi := &file_engine_auth_authorized_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverIdentity.ProtoReflect.Descriptor instead.
func (*DriverIdentity) Descriptor() ([]byte, []int) {
	return file_engine_auth_authorized_profile_proto_rawDescGZIP(), []int{1}
}

func (x *DriverIdentity) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type AuthorizedProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver  *DriverIdentity  `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Parties []*PartyIdentity `protobuf:"bytes,2,rep,name=parties,proto3" json:"parties,omitempty"`
}

func (x *AuthorizedProfile) Reset() {
	*x = AuthorizedProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_auth_authorized_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizedProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizedProfile) ProtoMessage() {}

func (x *AuthorizedProfile) ProtoReflect() protoreflect.Message {
	mi := &file_engine_auth_authorized_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizedProfile.ProtoReflect.Descriptor instead.
func (*AuthorizedProfile) Descriptor() ([]byte, []int) {
	return file_engine_auth_authorized_profile_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorizedProfile) GetDriver() *DriverIdentity {
	if x != nil {
		return x.Driver
	}
	return nil
}

func (x *AuthorizedProfile) GetParties() []*PartyIdentity {
	if x != nil {
		return x.Parties
	}
	return nil
}

var File_engine_auth_authorized_profile_proto protoreflect.FileDescriptor

var file_engine_auth_authorized_profile_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x62, 0x22, 0x4d, 0x0a, 0x0d, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x2f, 0x0a, 0x0e, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x8e, 0x01, 0x0a, 0x11,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x12, 0x3b, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x3c,
	0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x42, 0xd0, 0x01, 0x0a,
	0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x62, 0x42, 0x16, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x73, 0x63, 0x71, 0x6c, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0xa2, 0x02, 0x04, 0x53, 0x45, 0x41, 0x50, 0xaa, 0x02, 0x13, 0x53, 0x63, 0x71, 0x6c, 0x2e,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x62, 0xca, 0x02,
	0x13, 0x53, 0x63, 0x71, 0x6c, 0x5c, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x41, 0x75, 0x74,
	0x68, 0x5c, 0x50, 0x62, 0xe2, 0x02, 0x1f, 0x53, 0x63, 0x71, 0x6c, 0x5c, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x50, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x53, 0x63, 0x71, 0x6c, 0x3a, 0x3a, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x50, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_engine_auth_authorized_profile_proto_rawDescOnce sync.Once
	file_engine_auth_authorized_profile_proto_rawDescData = file_engine_auth_authorized_profile_proto_rawDesc
)

func file_engine_auth_authorized_profile_proto_rawDescGZIP() []byte {
	file_engine_auth_authorized_profile_proto_rawDescOnce.Do(func() {
		file_engine_auth_authorized_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_engine_auth_authorized_profile_proto_rawDescData)
	})
	return file_engine_auth_authorized_profile_proto_rawDescData
}

var file_engine_auth_authorized_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_engine_auth_authorized_profile_proto_goTypes = []interface{}{
	(*PartyIdentity)(nil),     // 0: scql.engine.auth.pb.PartyIdentity
	(*DriverIdentity)(nil),    // 1: scql.engine.auth.pb.DriverIdentity
	(*AuthorizedProfile)(nil), // 2: scql.engine.auth.pb.AuthorizedProfile
}
var file_engine_auth_authorized_profile_proto_depIdxs = []int32{
	1, // 0: scql.engine.auth.pb.AuthorizedProfile.driver:type_name -> scql.engine.auth.pb.DriverIdentity
	0, // 1: scql.engine.auth.pb.AuthorizedProfile.parties:type_name -> scql.engine.auth.pb.PartyIdentity
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_engine_auth_authorized_profile_proto_init() }
func file_engine_auth_authorized_profile_proto_init() {
	if File_engine_auth_authorized_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_engine_auth_authorized_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyIdentity); i {
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
		file_engine_auth_authorized_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverIdentity); i {
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
		file_engine_auth_authorized_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizedProfile); i {
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
			RawDescriptor: file_engine_auth_authorized_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_engine_auth_authorized_profile_proto_goTypes,
		DependencyIndexes: file_engine_auth_authorized_profile_proto_depIdxs,
		MessageInfos:      file_engine_auth_authorized_profile_proto_msgTypes,
	}.Build()
	File_engine_auth_authorized_profile_proto = out.File
	file_engine_auth_authorized_profile_proto_rawDesc = nil
	file_engine_auth_authorized_profile_proto_goTypes = nil
	file_engine_auth_authorized_profile_proto_depIdxs = nil
}
