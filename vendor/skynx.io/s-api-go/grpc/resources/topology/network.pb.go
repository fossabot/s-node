// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc3
// source: skynx/protobuf/resources/v1/topology/network.proto

package topology

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	resource "skynx.io/s-api-go/grpc/resources/resource"
	tenant "skynx.io/s-api-go/grpc/resources/tenant"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NewNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID     string `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID      string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	Description   string `protobuf:"bytes,21,opt,name=description,proto3" json:"description,omitempty"`
	NetworkCIDR   string `protobuf:"bytes,31,opt,name=networkCIDR,proto3" json:"networkCIDR,omitempty"`
	RoutedSubnets bool   `protobuf:"varint,41,opt,name=routedSubnets,proto3" json:"routedSubnets,omitempty"` // connect subnets within the network
	LocationID    string `protobuf:"bytes,51,opt,name=locationID,proto3" json:"locationID,omitempty"`
}

func (x *NewNetworkRequest) Reset() {
	*x = NewNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewNetworkRequest) ProtoMessage() {}

func (x *NewNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewNetworkRequest.ProtoReflect.Descriptor instead.
func (*NewNetworkRequest) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_topology_network_proto_rawDescGZIP(), []int{0}
}

func (x *NewNetworkRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *NewNetworkRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *NewNetworkRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewNetworkRequest) GetNetworkCIDR() string {
	if x != nil {
		return x.NetworkCIDR
	}
	return ""
}

func (x *NewNetworkRequest) GetRoutedSubnets() bool {
	if x != nil {
		return x.RoutedSubnets
	}
	return false
}

func (x *NewNetworkRequest) GetLocationID() string {
	if x != nil {
		return x.LocationID
	}
	return ""
}

type UpdateNetworkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID     string `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID      string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	NetID         string `protobuf:"bytes,3,opt,name=netID,proto3" json:"netID,omitempty"`
	Description   string `protobuf:"bytes,21,opt,name=description,proto3" json:"description,omitempty"`
	RoutedSubnets bool   `protobuf:"varint,41,opt,name=routedSubnets,proto3" json:"routedSubnets,omitempty"` // connect subnets within the network
}

func (x *UpdateNetworkRequest) Reset() {
	*x = UpdateNetworkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNetworkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNetworkRequest) ProtoMessage() {}

func (x *UpdateNetworkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNetworkRequest.ProtoReflect.Descriptor instead.
func (*UpdateNetworkRequest) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_topology_network_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateNetworkRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *UpdateNetworkRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *UpdateNetworkRequest) GetNetID() string {
	if x != nil {
		return x.NetID
	}
	return ""
}

func (x *UpdateNetworkRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateNetworkRequest) GetRoutedSubnets() bool {
	if x != nil {
		return x.RoutedSubnets
	}
	return false
}

type Network struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID     string   `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID      string   `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	NetID         string   `protobuf:"bytes,3,opt,name=netID,proto3" json:"netID,omitempty"`
	Description   string   `protobuf:"bytes,21,opt,name=description,proto3" json:"description,omitempty"`
	NetworkCIDR   string   `protobuf:"bytes,31,opt,name=networkCIDR,proto3" json:"networkCIDR,omitempty"`
	RoutedSubnets bool     `protobuf:"varint,41,opt,name=routedSubnets,proto3" json:"routedSubnets,omitempty"` // connect subnets within the network
	LocationID    string   `protobuf:"bytes,51,opt,name=locationID,proto3" json:"locationID,omitempty"`
	Tags          []string `protobuf:"bytes,1000,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Network) Reset() {
	*x = Network{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Network) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Network) ProtoMessage() {}

func (x *Network) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Network.ProtoReflect.Descriptor instead.
func (*Network) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_topology_network_proto_rawDescGZIP(), []int{2}
}

func (x *Network) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *Network) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *Network) GetNetID() string {
	if x != nil {
		return x.NetID
	}
	return ""
}

func (x *Network) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Network) GetNetworkCIDR() string {
	if x != nil {
		return x.NetworkCIDR
	}
	return ""
}

func (x *Network) GetRoutedSubnets() bool {
	if x != nil {
		return x.RoutedSubnets
	}
	return false
}

func (x *Network) GetLocationID() string {
	if x != nil {
		return x.LocationID
	}
	return ""
}

func (x *Network) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Networks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta     *resource.ListResponse `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Networks []*Network             `protobuf:"bytes,2,rep,name=networks,proto3" json:"networks,omitempty"`
}

func (x *Networks) Reset() {
	*x = Networks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Networks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Networks) ProtoMessage() {}

func (x *Networks) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Networks.ProtoReflect.Descriptor instead.
func (*Networks) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_topology_network_proto_rawDescGZIP(), []int{3}
}

func (x *Networks) GetMeta() *resource.ListResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Networks) GetNetworks() []*Network {
	if x != nil {
		return x.Networks
	}
	return nil
}

type ListNetworksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta   *resource.ListRequest `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Tenant *tenant.TenantReq     `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *ListNetworksRequest) Reset() {
	*x = ListNetworksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNetworksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNetworksRequest) ProtoMessage() {}

