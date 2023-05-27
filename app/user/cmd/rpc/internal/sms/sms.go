package sms

import (
	"context"
	"fmt"
	"math/rand"
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/common/rediskey"
	errTypes "social-im/common/types"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
)

type Sms struct {
	svcCtx *svc.ServiceContext
}

var sms *Sms

func NewSms(svcCtx *svc.ServiceContext) *Sms {
	if sms != nil {
		return sms
	}
	sms = &Sms{
		svcCtx: svcCtx,
		// Redis:  xcache.GetClient(svcCtx.Config.Redis.RedisConf, global.DB(svcCtx.Config.Redis.DB)),
	}
	return sms
}

func (l Sms) VerifyCode(phone string, code string, codeType string) error {
	redisConn := l.svcCtx.Redis
	// defer redisConn.Close()

	verifyKey := rediskey.CacheSocialImVerifyCodePrefix + phone + codeType
	codeInCache, err := redisConn.Get(verifyKey)

	if err != redis.Nil && err != nil {
		fmt.Printf("error is %v \n", err)
		return errTypes.ErrSysBusy
	}

	fmt.Printf("cache in redis is %v, in code is %v \n", codeInCache, code)
	if codeInCache != code {
		return errTypes.ErrCodeNotMatch
	}
	return nil

}

func (l Sms) SendSms(ctx context.Context, phone string, codeType string) (string, error) {
	/**
	1、查询手机号对应的短信频率限制
	2、发短信
	3、写入redis验证码内容
	4、更新手机号的频率限制
	*/

	redisConn := l.svcCtx.Redis

	//1 查询redis频率限制
	// limiKey := "cache111:socialIm:sms:limit:" + phone
	limiKey := rediskey.CacheSocialImVerifyLimitPrefix + phone
	// verifyKey := "cache111:socialIm:sms:verify:" + phone + codeType
	verifyKey := rediskey.CacheSocialImVerifyCodePrefix + phone + codeType
	fmt.Printf("verifyKey is %s \n", verifyKey)
	limit, err := redisConn.GetCtx(ctx, limiKey)
	if err != redis.Nil && limit != "" {
		return "", err
	}

	if limit != "" {
		return "", err
	}

	//2 发短信
	client, err := dysmsapi.NewClientWithAccessKey("im-develop", "LTAI5tSjv5e5JSm6sXUGv4ep", "IhkHh1YKbyNnlp5oo2RfquoaL3Pcs3")
	if err != nil {
		return "", err
	}

	verifyCode := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = "霖灵科技"
	request.TemplateCode = "SMS_241180034"
	request.TemplateParam = fmt.Sprintf("{\"code\":\"%s\"}", verifyCode)

	fmt.Println("response:", client)
	response, err := client.SendSms(request)
	fmt.Println("response:", response)
	if err != nil {
		fmt.Printf("sms response err is %v \n", err)
		// return "", err
	}

	//3 写redis
	err = redisConn.SetexCtx(ctx, verifyKey, verifyCode, 60*10)
	if err != nil {
		fmt.Printf("set redis err is %v \n", err)
		return "", err
	}
	//4 todo 设置redis发送频率
	logx.Info("this is test info")
	return verifyCode, nil
}
