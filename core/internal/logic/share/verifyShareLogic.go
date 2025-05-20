package share

import (
	"context"

	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyShareLogic {
	return &VerifyShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyShareLogic) VerifyShare(req *types.ShareVerifyReq) (resp *types.DataResponse[types.ShareVerifyResp], err error) {
	// todo: add your logic here and delete this line

	return
}
