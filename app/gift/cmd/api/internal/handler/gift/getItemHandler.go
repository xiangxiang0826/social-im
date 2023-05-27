package gift

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"social-im/app/gift/cmd/api/internal/logic/gift"
	"social-im/app/gift/cmd/api/internal/svc"
	"social-im/app/gift/cmd/api/internal/types"
)

func GetItemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetItemReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := gift.NewGetItemLogic(r.Context(), svcCtx)
		resp, err := l.GetItem(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
