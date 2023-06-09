// Code generated by goctl. DO NOT EDIT.
// Source: userRpc.proto

package userrpc

import (
	"context"

	"social-im/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AcceptAdminReq            = pb.AcceptAdminReq
	AcceptAdminResq           = pb.AcceptAdminResq
	AcceptMicReq              = pb.AcceptMicReq
	AcceptMicResq             = pb.AcceptMicResq
	ApplyMicReq               = pb.ApplyMicReq
	ApplyMicResq              = pb.ApplyMicResq
	ApplyingMicListReq        = pb.ApplyingMicListReq
	ApplyingMicListResp       = pb.ApplyingMicListResp
	ApproveMicReq             = pb.ApproveMicReq
	ApproveMicResp            = pb.ApproveMicResp
	AutoRegReq                = pb.AutoRegReq
	AutoRegResp               = pb.AutoRegResp
	CheckUserIdReq            = pb.CheckUserIdReq
	CheckUserIdResp           = pb.CheckUserIdResp
	CommonResp                = pb.CommonResp
	CommonRespNew             = pb.CommonRespNew
	DeclineAdminReq           = pb.DeclineAdminReq
	DeclineAdminResq          = pb.DeclineAdminResq
	DeclineMicReq             = pb.DeclineMicReq
	DeclineMicResq            = pb.DeclineMicResq
	GenerateTokenReq          = pb.GenerateTokenReq
	GenerateTokenResp         = pb.GenerateTokenResp
	GetAccessTokenReq         = pb.GetAccessTokenReq
	GetAccessTokenResp        = pb.GetAccessTokenResp
	GetRtcTokenReq            = pb.GetRtcTokenReq
	GetRtcTokenResp           = pb.GetRtcTokenResp
	GetRtmTokenReq            = pb.GetRtmTokenReq
	GetRtmTokenResp           = pb.GetRtmTokenResp
	GetUserAuthByAuthKeyReq   = pb.GetUserAuthByAuthKeyReq
	GetUserAuthByAuthKeyResp  = pb.GetUserAuthByAuthKeyResp
	GetUserAuthByUserIdReq    = pb.GetUserAuthByUserIdReq
	GetUserAuthyUserIdResp    = pb.GetUserAuthyUserIdResp
	GetUserBaseReq            = pb.GetUserBaseReq
	GetUserBaseResp           = pb.GetUserBaseResp
	GetUserInfoReq            = pb.GetUserInfoReq
	GetUserInfoResp           = pb.GetUserInfoResp
	IdentityReq               = pb.IdentityReq
	IdentityResp              = pb.IdentityResp
	InviteAdminReq            = pb.InviteAdminReq
	InviteAdminResq           = pb.InviteAdminResq
	InviteMicReq              = pb.InviteMicReq
	InviteMicResp             = pb.InviteMicResp
	LoginReq                  = pb.LoginReq
	LoginResp                 = pb.LoginResp
	MicStatus                 = pb.MicStatus
	OnMicersListReq           = pb.OnMicersListReq
	OnMicersListResp          = pb.OnMicersListResp
	RegisterReq               = pb.RegisterReq
	RegisterResp              = pb.RegisterResp
	RejectMicReq              = pb.RejectMicReq
	RejectMicResp             = pb.RejectMicResp
	RemoveAdminReq            = pb.RemoveAdminReq
	RemoveAdminResp           = pb.RemoveAdminResp
	RemoveMicReq              = pb.RemoveMicReq
	RemoveMicResp             = pb.RemoveMicResp
	RemoveUserRoomStatusReq   = pb.RemoveUserRoomStatusReq
	RemoveUserRoomStatusResq  = pb.RemoveUserRoomStatusResq
	ResumeMicReq              = pb.ResumeMicReq
	ResumeMicResq             = pb.ResumeMicResq
	RoomAdminListReq          = pb.RoomAdminListReq
	RoomAdminListResp         = pb.RoomAdminListResp
	RoomUser                  = pb.RoomUser
	RoomUsers                 = pb.RoomUsers
	SelectTagReq              = pb.SelectTagReq
	SelectTagResp             = pb.SelectTagResp
	SendRtmChannelReq         = pb.SendRtmChannelReq
	SendRtmChannelResp        = pb.SendRtmChannelResp
	SendRtmReq                = pb.SendRtmReq
	SendRtmResp               = pb.SendRtmResp
	SmsCodeReq                = pb.SmsCodeReq
	SmsCodeResp               = pb.SmsCodeResp
	SmsLoginReq               = pb.SmsLoginReq
	SmsLoginResp              = pb.SmsLoginResp
	SmsRegReq                 = pb.SmsRegReq
	SmsRegResp                = pb.SmsRegResp
	SmsReq                    = pb.SmsReq
	SmsResp                   = pb.SmsResp
	SmsVerifyReq              = pb.SmsVerifyReq
	SmsVerifyResp             = pb.SmsVerifyResp
	StopMicReq                = pb.StopMicReq
	StopMicResp               = pb.StopMicResp
	Token                     = pb.Token
	UpdateAboutMeReq          = pb.UpdateAboutMeReq
	UpdateAboutMeResp         = pb.UpdateAboutMeResp
	UpdateBackgroundImageReq  = pb.UpdateBackgroundImageReq
	UpdateBackgroundImageResp = pb.UpdateBackgroundImageResp
	UpdateBaseReq             = pb.UpdateBaseReq
	UpdateBaseResp            = pb.UpdateBaseResp
	UpdateInfoReq             = pb.UpdateInfoReq
	UpdateInfoResp            = pb.UpdateInfoResp
	UpdatePwdReq              = pb.UpdatePwdReq
	UpdatePwdResp             = pb.UpdatePwdResp
	User                      = pb.User
	UserAuth                  = pb.UserAuth
	UserMicTime               = pb.UserMicTime
	Users                     = pb.Users

	UserRpc interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		Sms(ctx context.Context, in *SmsReq, opts ...grpc.CallOption) (*SmsResp, error)
		SmsCode(ctx context.Context, in *SmsCodeReq, opts ...grpc.CallOption) (*SmsCodeResp, error)
		SmsReg(ctx context.Context, in *SmsRegReq, opts ...grpc.CallOption) (*SmsRegResp, error)
		Identity(ctx context.Context, in *IdentityReq, opts ...grpc.CallOption) (*IdentityResp, error)
		SmsLogin(ctx context.Context, in *SmsLoginReq, opts ...grpc.CallOption) (*SmsLoginResp, error)
		AutoReg(ctx context.Context, in *AutoRegReq, opts ...grpc.CallOption) (*AutoRegResp, error)
		CheckUserId(ctx context.Context, in *CheckUserIdReq, opts ...grpc.CallOption) (*CheckUserIdResp, error)
		SmsVerify(ctx context.Context, in *SmsVerifyReq, opts ...grpc.CallOption) (*SmsVerifyResp, error)
		UpdatePwd(ctx context.Context, in *UpdatePwdReq, opts ...grpc.CallOption) (*UpdatePwdResp, error)
		UpdateInfo(ctx context.Context, in *UpdateInfoReq, opts ...grpc.CallOption) (*UpdateInfoResp, error)
		GetRtmToken(ctx context.Context, in *GetRtmTokenReq, opts ...grpc.CallOption) (*GetRtmTokenResp, error)
		GetRtcToken(ctx context.Context, in *GetRtcTokenReq, opts ...grpc.CallOption) (*GetRtcTokenResp, error)
		GetAccessToken(ctx context.Context, in *GetAccessTokenReq, opts ...grpc.CallOption) (*GetAccessTokenResp, error)
		SendRtm(ctx context.Context, in *SendRtmReq, opts ...grpc.CallOption) (*SendRtmResp, error)
		SendRtmChannel(ctx context.Context, in *SendRtmChannelReq, opts ...grpc.CallOption) (*SendRtmChannelResp, error)
		ApplyMic(ctx context.Context, in *ApplyMicReq, opts ...grpc.CallOption) (*ApplyMicResq, error)
		ApproveMic(ctx context.Context, in *ApproveMicReq, opts ...grpc.CallOption) (*ApproveMicResp, error)
		RejectMic(ctx context.Context, in *RejectMicReq, opts ...grpc.CallOption) (*RejectMicResp, error)
		ApplyingMicList(ctx context.Context, in *ApplyingMicListReq, opts ...grpc.CallOption) (*ApplyingMicListResp, error)
		OnMicersList(ctx context.Context, in *OnMicersListReq, opts ...grpc.CallOption) (*OnMicersListResp, error)
		InviteMic(ctx context.Context, in *InviteMicReq, opts ...grpc.CallOption) (*InviteMicResp, error)
		AcceptMic(ctx context.Context, in *AcceptMicReq, opts ...grpc.CallOption) (*AcceptMicResq, error)
		DeclineMic(ctx context.Context, in *DeclineMicReq, opts ...grpc.CallOption) (*DeclineMicResq, error)
		RemoveMic(ctx context.Context, in *RemoveMicReq, opts ...grpc.CallOption) (*RemoveMicResp, error)
		StopMic(ctx context.Context, in *StopMicReq, opts ...grpc.CallOption) (*StopMicResp, error)
		ResumeMic(ctx context.Context, in *ResumeMicReq, opts ...grpc.CallOption) (*ResumeMicResq, error)
		InviteAdmin(ctx context.Context, in *InviteAdminReq, opts ...grpc.CallOption) (*InviteAdminResq, error)
		AcceptAdmin(ctx context.Context, in *AcceptAdminReq, opts ...grpc.CallOption) (*AcceptAdminResq, error)
		DeclineAdmin(ctx context.Context, in *DeclineAdminReq, opts ...grpc.CallOption) (*DeclineAdminResq, error)
		RemoveAdmin(ctx context.Context, in *RemoveAdminReq, opts ...grpc.CallOption) (*RemoveAdminResp, error)
		RoomAdminList(ctx context.Context, in *RoomAdminListReq, opts ...grpc.CallOption) (*RoomAdminListResp, error)
		RemoveUserRoomStatus(ctx context.Context, in *RemoveUserRoomStatusReq, opts ...grpc.CallOption) (*RemoveUserRoomStatusResq, error)
		UpdateBaseInfo(ctx context.Context, in *UpdateBaseReq, opts ...grpc.CallOption) (*UpdateBaseResp, error)
		UpdateAboutMe(ctx context.Context, in *UpdateAboutMeReq, opts ...grpc.CallOption) (*UpdateAboutMeResp, error)
		UpdateBackgroundImage(ctx context.Context, in *UpdateBackgroundImageReq, opts ...grpc.CallOption) (*UpdateBackgroundImageResp, error)
		GetUserBaseInfo(ctx context.Context, in *GetUserBaseReq, opts ...grpc.CallOption) (*GetUserBaseResp, error)
		SelectTag(ctx context.Context, in *SelectTagReq, opts ...grpc.CallOption) (*SelectTagResp, error)
		UserFollow(ctx context.Context, in *RoomUsers, opts ...grpc.CallOption) (*CommonRespNew, error)
		UserUnFollow(ctx context.Context, in *RoomUsers, opts ...grpc.CallOption) (*CommonRespNew, error)
		UserFollowers(ctx context.Context, in *RoomUser, opts ...grpc.CallOption) (*Users, error)
		UserFollows(ctx context.Context, in *RoomUser, opts ...grpc.CallOption) (*Users, error)
	}

	defaultUserRpc struct {
		cli zrpc.Client
	}
)

