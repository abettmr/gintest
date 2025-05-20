package share

import (
	"context"

	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPopularSharesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPopularSharesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPopularSharesLogic {
	return &GetPopularSharesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPopularSharesLogic) GetPopularShares(req *types.PopShareReq) (resp *types.DataResponse[[]types.ShareListItem], err error) {
	// todo: add your logic here and delete this line

	return
}
