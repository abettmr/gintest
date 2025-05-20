package share

import (
	"context"

	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateShareLogic {
	return &UpdateShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateShareLogic) UpdateShare(req *types.ShareBasicInfo) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
