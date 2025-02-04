// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0--rc3
// source: skynx/protobuf/rpc/v1/servicesAPI.proto

package rpc

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

// ServicesAPIClient is the client API for ServicesAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// [s-api] ServicesAPI Definition: Billed services
type ServicesAPIClient interface {
}

type servicesAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewServicesAPIClient(cc grpc.ClientConnInterface) ServicesAPIClient {
	return &servicesAPIClient{cc}
}

// ServicesAPIServer is the server API for ServicesAPI service.
// All implementations must embed UnimplementedServicesAPIServer
// for forward compatibility
//
// [s-api] ServicesAPI Definition: Billed services
type ServicesAPIServer interface {
	mustEmbedUnimplementedServicesAPIServer()
}

// UnimplementedServicesAPIServer must be embedded to have forward compatible implementations.
type UnimplementedServicesAPIServer struct {
}

func (UnimplementedServicesAPIServer) mustEmbedUnimplementedServicesAPIServer() {}

// UnsafeServicesAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServicesAPIServer will
// result in compilation errors.
type UnsafeServicesAPIServer interface {
	mustEmbedUnimplementedServicesAPIServer()
}

func RegisterServicesAPIServer(s grpc.ServiceRegistrar, srv ServicesAPIServer) {
	s.RegisterService(&ServicesAPI_ServiceDesc, srv)
}

// ServicesAPI_ServiceDesc is the grpc.ServiceDesc for ServicesAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServicesAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ServicesAPI",
	HandlerType: (*ServicesAPIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "skynx/protobuf/rpc/v1/servicesAPI.proto",
}
