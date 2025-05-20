package repository

import (
	"context"

	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.FileListReq) (resp *types.PageResponse[types.FileItem], err error) {
	// todo: add your logic here and delete this line
	resp = &types.PageResponse[types.FileItem]{
		BaseResponse: types.BaseResponse{
			Code:    200,
			Message: "获取文件列表成功",
		},
	}
	// repo, err := l.svcCtx.RepoModel.FindAllParent(l.ctx, uint64(*req.ParentId))
	if err != nil && err != sqlx.ErrNotFound {
		logx.Errorf("Error: %v", err)
		resp = &types.PageResponse[types.FileItem]{
			BaseResponse: types.BaseResponse{
				Code:    500,
				Message: "服务器内部错误",
			},
		}
		return resp, err
	}
	if err == sqlx.ErrNotFound {
		resp = &types.PageResponse[types.FileItem]{
			BaseResponse: types.BaseResponse{
				Code:    400,
				Message: "文件夹不存在",
			},
		}
		return resp, err
	}
	// resp.Data = make([]types.FileItem)

	// listRes, err := l.svcCtx.S3Client.ListObjectsV2(&s3.ListObjectsV2Input{
	// 	Bucket: aws.String("your-bucket-name"),
	// 	Prefix: aws.String(req.Path),
	// })
	return
}
