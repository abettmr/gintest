package repository

import (
	"context"

	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveLogic {
	return &MoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveLogic) Move(req *types.FileMoveReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
