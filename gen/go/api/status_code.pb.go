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
// source: api/status_code.proto

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

type Code int32

const (
	// Not an error; returned on success
	Code_OK Code = 0
	// General bad request
	Code_BAD_REQUEST Code = 100
	// Failed to authenticate user credential
	Code_UNAUTHENTICATED Code = 101
	// Parsing SQL error, returned on invalid SQL
	Code_SQL_PARSE_ERROR  Code = 102
	Code_INVALID_ARGUMENT Code = 103
	// Query result is not ready, please try again later
	Code_NOT_READY Code = 104
	// The user does not have permission to execute the DDL.
	// For example, the user try to drop other's table.
	Code_DDL_PERMISSION_DENIED Code = 131
	// General not found error
	Code_NOT_FOUND Code = 140
	// SCDB session not found
	Code_SESSION_NOT_FOUND Code = 141
	// Query CCL check failed
	Code_CCL_CHECK_FAILED Code = 160
	// The P2P internal exchange error
	// Local storage info is not equal to storage info in other parties
	Code_CHECKSUM_CHECK_FAILED Code = 170
	// SCDB DB error
	Code_STORAGE_ERROR Code = 201
	// UNKNOWN Server Internal Error
	Code_INTERNAL Code = 300
	// 320-339 executing execution graph error
	Code_UNKNOWN_ENGINE_ERROR Code = 320
	Code_ENGINE_RUNSQL_ERROR  Code = 332
	// Feature not supported
	Code_NOT_SUPPORTED Code = 340
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:   "OK",
		100: "BAD_REQUEST",
		101: "UNAUTHENTICATED",
		102: "SQL_PARSE_ERROR",
		103: "INVALID_ARGUMENT",
		104: "NOT_READY",
		131: "DDL_PERMISSION_DENIED",
		140: "NOT_FOUND",
		141: "SESSION_NOT_FOUND",
		160: "CCL_CHECK_FAILED",
		170: "CHECKSUM_CHECK_FAILED",
		201: "STORAGE_ERROR",
		300: "INTERNAL",
		320: "UNKNOWN_ENGINE_ERROR",
		332: "ENGINE_RUNSQL_ERROR",
		340: "NOT_SUPPORTED",
	}
	Code_value = map[string]int32{
		"OK":                    0,
		"BAD_REQUEST":           100,
		"UNAUTHENTICATED":       101,
		"SQL_PARSE_ERROR":       102,
		"INVALID_ARGUMENT":      103,
		"NOT_READY":             104,
		"DDL_PERMISSION_DENIED": 131,
		"NOT_FOUND":             140,
		"SESSION_NOT_FOUND":     141,
		"CCL_CHECK_FAILED":      160,
		"CHECKSUM_CHECK_FAILED": 170,
		"STORAGE_ERROR":         201,
		"INTERNAL":              300,
		"UNKNOWN_ENGINE_ERROR":  320,
		"ENGINE_RUNSQL_ERROR":   332,
		"NOT_SUPPORTED":         340,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_api_status_code_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_api_status_code_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_api_status_code_proto_rawDescGZIP(), []int{0}
}

var File_api_status_code_proto protoreflect.FileDescriptor

var file_api_status_code_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62,
	0x2a, 0xd1, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x41, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x10, 0x64, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49,
	0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x51, 0x4c, 0x5f, 0x50,
	0x41, 0x52, 0x53, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x66, 0x12, 0x14, 0x0a, 0x10,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x67, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10,
	0x68, 0x12, 0x1a, 0x0a, 0x15, 0x44, 0x44, 0x4c, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x83, 0x01, 0x12, 0x0e, 0x0a,
	0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x8c, 0x01, 0x12, 0x16, 0x0a,
	0x11, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x8d, 0x01, 0x12, 0x15, 0x0a, 0x10, 0x43, 0x43, 0x4c, 0x5f, 0x43, 0x48, 0x45,
	0x43, 0x4b, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xa0, 0x01, 0x12, 0x1a, 0x0a, 0x15,
	0x43, 0x48, 0x45, 0x43, 0x4b, 0x53, 0x55, 0x4d, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xaa, 0x01, 0x12, 0x12, 0x0a, 0x0d, 0x53, 0x54, 0x4f, 0x52,
	0x41, 0x47, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xc9, 0x01, 0x12, 0x0d, 0x0a, 0x08,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0xac, 0x02, 0x12, 0x19, 0x0a, 0x14, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0xc0, 0x02, 0x12, 0x18, 0x0a, 0x13, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45,
	0x5f, 0x52, 0x55, 0x4e, 0x53, 0x51, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xcc, 0x02,
	0x12, 0x12, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45,
	0x44, 0x10, 0xd4, 0x02, 0x42, 0x82, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x63, 0x71,
	0x6c, 0x2e, 0x70, 0x62, 0x42, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
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
	file_api_status_code_proto_rawDescOnce sync.Once
	file_api_status_code_proto_rawDescData = file_api_status_code_proto_rawDesc
)

func file_api_status_code_proto_rawDescGZIP() []byte {
	file_api_status_code_proto_rawDescOnce.Do(func() {
		file_api_status_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_status_code_proto_rawDescData)
	})
	return file_api_status_code_proto_rawDescData
}

var file_api_status_code_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_status_code_proto_goTypes = []interface{}{
	(Code)(0), // 0: scql.pb.Code
}
var file_api_status_code_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_status_code_proto_init() }
func file_api_status_code_proto_init() {
	if File_api_status_code_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_status_code_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_status_code_proto_goTypes,
		DependencyIndexes: file_api_status_code_proto_depIdxs,
		EnumInfos:         file_api_status_code_proto_enumTypes,
	}.Build()
	File_api_status_code_proto = out.File
	file_api_status_code_proto_rawDesc = nil
	file_api_status_code_proto_goTypes = nil
	file_api_status_code_proto_depIdxs = nil
}