func NewUserRpc(cli zrpc.Client) UserRpc {
	return &defaultUserRpc{
		cli: cli,
	}
}

func (m *defaultUserRpc) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserRpc) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserRpc) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUserRpc) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultUserRpc) Sms(ctx context.Context, in *SmsReq, opts ...grpc.CallOption) (*SmsResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.Sms(ctx, in, opts...)
}

func (m *defaultUserRpc) SmsCode(ctx context.Context, in *SmsCodeReq, opts ...grpc.CallOption) (*SmsCodeResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.SmsCode(ctx, in, opts...)
}

func (m *defaultUserRpc) SmsReg(ctx context.Context, in *SmsRegReq, opts ...grpc.CallOption) (*SmsRegResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.SmsReg(ctx, in, opts...)
}

func (m *defaultUserRpc) Identity(ctx context.Context, in *IdentityReq, opts ...grpc.CallOption) (*IdentityResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.Identity(ctx, in, opts...)
}

func (m *defaultUserRpc) SmsLogin(ctx context.Context, in *SmsLoginReq, opts ...grpc.CallOption) (*SmsLoginResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.SmsLogin(ctx, in, opts...)
}

func (m *defaultUserRpc) AutoReg(ctx context.Context, in *AutoRegReq, opts ...grpc.CallOption) (*AutoRegResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.AutoReg(ctx, in, opts...)
}

