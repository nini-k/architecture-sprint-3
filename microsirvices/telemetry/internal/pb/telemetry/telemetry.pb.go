// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: telemetry/telemetry.proto

package telemetry

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListLatestDeviceTelemetryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PageSize int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Cursor   *Cursor `protobuf:"bytes,3,opt,name=cursor,proto3,oneof" json:"cursor,omitempty"`
}

func (x *ListLatestDeviceTelemetryRequest) Reset() {
	*x = ListLatestDeviceTelemetryRequest{}
	mi := &file_telemetry_telemetry_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLatestDeviceTelemetryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLatestDeviceTelemetryRequest) ProtoMessage() {}

func (x *ListLatestDeviceTelemetryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_telemetry_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLatestDeviceTelemetryRequest.ProtoReflect.Descriptor instead.
func (*ListLatestDeviceTelemetryRequest) Descriptor() ([]byte, []int) {
	return file_telemetry_telemetry_proto_rawDescGZIP(), []int{0}
}

func (x *ListLatestDeviceTelemetryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListLatestDeviceTelemetryRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListLatestDeviceTelemetryRequest) GetCursor() *Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

type ListLatestDeviceTelemetryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*TelemetryData `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *ListLatestDeviceTelemetryResponse) Reset() {
	*x = ListLatestDeviceTelemetryResponse{}
	mi := &file_telemetry_telemetry_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListLatestDeviceTelemetryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLatestDeviceTelemetryResponse) ProtoMessage() {}

func (x *ListLatestDeviceTelemetryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_telemetry_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLatestDeviceTelemetryResponse.ProtoReflect.Descriptor instead.
func (*ListLatestDeviceTelemetryResponse) Descriptor() ([]byte, []int) {
	return file_telemetry_telemetry_proto_rawDescGZIP(), []int{1}
}

func (x *ListLatestDeviceTelemetryResponse) GetRecords() []*TelemetryData {
	if x != nil {
		return x.Records
	}
	return nil
}

type ListDeviceTelemetryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ListDeviceTelemetryRequest) Reset() {
	*x = ListDeviceTelemetryRequest{}
	mi := &file_telemetry_telemetry_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDeviceTelemetryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeviceTelemetryRequest) ProtoMessage() {}

func (x *ListDeviceTelemetryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_telemetry_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeviceTelemetryRequest.ProtoReflect.Descriptor instead.
func (*ListDeviceTelemetryRequest) Descriptor() ([]byte, []int) {
	return file_telemetry_telemetry_proto_rawDescGZIP(), []int{2}
}

func (x *ListDeviceTelemetryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListDeviceTelemetryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records    []*TelemetryData `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	NextCursor *Cursor          `protobuf:"bytes,2,opt,name=next_cursor,json=nextCursor,proto3" json:"next_cursor,omitempty"`
}

func (x *ListDeviceTelemetryResponse) Reset() {
	*x = ListDeviceTelemetryResponse{}
	mi := &file_telemetry_telemetry_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDeviceTelemetryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeviceTelemetryResponse) ProtoMessage() {}

func (x *ListDeviceTelemetryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_telemetry_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeviceTelemetryResponse.ProtoReflect.Descriptor instead.
func (*ListDeviceTelemetryResponse) Descriptor() ([]byte, []int) {
	return file_telemetry_telemetry_proto_rawDescGZIP(), []int{3}
}

func (x *ListDeviceTelemetryResponse) GetRecords() []*TelemetryData {
	if x != nil {
		return x.Records
	}
	return nil
}

func (x *ListDeviceTelemetryResponse) GetNextCursor() *Cursor {
	if x != nil {
		return x.NextCursor
	}
	return nil
}

type Cursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Entity:
	//
	//	*Cursor_Id
	Entity isCursor_Entity `protobuf_oneof:"entity"`
}

func (x *Cursor) Reset() {
	*x = Cursor{}
	mi := &file_telemetry_telemetry_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cursor) ProtoMessage() {}

func (x *Cursor) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_telemetry_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cursor.ProtoReflect.Descriptor instead.
func (*Cursor) Descriptor() ([]byte, []int) {
	return file_telemetry_telemetry_proto_rawDescGZIP(), []int{4}
}

func (m *Cursor) GetEntity() isCursor_Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (x *Cursor) GetId() int64 {
	if x, ok := x.GetEntity().(*Cursor_Id); ok {
		return x.Id
	}
	return 0
}

type isCursor_Entity interface {
	isCursor_Entity()
}

type Cursor_Id struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3,oneof"`
}

func (*Cursor_Id) isCursor_Entity() {}

