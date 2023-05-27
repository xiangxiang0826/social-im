package room

import (
	"errors"
	"fmt"
	"net/http"

	"social-im/app/room/cmd/api/internal/logic/room"
	"social-im/app/room/cmd/api/internal/svc"
	"social-im/app/room/cmd/api/internal/types"
	"social-im/common/result"
	"social-im/common/utils/validator"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func JoinHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PartyJoinReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if errMsg := validator.NewValidator().Validate(req, r.Header.Get("Accept-Language")); errMsg != "" {
			result.ParamErrorResult(r, w, errors.New(errMsg))
			return
		}

		l := room.NewJoinLogic(r.Context(), svcCtx)
		resp, err := l.Join(&req)
		fmt.Printf("resp is %v \n", resp)
		result.HttpResult(r, w, resp, err)
	}
}
