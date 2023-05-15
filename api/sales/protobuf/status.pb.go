// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/sales/proto/status.proto

package __

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

// =================
// Metadata
// =================
type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int32 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Total   int32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sales_proto_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_api_sales_proto_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_api_sales_proto_status_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Meta) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *Meta) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

// =================
// Parameter Request
// =================
type ParameterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page            int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage         int32  `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	OrderBy         string `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	OrderMethod     string `protobuf:"bytes,4,opt,name=order_method,json=orderMethod,proto3" json:"order_method,omitempty"`
	SearchCondition string `protobuf:"bytes,5,opt,name=search_condition,json=searchCondition,proto3" json:"search_condition,omitempty"`
	Equal           string `protobuf:"bytes,6,opt,name=equal,proto3" json:"equal,omitempty"`
	Not             string `protobuf:"bytes,7,opt,name=not,proto3" json:"not,omitempty"`
	Like            string `protobuf:"bytes,8,opt,name=like,proto3" json:"like,omitempty"`
	DateRangeBy     string `protobuf:"bytes,9,opt,name=date_range_by,json=dateRangeBy,proto3" json:"date_range_by,omitempty"`
	DateStart       string `protobuf:"bytes,10,opt,name=date_start,json=dateStart,proto3" json:"date_start,omitempty"`
	DateEnd         string `protobuf:"bytes,11,opt,name=date_end,json=dateEnd,proto3" json:"date_end,omitempty"`
}

func (x *ParameterRequest) Reset() {
	*x = ParameterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sales_proto_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParameterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParameterRequest) ProtoMessage() {}

func (x *ParameterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sales_proto_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParameterRequest.ProtoReflect.Descriptor instead.
func (*ParameterRequest) Descriptor() ([]byte, []int) {
	return file_api_sales_proto_status_proto_rawDescGZIP(), []int{1}
}

func (x *ParameterRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ParameterRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *ParameterRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *ParameterRequest) GetOrderMethod() string {
	if x != nil {
		return x.OrderMethod
	}
	return ""
}

func (x *ParameterRequest) GetSearchCondition() string {
	if x != nil {
		return x.SearchCondition
	}
	return ""
}

func (x *ParameterRequest) GetEqual() string {
	if x != nil {
		return x.Equal
	}
	return ""
}

func (x *ParameterRequest) GetNot() string {
	if x != nil {
		return x.Not
	}
	return ""
}

func (x *ParameterRequest) GetLike() string {
	if x != nil {
		return x.Like
	}
	return ""
}

func (x *ParameterRequest) GetDateRangeBy() string {
	if x != nil {
		return x.DateRangeBy
	}
	return ""
}

func (x *ParameterRequest) GetDateStart() string {
	if x != nil {
		return x.DateStart
	}
	return ""
}

func (x *ParameterRequest) GetDateEnd() string {
	if x != nil {
		return x.DateEnd
	}
	return ""
}

// =================
// Response
// =================
type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sales_proto_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_sales_proto_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_api_sales_proto_status_proto_rawDescGZIP(), []int{2}
}

