// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc3
// source: skynx/protobuf/resources/v1/ops/project.proto

package ops

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

type NewProjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID        string `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID         string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	Name             string `protobuf:"bytes,21,opt,name=name,proto3" json:"name,omitempty"`
	Description      string `protobuf:"bytes,22,opt,name=description,proto3" json:"description,omitempty"`
	ReviewRequired   bool   `protobuf:"varint,51,opt,name=reviewRequired,proto3" json:"reviewRequired,omitempty"`
	ApprovalRequired bool   `protobuf:"varint,52,opt,name=approvalRequired,proto3" json:"approvalRequired,omitempty"`
}

func (x *NewProjectRequest) Reset() {
	*x = NewProjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewProjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewProjectRequest) ProtoMessage() {}

func (x *NewProjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewProjectRequest.ProtoReflect.Descriptor instead.
func (*NewProjectRequest) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_ops_project_proto_rawDescGZIP(), []int{0}
}

func (x *NewProjectRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *NewProjectRequest) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *NewProjectRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewProjectRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NewProjectRequest) GetReviewRequired() bool {
	if x != nil {
		return x.ReviewRequired
	}
	return false
}

func (x *NewProjectRequest) GetApprovalRequired() bool {
	if x != nil {
		return x.ApprovalRequired
	}
	return false
}

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID        string `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID         string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	ProjectID        string `protobuf:"bytes,3,opt,name=projectID,proto3" json:"projectID,omitempty"`
	Name             string `protobuf:"bytes,31,opt,name=name,proto3" json:"name,omitempty"`
	Description      string `protobuf:"bytes,32,opt,name=description,proto3" json:"description,omitempty"`
	ReviewRequired   bool   `protobuf:"varint,51,opt,name=reviewRequired,proto3" json:"reviewRequired,omitempty"`
	ApprovalRequired bool   `protobuf:"varint,52,opt,name=approvalRequired,proto3" json:"approvalRequired,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_ops_project_proto_rawDescGZIP(), []int{1}
}

func (x *Project) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *Project) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *Project) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Project) GetReviewRequired() bool {
	if x != nil {
		return x.ReviewRequired
	}
	return false
}

func (x *Project) GetApprovalRequired() bool {
	if x != nil {
		return x.ApprovalRequired
	}
	return false
}

type Projects struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta     *resource.ListResponse `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Projects []*Project             `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
}

func (x *Projects) Reset() {
	*x = Projects{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Projects) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Projects) ProtoMessage() {}

func (x *Projects) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Projects.ProtoReflect.Descriptor instead.
func (*Projects) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_ops_project_proto_rawDescGZIP(), []int{2}
}

func (x *Projects) GetMeta() *resource.ListResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Projects) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

type ListProjectsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta   *resource.ListRequest `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Tenant *tenant.TenantReq     `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
}

func (x *ListProjectsRequest) Reset() {
	*x = ListProjectsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProjectsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProjectsRequest) ProtoMessage() {}

func (x *ListProjectsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProjectsRequest.ProtoReflect.Descriptor instead.
func (*ListProjectsRequest) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_ops_project_proto_rawDescGZIP(), []int{3}
}

func (x *ListProjectsRequest) GetMeta() *resource.ListRequest {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListProjectsRequest) GetTenant() *tenant.TenantReq {
	if x != nil {
		return x.Tenant
	}
	return nil
}

type ProjectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	TenantID  string `protobuf:"bytes,2,opt,name=tenantID,proto3" json:"tenantID,omitempty"`
	ProjectID string `protobuf:"bytes,3,opt,name=projectID,proto3" json:"projectID,omitempty"`
}

func (x *ProjectReq) Reset() {
	*x = ProjectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectReq) ProtoMessage() {}

func (x *ProjectReq) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectReq.ProtoReflect.Descriptor instead.
func (*ProjectReq) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_ops_project_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectReq) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *ProjectReq) GetTenantID() string {
	if x != nil {
		return x.TenantID
	}
	return ""
}

func (x *ProjectReq) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

var File_skynx_protobuf_resources_v1_ops_project_proto protoreflect.FileDescriptor

var file_skynx_protobuf_resources_v1_ops_project_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6f, 0x70, 0x73, 0x1a, 0x2f, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x33,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x34, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10,
	0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x22, 0xeb, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x34, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x60,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x22, 0x6b, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x29, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x54, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x22, 0x64, 0x0a,
	0x0a, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x42, 0x26, 0x5a, 0x24, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2e, 0x69, 0x6f, 0x2f,
	0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6f, 0x70, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_skynx_protobuf_resources_v1_ops_project_proto_rawDescOnce sync.Once
	file_skynx_protobuf_resources_v1_ops_project_proto_rawDescData = file_skynx_protobuf_resources_v1_ops_project_proto_rawDesc
)

func file_skynx_protobuf_resources_v1_ops_project_proto_rawDescGZIP() []byte {
	file_skynx_protobuf_resources_v1_ops_project_proto_rawDescOnce.Do(func() {
		file_skynx_protobuf_resources_v1_ops_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_skynx_protobuf_resources_v1_ops_project_proto_rawDescData)
	})
	return file_skynx_protobuf_resources_v1_ops_project_proto_rawDescData
}

var file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_skynx_protobuf_resources_v1_ops_project_proto_goTypes = []interface{}{
	(*NewProjectRequest)(nil),     // 0: ops.NewProjectRequest
	(*Project)(nil),               // 1: ops.Project
	(*Projects)(nil),              // 2: ops.Projects
	(*ListProjectsRequest)(nil),   // 3: ops.ListProjectsRequest
	(*ProjectReq)(nil),            // 4: ops.ProjectReq
	(*resource.ListResponse)(nil), // 5: resource.ListResponse
	(*resource.ListRequest)(nil),  // 6: resource.ListRequest
	(*tenant.TenantReq)(nil),      // 7: tenant.TenantReq
}
var file_skynx_protobuf_resources_v1_ops_project_proto_depIdxs = []int32{
	5, // 0: ops.Projects.meta:type_name -> resource.ListResponse
	1, // 1: ops.Projects.projects:type_name -> ops.Project
	6, // 2: ops.ListProjectsRequest.meta:type_name -> resource.ListRequest
	7, // 3: ops.ListProjectsRequest.tenant:type_name -> tenant.TenantReq
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_skynx_protobuf_resources_v1_ops_project_proto_init() }
func file_skynx_protobuf_resources_v1_ops_project_proto_init() {
	if File_skynx_protobuf_resources_v1_ops_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewProjectRequest); i {
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
		file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Projects); i {
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
		file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProjectsRequest); i {
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
		file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectReq); i {
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
			RawDescriptor: file_skynx_protobuf_resources_v1_ops_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_skynx_protobuf_resources_v1_ops_project_proto_goTypes,
		DependencyIndexes: file_skynx_protobuf_resources_v1_ops_project_proto_depIdxs,
		MessageInfos:      file_skynx_protobuf_resources_v1_ops_project_proto_msgTypes,
	}.Build()
	File_skynx_protobuf_resources_v1_ops_project_proto = out.File
	file_skynx_protobuf_resources_v1_ops_project_proto_rawDesc = nil
	file_skynx_protobuf_resources_v1_ops_project_proto_goTypes = nil
	file_skynx_protobuf_resources_v1_ops_project_proto_depIdxs = nil
}
