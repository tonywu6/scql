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
// source: api/engine.proto

package api

import (
	libspu "github.com/secretflow/scql/gen/go/libspu"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StopSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// session stop reason, it maybe
	// - "Finish". Session ended normally.
	// - "Timeout". Session timeout.
	// - "Canceled". Session canceled by client, maybe triggered by CTRL+C
	// - "Error". Exception caught when running.
	// - "Killed". Killed by client.
	// - or something else.
	Reason string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *StopSessionRequest) Reset() {
	*x = StopSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopSessionRequest) ProtoMessage() {}

func (x *StopSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopSessionRequest.ProtoReflect.Descriptor instead.
func (*StopSessionRequest) Descriptor() ([]byte, []int) {
	return file_api_engine_proto_rawDescGZIP(), []int{0}
}

func (x *StopSessionRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *StopSessionRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

// A message carries all session-level (execution plan level) information.
type SessionStartParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This party code.
	// It may be used to get this party information from <parties> below.
	PartyCode string `protobuf:"bytes,1,opt,name=party_code,json=partyCode,proto3" json:"party_code,omitempty"`
	// All parties that would jointly complete an execution plan.
	Parties []*SessionStartParams_Party `protobuf:"bytes,2,rep,name=parties,proto3" json:"parties,omitempty"`
	// The session id
	SessionId string `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// The spu runtime configuration.
	SpuRuntimeCfg *libspu.RuntimeConfig `protobuf:"bytes,4,opt,name=spu_runtime_cfg,json=spuRuntimeCfg,proto3" json:"spu_runtime_cfg,omitempty"`
	// The session time zone, only support time offset, like: '+08:00'
	TimeZone string `protobuf:"bytes,5,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
}

func (x *SessionStartParams) Reset() {
	*x = SessionStartParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionStartParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionStartParams) ProtoMessage() {}

func (x *SessionStartParams) ProtoReflect() protoreflect.Message {
	mi := &file_api_engine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionStartParams.ProtoReflect.Descriptor instead.
func (*SessionStartParams) Descriptor() ([]byte, []int) {
	return file_api_engine_proto_rawDescGZIP(), []int{1}
}

func (x *SessionStartParams) GetPartyCode() string {
	if x != nil {
		return x.PartyCode
	}
	return ""
}

func (x *SessionStartParams) GetParties() []*SessionStartParams_Party {
	if x != nil {
		return x.Parties
	}
	return nil
}

func (x *SessionStartParams) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *SessionStartParams) GetSpuRuntimeCfg() *libspu.RuntimeConfig {
	if x != nil {
		return x.SpuRuntimeCfg
	}
	return nil
}

func (x *SessionStartParams) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

type RunExecutionPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionParams *SessionStartParams `protobuf:"bytes,1,opt,name=session_params,json=sessionParams,proto3" json:"session_params,omitempty"`
	Graph         *SubGraph           `protobuf:"bytes,2,opt,name=graph,proto3" json:"graph,omitempty"`
	// Whether the whole execution plan on the engine should be executed
	// synchronously or asynchronously. By default, the execution plan is executed
	// synchronously.
	Async bool `protobuf:"varint,3,opt,name=async,proto3" json:"async,omitempty"`
	// Callback url, e.g.: "http://alice.com:8080/path/to/callback".
	CallbackUrl string `protobuf:"bytes,4,opt,name=callback_url,json=callbackUrl,proto3" json:"callback_url,omitempty"`
}

func (x *RunExecutionPlanRequest) Reset() {
	*x = RunExecutionPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunExecutionPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunExecutionPlanRequest) ProtoMessage() {}

