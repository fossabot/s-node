// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc3
// source: skynx/protobuf/resources/v1/services/thirdParty/saas.proto

package thirdParty

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

type GitHub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled        bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	AccessToken    string `protobuf:"bytes,11,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	WebhooksSecret string `protobuf:"bytes,12,opt,name=webhooksSecret,proto3" json:"webhooksSecret,omitempty"`
}

func (x *GitHub) Reset() {
	*x = GitHub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitHub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitHub) ProtoMessage() {}

func (x *GitHub) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitHub.ProtoReflect.Descriptor instead.
func (*GitHub) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_rawDescGZIP(), []int{0}
}

func (x *GitHub) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *GitHub) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *GitHub) GetWebhooksSecret() string {
	if x != nil {
		return x.WebhooksSecret
	}
	return ""
}

type PagerDuty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled        bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	IntegrationKey string `protobuf:"bytes,11,opt,name=integrationKey,proto3" json:"integrationKey,omitempty"`
}

func (x *PagerDuty) Reset() {
	*x = PagerDuty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagerDuty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagerDuty) ProtoMessage() {}

func (x *PagerDuty) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagerDuty.ProtoReflect.Descriptor instead.
func (*PagerDuty) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_rawDescGZIP(), []int{1}
}

func (x *PagerDuty) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *PagerDuty) GetIntegrationKey() string {
	if x != nil {
		return x.IntegrationKey
	}
	return ""
}

type Slack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled       bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	AlertsWebhook string `protobuf:"bytes,11,opt,name=alertsWebhook,proto3" json:"alertsWebhook,omitempty"`
	OpsWebhook    string `protobuf:"bytes,21,opt,name=opsWebhook,proto3" json:"opsWebhook,omitempty"`
}

func (x *Slack) Reset() {
	*x = Slack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Slack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Slack) ProtoMessage() {}

func (x *Slack) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Slack.ProtoReflect.Descriptor instead.
func (*Slack) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_rawDescGZIP(), []int{2}
}

func (x *Slack) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Slack) GetAlertsWebhook() string {
	if x != nil {
		return x.AlertsWebhook
	}
	return ""
}

func (x *Slack) GetOpsWebhook() string {
	if x != nil {
		return x.OpsWebhook
	}
	return ""
}

var File_skynx_protobuf_resources_v1_services_thirdParty_saas_proto protoreflect.FileDescriptor

var file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x2f, 0x73, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x68,
	0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x22, 0x6c, 0x0a, 0x06, 0x47, 0x69, 0x74, 0x48,
	0x75, 0x62, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26,
	0x0a, 0x0e, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x4d, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x72, 0x44,
	0x75, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x26, 0x0a,
	0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x22, 0x67, 0x0a, 0x05, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x1e,
	0x0a, 0x0a, 0x6f, 0x70, 0x73, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6f, 0x70, 0x73, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x42, 0x36,
	0x5a, 0x34, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x2d, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_rawDescOnce sync.Once
	file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_rawDescData = file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_rawDesc
)

func file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_rawDescGZIP() []byte {
	file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_rawDescOnce.Do(func() {
		file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_rawDescData = protoimpl.X.CompressGZIP(file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_rawDescData)
	})
	return file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_rawDescData
}

var file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_goTypes = []interface{}{
	(*GitHub)(nil),    // 0: thirdParty.GitHub
	(*PagerDuty)(nil), // 1: thirdParty.PagerDuty
	(*Slack)(nil),     // 2: thirdParty.Slack
}
var file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_init() }
func file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_init() {
	if File_skynx_protobuf_resources_v1_services_thirdParty_saas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitHub); i {
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
		file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagerDuty); i {
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
		file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Slack); i {
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
			RawDescriptor: file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_goTypes,
		DependencyIndexes: file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_depIdxs,
		MessageInfos:      file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_msgTypes,
	}.Build()
	File_skynx_protobuf_resources_v1_services_thirdParty_saas_proto = out.File
	file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_rawDesc = nil
	file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_goTypes = nil
	file_skynx_protobuf_resources_v1_services_thirdParty_saas_proto_depIdxs = nil
}
