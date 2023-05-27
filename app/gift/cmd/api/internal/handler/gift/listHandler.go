package gift

import (
	"errors"
	"net/http"
	"social-im/app/gift/cmd/api/internal/logic/gift"
	"social-im/app/gift/cmd/api/internal/svc"
	"social-im/app/gift/cmd/api/internal/types"
	"social-im/common/result"
	"social-im/common/utils/validator"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GiftListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if errMsg := validator.NewValidator().Validate(req, r.Header.Get("Accept-Language")); errMsg != "" {
			result.ParamErrorResult(r, w, errors.New(errMsg))
			return
		}

		l := gift.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		result.HttpResult(r, w, resp, err)
	}
}