func (x *RunExecutionPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunExecutionPlanRequest.ProtoReflect.Descriptor instead.
func (*RunExecutionPlanRequest) Descriptor() ([]byte, []int) {
	return file_api_engine_proto_rawDescGZIP(), []int{2}
}

func (x *RunExecutionPlanRequest) GetSessionParams() *SessionStartParams {
	if x != nil {
		return x.SessionParams
	}
	return nil
}

func (x *RunExecutionPlanRequest) GetGraph() *SubGraph {
	if x != nil {
		return x.Graph
	}
	return nil
}

func (x *RunExecutionPlanRequest) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

func (x *RunExecutionPlanRequest) GetCallbackUrl() string {
	if x != nil {
		return x.CallbackUrl
	}
	return ""
}

type RunExecutionPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// Output columns used to store result datas.
	OutColumns []*Tensor `protobuf:"bytes,2,rep,name=out_columns,json=outColumns,proto3" json:"out_columns,omitempty"`
	SessionId  string    `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// Code of party which finished the execution plan.
	PartyCode string `protobuf:"bytes,4,opt,name=party_code,json=partyCode,proto3" json:"party_code,omitempty"`
	// The number of rows affected by a select into, update, insert, or delete.
	NumRowsAffected int64 `protobuf:"varint,5,opt,name=num_rows_affected,json=numRowsAffected,proto3" json:"num_rows_affected,omitempty"`
}

func (x *RunExecutionPlanResponse) Reset() {
	*x = RunExecutionPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_engine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunExecutionPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunExecutionPlanResponse) ProtoMessage() {}

func (x *RunExecutionPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_engine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunExecutionPlanResponse.ProtoReflect.Descriptor instead.
func (*RunExecutionPlanResponse) Descriptor() ([]byte, []int) {
	return file_api_engine_proto_rawDescGZIP(), []int{3}
}

func (x *RunExecutionPlanResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *RunExecutionPlanResponse) GetOutColumns() []*Tensor {
	if x != nil {
		return x.OutColumns
	}
	return nil
}

func (x *RunExecutionPlanResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *RunExecutionPlanResponse) GetPartyCode() string {
	if x != nil {
		return x.PartyCode
	}
	return ""
}

func (x *RunExecutionPlanResponse) GetNumRowsAffected() int64 {
	if x != nil {
		return x.NumRowsAffected
	}
	return 0
}

// After finishing running an execution plan, the Engine sends a ReportRequest
// to result callback server (such as scdb/broker/...) to return the results.
type ReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status including whether the plan run was handled successfully,
	// if not, error message and some more information included.
	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// Output columns.
	OutColumns []*Tensor `protobuf:"bytes,2,rep,name=out_columns,json=outColumns,proto3" json:"out_columns,omitempty"`
	// The session_id that this execution plan belongs to.
	SessionId string `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// Code of party which finished the plan run.
	PartyCode string `protobuf:"bytes,4,opt,name=party_code,json=partyCode,proto3" json:"party_code,omitempty"`
	// The number of rows affected by a select into, update, insert, or delete.
	NumRowsAffected int64 `protobuf:"varint,5,opt,name=num_rows_affected,json=numRowsAffected,proto3" json:"num_rows_affected,omitempty"`
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_engine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_engine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRequest.ProtoReflect.Descriptor instead.
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return file_api_engine_proto_rawDescGZIP(), []int{4}
}

func (x *ReportRequest) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ReportRequest) GetOutColumns() []*Tensor {
	if x != nil {
		return x.OutColumns
	}
	return nil
}

func (x *ReportRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *ReportRequest) GetPartyCode() string {
	if x != nil {
		return x.PartyCode
	}
	return ""
}

func (x *ReportRequest) GetNumRowsAffected() int64 {
	if x != nil {
		return x.NumRowsAffected
	}
	return 0
}

