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
// source: api/common.proto

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

// RequestHeader carries the user custom headers.
type RequestHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Custom headers used to record custom information.
	CustomHeaders map[string]string `protobuf:"bytes,1,rep,name=custom_headers,json=customHeaders,proto3" json:"custom_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RequestHeader) Reset() {
	*x = RequestHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHeader) ProtoMessage() {}

func (x *RequestHeader) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHeader.ProtoReflect.Descriptor instead.
func (*RequestHeader) Descriptor() ([]byte, []int) {
	return file_api_common_proto_rawDescGZIP(), []int{0}
}

func (x *RequestHeader) GetCustomHeaders() map[string]string {
	if x != nil {
		return x.CustomHeaders
	}
	return nil
}

// TODO: move SQLWarning to a proper place
type SQLWarning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Description of the warning
	Reason string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *SQLWarning) Reset() {
	*x = SQLWarning{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLWarning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLWarning) ProtoMessage() {}

func (x *SQLWarning) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLWarning.ProtoReflect.Descriptor instead.
func (*SQLWarning) Descriptor() ([]byte, []int) {
	return file_api_common_proto_rawDescGZIP(), []int{1}
}

func (x *SQLWarning) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_api_common_proto protoreflect.FileDescriptor

var file_api_common_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x1a, 0x0e, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01,
	0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x50, 0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x1a, 0x40, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x24, 0x0a, 0x0a, 0x53, 0x51, 0x4c, 0x57, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x7e, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x42, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x73,
	0x63, 0x71, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0xa2, 0x02,
	0x03, 0x53, 0x50, 0x58, 0xaa, 0x02, 0x07, 0x53, 0x63, 0x71, 0x6c, 0x2e, 0x50, 0x62, 0xca, 0x02,
	0x07, 0x53, 0x63, 0x71, 0x6c, 0x5c, 0x50, 0x62, 0xe2, 0x02, 0x13, 0x53, 0x63, 0x71, 0x6c, 0x5c,
	0x50, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x08, 0x53, 0x63, 0x71, 0x6c, 0x3a, 0x3a, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_common_proto_rawDescOnce sync.Once
	file_api_common_proto_rawDescData = file_api_common_proto_rawDesc
)

func file_api_common_proto_rawDescGZIP() []byte {
	file_api_common_proto_rawDescOnce.Do(func() {
		file_api_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_common_proto_rawDescData)
	})
	return file_api_common_proto_rawDescData
}

var file_api_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_common_proto_goTypes = []interface{}{
	(*RequestHeader)(nil), // 0: scql.pb.RequestHeader
	(*SQLWarning)(nil),    // 1: scql.pb.SQLWarning
	nil,                   // 2: scql.pb.RequestHeader.CustomHeadersEntry
}
var file_api_common_proto_depIdxs = []int32{
	2, // 0: scql.pb.RequestHeader.custom_headers:type_name -> scql.pb.RequestHeader.CustomHeadersEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_common_proto_init() }
func file_api_common_proto_init() {
	if File_api_common_proto != nil {
		return
	}
	file_api_core_proto_init()
	file_api_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHeader); i {
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
		file_api_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLWarning); i {
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
			RawDescriptor: file_api_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_common_proto_goTypes,
		DependencyIndexes: file_api_common_proto_depIdxs,
		MessageInfos:      file_api_common_proto_msgTypes,
	}.Build()
	File_api_common_proto = out.File
	file_api_common_proto_rawDesc = nil
	file_api_common_proto_goTypes = nil
	file_api_common_proto_depIdxs = nil
}
