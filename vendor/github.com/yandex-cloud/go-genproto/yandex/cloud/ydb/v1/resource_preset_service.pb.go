// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: yandex/cloud/ydb/v1/resource_preset_service.proto

package ydb

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetResourcePresetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the resource preset to return.
	ResourcePresetId string `protobuf:"bytes,1,opt,name=resource_preset_id,json=resourcePresetId,proto3" json:"resource_preset_id,omitempty"`
}

func (x *GetResourcePresetRequest) Reset() {
	*x = GetResourcePresetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ydb_v1_resource_preset_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourcePresetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcePresetRequest) ProtoMessage() {}

func (x *GetResourcePresetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_resource_preset_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcePresetRequest.ProtoReflect.Descriptor instead.
func (*GetResourcePresetRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_resource_preset_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetResourcePresetRequest) GetResourcePresetId() string {
	if x != nil {
		return x.ResourcePresetId
	}
	return ""
}

type ListResourcePresetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of results per page that should be returned. If the number of available
	// results is larger than `page_size`, the service returns a `next_page_token` that can be used
	// to get the next page of results in subsequent ListResourcePresets requests.
	// Acceptable values are 0 to 1000, inclusive. Default value: 100.
	PageSize int64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. Set `page_token` to the `next_page_token` returned by a previous ListResourcePresets
	// request to get the next page of results.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListResourcePresetsRequest) Reset() {
	*x = ListResourcePresetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ydb_v1_resource_preset_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourcePresetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourcePresetsRequest) ProtoMessage() {}

func (x *ListResourcePresetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_resource_preset_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourcePresetsRequest.ProtoReflect.Descriptor instead.
func (*ListResourcePresetsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_resource_preset_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListResourcePresetsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListResourcePresetsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListResourcePresetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requested list of resource presets.
	ResourcePresets []*ResourcePreset `protobuf:"bytes,1,rep,name=resource_presets,json=resourcePresets,proto3" json:"resource_presets,omitempty"`
	// This token allows you to get the next page of results for ListResourcePresets requests,
	// if the number of results is larger than `page_size` specified in the request.
	// To get the next page, specify the value of `next_page_token` as a value for
	// the `page_token` parameter in the next ListResourcePresets request. Subsequent ListResourcePresets
	// requests will have their own `next_page_token` to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListResourcePresetsResponse) Reset() {
	*x = ListResourcePresetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ydb_v1_resource_preset_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourcePresetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourcePresetsResponse) ProtoMessage() {}

func (x *ListResourcePresetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_resource_preset_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourcePresetsResponse.ProtoReflect.Descriptor instead.
func (*ListResourcePresetsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_resource_preset_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListResourcePresetsResponse) GetResourcePresets() []*ResourcePreset {
	if x != nil {
		return x.ResourcePresets
	}
	return nil
}

func (x *ListResourcePresetsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_ydb_v1_resource_preset_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_ydb_v1_resource_preset_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x79,
	0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x79, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x12,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x10,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x49, 0x64,
	0x22, 0x6f, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06, 0x30, 0x2d, 0x31, 0x30, 0x30, 0x30, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31,
	0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x95, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4e, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xb6, 0x02, 0x0a, 0x15, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x8f, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2d, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x22,
	0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x12, 0x2c, 0x2f, 0x79, 0x64, 0x62, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73,
	0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x8a, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79,
	0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x79, 0x64, 0x62, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x73, 0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x79, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x79, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_yandex_cloud_ydb_v1_resource_preset_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_ydb_v1_resource_preset_service_proto_rawDescData = file_yandex_cloud_ydb_v1_resource_preset_service_proto_rawDesc
)

func file_yandex_cloud_ydb_v1_resource_preset_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ydb_v1_resource_preset_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ydb_v1_resource_preset_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ydb_v1_resource_preset_service_proto_rawDescData)
	})
	return file_yandex_cloud_ydb_v1_resource_preset_service_proto_rawDescData
}

var file_yandex_cloud_ydb_v1_resource_preset_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_ydb_v1_resource_preset_service_proto_goTypes = []interface{}{
	(*GetResourcePresetRequest)(nil),    // 0: yandex.cloud.ydb.v1.GetResourcePresetRequest
	(*ListResourcePresetsRequest)(nil),  // 1: yandex.cloud.ydb.v1.ListResourcePresetsRequest
	(*ListResourcePresetsResponse)(nil), // 2: yandex.cloud.ydb.v1.ListResourcePresetsResponse
	(*ResourcePreset)(nil),              // 3: yandex.cloud.ydb.v1.ResourcePreset
}
var file_yandex_cloud_ydb_v1_resource_preset_service_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.ydb.v1.ListResourcePresetsResponse.resource_presets:type_name -> yandex.cloud.ydb.v1.ResourcePreset
	0, // 1: yandex.cloud.ydb.v1.ResourcePresetService.Get:input_type -> yandex.cloud.ydb.v1.GetResourcePresetRequest
	1, // 2: yandex.cloud.ydb.v1.ResourcePresetService.List:input_type -> yandex.cloud.ydb.v1.ListResourcePresetsRequest
	3, // 3: yandex.cloud.ydb.v1.ResourcePresetService.Get:output_type -> yandex.cloud.ydb.v1.ResourcePreset
	2, // 4: yandex.cloud.ydb.v1.ResourcePresetService.List:output_type -> yandex.cloud.ydb.v1.ListResourcePresetsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ydb_v1_resource_preset_service_proto_init() }
func file_yandex_cloud_ydb_v1_resource_preset_service_proto_init() {
	if File_yandex_cloud_ydb_v1_resource_preset_service_proto != nil {
		return
	}
	file_yandex_cloud_ydb_v1_resource_preset_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ydb_v1_resource_preset_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourcePresetRequest); i {
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
		file_yandex_cloud_ydb_v1_resource_preset_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourcePresetsRequest); i {
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
		file_yandex_cloud_ydb_v1_resource_preset_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourcePresetsResponse); i {
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
			RawDescriptor: file_yandex_cloud_ydb_v1_resource_preset_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_ydb_v1_resource_preset_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ydb_v1_resource_preset_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ydb_v1_resource_preset_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ydb_v1_resource_preset_service_proto = out.File
	file_yandex_cloud_ydb_v1_resource_preset_service_proto_rawDesc = nil
	file_yandex_cloud_ydb_v1_resource_preset_service_proto_goTypes = nil
	file_yandex_cloud_ydb_v1_resource_preset_service_proto_depIdxs = nil
}
