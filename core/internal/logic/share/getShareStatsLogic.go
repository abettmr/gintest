package share

import (
	"context"

	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShareStatsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShareStatsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShareStatsLogic {
	return &GetShareStatsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShareStatsLogic) GetShareStats() (resp *types.DataResponse[types.ShareStats], err error) {
	// todo: add your logic here and delete this line

	return
}
