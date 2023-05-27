package cachedtoken

import (
	"fmt"
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/common/rediskey"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"

	// rtmtokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/RtmTokenBuilder"
	"social-im/common/agora"
)

type CacheToken struct {
	rtmToken    string
	rtcToken    string
	accessToken string
	svcCtx      *svc.ServiceContext
}

var (
	rtmExpireMinute  = 86400
	rtcExpireMinute  = 86400
	accessExpireHour = 120
)

func NewCacheToken(svcCtx *svc.ServiceContext) *CacheToken {
	return &CacheToken{
		svcCtx: svcCtx,
	}
}

// 读取redis中的值
// 存在则返回，
// 不存在则读接口，然后写redis缓存
func (l CacheToken) GetRtmTokenWithCache(appId, appCertificate, uidStr string) (string, error) {
	fmt.Println("GetRtmWithCache")
	//1 读取redis中的值
	redisConn := l.svcCtx.Redis
	rtmKey := rediskey.CacheSocialImRtmTokenPrefix + uidStr
	rtmToken, err := redisConn.Get(rtmKey)
	expireTimestamp := uint32(time.Now().Add(time.Minute * time.Duration(rtmExpireMinute)).Unix())
	if err != redis.Nil && err != nil {
		return "", err
	}

	//2 存在则返回，
	if err == nil && rtmToken != "" {
		return rtmToken, nil
	}

	//3 不存在则读接口，然后写redis缓存
	rtmToken, err = agora.GetRtmToken(appId, appCertificate, uidStr, expireTimestamp)

	if err != nil {
		return "", err
	}
	err = redisConn.Setex(rtmKey, rtmToken, rtmExpireMinute)
	if err != nil {
		return "", err
	}
	return rtmToken, err
}

// 获取redis的token
// 存在则返回
// 不存在则读声网接口，写redis缓存
func (l CacheToken) GetRtcTokenWithCache(appId, appCertificate, channelName, uidStr string) (string, error) {
	fmt.Printf("GetRtcTokenWithCache channel: %v; uid: %v \n", channelName, uidStr)

	//1 获取redis的token
	redisConn := l.svcCtx.Redis
	rtcKey := rediskey.CacheSocialImRtcTokenPrefix + uidStr + ":" + channelName
	fmt.Printf("GetRtcTokenWithCache rediskey: %v \n", rtcKey)
	rtcToken, err := redisConn.Get(rtcKey)
	expireTimestamp := uint32(time.Now().Add(time.Minute * time.Duration(rtcExpireMinute)).Unix())
	if err != redis.Nil && err != nil {
		fmt.Printf("redis err is %v \n", err)
		return "", err
	}

	//2 存在则返回
	if err == nil && rtcToken != "" {
		fmt.Printf("GetJwtTokenWithCache in redis: %v", rtcToken)
		return rtcToken, nil
	}

	//3 不存在则读声网接口，写redis缓存
	rtcToken, err = agora.GetRtcTokenWithAccount(appId, appCertificate, channelName, uidStr, expireTimestamp)
	fmt.Printf("GetRtcTokenWithCache in agora: %v \n", rtcToken)
	if err != nil {
		fmt.Printf("GetRtcTokenWithAccount err is %v \n", err)
		return "", err
	}

	err = redisConn.Setex(rtcKey, rtcToken, rtcExpireMinute)
	if err != nil {
		fmt.Printf("Setex err is %v \n", err)
		return "", err
	}
	return rtcToken, err
}

// 获取redis的token
// 存在则返回
// 不存在则读声网接口，写redis缓存
func (l CacheToken) GetJwtTokenWithCache(uidStr, plat, phone, tokenSecrect string) (string, error) {
	fmt.Println("GetJwtTokenWithCache")
	//1 获取redis的token
	redisConn := l.svcCtx.Redis
	jwtTokenKey := rediskey.CacheSocialImAccessTokenPrefix + uidStr
	jwtToken, err := redisConn.Get(jwtTokenKey)
	tokenExpireHour := accessExpireHour
	if err != redis.Nil && err != nil {
		fmt.Printf("get jwt token err %v \n", err)
		return "", err
	}

	//2 存在则返回
	if err == nil && jwtToken != "" {
		fmt.Println("GetJwtTokenWithCache in redis ")
		return jwtToken, err
	}

	//3 不存在则生成新token，写redis缓存
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = uidStr
	claims["phone"] = phone
	claims["plat"] = plat
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenExpireHour)).Unix()
	fmt.Println("new generate JwtTokenWithCache with mobile ")
	jwtToken, err = token.SignedString([]byte(tokenSecrect))
	if err != nil {
		fmt.Printf("generate jwt token err %v \n", err)
		return "", err
	}

	err = redisConn.Setex(jwtTokenKey, jwtToken, int(60*accessExpireHour))
	if err != nil {
		fmt.Printf("set jwt token err %v \n", err)
		return "", err
	}
	return jwtToken, nil
}

func (l CacheToken) GetLoginToken(uidStr, platId, phone string) (string, string, string) {

	//并发获取token
	var wg sync.WaitGroup
	accessChan := make(chan string)
	rtmChan := make(chan string)
	rtcChan := make(chan string)
	wg.Add(3)
	go func() {
		defer wg.Done()
		accessToken, _ := l.GetJwtTokenWithCache(uidStr, platId, phone, l.svcCtx.Config.JwtAuth.AccessSecret)
		fmt.Printf("sync1 end %s \n", accessToken)
		accessChan <- accessToken
	}()

	go func() {
		wg.Done()
		rtmToken, _ := l.GetRtmTokenWithCache(l.svcCtx.Config.AgoraConf.AppId, l.svcCtx.Config.AgoraConf.AppCertificate, uidStr)
		fmt.Printf("sync2 end %s \n", rtmToken)
		rtmChan <- rtmToken
	}()

	go func() {
		wg.Done()
		rtcToken, _ := l.GetRtcTokenWithCache(l.svcCtx.Config.AgoraConf.AppId, l.svcCtx.Config.AgoraConf.AppCertificate, l.svcCtx.Config.AgoraConf.PubChannel, uidStr)
		fmt.Printf("sync3 end %s \n", rtcToken)
		rtcChan <- rtcToken
	}()

	accessToken := <-accessChan
	rtcToken := <-rtcChan
	rtmToken := <-rtmChan
	wg.Wait()
	//并发获取token  end

	return accessToken, rtcToken, rtmToken
}
