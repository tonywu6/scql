//
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
// source: engine/datasource/datasource.proto

package datasource

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

type DataSourceKind int32

const (
	DataSourceKind_UNKNOWN    DataSourceKind = 0
	DataSourceKind_MYSQL      DataSourceKind = 1
	DataSourceKind_SQLITE     DataSourceKind = 2
	DataSourceKind_POSTGRESQL DataSourceKind = 3
	DataSourceKind_CSVDB      DataSourceKind = 4
	DataSourceKind_ARROWSQL   DataSourceKind = 5
	DataSourceKind_GRPC       DataSourceKind = 6
)

// Enum value maps for DataSourceKind.
var (
	DataSourceKind_name = map[int32]string{
		0: "UNKNOWN",
		1: "MYSQL",
		2: "SQLITE",
		3: "POSTGRESQL",
		4: "CSVDB",
		5: "ARROWSQL",
		6: "GRPC",
	}
	DataSourceKind_value = map[string]int32{
		"UNKNOWN":    0,
		"MYSQL":      1,
		"SQLITE":     2,
		"POSTGRESQL": 3,
		"CSVDB":      4,
		"ARROWSQL":   5,
		"GRPC":       6,
	}
)

func (x DataSourceKind) Enum() *DataSourceKind {
	p := new(DataSourceKind)
	*p = x
	return p
}

func (x DataSourceKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataSourceKind) Descriptor() protoreflect.EnumDescriptor {
	return file_engine_datasource_datasource_proto_enumTypes[0].Descriptor()
}

func (DataSourceKind) Type() protoreflect.EnumType {
	return &file_engine_datasource_datasource_proto_enumTypes[0]
}

func (x DataSourceKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataSourceKind.Descriptor instead.
func (DataSourceKind) EnumDescriptor() ([]byte, []int) {
	return file_engine_datasource_datasource_proto_rawDescGZIP(), []int{0}
}

type DataSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// datasource uuid
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// human-friendly datasource name
	Name string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Kind DataSourceKind `protobuf:"varint,3,opt,name=kind,proto3,enum=scql.engine.DataSourceKind" json:"kind,omitempty"`
	// concrete data source connection string
	// It is comprehend to related data source adaptor.
	ConnectionStr string `protobuf:"bytes,4,opt,name=connection_str,json=connectionStr,proto3" json:"connection_str,omitempty"`
}

func (x *DataSource) Reset() {
	*x = DataSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_datasource_datasource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSource) ProtoMessage() {}

func (x *DataSource) ProtoReflect() protoreflect.Message {
	mi := &file_engine_datasource_datasource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSource.ProtoReflect.Descriptor instead.
func (*DataSource) Descriptor() ([]byte, []int) {
	return file_engine_datasource_datasource_proto_rawDescGZIP(), []int{0}
}

func (x *DataSource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataSource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataSource) GetKind() DataSourceKind {
	if x != nil {
		return x.Kind
	}
	return DataSourceKind_UNKNOWN
}

func (x *DataSource) GetConnectionStr() string {
	if x != nil {
		return x.ConnectionStr
	}
	return ""
}

var File_engine_datasource_datasource_proto protoreflect.FileDescriptor

var file_engine_datasource_datasource_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x22, 0x88, 0x01, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x2a, 0x67, 0x0a, 0x0e,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0b,
	0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4d,
	0x59, 0x53, 0x51, 0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x51, 0x4c, 0x49, 0x54, 0x45,
	0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x52, 0x45, 0x53, 0x51, 0x4c,
	0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x53, 0x56, 0x44, 0x42, 0x10, 0x04, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x52, 0x52, 0x4f, 0x57, 0x53, 0x51, 0x4c, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x47,
	0x52, 0x50, 0x43, 0x10, 0x06, 0x42, 0xa4, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x63,
	0x71, 0x6c, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x42, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x73, 0x63, 0x71, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0xa2, 0x02, 0x03, 0x53, 0x45, 0x58, 0xaa, 0x02, 0x0b, 0x53, 0x63, 0x71, 0x6c, 0x2e, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0xca, 0x02, 0x0b, 0x53, 0x63, 0x71, 0x6c, 0x5c, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0xe2, 0x02, 0x17, 0x53, 0x63, 0x71, 0x6c, 0x5c, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c,
	0x53, 0x63, 0x71, 0x6c, 0x3a, 0x3a, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_engine_datasource_datasource_proto_rawDescOnce sync.Once
	file_engine_datasource_datasource_proto_rawDescData = file_engine_datasource_datasource_proto_rawDesc
)

func file_engine_datasource_datasource_proto_rawDescGZIP() []byte {
	file_engine_datasource_datasource_proto_rawDescOnce.Do(func() {
		file_engine_datasource_datasource_proto_rawDescData = protoimpl.X.CompressGZIP(file_engine_datasource_datasource_proto_rawDescData)
	})
	return file_engine_datasource_datasource_proto_rawDescData
}

var file_engine_datasource_datasource_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_engine_datasource_datasource_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_engine_datasource_datasource_proto_goTypes = []interface{}{
	(DataSourceKind)(0), // 0: scql.engine.DataSourceKind
	(*DataSource)(nil),  // 1: scql.engine.DataSource
}
var file_engine_datasource_datasource_proto_depIdxs = []int32{
	0, // 0: scql.engine.DataSource.kind:type_name -> scql.engine.DataSourceKind
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_engine_datasource_datasource_proto_init() }
func file_engine_datasource_datasource_proto_init() {
	if File_engine_datasource_datasource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_engine_datasource_datasource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSource); i {
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
			RawDescriptor: file_engine_datasource_datasource_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_engine_datasource_datasource_proto_goTypes,
		DependencyIndexes: file_engine_datasource_datasource_proto_depIdxs,
		EnumInfos:         file_engine_datasource_datasource_proto_enumTypes,
		MessageInfos:      file_engine_datasource_datasource_proto_msgTypes,
	}.Build()
	File_engine_datasource_datasource_proto = out.File
	file_engine_datasource_datasource_proto_rawDesc = nil
	file_engine_datasource_datasource_proto_goTypes = nil
	file_engine_datasource_datasource_proto_depIdxs = nil
}
