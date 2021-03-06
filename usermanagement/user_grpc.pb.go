// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package usermanagement

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserManagementClient is the client API for UserManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserManagementClient interface {
	// User methods
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	// Project methods
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error)
	AddUserToProject(ctx context.Context, in *AddUserToProjectRequest, opts ...grpc.CallOption) (*AddUserToProjectResponse, error)
	UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectResponse, error)
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error)
	GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectResponse, error)
	GetUserProjects(ctx context.Context, in *GetUserProjectsRequest, opts ...grpc.CallOption) (*GetUserProjectsResponse, error)
	GetUserRoleOnProject(ctx context.Context, in *GetUserRoleOnProjectRequest, opts ...grpc.CallOption) (*GetUserRoleOnProjectResponse, error)
	CheckProjectToken(ctx context.Context, in *CheckProjectTokenRequest, opts ...grpc.CallOption) (*CheckProjectTokenResponse, error)
}

type userManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagementClient(cc grpc.ClientConnInterface) UserManagementClient {
	return &userManagementClient{cc}
}

func (c *userManagementClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	out := new(ValidateTokenResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/ValidateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	out := new(UpdateUserInfoResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/UpdateUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error) {
	out := new(CreateProjectResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) AddUserToProject(ctx context.Context, in *AddUserToProjectRequest, opts ...grpc.CallOption) (*AddUserToProjectResponse, error) {
	out := new(AddUserToProjectResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/AddUserToProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) UpdateProject(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*UpdateProjectResponse, error) {
	out := new(UpdateProjectResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/UpdateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error) {
	out := new(DeleteProjectResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/DeleteProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) GetProject(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*GetProjectResponse, error) {
	out := new(GetProjectResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/GetProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) GetUserProjects(ctx context.Context, in *GetUserProjectsRequest, opts ...grpc.CallOption) (*GetUserProjectsResponse, error) {
	out := new(GetUserProjectsResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/GetUserProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) GetUserRoleOnProject(ctx context.Context, in *GetUserRoleOnProjectRequest, opts ...grpc.CallOption) (*GetUserRoleOnProjectResponse, error) {
	out := new(GetUserRoleOnProjectResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/GetUserRoleOnProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) CheckProjectToken(ctx context.Context, in *CheckProjectTokenRequest, opts ...grpc.CallOption) (*CheckProjectTokenResponse, error) {
	out := new(CheckProjectTokenResponse)
	err := c.cc.Invoke(ctx, "/UserManagement/CheckProjectToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagementServer is the server API for UserManagement service.
// All implementations must embed UnimplementedUserManagementServer
// for forward compatibility
type UserManagementServer interface {
	// User methods
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	// Project methods
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
	AddUserToProject(context.Context, *AddUserToProjectRequest) (*AddUserToProjectResponse, error)
	UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectResponse, error)
	DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error)
	GetProject(context.Context, *GetProjectRequest) (*GetProjectResponse, error)
	GetUserProjects(context.Context, *GetUserProjectsRequest) (*GetUserProjectsResponse, error)
	GetUserRoleOnProject(context.Context, *GetUserRoleOnProjectRequest) (*GetUserRoleOnProjectResponse, error)
	CheckProjectToken(context.Context, *CheckProjectTokenRequest) (*CheckProjectTokenResponse, error)
	mustEmbedUnimplementedUserManagementServer()
}

// UnimplementedUserManagementServer must be embedded to have forward compatible implementations.
type UnimplementedUserManagementServer struct {
}

func (UnimplementedUserManagementServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserManagementServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserManagementServer) ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedUserManagementServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserManagementServer) UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUserManagementServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserManagementServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedUserManagementServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserManagementServer) CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedUserManagementServer) AddUserToProject(context.Context, *AddUserToProjectRequest) (*AddUserToProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToProject not implemented")
}
func (UnimplementedUserManagementServer) UpdateProject(context.Context, *UpdateProjectRequest) (*UpdateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProject not implemented")
}
func (UnimplementedUserManagementServer) DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}
func (UnimplementedUserManagementServer) GetProject(context.Context, *GetProjectRequest) (*GetProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProject not implemented")
}
func (UnimplementedUserManagementServer) GetUserProjects(context.Context, *GetUserProjectsRequest) (*GetUserProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProjects not implemented")
}
func (UnimplementedUserManagementServer) GetUserRoleOnProject(context.Context, *GetUserRoleOnProjectRequest) (*GetUserRoleOnProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRoleOnProject not implemented")
}
func (UnimplementedUserManagementServer) CheckProjectToken(context.Context, *CheckProjectTokenRequest) (*CheckProjectTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckProjectToken not implemented")
}
func (UnimplementedUserManagementServer) mustEmbedUnimplementedUserManagementServer() {}

// UnsafeUserManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserManagementServer will
// result in compilation errors.
type UnsafeUserManagementServer interface {
	mustEmbedUnimplementedUserManagementServer()
}

func RegisterUserManagementServer(s grpc.ServiceRegistrar, srv UserManagementServer) {
	s.RegisterService(&UserManagement_ServiceDesc, srv)
}

func _UserManagement_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/ValidateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/UpdateUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_AddUserToProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserToProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).AddUserToProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/AddUserToProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).AddUserToProject(ctx, req.(*AddUserToProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_UpdateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).UpdateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/UpdateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).UpdateProject(ctx, req.(*UpdateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/DeleteProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).GetProject(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_GetUserProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).GetUserProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/GetUserProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).GetUserProjects(ctx, req.(*GetUserProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_GetUserRoleOnProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRoleOnProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).GetUserRoleOnProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/GetUserRoleOnProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).GetUserRoleOnProject(ctx, req.(*GetUserRoleOnProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_CheckProjectToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckProjectTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).CheckProjectToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserManagement/CheckProjectToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).CheckProjectToken(ctx, req.(*CheckProjectTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserManagement_ServiceDesc is the grpc.ServiceDesc for UserManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserManagement",
	HandlerType: (*UserManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserManagement_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _UserManagement_Logout_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _UserManagement_ValidateToken_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserManagement_Register_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _UserManagement_UpdateUserInfo_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserManagement_DeleteUser_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserManagement_ChangePassword_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserManagement_GetUserInfo_Handler,
		},
		{
			MethodName: "CreateProject",
			Handler:    _UserManagement_CreateProject_Handler,
		},
		{
			MethodName: "AddUserToProject",
			Handler:    _UserManagement_AddUserToProject_Handler,
		},
		{
			MethodName: "UpdateProject",
			Handler:    _UserManagement_UpdateProject_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _UserManagement_DeleteProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _UserManagement_GetProject_Handler,
		},
		{
			MethodName: "GetUserProjects",
			Handler:    _UserManagement_GetUserProjects_Handler,
		},
		{
			MethodName: "GetUserRoleOnProject",
			Handler:    _UserManagement_GetUserRoleOnProject_Handler,
		},
		{
			MethodName: "CheckProjectToken",
			Handler:    _UserManagement_CheckProjectToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
