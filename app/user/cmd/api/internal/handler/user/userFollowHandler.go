package user

import (
	"net/http"
	"social-im/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"social-im/app/user/cmd/api/internal/logic/user"
	"social-im/app/user/cmd/api/internal/svc"
	"social-im/app/user/cmd/api/internal/types"
)

func UserFollowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFollowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserFollowLogic(r.Context(), svcCtx)
		resp, err := l.UserFollow(&req)
		result.HttpResult(r, w, resp, err)
	}
}
