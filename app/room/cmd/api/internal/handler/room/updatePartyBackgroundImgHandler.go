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

func UpdatePartyBackgroundImgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PartyBackGroundImgUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		if errMsg := validator.NewValidator().Validate(req, r.Header.Get("Accept-Language")); errMsg != "" {
			result.ParamErrorResult(r, w, errors.New(errMsg))
			return
		}
		l := room.NewUpdatePartyBackgroundImgLogic(r.Context(), svcCtx)
		resp, err := l.UpdatePartyBackgroundImg(&req)
		result.HttpResult(r, w, resp, err)
	}
}
