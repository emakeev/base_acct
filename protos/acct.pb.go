//
// Copyright 2021 Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: acct.proto

package protos

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type QoS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DownloadMbps float32 `protobuf:"fixed32,1,opt,name=download_mbps,json=downloadMbps,proto3" json:"download_mbps,omitempty"`
	UploadMbps   float32 `protobuf:"fixed32,2,opt,name=upload_mbps,json=uploadMbps,proto3" json:"upload_mbps,omitempty"` // TBD
}

func (x *QoS) Reset() {
	*x = QoS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QoS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QoS) ProtoMessage() {}

func (x *QoS) ProtoReflect() protoreflect.Message {
	mi := &file_acct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QoS.ProtoReflect.Descriptor instead.
func (*QoS) Descriptor() ([]byte, []int) {
	return file_acct_proto_rawDescGZIP(), []int{0}
}

func (x *QoS) GetDownloadMbps() float32 {
	if x != nil {
		return x.DownloadMbps
	}
	return 0
}

func (x *QoS) GetUploadMbps() float32 {
	if x != nil {
		return x.UploadMbps
	}
	return 0
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to User:
	//	*Session_IMSI
	//	*Session_CertificateSerialNumber
	//	*Session_HardwareAddr
	//	*Session_Name
	User         isSession_User `protobuf_oneof:"user"`
	SessionId    string         `protobuf:"bytes,5,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	AvailableQos *QoS           `protobuf:"bytes,6,opt,name=available_qos,json=availableQos,proto3" json:"available_qos,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acct_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_acct_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_acct_proto_rawDescGZIP(), []int{1}
}

func (m *Session) GetUser() isSession_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (x *Session) GetIMSI() string {
	if x, ok := x.GetUser().(*Session_IMSI); ok {
		return x.IMSI
	}
	return ""
}

func (x *Session) GetCertificateSerialNumber() []byte {
	if x, ok := x.GetUser().(*Session_CertificateSerialNumber); ok {
		return x.CertificateSerialNumber
	}
	return nil
}

func (x *Session) GetHardwareAddr() []byte {
	if x, ok := x.GetUser().(*Session_HardwareAddr); ok {
		return x.HardwareAddr
	}
	return nil
}

func (x *Session) GetName() string {
	if x, ok := x.GetUser().(*Session_Name); ok {
		return x.Name
	}
	return ""
}

func (x *Session) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *Session) GetAvailableQos() *QoS {
	if x != nil {
		return x.AvailableQos
	}
	return nil
}

type isSession_User interface {
	isSession_User()
}

type Session_IMSI struct {
	IMSI string `protobuf:"bytes,1,opt,name=IMSI,proto3,oneof"`
}

type Session_CertificateSerialNumber struct {
	CertificateSerialNumber []byte `protobuf:"bytes,2,opt,name=certificate_serial_number,json=certificateSerialNumber,proto3,oneof"`
}

type Session_HardwareAddr struct {
	HardwareAddr []byte `protobuf:"bytes,3,opt,name=hardware_addr,json=hardwareAddr,proto3,oneof"`
}

type Session_Name struct {
	Name string `protobuf:"bytes,4,opt,name=name,proto3,oneof"`
}

func (*Session_IMSI) isSession_User() {}

func (*Session_CertificateSerialNumber) isSession_User() {}

func (*Session_HardwareAddr) isSession_User() {}

func (*Session_Name) isSession_User() {}

type SessionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReportingAdvisory *SessionRespReportLimits `protobuf:"bytes,1,opt,name=reporting_advisory,json=reportingAdvisory,proto3" json:"reporting_advisory,omitempty"`
	MinAcceptableQos  *QoS                     `protobuf:"bytes,2,opt,name=min_acceptable_qos,json=minAcceptableQos,proto3" json:"min_acceptable_qos,omitempty"`
}

func (x *SessionResp) Reset() {
	*x = SessionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acct_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionResp) ProtoMessage() {}

