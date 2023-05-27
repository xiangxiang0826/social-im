package repository

type UserInfo struct {
	Avatar string
}

type PartyInfo struct {
	Id                 int64
	Name               string                //名称
	Mark               string //房间唯一标识id
	BackgroundSmallUrl string //背景小图地址
	PartyType          int64
	CreatedAt          int64
	OnlineNums         int64
	IsHot              int64
	UserList           []*UserInfo
}

type UserProhibitionChannelMsg struct {
	UserId int64 `json:"user_id"`
	Status int64 `json:"status"`
	Id     int64 `json:"id"`
}