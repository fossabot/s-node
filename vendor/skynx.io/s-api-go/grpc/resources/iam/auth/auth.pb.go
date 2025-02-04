// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc3
// source: skynx/protobuf/resources/v1/iam/auth/auth.proto

package auth

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

type AuthKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *AuthKey) Reset() {
	*x = AuthKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthKey) ProtoMessage() {}

func (x *AuthKey) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthKey.ProtoReflect.Descriptor instead.
func (*AuthKey) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_iam_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type AuthToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subkey1          string `protobuf:"bytes,1,opt,name=subkey1,proto3" json:"subkey1,omitempty"`                     // uniqueID (userID)
	Subkey2          string `protobuf:"bytes,2,opt,name=subkey2,proto3" json:"subkey2,omitempty"`                     // realm (accountID)
	ValidationKey    string `protobuf:"bytes,21,opt,name=validationKey,proto3" json:"validationKey,omitempty"`        // encrypted expTime
	UserSessionToken string `protobuf:"bytes,101,opt,name=userSessionToken,proto3" json:"userSessionToken,omitempty"` // encrypted sessionToken
}

func (x *AuthToken) Reset() {
	*x = AuthToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthToken) ProtoMessage() {}

func (x *AuthToken) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthToken.ProtoReflect.Descriptor instead.
func (*AuthToken) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_iam_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthToken) GetSubkey1() string {
	if x != nil {
		return x.Subkey1
	}
	return ""
}

func (x *AuthToken) GetSubkey2() string {
	if x != nil {
		return x.Subkey2
	}
	return ""
}

func (x *AuthToken) GetValidationKey() string {
	if x != nil {
		return x.ValidationKey
	}
	return ""
}

func (x *AuthToken) GetUserSessionToken() string {
	if x != nil {
		return x.UserSessionToken
	}
	return ""
}

type AuthCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Realm        string `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"`                // subkey2 -- realm/scope (accountID)
	ID           string `protobuf:"bytes,11,opt,name=ID,proto3" json:"ID,omitempty"`                     // subkey1 -- uniqueID (userID)
	SessionToken string `protobuf:"bytes,21,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"` // user sessionToken
}

func (x *AuthCredentials) Reset() {
	*x = AuthCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthCredentials) ProtoMessage() {}

func (x *AuthCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthCredentials.ProtoReflect.Descriptor instead.
func (*AuthCredentials) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_iam_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *AuthCredentials) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *AuthCredentials) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *AuthCredentials) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

type AccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserRealm    string `protobuf:"bytes,1,opt,name=userRealm,proto3" json:"userRealm,omitempty"`
	UserID       string `protobuf:"bytes,11,opt,name=userID,proto3" json:"userID,omitempty"` // uniqueID
	IsAdmin      bool   `protobuf:"varint,12,opt,name=isAdmin,proto3" json:"isAdmin,omitempty"`
	SessionToken string `protobuf:"bytes,15,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"` // string sessionJWT = 16;
	Realm        string `protobuf:"bytes,21,opt,name=realm,proto3" json:"realm,omitempty"`
	TenantID     string `protobuf:"bytes,22,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	NetID        string `protobuf:"bytes,23,opt,name=netID,proto3" json:"netID,omitempty"`
	SubnetID     string `protobuf:"bytes,24,opt,name=subnetID,proto3" json:"subnetID,omitempty"`
	NodeID       string `protobuf:"bytes,25,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	LocationID   string `protobuf:"bytes,41,opt,name=locationID,proto3" json:"locationID,omitempty"`
	RouterID     string `protobuf:"bytes,42,opt,name=routerID,proto3" json:"routerID,omitempty"`
	Permission   int32  `protobuf:"varint,101,opt,name=permission,proto3" json:"permission,omitempty"`
	Endpoint     string `protobuf:"bytes,201,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *AccessRequest) Reset() {
	*x = AccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequest) ProtoMessage() {}

func (x *AccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequest.ProtoReflect.Descriptor instead.
func (*AccessRequest) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_iam_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *AccessRequest) GetUserRealm() string {
	if x != nil {
		return x.UserRealm
	}
	return ""
}

func (x *AccessRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *AccessRequest) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *AccessRequest) GetSessionToken() string {
	if x != nil {
		return x.SessionToken
	}
	return ""
}

func (x *AccessRequest) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (x *AccessRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *AccessRequest) GetNetID() string {
	if x != nil {
		return x.NetID
	}
	return ""
}

func (x *AccessRequest) GetSubnetID() string {
	if x != nil {
		return x.SubnetID
	}
	return ""
}

func (x *AccessRequest) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *AccessRequest) GetLocationID() string {
	if x != nil {
		return x.LocationID
	}
	return ""
}

func (x *AccessRequest) GetRouterID() string {
	if x != nil {
		return x.RouterID
	}
	return ""
}

func (x *AccessRequest) GetPermission() int32 {
	if x != nil {
		return x.Permission
	}
	return 0
}

func (x *AccessRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

var File_skynx_protobuf_resources_v1_iam_auth_auth_proto protoreflect.FileDescriptor

var file_skynx_protobuf_resources_v1_iam_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x1b, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x4b,
	0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x91, 0x01, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x31, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x6b, 0x65, 0x79, 0x32, 0x12, 0x24, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x10,
	0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5b, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x65, 0x61, 0x6c, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xf8, 0x02, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x61, 0x6c, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x61, 0x6c, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x65, 0x61, 0x6c, 0x6d, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c,
	0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x65,
	0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0xc9, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x42, 0x2b, 0x5a, 0x29, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_skynx_protobuf_resources_v1_iam_auth_auth_proto_rawDescOnce sync.Once
	file_skynx_protobuf_resources_v1_iam_auth_auth_proto_rawDescData = file_skynx_protobuf_resources_v1_iam_auth_auth_proto_rawDesc
)

func file_skynx_protobuf_resources_v1_iam_auth_auth_proto_rawDescGZIP() []byte {
	file_skynx_protobuf_resources_v1_iam_auth_auth_proto_rawDescOnce.Do(func() {
		file_skynx_protobuf_resources_v1_iam_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_skynx_protobuf_resources_v1_iam_auth_auth_proto_rawDescData)
	})
	return file_skynx_protobuf_resources_v1_iam_auth_auth_proto_rawDescData
}

var file_skynx_protobuf_resources_v1_iam_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_skynx_protobuf_resources_v1_iam_auth_auth_proto_goTypes = []interface{}{
	(*AuthKey)(nil),         // 0: auth.AuthKey
	(*AuthToken)(nil),       // 1: auth.AuthToken
	(*AuthCredentials)(nil), // 2: auth.AuthCredentials
	(*AccessRequest)(nil),   // 3: auth.AccessRequest
}
var file_skynx_protobuf_resources_v1_iam_auth_auth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_skynx_protobuf_resources_v1_iam_auth_auth_proto_init() }
func file_skynx_protobuf_resources_v1_iam_auth_auth_proto_init() {
	if File_skynx_protobuf_resources_v1_iam_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skynx_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthKey); i {
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
		file_skynx_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthToken); i {
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
		file_skynx_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthCredentials); i {
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
		file_skynx_protobuf_resources_v1_iam_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessRequest); i {
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
			RawDescriptor: file_skynx_protobuf_resources_v1_iam_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_skynx_protobuf_resources_v1_iam_auth_auth_proto_goTypes,
		DependencyIndexes: file_skynx_protobuf_resources_v1_iam_auth_auth_proto_depIdxs,
		MessageInfos:      file_skynx_protobuf_resources_v1_iam_auth_auth_proto_msgTypes,
	}.Build()
	File_skynx_protobuf_resources_v1_iam_auth_auth_proto = out.File
	file_skynx_protobuf_resources_v1_iam_auth_auth_proto_rawDesc = nil
	file_skynx_protobuf_resources_v1_iam_auth_auth_proto_goTypes = nil
	file_skynx_protobuf_resources_v1_iam_auth_auth_proto_depIdxs = nil
}
