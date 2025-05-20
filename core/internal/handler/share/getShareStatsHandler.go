package share

import (
	"net/http"

	"cdisk/core/internal/logic/share"
	"cdisk/core/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetShareStatsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := share.NewGetShareStatsLogic(r.Context(), svcCtx)
		resp, err := l.GetShareStats()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
