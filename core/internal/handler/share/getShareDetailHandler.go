package share

import (
	"net/http"

	"cdisk/core/internal/logic/share"
	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetShareDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdJsonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := share.NewGetShareDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetShareDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