func (m *defaultUserRpc) CheckUserId(ctx context.Context, in *CheckUserIdReq, opts ...grpc.CallOption) (*CheckUserIdResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.CheckUserId(ctx, in, opts...)
}

func (m *defaultUserRpc) SmsVerify(ctx context.Context, in *SmsVerifyReq, opts ...grpc.CallOption) (*SmsVerifyResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.SmsVerify(ctx, in, opts...)
}

func (m *defaultUserRpc) UpdatePwd(ctx context.Context, in *UpdatePwdReq, opts ...grpc.CallOption) (*UpdatePwdResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.UpdatePwd(ctx, in, opts...)
}

func (m *defaultUserRpc) UpdateInfo(ctx context.Context, in *UpdateInfoReq, opts ...grpc.CallOption) (*UpdateInfoResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.UpdateInfo(ctx, in, opts...)
}

func (m *defaultUserRpc) GetRtmToken(ctx context.Context, in *GetRtmTokenReq, opts ...grpc.CallOption) (*GetRtmTokenResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.GetRtmToken(ctx, in, opts...)
}

func (m *defaultUserRpc) GetRtcToken(ctx context.Context, in *GetRtcTokenReq, opts ...grpc.CallOption) (*GetRtcTokenResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.GetRtcToken(ctx, in, opts...)
}