type TelemetryData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Types that are assignable to Data:
	//
	//	*TelemetryData_TemperatureTelemetryData
	Data isTelemetryData_Data `protobuf_oneof:"data"`
}

func (x *TelemetryData) Reset() {
	*x = TelemetryData{}
	mi := &file_telemetry_telemetry_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TelemetryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemetryData) ProtoMessage() {}

func (x *TelemetryData) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_telemetry_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemetryData.ProtoReflect.Descriptor instead.
func (*TelemetryData) Descriptor() ([]byte, []int) {
	return file_telemetry_telemetry_proto_rawDescGZIP(), []int{5}
}

func (x *TelemetryData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TelemetryData) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (m *TelemetryData) GetData() isTelemetryData_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *TelemetryData) GetTemperatureTelemetryData() *TemperatureTelemetryData {
	if x, ok := x.GetData().(*TelemetryData_TemperatureTelemetryData); ok {
		return x.TemperatureTelemetryData
	}
	return nil
}

type isTelemetryData_Data interface {
	isTelemetryData_Data()
}

type TelemetryData_TemperatureTelemetryData struct {
	TemperatureTelemetryData *TemperatureTelemetryData `protobuf:"bytes,3,opt,name=temperature_telemetry_data,json=temperatureTelemetryData,proto3,oneof"`
}

func (*TelemetryData_TemperatureTelemetryData) isTelemetryData_Data() {}

