// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: api/proto/resources/public_ip.proto

package resources

import (
	common "github.com/multycloud/multy/api/proto/common"
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

type CreatePublicIpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []*CloudSpecificPublicIpArgs `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *CreatePublicIpRequest) Reset() {
	*x = CreatePublicIpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resources_public_ip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePublicIpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePublicIpRequest) ProtoMessage() {}

func (x *CreatePublicIpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resources_public_ip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePublicIpRequest.ProtoReflect.Descriptor instead.
func (*CreatePublicIpRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resources_public_ip_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePublicIpRequest) GetResources() []*CloudSpecificPublicIpArgs {
	if x != nil {
		return x.Resources
	}
	return nil
}

type ReadPublicIpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *ReadPublicIpRequest) Reset() {
	*x = ReadPublicIpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resources_public_ip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPublicIpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPublicIpRequest) ProtoMessage() {}

func (x *ReadPublicIpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resources_public_ip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPublicIpRequest.ProtoReflect.Descriptor instead.
func (*ReadPublicIpRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resources_public_ip_proto_rawDescGZIP(), []int{1}
}

func (x *ReadPublicIpRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type UpdatePublicIpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string                       `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Resources  []*CloudSpecificPublicIpArgs `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *UpdatePublicIpRequest) Reset() {
	*x = UpdatePublicIpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resources_public_ip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePublicIpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePublicIpRequest) ProtoMessage() {}

func (x *UpdatePublicIpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resources_public_ip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePublicIpRequest.ProtoReflect.Descriptor instead.
func (*UpdatePublicIpRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resources_public_ip_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePublicIpRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *UpdatePublicIpRequest) GetResources() []*CloudSpecificPublicIpArgs {
	if x != nil {
		return x.Resources
	}
	return nil
}

type DeletePublicIpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *DeletePublicIpRequest) Reset() {
	*x = DeletePublicIpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resources_public_ip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePublicIpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePublicIpRequest) ProtoMessage() {}

func (x *DeletePublicIpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resources_public_ip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePublicIpRequest.ProtoReflect.Descriptor instead.
func (*DeletePublicIpRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resources_public_ip_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePublicIpRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type CloudSpecificPublicIpArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters   *common.CloudSpecificResourceCommonArgs `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	Name               string                                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NetworkInterfaceId string                                  `protobuf:"bytes,3,opt,name=network_interface_id,json=networkInterfaceId,proto3" json:"network_interface_id,omitempty"`
}

func (x *CloudSpecificPublicIpArgs) Reset() {
	*x = CloudSpecificPublicIpArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resources_public_ip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudSpecificPublicIpArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudSpecificPublicIpArgs) ProtoMessage() {}

func (x *CloudSpecificPublicIpArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resources_public_ip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudSpecificPublicIpArgs.ProtoReflect.Descriptor instead.
func (*CloudSpecificPublicIpArgs) Descriptor() ([]byte, []int) {
	return file_api_proto_resources_public_ip_proto_rawDescGZIP(), []int{4}
}

func (x *CloudSpecificPublicIpArgs) GetCommonParameters() *common.CloudSpecificResourceCommonArgs {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *CloudSpecificPublicIpArgs) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CloudSpecificPublicIpArgs) GetNetworkInterfaceId() string {
	if x != nil {
		return x.NetworkInterfaceId
	}
	return ""
}

type CloudSpecificPublicIpResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters   *common.CloudSpecificCommonResourceParameters `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	Name               string                                        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NetworkInterfaceId string                                        `protobuf:"bytes,3,opt,name=network_interface_id,json=networkInterfaceId,proto3" json:"network_interface_id,omitempty"`
}

func (x *CloudSpecificPublicIpResource) Reset() {
	*x = CloudSpecificPublicIpResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resources_public_ip_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudSpecificPublicIpResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudSpecificPublicIpResource) ProtoMessage() {}

func (x *CloudSpecificPublicIpResource) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resources_public_ip_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudSpecificPublicIpResource.ProtoReflect.Descriptor instead.
func (*CloudSpecificPublicIpResource) Descriptor() ([]byte, []int) {
	return file_api_proto_resources_public_ip_proto_rawDescGZIP(), []int{5}
}

func (x *CloudSpecificPublicIpResource) GetCommonParameters() *common.CloudSpecificCommonResourceParameters {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *CloudSpecificPublicIpResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CloudSpecificPublicIpResource) GetNetworkInterfaceId() string {
	if x != nil {
		return x.NetworkInterfaceId
	}
	return ""
}

type PublicIpResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters *common.CommonResourceParameters `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	Resources        []*CloudSpecificPublicIpResource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *PublicIpResource) Reset() {
	*x = PublicIpResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resources_public_ip_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicIpResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicIpResource) ProtoMessage() {}

func (x *PublicIpResource) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resources_public_ip_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicIpResource.ProtoReflect.Descriptor instead.
func (*PublicIpResource) Descriptor() ([]byte, []int) {
	return file_api_proto_resources_public_ip_proto_rawDescGZIP(), []int{6}
}