func (x *StatusResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StatusResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StatusResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *StatusResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type StatusesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*StatusResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Meta *Meta             `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *StatusesResponse) Reset() {
	*x = StatusesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sales_proto_status_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusesResponse) ProtoMessage() {}

func (x *StatusesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_sales_proto_status_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusesResponse.ProtoReflect.Descriptor instead.
func (*StatusesResponse) Descriptor() ([]byte, []int) {
	return file_api_sales_proto_status_proto_rawDescGZIP(), []int{3}
}

func (x *StatusesResponse) GetData() []*StatusResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *StatusesResponse) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type ResultStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string          `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *StatusResponse `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ResultStatusResponse) Reset() {
	*x = ResultStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sales_proto_status_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResultStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResultStatusResponse) ProtoMessage() {}

func (x *ResultStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_sales_proto_status_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResultStatusResponse.ProtoReflect.Descriptor instead.
func (*ResultStatusResponse) Descriptor() ([]byte, []int) {
	return file_api_sales_proto_status_proto_rawDescGZIP(), []int{4}
}

func (x *ResultStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ResultStatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResultStatusResponse) GetData() *StatusResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

// =================
// Request
// =================
type CreateStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateStatusRequest) Reset() {
	*x = CreateStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sales_proto_status_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStatusRequest) ProtoMessage() {}

func (x *CreateStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sales_proto_status_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStatusRequest.ProtoReflect.Descriptor instead.
func (*CreateStatusRequest) Descriptor() ([]byte, []int) {
	return file_api_sales_proto_status_proto_rawDescGZIP(), []int{5}
}

func (x *CreateStatusRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateStatusRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FindStatusByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindStatusByIdRequest) Reset() {
	*x = FindStatusByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sales_proto_status_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindStatusByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindStatusByIdRequest) ProtoMessage() {}

func (x *FindStatusByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sales_proto_status_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindStatusByIdRequest.ProtoReflect.Descriptor instead.
func (*FindStatusByIdRequest) Descriptor() ([]byte, []int) {
	return file_api_sales_proto_status_proto_rawDescGZIP(), []int{6}
}

func (x *FindStatusByIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_sales_proto_status_proto protoreflect.FileDescriptor

var file_api_sales_proto_status_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4b, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0xc4, 0x02, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x6f, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6e, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6b, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x22, 0x0a, 0x0d,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x22, 0x72, 0x0a, 0x0e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x60,
	0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x22, 0x74, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x39, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x27, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0xb7, 0x01, 0x0a, 0x0d, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x13,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54,
	0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_sales_proto_status_proto_rawDescOnce sync.Once
	file_api_sales_proto_status_proto_rawDescData = file_api_sales_proto_status_proto_rawDesc
)

func file_api_sales_proto_status_proto_rawDescGZIP() []byte {
	file_api_sales_proto_status_proto_rawDescOnce.Do(func() {
		file_api_sales_proto_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_sales_proto_status_proto_rawDescData)
	})
	return file_api_sales_proto_status_proto_rawDescData
}

var file_api_sales_proto_status_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_sales_proto_status_proto_goTypes = []interface{}{
	(*Meta)(nil),                  // 0: status.Meta
	(*ParameterRequest)(nil),      // 1: status.ParameterRequest
	(*StatusResponse)(nil),        // 2: status.StatusResponse
	(*StatusesResponse)(nil),      // 3: status.StatusesResponse
	(*ResultStatusResponse)(nil),  // 4: status.ResultStatusResponse
	(*CreateStatusRequest)(nil),   // 5: status.CreateStatusRequest
	(*FindStatusByIdRequest)(nil), // 6: status.FindStatusByIdRequest
}
var file_api_sales_proto_status_proto_depIdxs = []int32{
	2, // 0: status.StatusesResponse.data:type_name -> status.StatusResponse
	0, // 1: status.StatusesResponse.meta:type_name -> status.Meta
	2, // 2: status.ResultStatusResponse.data:type_name -> status.StatusResponse
	5, // 3: status.AuthorService.ServiceCreateStatus:input_type -> status.CreateStatusRequest
	6, // 4: status.AuthorService.ServiceFindStatusById:input_type -> status.FindStatusByIdRequest
	4, // 5: status.AuthorService.ServiceCreateStatus:output_type -> status.ResultStatusResponse
	4, // 6: status.AuthorService.ServiceFindStatusById:output_type -> status.ResultStatusResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_sales_proto_status_proto_init() }
func file_api_sales_proto_status_proto_init() {
	if File_api_sales_proto_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_sales_proto_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
		file_api_sales_proto_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParameterRequest); i {
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
		file_api_sales_proto_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_api_sales_proto_status_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusesResponse); i {
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
		file_api_sales_proto_status_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResultStatusResponse); i {
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
		file_api_sales_proto_status_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStatusRequest); i {
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
		file_api_sales_proto_status_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindStatusByIdRequest); i {
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
			RawDescriptor: file_api_sales_proto_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_sales_proto_status_proto_goTypes,
		DependencyIndexes: file_api_sales_proto_status_proto_depIdxs,
		MessageInfos:      file_api_sales_proto_status_proto_msgTypes,
	}.Build()
	File_api_sales_proto_status_proto = out.File
	file_api_sales_proto_status_proto_rawDesc = nil
	file_api_sales_proto_status_proto_goTypes = nil
	file_api_sales_proto_status_proto_depIdxs = nil
}