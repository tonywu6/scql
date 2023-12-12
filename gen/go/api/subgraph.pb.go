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
// source: api/subgraph.proto

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

// SubGraph is the part of the whole execution graph seen
// from the perspective of the party.
// Each party could only see the ExecNode which it participates in.
type SubGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes  map[string]*ExecNode `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Policy *SchedulingPolicy    `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
	// checksum of the whole execution graph
	// It could be used to verify and ensure that engines execute the same graph.
	GraphChecksum []byte `protobuf:"bytes,3,opt,name=graph_checksum,json=graphChecksum,proto3" json:"graph_checksum,omitempty"`
}

func (x *SubGraph) Reset() {
	*x = SubGraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_subgraph_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubGraph) ProtoMessage() {}

func (x *SubGraph) ProtoReflect() protoreflect.Message {
	mi := &file_api_subgraph_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubGraph.ProtoReflect.Descriptor instead.
func (*SubGraph) Descriptor() ([]byte, []int) {
	return file_api_subgraph_proto_rawDescGZIP(), []int{0}
}

func (x *SubGraph) GetNodes() map[string]*ExecNode {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *SubGraph) GetPolicy() *SchedulingPolicy {
	if x != nil {
		return x.Policy
	}
	return nil
}

func (x *SubGraph) GetGraphChecksum() []byte {
	if x != nil {
		return x.GraphChecksum
	}
	return nil
}

type SubDAG struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs []*SubDAG_Job `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	// a barrier to sync among parties
	NeedCallBarrierAfterJobs bool `protobuf:"varint,2,opt,name=need_call_barrier_after_jobs,json=needCallBarrierAfterJobs,proto3" json:"need_call_barrier_after_jobs,omitempty"`
}

func (x *SubDAG) Reset() {
	*x = SubDAG{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_subgraph_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubDAG) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubDAG) ProtoMessage() {}

func (x *SubDAG) ProtoReflect() protoreflect.Message {
	mi := &file_api_subgraph_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubDAG.ProtoReflect.Descriptor instead.
func (*SubDAG) Descriptor() ([]byte, []int) {
	return file_api_subgraph_proto_rawDescGZIP(), []int{1}
}

func (x *SubDAG) GetJobs() []*SubDAG_Job {
	if x != nil {
		return x.Jobs
	}
	return nil
}

func (x *SubDAG) GetNeedCallBarrierAfterJobs() bool {
	if x != nil {
		return x.NeedCallBarrierAfterJobs
	}
	return false
}

type SchedulingPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerNum int32     `protobuf:"varint,1,opt,name=worker_num,json=workerNum,proto3" json:"worker_num,omitempty"`
	Subdags   []*SubDAG `protobuf:"bytes,2,rep,name=subdags,proto3" json:"subdags,omitempty"`
}

func (x *SchedulingPolicy) Reset() {
	*x = SchedulingPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_subgraph_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulingPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulingPolicy) ProtoMessage() {}

func (x *SchedulingPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_api_subgraph_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulingPolicy.ProtoReflect.Descriptor instead.
func (*SchedulingPolicy) Descriptor() ([]byte, []int) {
	return file_api_subgraph_proto_rawDescGZIP(), []int{2}
}

func (x *SchedulingPolicy) GetWorkerNum() int32 {
	if x != nil {
		return x.WorkerNum
	}
	return 0
}

func (x *SchedulingPolicy) GetSubdags() []*SubDAG {
	if x != nil {
		return x.Subdags
	}
	return nil
}

type SubDAG_Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerId int32    `protobuf:"varint,1,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	NodeIds  []string `protobuf:"bytes,2,rep,name=node_ids,json=nodeIds,proto3" json:"node_ids,omitempty"`
}

func (x *SubDAG_Job) Reset() {
	*x = SubDAG_Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_subgraph_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubDAG_Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubDAG_Job) ProtoMessage() {}

