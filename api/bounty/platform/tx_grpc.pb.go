// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: bounty/platform/tx.proto

package platform

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

const (
	Msg_UpdateParams_FullMethodName   = "/bounty.platform.Msg/UpdateParams"
	Msg_CreatePlatform_FullMethodName = "/bounty.platform.Msg/CreatePlatform"
	Msg_UpdatePlatform_FullMethodName = "/bounty.platform.Msg/UpdatePlatform"
	Msg_DeletePlatform_FullMethodName = "/bounty.platform.Msg/DeletePlatform"
	Msg_CreateClaim_FullMethodName    = "/bounty.platform.Msg/CreateClaim"
	Msg_UpdateClaim_FullMethodName    = "/bounty.platform.Msg/UpdateClaim"
	Msg_DeleteClaim_FullMethodName    = "/bounty.platform.Msg/DeleteClaim"
	Msg_CreateBounty_FullMethodName   = "/bounty.platform.Msg/CreateBounty"
	Msg_ClaimBounty_FullMethodName    = "/bounty.platform.Msg/ClaimBounty"
	Msg_CompleteBounty_FullMethodName = "/bounty.platform.Msg/CompleteBounty"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreatePlatform(ctx context.Context, in *MsgCreatePlatform, opts ...grpc.CallOption) (*MsgCreatePlatformResponse, error)
	UpdatePlatform(ctx context.Context, in *MsgUpdatePlatform, opts ...grpc.CallOption) (*MsgUpdatePlatformResponse, error)
	DeletePlatform(ctx context.Context, in *MsgDeletePlatform, opts ...grpc.CallOption) (*MsgDeletePlatformResponse, error)
	CreateClaim(ctx context.Context, in *MsgCreateClaim, opts ...grpc.CallOption) (*MsgCreateClaimResponse, error)
	UpdateClaim(ctx context.Context, in *MsgUpdateClaim, opts ...grpc.CallOption) (*MsgUpdateClaimResponse, error)
	DeleteClaim(ctx context.Context, in *MsgDeleteClaim, opts ...grpc.CallOption) (*MsgDeleteClaimResponse, error)
	CreateBounty(ctx context.Context, in *MsgCreateBounty, opts ...grpc.CallOption) (*MsgCreateBountyResponse, error)
	ClaimBounty(ctx context.Context, in *MsgClaimBounty, opts ...grpc.CallOption) (*MsgClaimBountyResponse, error)
	CompleteBounty(ctx context.Context, in *MsgCompleteBounty, opts ...grpc.CallOption) (*MsgCompleteBountyResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePlatform(ctx context.Context, in *MsgCreatePlatform, opts ...grpc.CallOption) (*MsgCreatePlatformResponse, error) {
	out := new(MsgCreatePlatformResponse)
	err := c.cc.Invoke(ctx, Msg_CreatePlatform_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdatePlatform(ctx context.Context, in *MsgUpdatePlatform, opts ...grpc.CallOption) (*MsgUpdatePlatformResponse, error) {
	out := new(MsgUpdatePlatformResponse)
	err := c.cc.Invoke(ctx, Msg_UpdatePlatform_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeletePlatform(ctx context.Context, in *MsgDeletePlatform, opts ...grpc.CallOption) (*MsgDeletePlatformResponse, error) {
	out := new(MsgDeletePlatformResponse)
	err := c.cc.Invoke(ctx, Msg_DeletePlatform_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateClaim(ctx context.Context, in *MsgCreateClaim, opts ...grpc.CallOption) (*MsgCreateClaimResponse, error) {
	out := new(MsgCreateClaimResponse)
	err := c.cc.Invoke(ctx, Msg_CreateClaim_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateClaim(ctx context.Context, in *MsgUpdateClaim, opts ...grpc.CallOption) (*MsgUpdateClaimResponse, error) {
	out := new(MsgUpdateClaimResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateClaim_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteClaim(ctx context.Context, in *MsgDeleteClaim, opts ...grpc.CallOption) (*MsgDeleteClaimResponse, error) {
	out := new(MsgDeleteClaimResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteClaim_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateBounty(ctx context.Context, in *MsgCreateBounty, opts ...grpc.CallOption) (*MsgCreateBountyResponse, error) {
	out := new(MsgCreateBountyResponse)
	err := c.cc.Invoke(ctx, Msg_CreateBounty_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClaimBounty(ctx context.Context, in *MsgClaimBounty, opts ...grpc.CallOption) (*MsgClaimBountyResponse, error) {
	out := new(MsgClaimBountyResponse)
	err := c.cc.Invoke(ctx, Msg_ClaimBounty_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CompleteBounty(ctx context.Context, in *MsgCompleteBounty, opts ...grpc.CallOption) (*MsgCompleteBountyResponse, error) {
	out := new(MsgCompleteBountyResponse)
	err := c.cc.Invoke(ctx, Msg_CompleteBounty_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreatePlatform(context.Context, *MsgCreatePlatform) (*MsgCreatePlatformResponse, error)
	UpdatePlatform(context.Context, *MsgUpdatePlatform) (*MsgUpdatePlatformResponse, error)
	DeletePlatform(context.Context, *MsgDeletePlatform) (*MsgDeletePlatformResponse, error)
	CreateClaim(context.Context, *MsgCreateClaim) (*MsgCreateClaimResponse, error)
	UpdateClaim(context.Context, *MsgUpdateClaim) (*MsgUpdateClaimResponse, error)
	DeleteClaim(context.Context, *MsgDeleteClaim) (*MsgDeleteClaimResponse, error)
	CreateBounty(context.Context, *MsgCreateBounty) (*MsgCreateBountyResponse, error)
	ClaimBounty(context.Context, *MsgClaimBounty) (*MsgClaimBountyResponse, error)
	CompleteBounty(context.Context, *MsgCompleteBounty) (*MsgCompleteBountyResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreatePlatform(context.Context, *MsgCreatePlatform) (*MsgCreatePlatformResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlatform not implemented")
}
func (UnimplementedMsgServer) UpdatePlatform(context.Context, *MsgUpdatePlatform) (*MsgUpdatePlatformResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlatform not implemented")
}
func (UnimplementedMsgServer) DeletePlatform(context.Context, *MsgDeletePlatform) (*MsgDeletePlatformResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlatform not implemented")
}
func (UnimplementedMsgServer) CreateClaim(context.Context, *MsgCreateClaim) (*MsgCreateClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClaim not implemented")
}
func (UnimplementedMsgServer) UpdateClaim(context.Context, *MsgUpdateClaim) (*MsgUpdateClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClaim not implemented")
}
func (UnimplementedMsgServer) DeleteClaim(context.Context, *MsgDeleteClaim) (*MsgDeleteClaimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClaim not implemented")
}
func (UnimplementedMsgServer) CreateBounty(context.Context, *MsgCreateBounty) (*MsgCreateBountyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBounty not implemented")
}
func (UnimplementedMsgServer) ClaimBounty(context.Context, *MsgClaimBounty) (*MsgClaimBountyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimBounty not implemented")
}
func (UnimplementedMsgServer) CompleteBounty(context.Context, *MsgCompleteBounty) (*MsgCompleteBountyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteBounty not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePlatform)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePlatform_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePlatform(ctx, req.(*MsgCreatePlatform))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdatePlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdatePlatform)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdatePlatform_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePlatform(ctx, req.(*MsgUpdatePlatform))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeletePlatform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeletePlatform)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeletePlatform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeletePlatform_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeletePlatform(ctx, req.(*MsgDeletePlatform))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateClaim)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateClaim_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateClaim(ctx, req.(*MsgCreateClaim))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateClaim)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateClaim_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateClaim(ctx, req.(*MsgUpdateClaim))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteClaim)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteClaim_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteClaim(ctx, req.(*MsgDeleteClaim))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateBounty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateBounty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateBounty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateBounty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateBounty(ctx, req.(*MsgCreateBounty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClaimBounty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClaimBounty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClaimBounty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ClaimBounty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClaimBounty(ctx, req.(*MsgClaimBounty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CompleteBounty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCompleteBounty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CompleteBounty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CompleteBounty_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CompleteBounty(ctx, req.(*MsgCompleteBounty))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bounty.platform.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreatePlatform",
			Handler:    _Msg_CreatePlatform_Handler,
		},
		{
			MethodName: "UpdatePlatform",
			Handler:    _Msg_UpdatePlatform_Handler,
		},
		{
			MethodName: "DeletePlatform",
			Handler:    _Msg_DeletePlatform_Handler,
		},
		{
			MethodName: "CreateClaim",
			Handler:    _Msg_CreateClaim_Handler,
		},
		{
			MethodName: "UpdateClaim",
			Handler:    _Msg_UpdateClaim_Handler,
		},
		{
			MethodName: "DeleteClaim",
			Handler:    _Msg_DeleteClaim_Handler,
		},
		{
			MethodName: "CreateBounty",
			Handler:    _Msg_CreateBounty_Handler,
		},
		{
			MethodName: "ClaimBounty",
			Handler:    _Msg_ClaimBounty_Handler,
		},
		{
			MethodName: "CompleteBounty",
			Handler:    _Msg_CompleteBounty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bounty/platform/tx.proto",
}
