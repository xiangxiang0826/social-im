package agora

// 消息发送者
const (
	ADMINUSER = "1"
)

// 消息类型
const (
	APPLYMIC   = "3"
	APPROVEMIC = "4"
	REJECTMIC  = "5"
	INVITEMIC  = "6"
	ACCEPTMIC  = "7"
	DECLINEMIC = "8"
	REMOVEMIC  = "9"
	STOPMIC    = "10"
	RESUMEMIC  = "11"

	INVITEADMIN  = "14"
	ACCEPTADMIN  = "15"
	DECLINEADMIN = "16"
	REMOVEADMIN  = "17"

	REMOVEROOM    = "2"
	TERMINATEROOM = "101"
	SHUTUPUSER    = "12"
	UNSHUTUPUSER  = "13"

	ToUsersGift = "18"
)

// 消息内容
const (
	REMOVEROOMMSG   = "您已被管理踢出%s房间"
	SHUTUPUSERMSG   = "您已被房主/管理禁言"
	UNSHUTUPUSERMSG = "您的禁言已被解除"
)
