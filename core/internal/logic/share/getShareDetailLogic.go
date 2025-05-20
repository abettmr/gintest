package share

import (
	"context"

	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShareDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShareDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShareDetailLogic {
	return &GetShareDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShareDetailLogic) GetShareDetail(req *types.IdJsonReq) (resp *types.DataResponse[types.ShareDetailInfo], err error) {
	// todo: add your logic here and delete this line

	return
}