type SessionStartParams_Party struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// party code
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// party name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// party host
	Host string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	// party rank
	Rank int32 `protobuf:"varint,4,opt,name=rank,proto3" json:"rank,omitempty"`
	// base64 encoded version of the DER encoded public key
	//
	// Example:
	// # 1. generate private key
	// $ openssl genpkey -algorithm ed25519 -out ed25519key.pem
	// # 2. generate public key based on above private key
	// $ openssl pkey -in ed25519key.pem -pubout
	//
	// # its output like below:
	//
	// -----BEGIN PUBLIC KEY-----
	// BASE64 ENCODED DATA
	// -----END PUBLIC KEY-----
	//
	// the field `public_key` should be the string between header "-----BEGIN
	// PUBLIC KEY-----" and footer "-----END PUBLIC KEY-----"
	PublicKey string `protobuf:"bytes,5,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *SessionStartParams_Party) Reset() {
	*x = SessionStartParams_Party{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_engine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionStartParams_Party) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionStartParams_Party) ProtoMessage() {}

func (x *SessionStartParams_Party) ProtoReflect() protoreflect.Message {
	mi := &file_api_engine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionStartParams_Party.ProtoReflect.Descriptor instead.
func (*SessionStartParams_Party) Descriptor() ([]byte, []int) {
	return file_api_engine_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SessionStartParams_Party) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SessionStartParams_Party) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SessionStartParams_Party) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *SessionStartParams_Party) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *SessionStartParams_Party) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

var File_api_engine_proto protoreflect.FileDescriptor

var file_api_engine_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6c, 0x69, 0x62, 0x73, 0x70, 0x75,
	0x2f, 0x73, 0x70, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x75, 0x62, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22,
	0xe0, 0x02, 0x0a, 0x12, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x3a, 0x0a, 0x0f, 0x73, 0x70, 0x75, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x63, 0x66, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x70, 0x75,
	0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d,
	0x73, 0x70, 0x75, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x66, 0x67, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x1a, 0x76, 0x0a, 0x05, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72,
	0x61, 0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x22, 0xbf, 0x01, 0x0a, 0x17, 0x52, 0x75, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42,
	0x0a, 0x0e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x52, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x52, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x73, 0x79, 0x6e, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x73, 0x79, 0x6e,
	0x63, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x55, 0x72, 0x6c, 0x22, 0xdf, 0x01, 0x0a, 0x18, 0x52, 0x75, 0x6e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a, 0x0b, 0x6f, 0x75,
	0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x52, 0x0a, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x72, 0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x75,
	0x6d, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x5f, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x52, 0x6f, 0x77, 0x73, 0x41, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0xd4, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x30, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62,
	0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x5f, 0x61, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6e, 0x75,
	0x6d, 0x52, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x32, 0xa9, 0x01,
	0x0a, 0x11, 0x53, 0x43, 0x51, 0x4c, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x52, 0x75, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x20, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x75, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x63, 0x71, 0x6c,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x75, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0b,
	0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x73, 0x63,
	0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x50, 0x0a, 0x14, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x12, 0x38, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x73, 0x63,
	0x71, 0x6c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x81, 0x01, 0x0a, 0x0b,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x63, 0x71, 0x6c, 0x2e, 0x70, 0x62, 0x42, 0x0b, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x73, 0x63, 0x71, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x80, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x53, 0x50, 0x58, 0xaa, 0x02, 0x07, 0x53, 0x63, 0x71,
	0x6c, 0x2e, 0x50, 0x62, 0xca, 0x02, 0x07, 0x53, 0x63, 0x71, 0x6c, 0x5c, 0x50, 0x62, 0xe2, 0x02,
	0x13, 0x53, 0x63, 0x71, 0x6c, 0x5c, 0x50, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x53, 0x63, 0x71, 0x6c, 0x3a, 0x3a, 0x50, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_engine_proto_rawDescOnce sync.Once
	file_api_engine_proto_rawDescData = file_api_engine_proto_rawDesc
)

func file_api_engine_proto_rawDescGZIP() []byte {
	file_api_engine_proto_rawDescOnce.Do(func() {
		file_api_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_engine_proto_rawDescData)
	})
	return file_api_engine_proto_rawDescData
}

var file_api_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_engine_proto_goTypes = []interface{}{
	(*StopSessionRequest)(nil),       // 0: scql.pb.StopSessionRequest
	(*SessionStartParams)(nil),       // 1: scql.pb.SessionStartParams
	(*RunExecutionPlanRequest)(nil),  // 2: scql.pb.RunExecutionPlanRequest
	(*RunExecutionPlanResponse)(nil), // 3: scql.pb.RunExecutionPlanResponse
	(*ReportRequest)(nil),            // 4: scql.pb.ReportRequest
	(*SessionStartParams_Party)(nil), // 5: scql.pb.SessionStartParams.Party
	(*libspu.RuntimeConfig)(nil),     // 6: spu.RuntimeConfig
	(*SubGraph)(nil),                 // 7: scql.pb.SubGraph
	(*Status)(nil),                   // 8: scql.pb.Status
	(*Tensor)(nil),                   // 9: scql.pb.Tensor
	(*emptypb.Empty)(nil),            // 10: google.protobuf.Empty
}
var file_api_engine_proto_depIdxs = []int32{
	5,  // 0: scql.pb.SessionStartParams.parties:type_name -> scql.pb.SessionStartParams.Party
	6,  // 1: scql.pb.SessionStartParams.spu_runtime_cfg:type_name -> spu.RuntimeConfig
	1,  // 2: scql.pb.RunExecutionPlanRequest.session_params:type_name -> scql.pb.SessionStartParams
	7,  // 3: scql.pb.RunExecutionPlanRequest.graph:type_name -> scql.pb.SubGraph
	8,  // 4: scql.pb.RunExecutionPlanResponse.status:type_name -> scql.pb.Status
	9,  // 5: scql.pb.RunExecutionPlanResponse.out_columns:type_name -> scql.pb.Tensor
	8,  // 6: scql.pb.ReportRequest.status:type_name -> scql.pb.Status
	9,  // 7: scql.pb.ReportRequest.out_columns:type_name -> scql.pb.Tensor
	2,  // 8: scql.pb.SCQLEngineService.RunExecutionPlan:input_type -> scql.pb.RunExecutionPlanRequest
	0,  // 9: scql.pb.SCQLEngineService.StopSession:input_type -> scql.pb.StopSessionRequest
	4,  // 10: scql.pb.EngineResultCallback.Report:input_type -> scql.pb.ReportRequest
	3,  // 11: scql.pb.SCQLEngineService.RunExecutionPlan:output_type -> scql.pb.RunExecutionPlanResponse
	8,  // 12: scql.pb.SCQLEngineService.StopSession:output_type -> scql.pb.Status
	10, // 13: scql.pb.EngineResultCallback.Report:output_type -> google.protobuf.Empty
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_engine_proto_init() }
func file_api_engine_proto_init() {
	if File_api_engine_proto != nil {
		return
	}
	file_api_common_proto_init()
	file_api_core_proto_init()
	file_api_status_proto_init()
	file_api_subgraph_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopSessionRequest); i {
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
		file_api_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionStartParams); i {
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
		file_api_engine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunExecutionPlanRequest); i {
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
		file_api_engine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunExecutionPlanResponse); i {
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
		file_api_engine_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRequest); i {
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
		file_api_engine_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionStartParams_Party); i {
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
			RawDescriptor: file_api_engine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_engine_proto_goTypes,
		DependencyIndexes: file_api_engine_proto_depIdxs,
		MessageInfos:      file_api_engine_proto_msgTypes,
	}.Build()
	File_api_engine_proto = out.File
	file_api_engine_proto_rawDesc = nil
	file_api_engine_proto_goTypes = nil
	file_api_engine_proto_depIdxs = nil
}
