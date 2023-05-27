// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: giftRpc.proto

package pb

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

// GiftRpcClient is the client API for GiftRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GiftRpcClient interface {
	Send(ctx context.Context, in *GiftSendReq, opts ...grpc.CallOption) (*GiftSendResp, error)
	UpdateBag(ctx context.Context, in *GiftUpdateBagReq, opts ...grpc.CallOption) (*GiftUpdateBagResp, error)
	UpdateAttr(ctx context.Context, in *GiftUpdateAttrReq, opts ...grpc.CallOption) (*GiftUpdateAttrResp, error)
}

type giftRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewGiftRpcClient(cc grpc.ClientConnInterface) GiftRpcClient {
	return &giftRpcClient{cc}
}

func (c *giftRpcClient) Send(ctx context.Context, in *GiftSendReq, opts ...grpc.CallOption) (*GiftSendResp, error) {
	out := new(GiftSendResp)
	err := c.cc.Invoke(ctx, "/pb.giftRpc/send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftRpcClient) UpdateBag(ctx context.Context, in *GiftUpdateBagReq, opts ...grpc.CallOption) (*GiftUpdateBagResp, error) {
	out := new(GiftUpdateBagResp)
	err := c.cc.Invoke(ctx, "/pb.giftRpc/updateBag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giftRpcClient) UpdateAttr(ctx context.Context, in *GiftUpdateAttrReq, opts ...grpc.CallOption) (*GiftUpdateAttrResp, error) {
	out := new(GiftUpdateAttrResp)
	err := c.cc.Invoke(ctx, "/pb.giftRpc/updateAttr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GiftRpcServer is the server API for GiftRpc service.
// All implementations must embed UnimplementedGiftRpcServer
// for forward compatibility
type GiftRpcServer interface {
	Send(context.Context, *GiftSendReq) (*GiftSendResp, error)
	UpdateBag(context.Context, *GiftUpdateBagReq) (*GiftUpdateBagResp, error)
	UpdateAttr(context.Context, *GiftUpdateAttrReq) (*GiftUpdateAttrResp, error)
	mustEmbedUnimplementedGiftRpcServer()
}

// UnimplementedGiftRpcServer must be embedded to have forward compatible implementations.
type UnimplementedGiftRpcServer struct {
}

func (UnimplementedGiftRpcServer) Send(context.Context, *GiftSendReq) (*GiftSendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedGiftRpcServer) UpdateBag(context.Context, *GiftUpdateBagReq) (*GiftUpdateBagResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBag not implemented")
}
func (UnimplementedGiftRpcServer) UpdateAttr(context.Context, *GiftUpdateAttrReq) (*GiftUpdateAttrResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAttr not implemented")
}
func (UnimplementedGiftRpcServer) mustEmbedUnimplementedGiftRpcServer() {}

// UnsafeGiftRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GiftRpcServer will
// result in compilation errors.
type UnsafeGiftRpcServer interface {
	mustEmbedUnimplementedGiftRpcServer()
}

func RegisterGiftRpcServer(s grpc.ServiceRegistrar, srv GiftRpcServer) {
	s.RegisterService(&GiftRpc_ServiceDesc, srv)
}

func _GiftRpc_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GiftSendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftRpcServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.giftRpc/send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftRpcServer).Send(ctx, req.(*GiftSendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GiftRpc_UpdateBag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GiftUpdateBagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftRpcServer).UpdateBag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.giftRpc/updateBag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftRpcServer).UpdateBag(ctx, req.(*GiftUpdateBagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GiftRpc_UpdateAttr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GiftUpdateAttrReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiftRpcServer).UpdateAttr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.giftRpc/updateAttr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiftRpcServer).UpdateAttr(ctx, req.(*GiftUpdateAttrReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GiftRpc_ServiceDesc is the grpc.ServiceDesc for GiftRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GiftRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.giftRpc",
	HandlerType: (*GiftRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "send",
			Handler:    _GiftRpc_Send_Handler,
		},
		{
			MethodName: "updateBag",
			Handler:    _GiftRpc_UpdateBag_Handler,
		},
		{
			MethodName: "updateAttr",
			Handler:    _GiftRpc_UpdateAttr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "giftRpc.proto",
}
