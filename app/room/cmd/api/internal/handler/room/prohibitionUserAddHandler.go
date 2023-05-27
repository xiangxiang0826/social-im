package room

import (
	"errors"
	"net/http"
	"social-im/common/result"
	"social-im/common/utils/validator"

	"github.com/zeromicro/go-zero/rest/httpx"
	"social-im/app/room/cmd/api/internal/logic/room"
	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"
)

func ProhibitionUserAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProhibitionCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		if errMsg := validator.NewValidator().Validate(req, r.Header.Get("Accept-Language")); errMsg != "" {
			result.ParamErrorResult(r, w, errors.New(errMsg))
			return
		}
		l := room.NewProhibitionUserAddLogic(r.Context(), svcCtx)
		resp, err := l.ProhibitionUserAdd(&req)
		result.HttpResult(r, w, resp, err)
	}
}
