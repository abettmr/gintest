package share

import (
	"context"

	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadShareLogic {
	return &DownloadShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadShareLogic) DownloadShare(req *types.IdJsonReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
