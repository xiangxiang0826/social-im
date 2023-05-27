package admin

import (
	"errors"
	"net/http"
	"social-im/common/result"
	"social-im/common/utils/validator"

	"github.com/zeromicro/go-zero/rest/httpx"
	"social-im/app/admin/cmd/api/internal/logic/admin"
	"social-im/app/admin/cmd/api/internal/svc"
	"social-im/app/admin/cmd/api/internal/types"
)

func BackgroundImageListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BackgroundImgReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		if errMsg := validator.NewValidator().Validate(req, r.Header.Get("Accept-Language")); errMsg != "" {
			result.ParamErrorResult(r, w, errors.New(errMsg))
			return
		}
		l := admin.NewBackgroundImageListLogic(r.Context(), svcCtx)
		resp, err := l.BackgroundImageList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
