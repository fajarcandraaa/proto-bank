// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/document.proto

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

type Documents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnsignedDoc string `protobuf:"bytes,1,opt,name=unsigned_doc,json=unsignedDoc,proto3" json:"unsigned_doc,omitempty"`
	SignedDoc   string `protobuf:"bytes,2,opt,name=signed_doc,json=signedDoc,proto3" json:"signed_doc,omitempty"`
}

func (x *Documents) Reset() {
	*x = Documents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_document_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Documents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Documents) ProtoMessage() {}

func (x *Documents) ProtoReflect() protoreflect.Message {
	mi := &file_proto_document_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Documents.ProtoReflect.Descriptor instead.
func (*Documents) Descriptor() ([]byte, []int) {
	return file_proto_document_proto_rawDescGZIP(), []int{0}
}

func (x *Documents) GetUnsignedDoc() string {
	if x != nil {
		return x.UnsignedDoc
	}
	return ""
}

func (x *Documents) GetSignedDoc() string {
	if x != nil {
		return x.SignedDoc
	}
	return ""
}

type UpdateDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReffNumber string `protobuf:"bytes,1,opt,name=reff_number,json=reffNumber,proto3" json:"reff_number,omitempty"`
	Filename   string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *UpdateDocumentRequest) Reset() {
	*x = UpdateDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_document_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDocumentRequest) ProtoMessage() {}

