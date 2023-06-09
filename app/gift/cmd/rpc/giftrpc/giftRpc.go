// Code generated by goctl. DO NOT EDIT.
// Source: giftRpc.proto

package giftrpc

import (
	"context"

	"social-im/app/gift/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetItemReq         = pb.GetItemReq
	GetItemResp        = pb.GetItemResp
	GiftSendReq        = pb.GiftSendReq
	GiftSendResp       = pb.GiftSendResp
	GiftUpdateAttrReq  = pb.GiftUpdateAttrReq
	GiftUpdateAttrResp = pb.GiftUpdateAttrResp
	GiftUpdateBagReq   = pb.GiftUpdateBagReq
	GiftUpdateBagResp  = pb.GiftUpdateBagResp

	GiftRpc interface {
		Send(ctx context.Context, in *GiftSendReq, opts ...grpc.CallOption) (*GiftSendResp, error)
		UpdateBag(ctx context.Context, in *GiftUpdateBagReq, opts ...grpc.CallOption) (*GiftUpdateBagResp, error)
		UpdateAttr(ctx context.Context, in *GiftUpdateAttrReq, opts ...grpc.CallOption) (*GiftUpdateAttrResp, error)
	}

	defaultGiftRpc struct {
		cli zrpc.Client
	}
)

func NewGiftRpc(cli zrpc.Client) GiftRpc {
	return &defaultGiftRpc{
		cli: cli,
	}
}

func (m *defaultGiftRpc) Send(ctx context.Context, in *GiftSendReq, opts ...grpc.CallOption) (*GiftSendResp, error) {
	client := pb.NewGiftRpcClient(m.cli.Conn())
	return client.Send(ctx, in, opts...)
}

func (m *defaultGiftRpc) UpdateBag(ctx context.Context, in *GiftUpdateBagReq, opts ...grpc.CallOption) (*GiftUpdateBagResp, error) {
	client := pb.NewGiftRpcClient(m.cli.Conn())
	return client.UpdateBag(ctx, in, opts...)
}

func (m *defaultGiftRpc) UpdateAttr(ctx context.Context, in *GiftUpdateAttrReq, opts ...grpc.CallOption) (*GiftUpdateAttrResp, error) {
	client := pb.NewGiftRpcClient(m.cli.Conn())
	return client.UpdateAttr(ctx, in, opts...)
}