func (x *ListNetworksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNetworksRequest.ProtoReflect.Descriptor instead.
func (*ListNetworksRequest) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_topology_network_proto_rawDescGZIP(), []int{4}
}

func (x *ListNetworksRequest) GetMeta() *resource.ListRequest {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListNetworksRequest) GetTenant() *tenant.TenantReq {
	if x != nil {
		return x.Tenant
	}
	return nil
}

type NetworkReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID  string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	NetID     string `protobuf:"bytes,3,opt,name=netID,proto3" json:"netID,omitempty"`
}

func (x *NetworkReq) Reset() {
	*x = NetworkReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkReq) ProtoMessage() {}

func (x *NetworkReq) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkReq.ProtoReflect.Descriptor instead.
func (*NetworkReq) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_topology_network_proto_rawDescGZIP(), []int{5}
}

func (x *NetworkReq) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *NetworkReq) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *NetworkReq) GetNetID() string {
	if x != nil {
		return x.NetID
	}
	return ""
}

var File_skynx_protobuf_resources_v1_topology_network_proto protoreflect.FileDescriptor

var file_skynx_protobuf_resources_v1_topology_network_proto_rawDesc = []byte{
	0x0a, 0x32, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f,
	0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x1a, 0x2f,
	0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2f, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd7, 0x01, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x49, 0x44,
	0x52, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x43, 0x49, 0x44, 0x52, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x64, 0x53, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x29, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x64, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0xae, 0x01, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x65,
	0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x64, 0x53,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x29, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x64, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x22, 0xf8, 0x01, 0x0a, 0x07,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x43, 0x49, 0x44, 0x52, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x49, 0x44, 0x52, 0x12, 0x24, 0x0a, 0x0d, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x64, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x29, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x64, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x13, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0xe8, 0x07, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x65, 0x0a, 0x08, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x2d,
	0x0a, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x52, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x22, 0x6b, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12,
	0x29, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x5c, 0x0a, 0x0a, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x42, 0x2b, 0x5a, 0x29, 0x73, 0x6b, 0x79, 0x6e,
	0x78, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x6f, 0x70,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_skynx_protobuf_resources_v1_topology_network_proto_rawDescOnce sync.Once
	file_skynx_protobuf_resources_v1_topology_network_proto_rawDescData = file_skynx_protobuf_resources_v1_topology_network_proto_rawDesc
)

func file_skynx_protobuf_resources_v1_topology_network_proto_rawDescGZIP() []byte {
	file_skynx_protobuf_resources_v1_topology_network_proto_rawDescOnce.Do(func() {
		file_skynx_protobuf_resources_v1_topology_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_skynx_protobuf_resources_v1_topology_network_proto_rawDescData)
	})
	return file_skynx_protobuf_resources_v1_topology_network_proto_rawDescData
}

var file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_skynx_protobuf_resources_v1_topology_network_proto_goTypes = []interface{}{
	(*NewNetworkRequest)(nil),     // 0: topology.NewNetworkRequest
	(*UpdateNetworkRequest)(nil),  // 1: topology.UpdateNetworkRequest
	(*Network)(nil),               // 2: topology.Network
	(*Networks)(nil),              // 3: topology.Networks
	(*ListNetworksRequest)(nil),   // 4: topology.ListNetworksRequest
	(*NetworkReq)(nil),            // 5: topology.NetworkReq
	(*resource.ListResponse)(nil), // 6: resource.ListResponse
	(*resource.ListRequest)(nil),  // 7: resource.ListRequest
	(*tenant.TenantReq)(nil),      // 8: tenant.TenantReq
}
var file_skynx_protobuf_resources_v1_topology_network_proto_depIdxs = []int32{
	6, // 0: topology.Networks.meta:type_name -> resource.ListResponse
	2, // 1: topology.Networks.networks:type_name -> topology.Network
	7, // 2: topology.ListNetworksRequest.meta:type_name -> resource.ListRequest
	8, // 3: topology.ListNetworksRequest.tenant:type_name -> tenant.TenantReq
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_skynx_protobuf_resources_v1_topology_network_proto_init() }
func file_skynx_protobuf_resources_v1_topology_network_proto_init() {
	if File_skynx_protobuf_resources_v1_topology_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewNetworkRequest); i {
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
		file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNetworkRequest); i {
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
		file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Network); i {
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
		file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Networks); i {
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
		file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNetworksRequest); i {
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
		file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkReq); i {
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
			RawDescriptor: file_skynx_protobuf_resources_v1_topology_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_skynx_protobuf_resources_v1_topology_network_proto_goTypes,
		DependencyIndexes: file_skynx_protobuf_resources_v1_topology_network_proto_depIdxs,
		MessageInfos:      file_skynx_protobuf_resources_v1_topology_network_proto_msgTypes,
	}.Build()
	File_skynx_protobuf_resources_v1_topology_network_proto = out.File
	file_skynx_protobuf_resources_v1_topology_network_proto_rawDesc = nil
	file_skynx_protobuf_resources_v1_topology_network_proto_goTypes = nil
	file_skynx_protobuf_resources_v1_topology_network_proto_depIdxs = nil
}
