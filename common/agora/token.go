package agora

import (
	"fmt"
	"time"

	rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/RtcTokenBuilder"
	"github.com/dgrijalva/jwt-go"

	// rtmtokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/RtmTokenBuilder"
	rtmtokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtmtokenbuilder2"
)

func GetRtmToken(appId, appCertificate, uidStr string, expireTimestamp uint32) (string, error) {
	fmt.Println("Generating RTM token")

	rtmToken, err := rtmtokenbuilder.BuildToken(appId, appCertificate, uidStr, expireTimestamp)

	if err != nil {
		return "", err
	} else {
		return rtmToken, err
	}
}

func GetRtcTokenWithUid(appId, appCertificate, channelName string, uid, expireTimestamp uint32) (string, error) {
	fmt.Println("Generating RTM token")

	rtcToken, err := rtctokenbuilder.BuildTokenWithUID(appId, appCertificate, channelName, uid, rtctokenbuilder.RoleAttendee, expireTimestamp)

	if err != nil {
		return "", err
	} else {
		return rtcToken, err
	}
}

func GetRtcTokenWithAccount(appId, appCertificate, channelName, uidStr string, expireTimestamp uint32) (string, error) {
	fmt.Println("Generating RTM token")

	rtcToken, err := rtctokenbuilder.BuildTokenWithUserAccount(appId, appCertificate, channelName, uidStr, rtctokenbuilder.RoleAttendee, expireTimestamp)

	if err != nil {
		return "", err
	} else {
		return rtcToken, err
	}
}

func GetJwtToken(uid, platid, accessSecret string, accessExpire int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = uid
	claims["plat"] = "0"
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(accessExpire)).Unix()

	// tokenString, err := token.SignedString([]byte("social_im"))
	tokenString, err := token.SignedString([]byte(accessSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
