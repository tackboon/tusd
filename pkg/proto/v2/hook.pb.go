// If this file gets changed, you must recompile the generate package in pkg/proto.
// To do this, install the Go protobuf toolchain as mentioned in
// https://github.com/golang/protobuf#installation.
// Then use following command to recompile it with gRPC support:
//   protoc --go_out=plugins=grpc:../../../../../pkg/proto/ v1/hook.proto
// In addition, it may be necessary to update the protobuf or gRPC dependencies as well.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: cmd/tusd/cli/hooks/proto/v2/hook.proto

package v2

import (
	any1 "github.com/golang/protobuf/ptypes/any"
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

// Uploaded data
type Upload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique integer identifier of the uploaded file
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Total file size in bytes specified in the NewUpload call
	Size int64 `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	// Indicates whether the total file size is deferred until later
	SizeIsDeferred bool `protobuf:"varint,3,opt,name=SizeIsDeferred,proto3" json:"SizeIsDeferred,omitempty"`
	// Offset in bytes (zero-based)
	Offset   int64             `protobuf:"varint,4,opt,name=Offset,proto3" json:"Offset,omitempty"`
	MetaData map[string]string `protobuf:"bytes,5,rep,name=metaData,proto3" json:"metaData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Indicates that this is a partial upload which will later be used to form
	// a final upload by concatenation. Partial uploads should not be processed
	// when they are finished since they are only incomplete chunks of files.
	IsPartial bool `protobuf:"varint,6,opt,name=isPartial,proto3" json:"isPartial,omitempty"`
	// Indicates that this is a final upload
	IsFinal bool `protobuf:"varint,7,opt,name=isFinal,proto3" json:"isFinal,omitempty"`
	// If the upload is a final one (see IsFinal) this will be a non-empty
	// ordered slice containing the ids of the uploads of which the final upload
	// will consist after concatenation.
	PartialUploads []string `protobuf:"bytes,8,rep,name=partialUploads,proto3" json:"partialUploads,omitempty"`
	// Storage contains information about where the data storage saves the upload,
	// for example a file path. The available values vary depending on what data
	// store is used. This map may also be nil.
	Storage map[string]string `protobuf:"bytes,9,rep,name=storage,proto3" json:"storage,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Upload) Reset() {
	*x = Upload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Upload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Upload) ProtoMessage() {}

func (x *Upload) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Upload.ProtoReflect.Descriptor instead.
func (*Upload) Descriptor() ([]byte, []int) {
	return file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDescGZIP(), []int{0}
}

func (x *Upload) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Upload) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Upload) GetSizeIsDeferred() bool {
	if x != nil {
		return x.SizeIsDeferred
	}
	return false
}

func (x *Upload) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Upload) GetMetaData() map[string]string {
	if x != nil {
		return x.MetaData
	}
	return nil
}

func (x *Upload) GetIsPartial() bool {
	if x != nil {
		return x.IsPartial
	}
	return false
}

func (x *Upload) GetIsFinal() bool {
	if x != nil {
		return x.IsFinal
	}
	return false
}

func (x *Upload) GetPartialUploads() []string {
	if x != nil {
		return x.PartialUploads
	}
	return nil
}

func (x *Upload) GetStorage() map[string]string {
	if x != nil {
		return x.Storage
	}
	return nil
}

type HTTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Method is the HTTP method, e.g. POST or PATCH
	Method string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// URI is the full HTTP request URI, e.g. /files/fooo
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// RemoteAddr contains the network address that sent the request
	RemoteAddr string `protobuf:"bytes,3,opt,name=remoteAddr,proto3" json:"remoteAddr,omitempty"`
	// BearerAuth contains the HTTP Authorization header value
	BearerAuth string `protobuf:"bytes,4,opt,name=bearerAuth,proto3" json:"bearerAuth,omitempty"`
}

func (x *HTTPRequest) Reset() {
	*x = HTTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPRequest) ProtoMessage() {}

func (x *HTTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPRequest.ProtoReflect.Descriptor instead.
func (*HTTPRequest) Descriptor() ([]byte, []int) {
	return file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDescGZIP(), []int{1}
}

func (x *HTTPRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *HTTPRequest) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *HTTPRequest) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *HTTPRequest) GetBearerAuth() string {
	if x != nil {
		return x.BearerAuth
	}
	return ""
}

// Hook's data
type Hook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Upload contains information about the upload that caused this hook
	// to be fired.
	Upload *Upload `protobuf:"bytes,1,opt,name=upload,proto3" json:"upload,omitempty"`
	// HTTPRequest contains details about the HTTP request that reached
	// tusd.
	HttpRequest *HTTPRequest `protobuf:"bytes,2,opt,name=httpRequest,proto3" json:"httpRequest,omitempty"`
	// The hook name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Hook) Reset() {
	*x = Hook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hook) ProtoMessage() {}

func (x *Hook) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hook.ProtoReflect.Descriptor instead.
func (*Hook) Descriptor() ([]byte, []int) {
	return file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDescGZIP(), []int{2}
}

func (x *Hook) GetUpload() *Upload {
	if x != nil {
		return x.Upload
	}
	return nil
}

func (x *Hook) GetHttpRequest() *HTTPRequest {
	if x != nil {
		return x.HttpRequest
	}
	return nil
}

func (x *Hook) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request data to send hook
type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hook data
	Hook *Hook `protobuf:"bytes,1,opt,name=hook,proto3" json:"hook,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDescGZIP(), []int{3}
}