type TemperatureTelemetryData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature int32 `protobuf:"varint,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
}

func (x *TemperatureTelemetryData) Reset() {
	*x = TemperatureTelemetryData{}
	mi := &file_telemetry_telemetry_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TemperatureTelemetryData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemperatureTelemetryData) ProtoMessage() {}

func (x *TemperatureTelemetryData) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_telemetry_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemperatureTelemetryData.ProtoReflect.Descriptor instead.
func (*TemperatureTelemetryData) Descriptor() ([]byte, []int) {
	return file_telemetry_telemetry_proto_rawDescGZIP(), []int{6}
}

func (x *TemperatureTelemetryData) GetTemperature() int32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

var File_telemetry_telemetry_proto protoreflect.FileDescriptor

var file_telemetry_telemetry_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x06, 0x63, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x3a, 0x16, 0x92, 0x41, 0x13, 0x0a, 0x11, 0xd2, 0x01,
	0x02, 0x69, 0x64, 0xd2, 0x01, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x57, 0x0a, 0x21, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x22, 0x38, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x3a, 0x0a, 0x92, 0x41, 0x07, 0x0a, 0x05, 0xd2, 0x01, 0x02, 0x69, 0x64, 0x22, 0x85, 0x01,
	0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x32, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x43,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x24, 0x0a, 0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12,
	0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x42, 0x08, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0xc7, 0x01, 0x0a, 0x0d,
	0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x63, 0x0a, 0x1a, 0x74, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x48, 0x00, 0x52, 0x18, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3c, 0x0a, 0x18, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x32, 0xdd, 0x05, 0x0a, 0x10, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xf8, 0x02, 0x0a, 0x19, 0x4c, 0x69, 0x73,
	0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x2b, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xff, 0x01, 0x92, 0x41, 0xd5, 0x01, 0x12, 0x78, 0xd0, 0x9f, 0xd0, 0xbe, 0xd0, 0xbb,
	0xd1, 0x83, 0xd1, 0x87, 0xd0, 0xb8, 0xd1, 0x82, 0xd1, 0x8c, 0x20, 0xd0, 0xbf, 0xd0, 0xbe, 0xd1,
	0x81, 0xd0, 0xbb, 0xd0, 0xb5, 0xd0, 0xb4, 0xd0, 0xbd, 0xd0, 0xb8, 0xd0, 0xb5, 0x20, 0xd1, 0x82,
	0xd0, 0xb5, 0xd0, 0xbb, 0xd0, 0xb5, 0xd0, 0xbc, 0xd0, 0xb5, 0xd1, 0x82, 0xd1, 0x80, 0xd0, 0xb8,
	0xd1, 0x87, 0xd0, 0xb5, 0xd1, 0x81, 0xd0, 0xba, 0xd0, 0xb8, 0xd0, 0xb5, 0x20, 0xd0, 0xb4, 0xd0,
	0xb0, 0xd0, 0xbd, 0xd0, 0xbd, 0xd1, 0x8b, 0xd0, 0xb5, 0x20, 0xd1, 0x81, 0x20, 0xd0, 0xb4, 0xd0,
	0xb0, 0xd1, 0x82, 0xd1, 0x87, 0xd0, 0xb8, 0xd0, 0xba, 0xd0, 0xbe, 0xd0, 0xb2, 0x20, 0xd1, 0x83,
	0xd1, 0x81, 0xd1, 0x82, 0xd1, 0x80, 0xd0, 0xbe, 0xd0, 0xb9, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb2,
	0xd0, 0xb0, 0x4a, 0x59, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x52, 0x0a, 0x33, 0xd0, 0x9f, 0xd1,
	0x80, 0xd0, 0xb8, 0x20, 0xd0, 0xbd, 0xd0, 0xb5, 0xd0, 0xba, 0xd0, 0xbe, 0xd1, 0x80, 0xd1, 0x80,
	0xd0, 0xb5, 0xd1, 0x82, 0xd0, 0xbd, 0xd1, 0x8b, 0xd1, 0x85, 0x20, 0xd0, 0xb7, 0xd0, 0xbd, 0xd0,
	0xb0, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8, 0xd1, 0x8f, 0xd1, 0x85, 0x20, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x12, 0xcd, 0x02, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x25, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x26, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe6, 0x01, 0x92, 0x41, 0xc3,
	0x01, 0x12, 0x55, 0xd0, 0x9f, 0xd0, 0xbe, 0xd0, 0xbb, 0xd1, 0x83, 0xd1, 0x87, 0xd0, 0xb8, 0xd1,
	0x82, 0xd1, 0x8c, 0x20, 0xd0, 0xb8, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xbe, 0xd1, 0x80, 0xd0, 0xb8,
	0xd1, 0x87, 0xd0, 0xb5, 0xd1, 0x81, 0xd0, 0xba, 0xd0, 0xb8, 0xd0, 0xb5, 0x20, 0xd1, 0x82, 0xd0,
	0xb5, 0xd0, 0xbb, 0xd0, 0xb5, 0xd0, 0xbc, 0xd0, 0xb5, 0xd1, 0x82, 0xd1, 0x80, 0xd0, 0xb8, 0xd1,
	0x87, 0xd0, 0xb5, 0xd1, 0x81, 0xd0, 0xba, 0xd0, 0xb8, 0xd0, 0xb5, 0x20, 0xd0, 0xb4, 0xd0, 0xb0,
	0xd0, 0xbd, 0xd0, 0xbd, 0xd1, 0x8b, 0xd0, 0xb5, 0x4a, 0x6a, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12,
	0x63, 0x0a, 0x44, 0xd0, 0x9f, 0xd1, 0x80, 0xd0, 0xb8, 0x20, 0xd0, 0xbd, 0xd0, 0xb5, 0xd0, 0xba,
	0xd0, 0xbe, 0xd1, 0x80, 0xd1, 0x80, 0xd0, 0xb5, 0xd1, 0x82, 0xd0, 0xbd, 0xd1, 0x8b, 0xd1, 0x85,
	0x20, 0xd0, 0xb7, 0xd0, 0xbd, 0xd0, 0xb0, 0xd1, 0x87, 0xd0, 0xb5, 0xd0, 0xbd, 0xd0, 0xb8, 0xd1,
	0x8f, 0xd1, 0x85, 0x20, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x20, 0xd0, 0xb8,
	0xd0, 0xbb, 0xd0, 0xb8, 0x20, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x42, 0xd5, 0x02, 0x92, 0x41, 0x90, 0x02, 0x12, 0xe6, 0x01, 0x0a, 0x0d, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x20, 0x41, 0x50, 0x49, 0x12, 0x6a, 0xd0, 0x9c,
	0xd0, 0xb8, 0xd0, 0xba, 0xd1, 0x80, 0xd0, 0xbe, 0xd1, 0x81, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb2,
	0xd0, 0xb8, 0xd1, 0x81, 0x20, 0xd0, 0xb4, 0xd0, 0xbb, 0xd1, 0x8f, 0x20, 0xd1, 0x81, 0xd0, 0xb1,
	0xd0, 0xbe, 0xd1, 0x80, 0xd0, 0xb0, 0x20, 0xd1, 0x82, 0xd0, 0xb5, 0xd0, 0xbb, 0xd0, 0xb5, 0xd0,
	0xbc, 0xd0, 0xb5, 0xd1, 0x82, 0xd1, 0x80, 0xd0, 0xb8, 0xd1, 0x87, 0xd0, 0xb5, 0xd1, 0x81, 0xd0,
	0xba, 0xd0, 0xb8, 0xd1, 0x85, 0x20, 0xd0, 0xb4, 0xd0, 0xb0, 0xd0, 0xbd, 0xd0, 0xbd, 0xd1, 0x8b,
	0xd1, 0x85, 0x20, 0xd1, 0x81, 0x20, 0xd1, 0x83, 0xd1, 0x81, 0xd1, 0x82, 0xd1, 0x80, 0xd0, 0xbe,
	0xd0, 0xb9, 0xd1, 0x81, 0xd1, 0x82, 0xd0, 0xb2, 0x22, 0x62, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x3e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x6e, 0x69, 0x2d, 0x6b, 0x2f, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x73, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x2d, 0x33, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x69, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x74, 0x72, 0x79, 0x1a, 0x10, 0x6e, 0x6f, 0x6e, 0x65,
	0x40, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x00, 0x32, 0x03,
	0x31, 0x2e, 0x30, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x6e, 0x69, 0x2d, 0x6b, 0x2f, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x73, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x2d, 0x33, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x69, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_telemetry_telemetry_proto_rawDescOnce sync.Once
	file_telemetry_telemetry_proto_rawDescData = file_telemetry_telemetry_proto_rawDesc
)

func file_telemetry_telemetry_proto_rawDescGZIP() []byte {
	file_telemetry_telemetry_proto_rawDescOnce.Do(func() {
		file_telemetry_telemetry_proto_rawDescData = protoimpl.X.CompressGZIP(file_telemetry_telemetry_proto_rawDescData)
	})
	return file_telemetry_telemetry_proto_rawDescData
}

var file_telemetry_telemetry_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_telemetry_telemetry_proto_goTypes = []any{
	(*ListLatestDeviceTelemetryRequest)(nil),  // 0: telemetry.ListLatestDeviceTelemetryRequest
	(*ListLatestDeviceTelemetryResponse)(nil), // 1: telemetry.ListLatestDeviceTelemetryResponse
	(*ListDeviceTelemetryRequest)(nil),        // 2: telemetry.ListDeviceTelemetryRequest
	(*ListDeviceTelemetryResponse)(nil),       // 3: telemetry.ListDeviceTelemetryResponse
	(*Cursor)(nil),                            // 4: telemetry.Cursor
	(*TelemetryData)(nil),                     // 5: telemetry.TelemetryData
	(*TemperatureTelemetryData)(nil),          // 6: telemetry.TemperatureTelemetryData
	(*timestamppb.Timestamp)(nil),             // 7: google.protobuf.Timestamp
}
var file_telemetry_telemetry_proto_depIdxs = []int32{
	4, // 0: telemetry.ListLatestDeviceTelemetryRequest.cursor:type_name -> telemetry.Cursor
	5, // 1: telemetry.ListLatestDeviceTelemetryResponse.records:type_name -> telemetry.TelemetryData
	5, // 2: telemetry.ListDeviceTelemetryResponse.records:type_name -> telemetry.TelemetryData
	4, // 3: telemetry.ListDeviceTelemetryResponse.next_cursor:type_name -> telemetry.Cursor
	7, // 4: telemetry.TelemetryData.created_at:type_name -> google.protobuf.Timestamp
	6, // 5: telemetry.TelemetryData.temperature_telemetry_data:type_name -> telemetry.TemperatureTelemetryData
	0, // 6: telemetry.TelemetryService.ListLatestDeviceTelemetry:input_type -> telemetry.ListLatestDeviceTelemetryRequest
	2, // 7: telemetry.TelemetryService.ListDeviceTelemetry:input_type -> telemetry.ListDeviceTelemetryRequest
	1, // 8: telemetry.TelemetryService.ListLatestDeviceTelemetry:output_type -> telemetry.ListLatestDeviceTelemetryResponse
	3, // 9: telemetry.TelemetryService.ListDeviceTelemetry:output_type -> telemetry.ListDeviceTelemetryResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_telemetry_telemetry_proto_init() }
func file_telemetry_telemetry_proto_init() {
	if File_telemetry_telemetry_proto != nil {
		return
	}
	file_telemetry_telemetry_proto_msgTypes[0].OneofWrappers = []any{}
	file_telemetry_telemetry_proto_msgTypes[4].OneofWrappers = []any{
		(*Cursor_Id)(nil),
	}
	file_telemetry_telemetry_proto_msgTypes[5].OneofWrappers = []any{
		(*TelemetryData_TemperatureTelemetryData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_telemetry_telemetry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_telemetry_telemetry_proto_goTypes,
		DependencyIndexes: file_telemetry_telemetry_proto_depIdxs,
		MessageInfos:      file_telemetry_telemetry_proto_msgTypes,
	}.Build()
	File_telemetry_telemetry_proto = out.File
	file_telemetry_telemetry_proto_rawDesc = nil
	file_telemetry_telemetry_proto_goTypes = nil
	file_telemetry_telemetry_proto_depIdxs = nil
}
