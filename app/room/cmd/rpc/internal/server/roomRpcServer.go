// Code generated by goctl. DO NOT EDIT.
// Source: roomRpc.proto

package server

import (
	"context"

	"social-im/app/room/cmd/rpc/internal/logic"
	"social-im/app/room/cmd/rpc/internal/svc"
	"social-im/app/room/cmd/rpc/pb"
)

type RoomRpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedRoomRpcServer
}

func NewRoomRpcServer(svcCtx *svc.ServiceContext) *RoomRpcServer {
	return &RoomRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *RoomRpcServer) Create(ctx context.Context, in *pb.PartyCreateReq) (*pb.PartyCreateResp, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *RoomRpcServer) RoomLimitGet(ctx context.Context, in *pb.RoomLimitReq) (*pb.RoomLimitResp, error) {
	l := logic.NewRoomLimitGetLogic(ctx, s.svcCtx)
	return l.RoomLimitGet(in)
}

func (s *RoomRpcServer) Join(ctx context.Context, in *pb.PartyJoinReq) (*pb.PartyJoinResp, error) {
	l := logic.NewJoinLogic(ctx, s.svcCtx)
	return l.Join(in)
}

func (s *RoomRpcServer) Leave(ctx context.Context, in *pb.PartyLeaveReq) (*pb.PartyLeaveResp, error) {
	l := logic.NewLeaveLogic(ctx, s.svcCtx)
	return l.Leave(in)
}

func (s *RoomRpcServer) Remove(ctx context.Context, in *pb.PartyRemoveReq) (*pb.PartyRemoveResp, error) {
	l := logic.NewRemoveLogic(ctx, s.svcCtx)
	return l.Remove(in)
}

func (s *RoomRpcServer) PartyList(ctx context.Context, in *pb.PartyListReq) (*pb.PartyListResp, error) {
	l := logic.NewPartyListLogic(ctx, s.svcCtx)
	return l.PartyList(in)
}

func (s *RoomRpcServer) UserOnlineNumReport(ctx context.Context, in *pb.UserNumReportReq) (*pb.UserNumReportResp, error) {
	l := logic.NewUserOnlineNumReportLogic(ctx, s.svcCtx)
	return l.UserOnlineNumReport(in)
}

func (s *RoomRpcServer) RoominfoById(ctx context.Context, in *pb.RoominfoReq) (*pb.RoominfoResq, error) {
	l := logic.NewRoominfoByIdLogic(ctx, s.svcCtx)
	return l.RoominfoById(in)
}

func (s *RoomRpcServer) UpdatePartyName(ctx context.Context, in *pb.PartyNameUpdateReq) (*pb.PartyNameUpdateResp, error) {
	l := logic.NewUpdatePartyNameLogic(ctx, s.svcCtx)
	return l.UpdatePartyName(in)
}

func (s *RoomRpcServer) UpdatePartyBackgroundImg(ctx context.Context, in *pb.PartyBackGroundImgUpdateReq) (*pb.PartyBackGroundImgUpdateResp, error) {
	l := logic.NewUpdatePartyBackgroundImgLogic(ctx, s.svcCtx)
	return l.UpdatePartyBackgroundImg(in)
}

func (s *RoomRpcServer) ProhibitionUserAdd(ctx context.Context, in *pb.ProhibitionCreateReq) (*pb.ProhibitionCreateResp, error) {
	l := logic.NewProhibitionUserAddLogic(ctx, s.svcCtx)
	return l.ProhibitionUserAdd(in)
}

func (s *RoomRpcServer) ProhibitionUserRemove(ctx context.Context, in *pb.ProhibitionRemoveReq) (*pb.ProhibitionRemoveResp, error) {
	l := logic.NewProhibitionUserRemoveLogic(ctx, s.svcCtx)
	return l.ProhibitionUserRemove(in)
}

func (s *RoomRpcServer) ProhibitionList(ctx context.Context, in *pb.ProhibitionListReq) (*pb.ProhibitionListResp, error) {
	l := logic.NewProhibitionListLogic(ctx, s.svcCtx)
	return l.ProhibitionList(in)
}

func (s *RoomRpcServer) Terminate(ctx context.Context, in *pb.PartyTerminateReq) (*pb.PartyTerminateResp, error) {
	l := logic.NewTerminateLogic(ctx, s.svcCtx)
	return l.Terminate(in)
}

func (s *RoomRpcServer) ProhibitionUserInfo(ctx context.Context, in *pb.ProhibitionGetReq) (*pb.ProhibitionGetResp, error) {
	l := logic.NewProhibitionUserInfoLogic(ctx, s.svcCtx)
	return l.ProhibitionUserInfo(in)
}