func (x *UpdateDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_document_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDocumentRequest.ProtoReflect.Descriptor instead.
func (*UpdateDocumentRequest) Descriptor() ([]byte, []int) {
	return file_proto_document_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateDocumentRequest) GetReffNumber() string {
	if x != nil {
		return x.ReffNumber
	}
	return ""
}

func (x *UpdateDocumentRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type UpdateDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *BaseResponse `protobuf:"bytes,1,opt,name=base_response,json=baseResponse,proto3" json:"base_response,omitempty"`
	ReffNumber   string        `protobuf:"bytes,2,opt,name=reff_number,json=reffNumber,proto3" json:"reff_number,omitempty"`
	Filename     string        `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *UpdateDocumentResponse) Reset() {
	*x = UpdateDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_document_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDocumentResponse) ProtoMessage() {}

func (x *UpdateDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_document_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDocumentResponse.ProtoReflect.Descriptor instead.
func (*UpdateDocumentResponse) Descriptor() ([]byte, []int) {
	return file_proto_document_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateDocumentResponse) GetBaseResponse() *BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *UpdateDocumentResponse) GetReffNumber() string {
	if x != nil {
		return x.ReffNumber
	}
	return ""
}

func (x *UpdateDocumentResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type GetDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReffNumber string `protobuf:"bytes,1,opt,name=reff_number,json=reffNumber,proto3" json:"reff_number,omitempty"`
}

func (x *GetDocumentRequest) Reset() {
	*x = GetDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_document_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDocumentRequest) ProtoMessage() {}

func (x *GetDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_document_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDocumentRequest.ProtoReflect.Descriptor instead.
func (*GetDocumentRequest) Descriptor() ([]byte, []int) {
	return file_proto_document_proto_rawDescGZIP(), []int{3}
}

func (x *GetDocumentRequest) GetReffNumber() string {
	if x != nil {
		return x.ReffNumber
	}
	return ""
}

type GetDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *BaseResponse `protobuf:"bytes,1,opt,name=base_response,json=baseResponse,proto3" json:"base_response,omitempty"`
	ReffNumber   string        `protobuf:"bytes,2,opt,name=reff_number,json=reffNumber,proto3" json:"reff_number,omitempty"`
	Documents    *Documents    `protobuf:"bytes,3,opt,name=documents,proto3" json:"documents,omitempty"`
}

func (x *GetDocumentResponse) Reset() {
	*x = GetDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_document_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDocumentResponse) ProtoMessage() {}

func (x *GetDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_document_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDocumentResponse.ProtoReflect.Descriptor instead.
func (*GetDocumentResponse) Descriptor() ([]byte, []int) {
	return file_proto_document_proto_rawDescGZIP(), []int{4}
}

func (x *GetDocumentResponse) GetBaseResponse() *BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *GetDocumentResponse) GetReffNumber() string {
	if x != nil {
		return x.ReffNumber
	}
	return ""
}

func (x *GetDocumentResponse) GetDocuments() *Documents {
	if x != nil {
		return x.Documents
	}
	return nil
}

type UploadDocRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationCode string `protobuf:"bytes,1,opt,name=application_code,json=applicationCode,proto3" json:"application_code,omitempty"`
	Filename        string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	DocFile         []byte `protobuf:"bytes,3,opt,name=doc_file,json=docFile,proto3" json:"doc_file,omitempty"`
	SignType        int32  `protobuf:"varint,4,opt,name=sign_type,json=signType,proto3" json:"sign_type,omitempty"`
	DocFlow         int32  `protobuf:"varint,5,opt,name=doc_flow,json=docFlow,proto3" json:"doc_flow,omitempty"`
	DocCategory     string `protobuf:"bytes,6,opt,name=doc_category,json=docCategory,proto3" json:"doc_category,omitempty"`
}

func (x *UploadDocRequest) Reset() {
	*x = UploadDocRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_document_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadDocRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadDocRequest) ProtoMessage() {}

func (x *UploadDocRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_document_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadDocRequest.ProtoReflect.Descriptor instead.
func (*UploadDocRequest) Descriptor() ([]byte, []int) {
	return file_proto_document_proto_rawDescGZIP(), []int{5}
}

func (x *UploadDocRequest) GetApplicationCode() string {
	if x != nil {
		return x.ApplicationCode
	}
	return ""
}

func (x *UploadDocRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UploadDocRequest) GetDocFile() []byte {
	if x != nil {
		return x.DocFile
	}
	return nil
}

func (x *UploadDocRequest) GetSignType() int32 {
	if x != nil {
		return x.SignType
	}
	return 0
}

func (x *UploadDocRequest) GetDocFlow() int32 {
	if x != nil {
		return x.DocFlow
	}
	return 0
}

func (x *UploadDocRequest) GetDocCategory() string {
	if x != nil {
		return x.DocCategory
	}
	return ""
}

type UploadDocResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *BaseResponse `protobuf:"bytes,1,opt,name=base_response,json=baseResponse,proto3" json:"base_response,omitempty"`
	Id           int32         `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	ReffNumber   string        `protobuf:"bytes,3,opt,name=reff_number,json=reffNumber,proto3" json:"reff_number,omitempty"`
	DocName      string        `protobuf:"bytes,4,opt,name=doc_name,json=docName,proto3" json:"doc_name,omitempty"`
	SignType     string        `protobuf:"bytes,5,opt,name=sign_type,json=signType,proto3" json:"sign_type,omitempty"`
	DocCategory  string        `protobuf:"bytes,6,opt,name=doc_category,json=docCategory,proto3" json:"doc_category,omitempty"`
	DocFlow      string        `protobuf:"bytes,7,opt,name=doc_flow,json=docFlow,proto3" json:"doc_flow,omitempty"`
}

func (x *UploadDocResponse) Reset() {
	*x = UploadDocResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_document_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadDocResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadDocResponse) ProtoMessage() {}

func (x *UploadDocResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_document_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadDocResponse.ProtoReflect.Descriptor instead.
func (*UploadDocResponse) Descriptor() ([]byte, []int) {
	return file_proto_document_proto_rawDescGZIP(), []int{6}
}

func (x *UploadDocResponse) GetBaseResponse() *BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *UploadDocResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UploadDocResponse) GetReffNumber() string {
	if x != nil {
		return x.ReffNumber
	}
	return ""
}

func (x *UploadDocResponse) GetDocName() string {
	if x != nil {
		return x.DocName
	}
	return ""
}

func (x *UploadDocResponse) GetSignType() string {
	if x != nil {
		return x.SignType
	}
	return ""
}

func (x *UploadDocResponse) GetDocCategory() string {
	if x != nil {
		return x.DocCategory
	}
	return ""
}

func (x *UploadDocResponse) GetDocFlow() string {
	if x != nil {
		return x.DocFlow
	}
	return ""
}

type DeleteDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReffNumber string `protobuf:"bytes,1,opt,name=reff_number,json=reffNumber,proto3" json:"reff_number,omitempty"`
}

func (x *DeleteDocumentRequest) Reset() {
	*x = DeleteDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_document_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDocumentRequest) ProtoMessage() {}

