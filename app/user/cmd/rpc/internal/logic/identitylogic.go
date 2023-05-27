package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"social-im/app/user/cmd/rpc/internal/repository"
	"social-im/app/user/cmd/rpc/internal/svc"
	"social-im/app/user/cmd/rpc/pb"
	"social-im/common/fetch"
	"social-im/common/xerr"
	"social-im/common/xorm/errs"
	"strconv"
	"strings"
	"time"
)

type IdentityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rep *repository.Rep
}

func NewIdentityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IdentityLogic {
	return &IdentityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		rep:    repository.NewRep(svcCtx),
	}
}

// 获取今天开始与结束的时间戳
func IsSameDay(t1 int64, t2 int64) bool {
	t11 := time.Unix(t1, 0)
	t22 := time.Unix(t2, 0)
	if t11.Day() == t22.Day() {
		return true
	}
	return false
}

func (l *IdentityLogic) Identity(in *pb.IdentityReq) (*pb.IdentityResp, error) {
	// todo: add your logic here and delete this line
	in.Identity = strings.ToUpper(strings.TrimSpace(in.Identity))
	in.RealName = strings.TrimSpace(in.RealName)

	user, err := l.rep.UserModel.FindOneByUserIdentity(l.ctx, in.RealName, in.Identity)
	if err != nil && !errs.RecordNotFound(err) {
		return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "数据库错误", err)
	}

	if err == nil {
		if user.Id == in.UserId && user.Authtime.Valid && user.Authtime.Int64 > 0 {
			return nil, xerr.NewErrWithFormatMsg(xerr.DB_INFO_ERRO, "用户已用该信息实名认证过，无需再次认证:")
		} else {
			//别人已经用该身份证认证过， 需要记录这个错误
			nowTime := time.Now().Unix()
			nowStr := strconv.FormatInt(nowTime, 10)
			myUserInfo, err := l.rep.UserModel.FindOne(l.ctx, in.UserId)
			if err == nil {
				if myUserInfo.Authtime.Valid && myUserInfo.Authtime.Int64 > 0 {
					return nil, xerr.NewErrWithFormatMsg(xerr.DB_INFO_ERRO, "该用户已经实名认证过:")
				}
				if myUserInfo.RealName.Valid && myUserInfo.RealName.String != "" { //用户实名认证出现过错误
					timeNumStr := strings.Split(myUserInfo.RealName.String, "-")
					errTimes, _ := strconv.Atoi(timeNumStr[1])
					errTime, _ := strconv.Atoi(timeNumStr[0])
					if IsSameDay((int64)(errTime), nowTime) { //错误次数超过3次，提示明天再试
						if errTimes >= 3 {
							return nil, xerr.NewErrWithFormatMsg(xerr.DB_INFO_ERRO, "今天认证错误超过3次， 请明天再试：")
						}
						errTimes = errTimes + 1
						errTimesStr := strconv.Itoa(errTimes)
						realNameTimeErr := nowStr + "-" + errTimesStr

						myUserInfo.RealName.String = realNameTimeErr
						err = l.rep.UserModel.Update(l.ctx, l.rep.Mysql, myUserInfo)
						return nil, xerr.NewErrWithFormatMsg(xerr.DB_INFO_ERRO, "该身份信息已被他人认证，请更换身份信息:")
					} else {
						realNameTimeErr := nowStr + "-" + strconv.Itoa(1)
						myUserInfo.RealName.String = realNameTimeErr
						err = l.rep.UserModel.Update(l.ctx, l.rep.Mysql, myUserInfo)
						return nil, xerr.NewErrWithFormatMsg(xerr.DB_INFO_ERRO, "该身份信息已被他人认证，请更换身份信息:")
					}

				} else {
					realNameTimeErr := nowStr + "-" + strconv.Itoa(1)
					myUserInfo.RealName.String = realNameTimeErr
					myUserInfo.RealName.Valid = true
					err = l.rep.UserModel.Update(l.ctx, l.rep.Mysql, myUserInfo)
					return nil, xerr.NewErrWithFormatMsg(xerr.DB_INFO_ERRO, "该身份信息已被他人认证，请更换身份信息:")

				}
			} else {
				return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "数据库错误", err)
			}
		}
	}

	if errs.RecordNotFound(err) { // 该身份信息没有被其他人验证过
		myUserInfo, err := l.rep.UserModel.FindOne(l.ctx, in.UserId)
		if err == nil {
			if myUserInfo.Authtime.Valid && myUserInfo.Authtime.Int64 > 0 { //用户实名认证过
				return nil, xerr.NewErrWithFormatMsg(xerr.DB_INFO_ERRO, "用户已实名认证过，无需再认证:")
			} else { //用户没有实名认证过， 需要首先检查当日是否有实名认证过， 错误超过三次不给认证
				if myUserInfo.RealName.Valid && myUserInfo.RealName.String != "" {
					nowTime := time.Now().Unix()
					nowStr := strconv.FormatInt(nowTime, 10)

					timeNumStr := strings.Split(myUserInfo.RealName.String, "-")
					errTimes, _ := strconv.Atoi(timeNumStr[1])
					errTime, _ := strconv.Atoi(timeNumStr[0])
					if (IsSameDay(int64(errTime), nowTime) && errTimes < 3) || (!IsSameDay(int64(errTime), nowTime)) { //同一天 的错误次数小于3次或当天没有认证过 ，调用身份认证接口， 成功则更新user信息， 失败更新错误次数

						_, err = Authentication(in.RealName, in.Identity)
						if err != nil { //记录认证错误
							if !(IsSameDay(int64(errTime), nowTime)) { //当天第一次错误
								realNameTimeErr := nowStr + "-" + strconv.Itoa(1)
								myUserInfo.RealName.String = realNameTimeErr
								myUserInfo.RealName.Valid = true
								errs := l.rep.UserModel.Update(l.ctx, l.rep.Mysql, myUserInfo)
								if errs != nil {
									return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "数据库错误", errs)
								} else {
									return nil, xerr.NewErrWithFormatMsg(xerr.DB_INFO_ERRO, "提交的身份证信息实名认证不通过:", err)
								}
							} else { //当天错误累加
								errTimes = errTimes + 1
								errTimesStr := strconv.Itoa(errTimes)
								realNameTimeErr := nowStr + "-" + errTimesStr
								myUserInfo.RealName.String = realNameTimeErr

								errs := l.rep.UserModel.Update(l.ctx, l.rep.Mysql, myUserInfo)
								if errs != nil {
									return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "数据库错误", errs)
								} else {
									return nil, xerr.NewErrWithFormatMsg(xerr.DB_INFO_ERRO, "提交的身份证信息实名认证不通过:", err)
								}

							}
						} else { //认证成功更新信息
							myUserInfo.RealName.String = in.RealName
							myUserInfo.Identity.Valid = true
							myUserInfo.Identity.String = in.Identity
							myUserInfo.Authtime.Valid = true
							myUserInfo.Authtime.Int64 = nowTime
							fmt.Println("will update :", myUserInfo)
							err = l.rep.UserModel.Update(l.ctx, l.rep.Mysql, myUserInfo)
							if err == nil {
								return nil, xerr.NewErrWithFormatMsg(xerr.OK, "认证成功：", err)
							} else {
								return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "数据库错误", err)
							}

						}
					} else {
						return nil, xerr.NewErrWithFormatMsg(xerr.DB_INFO_ERRO, "今天认证错误超过3次， 请明天再次：")
					}
				} else { //第一次实名认证
					nowTime := time.Now().Unix()
					nowStr := strconv.FormatInt(nowTime, 10)
					_, err = Authentication(in.RealName, in.Identity)

					if err == nil {
						myUserInfo.RealName.Valid = true
						myUserInfo.RealName.String = in.RealName
						myUserInfo.Identity.Valid = true
						myUserInfo.Identity.String = in.Identity
						myUserInfo.Authtime.Valid = true
						myUserInfo.Authtime.Int64 = nowTime
						fmt.Println("will update :", myUserInfo)

						err = l.rep.UserModel.Update(l.ctx, l.rep.Mysql, myUserInfo)
						if err != nil {
							return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "数据库错误", err)
						} else {
							return nil, xerr.NewErrWithFormatMsg(xerr.OK, "认证成功：", err)
						}
					} else { //认证失败
						realNameTimeErr := nowStr + "-" + strconv.Itoa(1)
						myUserInfo.RealName.String = realNameTimeErr
						myUserInfo.RealName.Valid = true
						errs := l.rep.UserModel.Update(l.ctx, l.rep.Mysql, myUserInfo)
						if errs != nil {
							return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "数据库错误", errs)
						} else {
							return nil, xerr.NewErrWithFormatMsg(xerr.DB_INFO_ERRO, "提交的身份证信息实名认证不通过:", err)
						}

					}

				}
			}

		} else {
			return nil, xerr.NewErrWithFormatMsg(xerr.DB_ERROR, "数据库错误", err)
		}

	}

	return &pb.IdentityResp{}, nil
}

type AuthResponse struct {
	RequestId     string `json:"request_id""`
	Status        string `json:"status"`
	State         int    `json:"state"`
	ResultCode    int    `json:"result_code"`
	ResultMessage string `json:"result_message"`
}

// 实名认证
func Authentication(name, idCard string) (bool, error) {

	appCode := "6957504054e0446ebe797fe5b392ffe4"
	url := "http://dfidveri.market.alicloudapi.com/verify_id_name"

	header := http.Header{}
	header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	header.Add("Authorization", fmt.Sprintf(`APPCODE %v`, appCode))

	result, err := fetch.Cmd(fetch.Request{
		Method: "POST",
		URL:    url,
		Body:   []byte(fmt.Sprintf(`id_number=%v&name=%v`, idCard, name)),
		Header: header,
	})

	var res AuthResponse
	err = json.Unmarshal(result, &res)
	if err == nil {
		if res.State == 1 {
			return true, nil
		} else {
			return false, xerr.NewErrWithFormatMsg(xerr.IDENTITY_INFO_ERRO, "身份信息实名验证不通过：", err)
		}
	} else {
		return false, xerr.NewErrWithFormatMsg(xerr.IDENTITY_INFO_ERRO, "无效或错误的身份证信息：", err)
	}
}