func (x *SessionResp) ProtoReflect() protoreflect.Message {
	mi := &file_acct_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionResp.ProtoReflect.Descriptor instead.
func (*SessionResp) Descriptor() ([]byte, []int) {
	return file_acct_proto_rawDescGZIP(), []int{2}
}

func (x *SessionResp) GetReportingAdvisory() *SessionRespReportLimits {
	if x != nil {
		return x.ReportingAdvisory
	}
	return nil
}

func (x *SessionResp) GetMinAcceptableQos() *QoS {
	if x != nil {
		return x.MinAcceptableQos
	}
	return nil
}

type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session   *Session `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	OctetsIn  uint32   `protobuf:"varint,2,opt,name=octets_in,json=octetsIn,proto3" json:"octets_in,omitempty"`
	OctetsOut uint32   `protobuf:"varint,3,opt,name=octets_out,json=octetsOut,proto3" json:"octets_out,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acct_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_acct_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_acct_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateReq) GetSession() *Session {
	if x != nil {
		return x.Session
	}
	return nil
}

func (x *UpdateReq) GetOctetsIn() uint32 {
	if x != nil {
		return x.OctetsIn
	}
	return 0
}

func (x *UpdateReq) GetOctetsOut() uint32 {
	if x != nil {
		return x.OctetsOut
	}
	return 0
}

type StopResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopResp) Reset() {
	*x = StopResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acct_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopResp) ProtoMessage() {}

func (x *StopResp) ProtoReflect() protoreflect.Message {
	mi := &file_acct_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopResp.ProtoReflect.Descriptor instead.
func (*StopResp) Descriptor() ([]byte, []int) {
	return file_acct_proto_rawDescGZIP(), []int{4}
}

type SessionRespReportLimits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OctetsIn       uint32 `protobuf:"varint,1,opt,name=octets_in,json=octetsIn,proto3" json:"octets_in,omitempty"`
	OctetsOut      uint32 `protobuf:"varint,2,opt,name=octets_out,json=octetsOut,proto3" json:"octets_out,omitempty"`
	ElapsedTimeSec uint32 `protobuf:"varint,3,opt,name=elapsed_time_sec,json=elapsedTimeSec,proto3" json:"elapsed_time_sec,omitempty"`
}

func (x *SessionRespReportLimits) Reset() {
	*x = SessionRespReportLimits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acct_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionRespReportLimits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionRespReportLimits) ProtoMessage() {}

