package room

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"social-im/app/room/cmd/api/internal/logic/room"
	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"
	"social-im/common/result"
	"social-im/common/utils/validator"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PartyCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		if errMsg := validator.NewValidator().Validate(req, r.Header.Get("Accept-Language")); errMsg != "" {
			result.ParamErrorResult(r, w, errors.New(errMsg))
			return
		}
		l := room.NewCreateLogic(r.Context(), svcCtx)
		resp, err := l.Create(&req)
		result.HttpResult(r, w, resp, err)
	}
}
