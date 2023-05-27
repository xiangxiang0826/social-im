package room

import (
	"net/http"

	"social-im/app/room/cmd/api/internal/logic/room"
	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"
	"social-im/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func TerminateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PartyTerminateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := room.NewTerminateLogic(r.Context(), svcCtx)
		resp, err := l.Terminate(&req)
		result.HttpResult(r, w, resp, err)
	}
}