func (m *defaultUserRpc) GetAccessToken(ctx context.Context, in *GetAccessTokenReq, opts ...grpc.CallOption) (*GetAccessTokenResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.GetAccessToken(ctx, in, opts...)
}

func (m *defaultUserRpc) SendRtm(ctx context.Context, in *SendRtmReq, opts ...grpc.CallOption) (*SendRtmResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.SendRtm(ctx, in, opts...)
}

func (m *defaultUserRpc) SendRtmChannel(ctx context.Context, in *SendRtmChannelReq, opts ...grpc.CallOption) (*SendRtmChannelResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.SendRtmChannel(ctx, in, opts...)
}

func (m *defaultUserRpc) ApplyMic(ctx context.Context, in *ApplyMicReq, opts ...grpc.CallOption) (*ApplyMicResq, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.ApplyMic(ctx, in, opts...)
}

func (m *defaultUserRpc) ApproveMic(ctx context.Context, in *ApproveMicReq, opts ...grpc.CallOption) (*ApproveMicResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.ApproveMic(ctx, in, opts...)
}

func (m *defaultUserRpc) RejectMic(ctx context.Context, in *RejectMicReq, opts ...grpc.CallOption) (*RejectMicResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.RejectMic(ctx, in, opts...)
}

func (m *defaultUserRpc) ApplyingMicList(ctx context.Context, in *ApplyingMicListReq, opts ...grpc.CallOption) (*ApplyingMicListResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.ApplyingMicList(ctx, in, opts...)
}

func (m *defaultUserRpc) OnMicersList(ctx context.Context, in *OnMicersListReq, opts ...grpc.CallOption) (*OnMicersListResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.OnMicersList(ctx, in, opts...)
}

func (m *defaultUserRpc) InviteMic(ctx context.Context, in *InviteMicReq, opts ...grpc.CallOption) (*InviteMicResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.InviteMic(ctx, in, opts...)
}

func (m *defaultUserRpc) AcceptMic(ctx context.Context, in *AcceptMicReq, opts ...grpc.CallOption) (*AcceptMicResq, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.AcceptMic(ctx, in, opts...)
}

func (m *defaultUserRpc) DeclineMic(ctx context.Context, in *DeclineMicReq, opts ...grpc.CallOption) (*DeclineMicResq, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.DeclineMic(ctx, in, opts...)
}

