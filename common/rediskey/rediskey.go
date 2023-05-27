package rediskey

var (
	CacheSocialImUserIdPrefix      = "cache111:socialIm:user:id:"
	CacheSocialImUserMobilePrefix  = "cache111:socialIm:user:mobile:"
	CacheSocialImVerifyCodePrefix  = "cache111:socialIm:sms:verify:"
	CacheSocialImVerifyLimitPrefix = "cache111:socialIm:sms:limit:"
	CacheSocialImRtmTokenPrefix    = "cache111:socialIm:sms:rtmtoken:"
	CacheSocialImRtcTokenPrefix    = "cache111:socialIm:sms:rtctoken:"
	CacheSocialImAccessTokenPrefix = "cache111:socialIm:sms:accesstoken:test1"

	CacheSocialImRoomAdminPrefix           = "cache:socialIm:roomAdmin:"
	CacheSocialImRoomUserAdminStatusPrefix = "cache:socialIm:roomId:adminStatus:"

	CacheSocialImUserFollowersPrefix = "cache:socialIm:roomId:userFollowers:"
	CacheSocialImUserFollowsPrefix   = "cache:socialIm:roomId:userFollows:"

	//room
	CacheSocialImRoomOnLineHash = "cache111:socialIm:roomonline:id:"
	CacheSocialImRoomUser       = "cache:room:appRoomUser:id:"       //进入房间的用户
	CacheSocialImRoomBloom      = "cache:room:id:bloom:"             //统计房间uv的布隆过滤器
	CacheSocialImRoomUV         = "cache:room:id:uv:"                //房间uv
	CacheSocialImRoomAvatarList = "cache111:socialIm:roomavatar:id:" //房间的头像列表

	//gift
	CacheSocialImGiftListKey = "cache:Gift:list:"
)
