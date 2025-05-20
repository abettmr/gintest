package repository

import (
	"context"

	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"
	"cdisk/mysql/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(rp *model.RepoPool) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line
	resp = &types.BaseResponse{
		Code:    200,
		Message: "OK",
	}
	// m := &model.RepoPool{
	// 	Identity:  uuid.New().String(),
	// 	Hash:      req.File,
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// 	Name:      req.File,
	// }
	_, err = l.svcCtx.RepoModel.Insert(l.ctx, rp)
	if err != nil {
		resp.Code = 500
		resp.Message = "服务器内部错误"
	}
	return resp, err
}
