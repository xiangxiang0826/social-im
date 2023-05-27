package user

import (
	"net/http"
	"social-im/app/user/cmd/api/internal/logic/user"
	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
	"social-im/common/result"
	"social-im/common/xhttp"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SmsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SmsReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewSmsLogic(r.Context(), svcCtx)
		resp, err := l.Sms(&req)
		result.HttpResult(r, w, resp, err)

	}
}
