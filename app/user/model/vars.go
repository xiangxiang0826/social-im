package model

import "gorm.io/gorm"

var ErrNotFound = gorm.ErrRecordNotFound
var UserAuthTypeSystem string = "system"  //平台内部
var UserAuthTypeSmallWX string = "wxMini" //微信小程序