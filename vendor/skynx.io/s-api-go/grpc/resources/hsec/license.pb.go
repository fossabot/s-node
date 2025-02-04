// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc3
// source: skynx/protobuf/resources/v1/hsec/license.proto

package hsec

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

type DetectedLicense struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Severity   string  `protobuf:"bytes,11,opt,name=severity,proto3" json:"severity,omitempty"`
	Category   string  `protobuf:"bytes,21,opt,name=category,proto3" json:"category,omitempty"` // licenseCategory
	PkgName    string  `protobuf:"bytes,31,opt,name=pkgName,proto3" json:"pkgName,omitempty"`
	FilePath   string  `protobuf:"bytes,41,opt,name=filePath,proto3" json:"filePath,omitempty"`
	Name       string  `protobuf:"bytes,51,opt,name=name,proto3" json:"name,omitempty"`
	Confidence float64 `protobuf:"fixed64,61,opt,name=confidence,proto3" json:"confidence,omitempty"`
	Link       string  `protobuf:"bytes,71,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *DetectedLicense) Reset() {
	*x = DetectedLicense{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_resources_v1_hsec_license_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetectedLicense) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetectedLicense) ProtoMessage() {}

func (x *DetectedLicense) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_resources_v1_hsec_license_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetectedLicense.ProtoReflect.Descriptor instead.
func (*DetectedLicense) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_resources_v1_hsec_license_proto_rawDescGZIP(), []int{0}
}

func (x *DetectedLicense) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *DetectedLicense) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *DetectedLicense) GetPkgName() string {
	if x != nil {
		return x.PkgName
	}
	return ""
}

func (x *DetectedLicense) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *DetectedLicense) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DetectedLicense) GetConfidence() float64 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

func (x *DetectedLicense) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_skynx_protobuf_resources_v1_hsec_license_proto protoreflect.FileDescriptor

var file_skynx_protobuf_resources_v1_hsec_license_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x73,
	0x65, 0x63, 0x2f, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x68, 0x73, 0x65, 0x63, 0x22, 0xc7, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6b, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6b, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x47, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x42, 0x27, 0x5a, 0x25, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x2d, 0x61,
	0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x68, 0x73, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_skynx_protobuf_resources_v1_hsec_license_proto_rawDescOnce sync.Once
	file_skynx_protobuf_resources_v1_hsec_license_proto_rawDescData = file_skynx_protobuf_resources_v1_hsec_license_proto_rawDesc
)

func file_skynx_protobuf_resources_v1_hsec_license_proto_rawDescGZIP() []byte {
	file_skynx_protobuf_resources_v1_hsec_license_proto_rawDescOnce.Do(func() {
		file_skynx_protobuf_resources_v1_hsec_license_proto_rawDescData = protoimpl.X.CompressGZIP(file_skynx_protobuf_resources_v1_hsec_license_proto_rawDescData)
	})
	return file_skynx_protobuf_resources_v1_hsec_license_proto_rawDescData
}

var file_skynx_protobuf_resources_v1_hsec_license_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_skynx_protobuf_resources_v1_hsec_license_proto_goTypes = []interface{}{
	(*DetectedLicense)(nil), // 0: hsec.DetectedLicense
}
var file_skynx_protobuf_resources_v1_hsec_license_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_skynx_protobuf_resources_v1_hsec_license_proto_init() }
func file_skynx_protobuf_resources_v1_hsec_license_proto_init() {
	if File_skynx_protobuf_resources_v1_hsec_license_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skynx_protobuf_resources_v1_hsec_license_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetectedLicense); i {
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
			RawDescriptor: file_skynx_protobuf_resources_v1_hsec_license_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_skynx_protobuf_resources_v1_hsec_license_proto_goTypes,
		DependencyIndexes: file_skynx_protobuf_resources_v1_hsec_license_proto_depIdxs,
		MessageInfos:      file_skynx_protobuf_resources_v1_hsec_license_proto_msgTypes,
	}.Build()
	File_skynx_protobuf_resources_v1_hsec_license_proto = out.File
	file_skynx_protobuf_resources_v1_hsec_license_proto_rawDesc = nil
	file_skynx_protobuf_resources_v1_hsec_license_proto_goTypes = nil
	file_skynx_protobuf_resources_v1_hsec_license_proto_depIdxs = nil
}