func (m *defaultUserRpc) RemoveMic(ctx context.Context, in *RemoveMicReq, opts ...grpc.CallOption) (*RemoveMicResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.RemoveMic(ctx, in, opts...)
}

func (m *defaultUserRpc) StopMic(ctx context.Context, in *StopMicReq, opts ...grpc.CallOption) (*StopMicResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.StopMic(ctx, in, opts...)
}

func (m *defaultUserRpc) ResumeMic(ctx context.Context, in *ResumeMicReq, opts ...grpc.CallOption) (*ResumeMicResq, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.ResumeMic(ctx, in, opts...)
}

func (m *defaultUserRpc) InviteAdmin(ctx context.Context, in *InviteAdminReq, opts ...grpc.CallOption) (*InviteAdminResq, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.InviteAdmin(ctx, in, opts...)
}

func (m *defaultUserRpc) AcceptAdmin(ctx context.Context, in *AcceptAdminReq, opts ...grpc.CallOption) (*AcceptAdminResq, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.AcceptAdmin(ctx, in, opts...)
}

func (m *defaultUserRpc) DeclineAdmin(ctx context.Context, in *DeclineAdminReq, opts ...grpc.CallOption) (*DeclineAdminResq, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.DeclineAdmin(ctx, in, opts...)
}

func (m *defaultUserRpc) RemoveAdmin(ctx context.Context, in *RemoveAdminReq, opts ...grpc.CallOption) (*RemoveAdminResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.RemoveAdmin(ctx, in, opts...)
}

func (m *defaultUserRpc) RoomAdminList(ctx context.Context, in *RoomAdminListReq, opts ...grpc.CallOption) (*RoomAdminListResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.RoomAdminList(ctx, in, opts...)
}

func (m *defaultUserRpc) RemoveUserRoomStatus(ctx context.Context, in *RemoveUserRoomStatusReq, opts ...grpc.CallOption) (*RemoveUserRoomStatusResq, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.RemoveUserRoomStatus(ctx, in, opts...)
}

func (m *defaultUserRpc) UpdateBaseInfo(ctx context.Context, in *UpdateBaseReq, opts ...grpc.CallOption) (*UpdateBaseResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.UpdateBaseInfo(ctx, in, opts...)
}

func (m *defaultUserRpc) UpdateAboutMe(ctx context.Context, in *UpdateAboutMeReq, opts ...grpc.CallOption) (*UpdateAboutMeResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.UpdateAboutMe(ctx, in, opts...)
}

func (m *defaultUserRpc) UpdateBackgroundImage(ctx context.Context, in *UpdateBackgroundImageReq, opts ...grpc.CallOption) (*UpdateBackgroundImageResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.UpdateBackgroundImage(ctx, in, opts...)
}

func (m *defaultUserRpc) GetUserBaseInfo(ctx context.Context, in *GetUserBaseReq, opts ...grpc.CallOption) (*GetUserBaseResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.GetUserBaseInfo(ctx, in, opts...)
}

func (m *defaultUserRpc) SelectTag(ctx context.Context, in *SelectTagReq, opts ...grpc.CallOption) (*SelectTagResp, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.SelectTag(ctx, in, opts...)
}

func (m *defaultUserRpc) UserFollow(ctx context.Context, in *RoomUsers, opts ...grpc.CallOption) (*CommonRespNew, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.UserFollow(ctx, in, opts...)
}

func (m *defaultUserRpc) UserUnFollow(ctx context.Context, in *RoomUsers, opts ...grpc.CallOption) (*CommonRespNew, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.UserUnFollow(ctx, in, opts...)
}

func (m *defaultUserRpc) UserFollowers(ctx context.Context, in *RoomUser, opts ...grpc.CallOption) (*Users, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.UserFollowers(ctx, in, opts...)
}

func (m *defaultUserRpc) UserFollows(ctx context.Context, in *RoomUser, opts ...grpc.CallOption) (*Users, error) {
	client := pb.NewUserRpcClient(m.cli.Conn())
	return client.UserFollows(ctx, in, opts...)
}