func (x *SubDAG_Job) ProtoReflect() protoreflect.Message {
	mi := &file_api_subgraph_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubDAG_Job.ProtoReflect.Descriptor instead.
func (*SubDAG_Job) Descriptor() ([]byte, []int) {
	return file_api_subgraph_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SubDAG_Job) GetWorkerId() int32 {
	if x != nil {
		return x.WorkerId
	}
	return 0
}

func (x *SubDAG_Job) GetNodeIds() []string {
	if x != nil {
		return x.NodeIds
	}
	return nil
}

var File_api_subgraph_proto protoreflect.FileDescriptor

var file_api_subgraph_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x62, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x1a, 0x0e, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01,
	0x0a, 0x08, 0x53, 0x75, 0x62, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x32, 0x0a, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x63, 0x71, 0x6c,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x47, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x31,
	0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x25, 0x0a, 0x0e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x1a, 0x4b, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70,
	0x62, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb0, 0x01, 0x0a, 0x06, 0x53, 0x75, 0x62, 0x44, 0x41, 0x47,
	0x12, 0x27, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x44, 0x41, 0x47, 0x2e,
	0x4a, 0x6f, 0x62, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x12, 0x3e, 0x0a, 0x1c, 0x6e, 0x65, 0x65,
	0x64, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x62, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x5f, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x18, 0x6e, 0x65, 0x65, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72,
	0x41, 0x66, 0x74, 0x65, 0x72, 0x4a, 0x6f, 0x62, 0x73, 0x1a, 0x3d, 0x0a, 0x03, 0x4a, 0x6f, 0x62,
	0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x73, 0x22, 0x5c, 0x0a, 0x10, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x29, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x64, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73,
	0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x44, 0x41, 0x47, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x64, 0x61, 0x67, 0x73, 0x42, 0x80, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x73,
	0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x42, 0x0d, 0x53, 0x75, 0x62, 0x67, 0x72, 0x61, 0x70, 0x68,
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
	file_api_subgraph_proto_rawDescOnce sync.Once
	file_api_subgraph_proto_rawDescData = file_api_subgraph_proto_rawDesc
)

func file_api_subgraph_proto_rawDescGZIP() []byte {
	file_api_subgraph_proto_rawDescOnce.Do(func() {
		file_api_subgraph_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_subgraph_proto_rawDescData)
	})
	return file_api_subgraph_proto_rawDescData
}

var file_api_subgraph_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_subgraph_proto_goTypes = []interface{}{
	(*SubGraph)(nil),         // 0: scql.pb.SubGraph
	(*SubDAG)(nil),           // 1: scql.pb.SubDAG
	(*SchedulingPolicy)(nil), // 2: scql.pb.SchedulingPolicy
	nil,                      // 3: scql.pb.SubGraph.NodesEntry
	(*SubDAG_Job)(nil),       // 4: scql.pb.SubDAG.Job
	(*ExecNode)(nil),         // 5: scql.pb.ExecNode
}
var file_api_subgraph_proto_depIdxs = []int32{
	3, // 0: scql.pb.SubGraph.nodes:type_name -> scql.pb.SubGraph.NodesEntry
	2, // 1: scql.pb.SubGraph.policy:type_name -> scql.pb.SchedulingPolicy
	4, // 2: scql.pb.SubDAG.jobs:type_name -> scql.pb.SubDAG.Job
	1, // 3: scql.pb.SchedulingPolicy.subdags:type_name -> scql.pb.SubDAG
	5, // 4: scql.pb.SubGraph.NodesEntry.value:type_name -> scql.pb.ExecNode
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_subgraph_proto_init() }
func file_api_subgraph_proto_init() {
	if File_api_subgraph_proto != nil {
		return
	}
	file_api_core_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_subgraph_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubGraph); i {
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
		file_api_subgraph_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubDAG); i {
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
		file_api_subgraph_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulingPolicy); i {
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
		file_api_subgraph_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubDAG_Job); i {
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
			RawDescriptor: file_api_subgraph_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_subgraph_proto_goTypes,
		DependencyIndexes: file_api_subgraph_proto_depIdxs,
		MessageInfos:      file_api_subgraph_proto_msgTypes,
	}.Build()
	File_api_subgraph_proto = out.File
	file_api_subgraph_proto_rawDesc = nil
	file_api_subgraph_proto_goTypes = nil
	file_api_subgraph_proto_depIdxs = nil
}
