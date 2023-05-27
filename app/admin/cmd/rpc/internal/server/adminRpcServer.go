// Code generated by goctl. DO NOT EDIT.
// Source: adminRpc.proto

package server

import (
	"context"

	"social-im/app/admin/cmd/rpc/internal/logic"
	"social-im/app/admin/cmd/rpc/internal/svc"
	"social-im/app/admin/cmd/rpc/pb"
)

type AdminRpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedAdminRpcServer
}

func NewAdminRpcServer(svcCtx *svc.ServiceContext) *AdminRpcServer {
	return &AdminRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *AdminRpcServer) BackgroundImageList(ctx context.Context, in *pb.BackgroundImgReq) (*pb.BackgroundImgResp, error) {
	l := logic.NewBackgroundImageListLogic(ctx, s.svcCtx)
	return l.BackgroundImageList(in)
}

func (s *AdminRpcServer) DictionaryList(ctx context.Context, in *pb.DictionaryGetReq) (*pb.DictionaryGetResp, error) {
	l := logic.NewDictionaryListLogic(ctx, s.svcCtx)
	return l.DictionaryList(in)
}

func (s *AdminRpcServer) ProjectConfigDetail(ctx context.Context, in *pb.ProjectConfigDetailReq) (*pb.ProjectConfigDetailResp, error) {
	l := logic.NewProjectConfigDetailLogic(ctx, s.svcCtx)
	return l.ProjectConfigDetail(in)
}

func (s *AdminRpcServer) GiftList(ctx context.Context, in *pb.GiftListReq) (*pb.GiftListResp, error) {
	l := logic.NewGiftListLogic(ctx, s.svcCtx)
	return l.GiftList(in)
}

func (s *AdminRpcServer) GiftItem(ctx context.Context, in *pb.GiftItemReq) (*pb.GiftItemResp, error) {
	l := logic.NewGiftItemLogic(ctx, s.svcCtx)
	return l.GiftItem(in)
}

func (s *AdminRpcServer) AreaList(ctx context.Context, in *pb.AreaListReq) (*pb.AreaListResp, error) {
	l := logic.NewAreaListLogic(ctx, s.svcCtx)
	return l.AreaList(in)
}