func (x *DeleteDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_document_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDocumentRequest.ProtoReflect.Descriptor instead.
func (*DeleteDocumentRequest) Descriptor() ([]byte, []int) {
	return file_proto_document_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteDocumentRequest) GetReffNumber() string {
	if x != nil {
		return x.ReffNumber
	}
	return ""
}

type DeleteDocumentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResponse *BaseResponse `protobuf:"bytes,1,opt,name=base_response,json=baseResponse,proto3" json:"base_response,omitempty"`
	Result       bool          `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DeleteDocumentResponse) Reset() {
	*x = DeleteDocumentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_document_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDocumentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDocumentResponse) ProtoMessage() {}

func (x *DeleteDocumentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_document_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDocumentResponse.ProtoReflect.Descriptor instead.
func (*DeleteDocumentResponse) Descriptor() ([]byte, []int) {
	return file_proto_document_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteDocumentResponse) GetBaseResponse() *BaseResponse {
	if x != nil {
		return x.BaseResponse
	}
	return nil
}

func (x *DeleteDocumentResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type BaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_document_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_document_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_document_proto_rawDescGZIP(), []int{9}
}

func (x *BaseResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BaseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_document_proto protoreflect.FileDescriptor

var file_proto_document_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x09, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f,
	0x64, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x6e, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x44, 0x6f, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x5f, 0x64, 0x6f, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x44, 0x6f, 0x63, 0x22, 0x54, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x66, 0x66, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x66, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x16,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x62, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x66, 0x66, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x66, 0x66, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x66, 0x66, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x66, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x94,
	0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x62, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x66, 0x66, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x66, 0x66, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x09, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x09, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xcf, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x44, 0x6f, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x64, 0x6f, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x6f, 0x63,
	0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x6f, 0x63,
	0x46, 0x6c, 0x6f, 0x77, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x63, 0x5f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x6f, 0x63, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0xee, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x66, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x66, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f,
	0x63, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x6f, 0x63, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x0a,
	0x08, 0x64, 0x6f, 0x63, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x6f, 0x63, 0x46, 0x6c, 0x6f, 0x77, 0x22, 0x38, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x66, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x66, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x64, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0d,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x49, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xa8, 0x02, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x38, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x11, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x6f, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x44, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x52, 0x65, 0x66, 0x66, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4d, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x42, 0x79, 0x52, 0x65, 0x66, 0x66, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4d, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0x79, 0x52, 0x65, 0x66, 0x66, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03,
	0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_document_proto_rawDescOnce sync.Once
	file_proto_document_proto_rawDescData = file_proto_document_proto_rawDesc
)

func file_proto_document_proto_rawDescGZIP() []byte {
	file_proto_document_proto_rawDescOnce.Do(func() {
		file_proto_document_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_document_proto_rawDescData)
	})
	return file_proto_document_proto_rawDescData
}

var file_proto_document_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_document_proto_goTypes = []interface{}{
	(*Documents)(nil),              // 0: Documents
	(*UpdateDocumentRequest)(nil),  // 1: UpdateDocumentRequest
	(*UpdateDocumentResponse)(nil), // 2: UpdateDocumentResponse
	(*GetDocumentRequest)(nil),     // 3: GetDocumentRequest
	(*GetDocumentResponse)(nil),    // 4: GetDocumentResponse
	(*UploadDocRequest)(nil),       // 5: UploadDocRequest
	(*UploadDocResponse)(nil),      // 6: UploadDocResponse
	(*DeleteDocumentRequest)(nil),  // 7: DeleteDocumentRequest
	(*DeleteDocumentResponse)(nil), // 8: DeleteDocumentResponse
	(*BaseResponse)(nil),           // 9: BaseResponse
}
var file_proto_document_proto_depIdxs = []int32{
	9, // 0: UpdateDocumentResponse.base_response:type_name -> BaseResponse
	9, // 1: GetDocumentResponse.base_response:type_name -> BaseResponse
	0, // 2: GetDocumentResponse.documents:type_name -> Documents
	9, // 3: UploadDocResponse.base_response:type_name -> BaseResponse
	9, // 4: DeleteDocumentResponse.base_response:type_name -> BaseResponse
	5, // 5: Document.StoreDocument:input_type -> UploadDocRequest
	3, // 6: Document.GetDocumentByReffNumber:input_type -> GetDocumentRequest
	1, // 7: Document.UpdateDocumentByReffNumber:input_type -> UpdateDocumentRequest
	7, // 8: Document.DeleteDocumentByReffNumber:input_type -> DeleteDocumentRequest
	6, // 9: Document.StoreDocument:output_type -> UploadDocResponse
	4, // 10: Document.GetDocumentByReffNumber:output_type -> GetDocumentResponse
	2, // 11: Document.UpdateDocumentByReffNumber:output_type -> UpdateDocumentResponse
	8, // 12: Document.DeleteDocumentByReffNumber:output_type -> DeleteDocumentResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_document_proto_init() }
func file_proto_document_proto_init() {
	if File_proto_document_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_document_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Documents); i {
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
		file_proto_document_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDocumentRequest); i {
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
		file_proto_document_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDocumentResponse); i {
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
		file_proto_document_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDocumentRequest); i {
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
		file_proto_document_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDocumentResponse); i {
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
		file_proto_document_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadDocRequest); i {
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
		file_proto_document_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadDocResponse); i {
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
		file_proto_document_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDocumentRequest); i {
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
		file_proto_document_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDocumentResponse); i {
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
		file_proto_document_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResponse); i {
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
			RawDescriptor: file_proto_document_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_document_proto_goTypes,
		DependencyIndexes: file_proto_document_proto_depIdxs,
		MessageInfos:      file_proto_document_proto_msgTypes,
	}.Build()
	File_proto_document_proto = out.File
	file_proto_document_proto_rawDesc = nil
	file_proto_document_proto_goTypes = nil
	file_proto_document_proto_depIdxs = nil
}