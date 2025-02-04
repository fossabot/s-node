// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0--rc3
// source: skynx/protobuf/rpc/v1/iamAPI.proto

package rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	empty "skynx.io/s-api-go/grpc/common/empty"
	iam "skynx.io/s-api-go/grpc/resources/iam"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	IAMAPI_GetRBAC_FullMethodName             = "/api.IAMAPI/GetRBAC"
	IAMAPI_ListIAMPermissions_FullMethodName  = "/api.IAMAPI/ListIAMPermissions"
	IAMAPI_CreateUser_FullMethodName          = "/api.IAMAPI/CreateUser"
	IAMAPI_ListUsers_FullMethodName           = "/api.IAMAPI/ListUsers"
	IAMAPI_GetUser_FullMethodName             = "/api.IAMAPI/GetUser"
	IAMAPI_DeleteUser_FullMethodName          = "/api.IAMAPI/DeleteUser"
	IAMAPI_EnableUser_FullMethodName          = "/api.IAMAPI/EnableUser"
	IAMAPI_DisableUser_FullMethodName         = "/api.IAMAPI/DisableUser"
	IAMAPI_NewUserToken_FullMethodName        = "/api.IAMAPI/NewUserToken"
	IAMAPI_SetUserPermissions_FullMethodName  = "/api.IAMAPI/SetUserPermissions"
	IAMAPI_ListSecurityGroups_FullMethodName  = "/api.IAMAPI/ListSecurityGroups"
	IAMAPI_GetSecurityGroup_FullMethodName    = "/api.IAMAPI/GetSecurityGroup"
	IAMAPI_SetSecurityGroup_FullMethodName    = "/api.IAMAPI/SetSecurityGroup"
	IAMAPI_DeleteSecurityGroup_FullMethodName = "/api.IAMAPI/DeleteSecurityGroup"
	IAMAPI_ListRoles_FullMethodName           = "/api.IAMAPI/ListRoles"
	IAMAPI_GetRole_FullMethodName             = "/api.IAMAPI/GetRole"
	IAMAPI_SetRole_FullMethodName             = "/api.IAMAPI/SetRole"
	IAMAPI_DeleteRole_FullMethodName          = "/api.IAMAPI/DeleteRole"
	IAMAPI_ListACLs_FullMethodName            = "/api.IAMAPI/ListACLs"
	IAMAPI_GetACL_FullMethodName              = "/api.IAMAPI/GetACL"
	IAMAPI_SetACL_FullMethodName              = "/api.IAMAPI/SetACL"
	IAMAPI_DeleteACL_FullMethodName           = "/api.IAMAPI/DeleteACL"
)

// IAMAPIClient is the client API for IAMAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// [iam-api] iamAPI Definition: IAM Resources
type IAMAPIClient interface {
	// iam
	GetRBAC(ctx context.Context, in *iam.RBACRequest, opts ...grpc.CallOption) (*iam.RBAC, error)
	// permissions
	ListIAMPermissions(ctx context.Context, in *empty.Request, opts ...grpc.CallOption) (*iam.Permissions, error)
	// user
	CreateUser(ctx context.Context, in *iam.NewUserRequest, opts ...grpc.CallOption) (*iam.User, error)
	ListUsers(ctx context.Context, in *iam.ListUsersRequest, opts ...grpc.CallOption) (*iam.Users, error)
	GetUser(ctx context.Context, in *iam.UserReq, opts ...grpc.CallOption) (*iam.User, error)
	DeleteUser(ctx context.Context, in *iam.UserReq, opts ...grpc.CallOption) (*empty.Response, error)
	EnableUser(ctx context.Context, in *iam.UserReq, opts ...grpc.CallOption) (*iam.User, error)
	DisableUser(ctx context.Context, in *iam.UserReq, opts ...grpc.CallOption) (*iam.User, error)
	NewUserToken(ctx context.Context, in *iam.UserReq, opts ...grpc.CallOption) (*iam.User, error)
	//	rpc SetUserEmail(iam.SetUserEmailRequest) returns (iam.User) {
	//	  option (google.api.http) = {
	//	    patch: "/api/v1/accounts/{accountID}/iam/users/{userID}:email"
	//	    body: "*"
	//	  };
	//	}
	SetUserPermissions(ctx context.Context, in *iam.SetUserPermissionsRequest, opts ...grpc.CallOption) (*iam.User, error)
	// security-group
	ListSecurityGroups(ctx context.Context, in *iam.ListSecurityGroupsRequest, opts ...grpc.CallOption) (*iam.SecurityGroups, error)
	GetSecurityGroup(ctx context.Context, in *iam.SecurityGroup, opts ...grpc.CallOption) (*iam.SecurityGroup, error)
	SetSecurityGroup(ctx context.Context, in *iam.SecurityGroup, opts ...grpc.CallOption) (*iam.SecurityGroup, error)
	DeleteSecurityGroup(ctx context.Context, in *iam.SecurityGroup, opts ...grpc.CallOption) (*empty.Response, error)
	// role
	ListRoles(ctx context.Context, in *iam.ListRolesRequest, opts ...grpc.CallOption) (*iam.Roles, error)
	GetRole(ctx context.Context, in *iam.Role, opts ...grpc.CallOption) (*iam.Role, error)
	SetRole(ctx context.Context, in *iam.Role, opts ...grpc.CallOption) (*iam.Role, error)
	DeleteRole(ctx context.Context, in *iam.Role, opts ...grpc.CallOption) (*empty.Response, error)
	// acl
	ListACLs(ctx context.Context, in *iam.ListACLsRequest, opts ...grpc.CallOption) (*iam.ACLs, error)
	GetACL(ctx context.Context, in *iam.ACL, opts ...grpc.CallOption) (*iam.ACL, error)
	SetACL(ctx context.Context, in *iam.ACL, opts ...grpc.CallOption) (*iam.ACL, error)
	DeleteACL(ctx context.Context, in *iam.ACL, opts ...grpc.CallOption) (*empty.Response, error)
}

type iAMAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewIAMAPIClient(cc grpc.ClientConnInterface) IAMAPIClient {
	return &iAMAPIClient{cc}
}

func (c *iAMAPIClient) GetRBAC(ctx context.Context, in *iam.RBACRequest, opts ...grpc.CallOption) (*iam.RBAC, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.RBAC)
	err := c.cc.Invoke(ctx, IAMAPI_GetRBAC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) ListIAMPermissions(ctx context.Context, in *empty.Request, opts ...grpc.CallOption) (*iam.Permissions, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.Permissions)
	err := c.cc.Invoke(ctx, IAMAPI_ListIAMPermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) CreateUser(ctx context.Context, in *iam.NewUserRequest, opts ...grpc.CallOption) (*iam.User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.User)
	err := c.cc.Invoke(ctx, IAMAPI_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) ListUsers(ctx context.Context, in *iam.ListUsersRequest, opts ...grpc.CallOption) (*iam.Users, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.Users)
	err := c.cc.Invoke(ctx, IAMAPI_ListUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) GetUser(ctx context.Context, in *iam.UserReq, opts ...grpc.CallOption) (*iam.User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.User)
	err := c.cc.Invoke(ctx, IAMAPI_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) DeleteUser(ctx context.Context, in *iam.UserReq, opts ...grpc.CallOption) (*empty.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Response)
	err := c.cc.Invoke(ctx, IAMAPI_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) EnableUser(ctx context.Context, in *iam.UserReq, opts ...grpc.CallOption) (*iam.User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.User)
	err := c.cc.Invoke(ctx, IAMAPI_EnableUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) DisableUser(ctx context.Context, in *iam.UserReq, opts ...grpc.CallOption) (*iam.User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.User)
	err := c.cc.Invoke(ctx, IAMAPI_DisableUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) NewUserToken(ctx context.Context, in *iam.UserReq, opts ...grpc.CallOption) (*iam.User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.User)
	err := c.cc.Invoke(ctx, IAMAPI_NewUserToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) SetUserPermissions(ctx context.Context, in *iam.SetUserPermissionsRequest, opts ...grpc.CallOption) (*iam.User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.User)
	err := c.cc.Invoke(ctx, IAMAPI_SetUserPermissions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) ListSecurityGroups(ctx context.Context, in *iam.ListSecurityGroupsRequest, opts ...grpc.CallOption) (*iam.SecurityGroups, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.SecurityGroups)
	err := c.cc.Invoke(ctx, IAMAPI_ListSecurityGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) GetSecurityGroup(ctx context.Context, in *iam.SecurityGroup, opts ...grpc.CallOption) (*iam.SecurityGroup, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.SecurityGroup)
	err := c.cc.Invoke(ctx, IAMAPI_GetSecurityGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) SetSecurityGroup(ctx context.Context, in *iam.SecurityGroup, opts ...grpc.CallOption) (*iam.SecurityGroup, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.SecurityGroup)
	err := c.cc.Invoke(ctx, IAMAPI_SetSecurityGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) DeleteSecurityGroup(ctx context.Context, in *iam.SecurityGroup, opts ...grpc.CallOption) (*empty.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Response)
	err := c.cc.Invoke(ctx, IAMAPI_DeleteSecurityGroup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) ListRoles(ctx context.Context, in *iam.ListRolesRequest, opts ...grpc.CallOption) (*iam.Roles, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.Roles)
	err := c.cc.Invoke(ctx, IAMAPI_ListRoles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) GetRole(ctx context.Context, in *iam.Role, opts ...grpc.CallOption) (*iam.Role, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.Role)
	err := c.cc.Invoke(ctx, IAMAPI_GetRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) SetRole(ctx context.Context, in *iam.Role, opts ...grpc.CallOption) (*iam.Role, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.Role)
	err := c.cc.Invoke(ctx, IAMAPI_SetRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) DeleteRole(ctx context.Context, in *iam.Role, opts ...grpc.CallOption) (*empty.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Response)
	err := c.cc.Invoke(ctx, IAMAPI_DeleteRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) ListACLs(ctx context.Context, in *iam.ListACLsRequest, opts ...grpc.CallOption) (*iam.ACLs, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.ACLs)
	err := c.cc.Invoke(ctx, IAMAPI_ListACLs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) GetACL(ctx context.Context, in *iam.ACL, opts ...grpc.CallOption) (*iam.ACL, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.ACL)
	err := c.cc.Invoke(ctx, IAMAPI_GetACL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) SetACL(ctx context.Context, in *iam.ACL, opts ...grpc.CallOption) (*iam.ACL, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(iam.ACL)
	err := c.cc.Invoke(ctx, IAMAPI_SetACL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAPIClient) DeleteACL(ctx context.Context, in *iam.ACL, opts ...grpc.CallOption) (*empty.Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Response)
	err := c.cc.Invoke(ctx, IAMAPI_DeleteACL_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IAMAPIServer is the server API for IAMAPI service.
// All implementations must embed UnimplementedIAMAPIServer
// for forward compatibility
//
// [iam-api] iamAPI Definition: IAM Resources
type IAMAPIServer interface {
	// iam
	GetRBAC(context.Context, *iam.RBACRequest) (*iam.RBAC, error)
	// permissions
	ListIAMPermissions(context.Context, *empty.Request) (*iam.Permissions, error)
	// user
	CreateUser(context.Context, *iam.NewUserRequest) (*iam.User, error)
	ListUsers(context.Context, *iam.ListUsersRequest) (*iam.Users, error)
	GetUser(context.Context, *iam.UserReq) (*iam.User, error)
	DeleteUser(context.Context, *iam.UserReq) (*empty.Response, error)
	EnableUser(context.Context, *iam.UserReq) (*iam.User, error)
	DisableUser(context.Context, *iam.UserReq) (*iam.User, error)
	NewUserToken(context.Context, *iam.UserReq) (*iam.User, error)
	//	rpc SetUserEmail(iam.SetUserEmailRequest) returns (iam.User) {
	//	  option (google.api.http) = {
	//	    patch: "/api/v1/accounts/{accountID}/iam/users/{userID}:email"
	//	    body: "*"
	//	  };
	//	}
	SetUserPermissions(context.Context, *iam.SetUserPermissionsRequest) (*iam.User, error)
	// security-group
	ListSecurityGroups(context.Context, *iam.ListSecurityGroupsRequest) (*iam.SecurityGroups, error)
	GetSecurityGroup(context.Context, *iam.SecurityGroup) (*iam.SecurityGroup, error)
	SetSecurityGroup(context.Context, *iam.SecurityGroup) (*iam.SecurityGroup, error)
	DeleteSecurityGroup(context.Context, *iam.SecurityGroup) (*empty.Response, error)
	// role
	ListRoles(context.Context, *iam.ListRolesRequest) (*iam.Roles, error)
	GetRole(context.Context, *iam.Role) (*iam.Role, error)
	SetRole(context.Context, *iam.Role) (*iam.Role, error)
	DeleteRole(context.Context, *iam.Role) (*empty.Response, error)
	// acl
	ListACLs(context.Context, *iam.ListACLsRequest) (*iam.ACLs, error)
	GetACL(context.Context, *iam.ACL) (*iam.ACL, error)
	SetACL(context.Context, *iam.ACL) (*iam.ACL, error)
	DeleteACL(context.Context, *iam.ACL) (*empty.Response, error)
	mustEmbedUnimplementedIAMAPIServer()
}

// UnimplementedIAMAPIServer must be embedded to have forward compatible implementations.
type UnimplementedIAMAPIServer struct {
}

func (UnimplementedIAMAPIServer) GetRBAC(context.Context, *iam.RBACRequest) (*iam.RBAC, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRBAC not implemented")
}
func (UnimplementedIAMAPIServer) ListIAMPermissions(context.Context, *empty.Request) (*iam.Permissions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIAMPermissions not implemented")
}
func (UnimplementedIAMAPIServer) CreateUser(context.Context, *iam.NewUserRequest) (*iam.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedIAMAPIServer) ListUsers(context.Context, *iam.ListUsersRequest) (*iam.Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedIAMAPIServer) GetUser(context.Context, *iam.UserReq) (*iam.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedIAMAPIServer) DeleteUser(context.Context, *iam.UserReq) (*empty.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedIAMAPIServer) EnableUser(context.Context, *iam.UserReq) (*iam.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableUser not implemented")
}
func (UnimplementedIAMAPIServer) DisableUser(context.Context, *iam.UserReq) (*iam.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableUser not implemented")
}
func (UnimplementedIAMAPIServer) NewUserToken(context.Context, *iam.UserReq) (*iam.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewUserToken not implemented")
}
func (UnimplementedIAMAPIServer) SetUserPermissions(context.Context, *iam.SetUserPermissionsRequest) (*iam.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserPermissions not implemented")
}
func (UnimplementedIAMAPIServer) ListSecurityGroups(context.Context, *iam.ListSecurityGroupsRequest) (*iam.SecurityGroups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSecurityGroups not implemented")
}
func (UnimplementedIAMAPIServer) GetSecurityGroup(context.Context, *iam.SecurityGroup) (*iam.SecurityGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecurityGroup not implemented")
}
func (UnimplementedIAMAPIServer) SetSecurityGroup(context.Context, *iam.SecurityGroup) (*iam.SecurityGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSecurityGroup not implemented")
}
func (UnimplementedIAMAPIServer) DeleteSecurityGroup(context.Context, *iam.SecurityGroup) (*empty.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecurityGroup not implemented")
}
func (UnimplementedIAMAPIServer) ListRoles(context.Context, *iam.ListRolesRequest) (*iam.Roles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoles not implemented")
}
func (UnimplementedIAMAPIServer) GetRole(context.Context, *iam.Role) (*iam.Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedIAMAPIServer) SetRole(context.Context, *iam.Role) (*iam.Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRole not implemented")
}
func (UnimplementedIAMAPIServer) DeleteRole(context.Context, *iam.Role) (*empty.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedIAMAPIServer) ListACLs(context.Context, *iam.ListACLsRequest) (*iam.ACLs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListACLs not implemented")
}
func (UnimplementedIAMAPIServer) GetACL(context.Context, *iam.ACL) (*iam.ACL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetACL not implemented")
}
func (UnimplementedIAMAPIServer) SetACL(context.Context, *iam.ACL) (*iam.ACL, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetACL not implemented")
}
func (UnimplementedIAMAPIServer) DeleteACL(context.Context, *iam.ACL) (*empty.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteACL not implemented")
}
func (UnimplementedIAMAPIServer) mustEmbedUnimplementedIAMAPIServer() {}

// UnsafeIAMAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IAMAPIServer will
// result in compilation errors.
type UnsafeIAMAPIServer interface {
	mustEmbedUnimplementedIAMAPIServer()
}

func RegisterIAMAPIServer(s grpc.ServiceRegistrar, srv IAMAPIServer) {
	s.RegisterService(&IAMAPI_ServiceDesc, srv)
}

func _IAMAPI_GetRBAC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.RBACRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).GetRBAC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_GetRBAC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).GetRBAC(ctx, req.(*iam.RBACRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_ListIAMPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).ListIAMPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_ListIAMPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).ListIAMPermissions(ctx, req.(*empty.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.NewUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).CreateUser(ctx, req.(*iam.NewUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_ListUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).ListUsers(ctx, req.(*iam.ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).GetUser(ctx, req.(*iam.UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).DeleteUser(ctx, req.(*iam.UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_EnableUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).EnableUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_EnableUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).EnableUser(ctx, req.(*iam.UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_DisableUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).DisableUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_DisableUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).DisableUser(ctx, req.(*iam.UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_NewUserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.UserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).NewUserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_NewUserToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).NewUserToken(ctx, req.(*iam.UserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_SetUserPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.SetUserPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).SetUserPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_SetUserPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).SetUserPermissions(ctx, req.(*iam.SetUserPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_ListSecurityGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.ListSecurityGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).ListSecurityGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_ListSecurityGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).ListSecurityGroups(ctx, req.(*iam.ListSecurityGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_GetSecurityGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.SecurityGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).GetSecurityGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_GetSecurityGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).GetSecurityGroup(ctx, req.(*iam.SecurityGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_SetSecurityGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.SecurityGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).SetSecurityGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_SetSecurityGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).SetSecurityGroup(ctx, req.(*iam.SecurityGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_DeleteSecurityGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.SecurityGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).DeleteSecurityGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_DeleteSecurityGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).DeleteSecurityGroup(ctx, req.(*iam.SecurityGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_ListRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.ListRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).ListRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_ListRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).ListRoles(ctx, req.(*iam.ListRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_GetRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).GetRole(ctx, req.(*iam.Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_SetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).SetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_SetRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).SetRole(ctx, req.(*iam.Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_DeleteRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).DeleteRole(ctx, req.(*iam.Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_ListACLs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.ListACLsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).ListACLs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_ListACLs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).ListACLs(ctx, req.(*iam.ListACLsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_GetACL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.ACL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).GetACL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_GetACL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).GetACL(ctx, req.(*iam.ACL))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_SetACL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.ACL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).SetACL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_SetACL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).SetACL(ctx, req.(*iam.ACL))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAPI_DeleteACL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(iam.ACL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAPIServer).DeleteACL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IAMAPI_DeleteACL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAPIServer).DeleteACL(ctx, req.(*iam.ACL))
	}
	return interceptor(ctx, in, info, handler)
}

// IAMAPI_ServiceDesc is the grpc.ServiceDesc for IAMAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IAMAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.IAMAPI",
	HandlerType: (*IAMAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRBAC",
			Handler:    _IAMAPI_GetRBAC_Handler,
		},
		{
			MethodName: "ListIAMPermissions",
			Handler:    _IAMAPI_ListIAMPermissions_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _IAMAPI_CreateUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _IAMAPI_ListUsers_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _IAMAPI_GetUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _IAMAPI_DeleteUser_Handler,
		},
		{
			MethodName: "EnableUser",
			Handler:    _IAMAPI_EnableUser_Handler,
		},
		{
			MethodName: "DisableUser",
			Handler:    _IAMAPI_DisableUser_Handler,
		},
		{
			MethodName: "NewUserToken",
			Handler:    _IAMAPI_NewUserToken_Handler,
		},
		{
			MethodName: "SetUserPermissions",
			Handler:    _IAMAPI_SetUserPermissions_Handler,
		},
		{
			MethodName: "ListSecurityGroups",
			Handler:    _IAMAPI_ListSecurityGroups_Handler,
		},
		{
			MethodName: "GetSecurityGroup",
			Handler:    _IAMAPI_GetSecurityGroup_Handler,
		},
		{
			MethodName: "SetSecurityGroup",
			Handler:    _IAMAPI_SetSecurityGroup_Handler,
		},
		{
			MethodName: "DeleteSecurityGroup",
			Handler:    _IAMAPI_DeleteSecurityGroup_Handler,
		},
		{
			MethodName: "ListRoles",
			Handler:    _IAMAPI_ListRoles_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _IAMAPI_GetRole_Handler,
		},
		{
			MethodName: "SetRole",
			Handler:    _IAMAPI_SetRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _IAMAPI_DeleteRole_Handler,
		},
		{
			MethodName: "ListACLs",
			Handler:    _IAMAPI_ListACLs_Handler,
		},
		{
			MethodName: "GetACL",
			Handler:    _IAMAPI_GetACL_Handler,
		},
		{
			MethodName: "SetACL",
			Handler:    _IAMAPI_SetACL_Handler,
		},
		{
			MethodName: "DeleteACL",
			Handler:    _IAMAPI_DeleteACL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "skynx/protobuf/rpc/v1/iamAPI.proto",
}
