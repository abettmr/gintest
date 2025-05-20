package share

import (
	"context"

	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMySharesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListMySharesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMySharesLogic {
	return &ListMySharesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMySharesLogic) ListMyShares() (resp *types.PageResponse[types.ShareListItem], err error) {
	// todo: add your logic here and delete this line

	return
}