func (x *SendRequest) GetHook() *Hook {
	if x != nil {
		return x.Hook
	}
	return nil
}

// Response that contains data for sended hook
type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The response of the hook.
	Response *any1.Any `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDescGZIP(), []int{4}
}

func (x *SendResponse) GetResponse() *any1.Any {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_cmd_tusd_cli_hooks_proto_v2_hook_proto protoreflect.FileDescriptor

var file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6d, 0x64, 0x2f, 0x74, 0x75, 0x73, 0x64, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x68,
	0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x68, 0x6f,
	0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x32, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x03, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x53, 0x69, 0x7a, 0x65, 0x49, 0x73,
	0x44, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e,
	0x53, 0x69, 0x7a, 0x65, 0x49, 0x73, 0x44, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x73, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73,
	0x46, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x46,
	0x69, 0x6e, 0x61, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x31, 0x0a, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x76, 0x32, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a,
	0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x77, 0x0a, 0x0b, 0x48, 0x54, 0x54, 0x50,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x69, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74,
	0x68, 0x22, 0x71, 0x0a, 0x04, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x22, 0x0a, 0x06, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x76, 0x32, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x31, 0x0a,
	0x0b, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x32, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x76, 0x32, 0x2e, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x68, 0x6f, 0x6f,
	0x6b, 0x22, 0x40, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x3a, 0x0a, 0x0b, 0x48, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x0f, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x76, 0x32,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDescOnce sync.Once
	file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDescData = file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDesc
)

func file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDescGZIP() []byte {
	file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDescOnce.Do(func() {
		file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDescData)
	})
	return file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDescData
}

var file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cmd_tusd_cli_hooks_proto_v2_hook_proto_goTypes = []interface{}{
	(*Upload)(nil),       // 0: v2.Upload
	(*HTTPRequest)(nil),  // 1: v2.HTTPRequest
	(*Hook)(nil),         // 2: v2.Hook
	(*SendRequest)(nil),  // 3: v2.SendRequest
	(*SendResponse)(nil), // 4: v2.SendResponse
	nil,                  // 5: v2.Upload.MetaDataEntry
	nil,                  // 6: v2.Upload.StorageEntry
	(*any1.Any)(nil),     // 7: google.protobuf.Any
}
var file_cmd_tusd_cli_hooks_proto_v2_hook_proto_depIdxs = []int32{
	5, // 0: v2.Upload.metaData:type_name -> v2.Upload.MetaDataEntry
	6, // 1: v2.Upload.storage:type_name -> v2.Upload.StorageEntry
	0, // 2: v2.Hook.upload:type_name -> v2.Upload
	1, // 3: v2.Hook.httpRequest:type_name -> v2.HTTPRequest
	2, // 4: v2.SendRequest.hook:type_name -> v2.Hook
	7, // 5: v2.SendResponse.response:type_name -> google.protobuf.Any
	3, // 6: v2.HookService.Send:input_type -> v2.SendRequest
	4, // 7: v2.HookService.Send:output_type -> v2.SendResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cmd_tusd_cli_hooks_proto_v2_hook_proto_init() }
func file_cmd_tusd_cli_hooks_proto_v2_hook_proto_init() {
	if File_cmd_tusd_cli_hooks_proto_v2_hook_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Upload); i {
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
		file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPRequest); i {
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
		file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hook); i {
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
		file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
			RawDescriptor: file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_tusd_cli_hooks_proto_v2_hook_proto_goTypes,
		DependencyIndexes: file_cmd_tusd_cli_hooks_proto_v2_hook_proto_depIdxs,
		MessageInfos:      file_cmd_tusd_cli_hooks_proto_v2_hook_proto_msgTypes,
	}.Build()
	File_cmd_tusd_cli_hooks_proto_v2_hook_proto = out.File
	file_cmd_tusd_cli_hooks_proto_v2_hook_proto_rawDesc = nil
	file_cmd_tusd_cli_hooks_proto_v2_hook_proto_goTypes = nil
	file_cmd_tusd_cli_hooks_proto_v2_hook_proto_depIdxs = nil
}