func (x *PublicIpResource) GetCommonParameters() *common.CommonResourceParameters {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *PublicIpResource) GetResources() []*CloudSpecificPublicIpResource {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_api_proto_resources_public_ip_proto protoreflect.FileDescriptor

var file_api_proto_resources_public_ip_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x4c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74,
	0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49,
	0x70, 0x41, 0x72, 0x67, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x22, 0x36, 0x0a, 0x13, 0x52, 0x65, 0x61, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c,
	0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x49, 0x70, 0x41, 0x72, 0x67, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x22, 0x38, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x49, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0xc1, 0x01, 0x0a, 0x19,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x49, 0x70, 0x41, 0x72, 0x67, 0x73, 0x12, 0x5e, 0x0a, 0x11, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a,
	0x14, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22,
	0xcb, 0x01, 0x0a, 0x1d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x64, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x64,
	0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0xbd, 0x01,
	0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x57, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x50, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x5c, 0x0a,
	0x17, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_proto_resources_public_ip_proto_rawDescOnce sync.Once
	file_api_proto_resources_public_ip_proto_rawDescData = file_api_proto_resources_public_ip_proto_rawDesc
)

func file_api_proto_resources_public_ip_proto_rawDescGZIP() []byte {
	file_api_proto_resources_public_ip_proto_rawDescOnce.Do(func() {
		file_api_proto_resources_public_ip_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_resources_public_ip_proto_rawDescData)
	})
	return file_api_proto_resources_public_ip_proto_rawDescData
}

var file_api_proto_resources_public_ip_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_resources_public_ip_proto_goTypes = []interface{}{
	(*CreatePublicIpRequest)(nil),                        // 0: dev.multy.resources.CreatePublicIpRequest
	(*ReadPublicIpRequest)(nil),                          // 1: dev.multy.resources.ReadPublicIpRequest
	(*UpdatePublicIpRequest)(nil),                        // 2: dev.multy.resources.UpdatePublicIpRequest
	(*DeletePublicIpRequest)(nil),                        // 3: dev.multy.resources.DeletePublicIpRequest
	(*CloudSpecificPublicIpArgs)(nil),                    // 4: dev.multy.resources.CloudSpecificPublicIpArgs
	(*CloudSpecificPublicIpResource)(nil),                // 5: dev.multy.resources.CloudSpecificPublicIpResource
	(*PublicIpResource)(nil),                             // 6: dev.multy.resources.PublicIpResource
	(*common.CloudSpecificResourceCommonArgs)(nil),       // 7: dev.multy.common.CloudSpecificResourceCommonArgs
	(*common.CloudSpecificCommonResourceParameters)(nil), // 8: dev.multy.common.CloudSpecificCommonResourceParameters
	(*common.CommonResourceParameters)(nil),              // 9: dev.multy.common.CommonResourceParameters
}
var file_api_proto_resources_public_ip_proto_depIdxs = []int32{
	4, // 0: dev.multy.resources.CreatePublicIpRequest.resources:type_name -> dev.multy.resources.CloudSpecificPublicIpArgs
	4, // 1: dev.multy.resources.UpdatePublicIpRequest.resources:type_name -> dev.multy.resources.CloudSpecificPublicIpArgs
	7, // 2: dev.multy.resources.CloudSpecificPublicIpArgs.common_parameters:type_name -> dev.multy.common.CloudSpecificResourceCommonArgs
	8, // 3: dev.multy.resources.CloudSpecificPublicIpResource.common_parameters:type_name -> dev.multy.common.CloudSpecificCommonResourceParameters
	9, // 4: dev.multy.resources.PublicIpResource.common_parameters:type_name -> dev.multy.common.CommonResourceParameters
	5, // 5: dev.multy.resources.PublicIpResource.resources:type_name -> dev.multy.resources.CloudSpecificPublicIpResource
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_proto_resources_public_ip_proto_init() }
func file_api_proto_resources_public_ip_proto_init() {
	if File_api_proto_resources_public_ip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_resources_public_ip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePublicIpRequest); i {
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
		file_api_proto_resources_public_ip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPublicIpRequest); i {
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
		file_api_proto_resources_public_ip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePublicIpRequest); i {
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
		file_api_proto_resources_public_ip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePublicIpRequest); i {
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
		file_api_proto_resources_public_ip_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudSpecificPublicIpArgs); i {
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
		file_api_proto_resources_public_ip_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudSpecificPublicIpResource); i {
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
		file_api_proto_resources_public_ip_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicIpResource); i {
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
			RawDescriptor: file_api_proto_resources_public_ip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_resources_public_ip_proto_goTypes,
		DependencyIndexes: file_api_proto_resources_public_ip_proto_depIdxs,
		MessageInfos:      file_api_proto_resources_public_ip_proto_msgTypes,
	}.Build()
	File_api_proto_resources_public_ip_proto = out.File
	file_api_proto_resources_public_ip_proto_rawDesc = nil
	file_api_proto_resources_public_ip_proto_goTypes = nil
	file_api_proto_resources_public_ip_proto_depIdxs = nil
}