func (x *SessionRespReportLimits) ProtoReflect() protoreflect.Message {
	mi := &file_acct_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionRespReportLimits.ProtoReflect.Descriptor instead.
func (*SessionRespReportLimits) Descriptor() ([]byte, []int) {
	return file_acct_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SessionRespReportLimits) GetOctetsIn() uint32 {
	if x != nil {
		return x.OctetsIn
	}
	return 0
}

func (x *SessionRespReportLimits) GetOctetsOut() uint32 {
	if x != nil {
		return x.OctetsOut
	}
	return 0
}

func (x *SessionRespReportLimits) GetElapsedTimeSec() uint32 {
	if x != nil {
		return x.ElapsedTimeSec
	}
	return 0
}

var File_acct_proto protoreflect.FileDescriptor

var file_acct_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x22, 0x4b, 0x0a, 0x03, 0x51, 0x6f, 0x53, 0x12, 0x23, 0x0a, 0x0d, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d, 0x62, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x62, 0x70, 0x73,
	0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d, 0x62, 0x70, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x62, 0x70,
	0x73, 0x22, 0xf3, 0x01, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x04, 0x49, 0x4d, 0x53, 0x49, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x49,
	0x4d, 0x53, 0x49, 0x12, 0x3c, 0x0a, 0x19, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x17, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x25, 0x0a, 0x0d, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0c, 0x68, 0x61, 0x72, 0x64,
	0x77, 0x61, 0x72, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x30, 0x0a,
	0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x71, 0x6f, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x51, 0x6f,
	0x53, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x51, 0x6f, 0x73, 0x42,
	0x06, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x93, 0x02, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x12, 0x51, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x11, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x69, 0x6e, 0x67, 0x41, 0x64, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x79, 0x12, 0x39, 0x0a, 0x12, 0x6d,
	0x69, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x71, 0x6f,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x51, 0x6f, 0x53, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x51, 0x6f, 0x73, 0x1a, 0x75, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x63, 0x74, 0x65, 0x74,
	0x73, 0x5f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x63, 0x74, 0x65,
	0x74, 0x73, 0x49, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x63, 0x74, 0x65, 0x74, 0x73, 0x5f, 0x6f,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6f, 0x63, 0x74, 0x65, 0x74, 0x73,
	0x4f, 0x75, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x65,
	0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x22, 0x73, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x07, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x63, 0x74, 0x65, 0x74, 0x73,
	0x5f, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x63, 0x74, 0x65, 0x74,
	0x73, 0x49, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x63, 0x74, 0x65, 0x74, 0x73, 0x5f, 0x6f, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6f, 0x63, 0x74, 0x65, 0x74, 0x73, 0x4f,
	0x75, 0x74, 0x22, 0x0b, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x32,
	0x9f, 0x01, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2e,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x12, 0x32,
	0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x12, 0x2d, 0x0a, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x1a, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6d, 0x61, 0x6b, 0x65, 0x65, 0x76, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x61, 0x63, 0x63,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_acct_proto_rawDescOnce sync.Once
	file_acct_proto_rawDescData = file_acct_proto_rawDesc
)

func file_acct_proto_rawDescGZIP() []byte {
	file_acct_proto_rawDescOnce.Do(func() {
		file_acct_proto_rawDescData = protoimpl.X.CompressGZIP(file_acct_proto_rawDescData)
	})
	return file_acct_proto_rawDescData
}

var file_acct_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_acct_proto_goTypes = []interface{}{
	(*QoS)(nil),                     // 0: protos.QoS
	(*Session)(nil),                 // 1: protos.session
	(*SessionResp)(nil),             // 2: protos.session_resp
	(*UpdateReq)(nil),               // 3: protos.update_req
	(*StopResp)(nil),                // 4: protos.stop_resp
	(*SessionRespReportLimits)(nil), // 5: protos.session_resp.report_limits
}
var file_acct_proto_depIdxs = []int32{
	0, // 0: protos.session.available_qos:type_name -> protos.QoS
	5, // 1: protos.session_resp.reporting_advisory:type_name -> protos.session_resp.report_limits
	0, // 2: protos.session_resp.min_acceptable_qos:type_name -> protos.QoS
	1, // 3: protos.update_req.session:type_name -> protos.session
	1, // 4: protos.accounting.start:input_type -> protos.session
	3, // 5: protos.accounting.update:input_type -> protos.update_req
	3, // 6: protos.accounting.stop:input_type -> protos.update_req
	2, // 7: protos.accounting.start:output_type -> protos.session_resp
	2, // 8: protos.accounting.update:output_type -> protos.session_resp
	4, // 9: protos.accounting.stop:output_type -> protos.stop_resp
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_acct_proto_init() }
func file_acct_proto_init() {
	if File_acct_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_acct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QoS); i {
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
		file_acct_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_acct_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionResp); i {
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
		file_acct_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_acct_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopResp); i {
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
		file_acct_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionRespReportLimits); i {
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
	file_acct_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Session_IMSI)(nil),
		(*Session_CertificateSerialNumber)(nil),
		(*Session_HardwareAddr)(nil),
		(*Session_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_acct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_acct_proto_goTypes,
		DependencyIndexes: file_acct_proto_depIdxs,
		MessageInfos:      file_acct_proto_msgTypes,
	}.Build()
	File_acct_proto = out.File
	file_acct_proto_rawDesc = nil
	file_acct_proto_goTypes = nil
	file_acct_proto_depIdxs = nil
}