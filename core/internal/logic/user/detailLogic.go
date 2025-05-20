package user

import (
	"context"
	"log"

	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail() (resp *types.DataResponse[types.UserDetailResp], err error) {
	// todo: add your logic here and delete this line
	resp = &types.DataResponse[types.UserDetailResp]{
		BaseResponse: types.BaseResponse{
			Code:    200,
			Message: "获取用户详情成功",
		},
		Data: types.UserDetailResp{
			CreatedAt:     "",
			UserBasicInfo: types.UserBasicInfo{}, // Initialize with an empty UserBasicInfo struct
		},
	}
	userIdStr := l.ctx.Value("userId").(string)
	// userId, err := strconv.ParseUint(userIdStr, 10, 64)
	user, err := l.svcCtx.UserModel.FindOneUuid(l.ctx, userIdStr)
	if err != nil && err != sqlx.ErrNotFound {
		resp.Code = 500
		resp.Message = "服务器内部错误"
		log.Printf("Error finding user: %v\n", err)
		return resp, err
	}
	if user == nil {
		resp.Code = 400
		resp.Message = "用户不存在"
		return resp, nil
	}
	resp.Data.CreatedAt = user.CreatedAt.Format("2006-01-02 15:04:05")
	resp.Data.UserBasicInfo = types.UserBasicInfo{
		Identity: user.Identity,
		Name:     user.Name,
		Email:    user.Email,
	}
	return resp, nil
}
