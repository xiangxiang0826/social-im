package admin

import (
	"net/http"
	"social-im/common/result"

	"social-im/app/admin/cmd/api/internal/logic/admin"
	"social-im/app/admin/cmd/api/internal/svc"
	"social-im/app/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DictionaryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictionaryGetReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		// if errMsg := validator.NewValidator().Validate(req, r.Header.Get("Accept-Language")); errMsg != "" {
		// 	result.ParamErrorResult(r, w, errors.New(errMsg))
		// 	return
		// }
		l := admin.NewDictionaryListLogic(r.Context(), svcCtx)
		resp, err := l.DictionaryList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
