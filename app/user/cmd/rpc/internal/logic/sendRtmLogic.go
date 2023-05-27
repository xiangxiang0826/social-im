package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"social-im/app/user/cmd/rpc/internal/svc"
	cachedtoken "social-im/app/user/cmd/rpc/internal/token"
	"social-im/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendRtmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	cachedtoken *cachedtoken.CacheToken
}

func NewSendRtmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendRtmLogic {
	return &SendRtmLogic{
		ctx:         ctx,
		svcCtx:      svcCtx,
		Logger:      logx.WithContext(ctx),
		cachedtoken: cachedtoken.NewCacheToken(svcCtx),
	}
}

func (l *SendRtmLogic) SendRtm(in *pb.SendRtmReq) (*pb.SendRtmResp, error) {
	// todo: add your logic here and delete this line
	tokenValue, err := l.cachedtoken.GetRtmTokenWithCache(l.svcCtx.Config.AgoraConf.AppId, l.svcCtx.Config.AgoraConf.AppCertificate, in.From)
	if err != nil {
		return &pb.SendRtmResp{
			Iret: 1,
			Smsg: err.Error(),
		}, err
	}

	appID := l.svcCtx.Config.AgoraConf.AppId
	host := l.svcCtx.Config.AgoraConf.Server
	urlstr := fmt.Sprintf("https://%s/dev/v2/project/%s/rtm/users/%s/peer_messages?wait_for_ack=true", host, appID, in.From)
	fmt.Printf("request url is %s \n", urlstr)
	authValue := fmt.Sprintf("agora token=%s", tokenValue)

	messageBody := map[string]interface{}{
		"message_type": in.MessageType,
		"message_body": in.MessageBody,
	}
	messageBodyByte, _ := json.Marshal(messageBody)
	messageBodyStr := string(messageBodyByte)
	fmt.Sprintf("messageBody is %s \n", messageBodyStr)

	requestBody := map[string]interface{}{
		"destination":                 in.To,
		"enable_offline_messaging":    false,
		"enable_historical_messaging": false,
		"payload":                     messageBodyStr,
	}

	// requestBody := map[string]interface{}{
	// 	"destination":                 in.To,
	// 	"enable_offline_messaging":    false,
	// 	"enable_historical_messaging": false,
	// 	"payload":                     "hello from xiangxiang",
	// }

	headers := map[string]string{
		"Content-Type":  "application/json; charset=utf-8",
		"Authorization": authValue,
	}
	fmt.Printf("request header is %v \n", headers)

	requestBodyStr, _ := json.Marshal(requestBody)
	fmt.Printf("request body is %v \n", string(requestBodyStr))

	client := &http.Client{}
	req, err := http.NewRequest("POST", urlstr, bytes.NewBuffer(requestBodyStr))
	if err != nil {
		return &pb.SendRtmResp{
			Iret: 1,
			Smsg: err.Error(),
		}, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return &pb.SendRtmResp{
			Iret: 1,
			Smsg: err.Error(),
		}, err
		// return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	return &pb.SendRtmResp{}, nil
}
