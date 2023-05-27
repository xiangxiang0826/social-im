package user

import (
	"net/http"
	"social-im/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"social-im/app/user/cmd/api/internal/logic/user"
	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
)

func DeclineMicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeclineMicReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewDeclineMicLogic(r.Context(), svcCtx)
		resp, err := l.DeclineMic(&req)
		result.HttpResult(r, w, resp, err)
	}
}
