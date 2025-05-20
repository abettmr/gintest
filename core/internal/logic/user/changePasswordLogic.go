package user

import (
	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"
	"cdisk/util"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.ChangePasswordReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line
	resp = &types.BaseResponse{
		Code:    200,
		Message: "修改密码成功",
	}

	userId := l.ctx.Value("userId").(string)
	user, err := l.svcCtx.UserModel.FindOneName(l.ctx, userId)
	if err != nil && err != sqlx.ErrNotFound {
		// log.Fatalln(err)
	}
	if user == nil {
		resp.Code = 400
		resp.Message = "用户不存在"
		return resp, nil
	}
	if user.Password != util.Md5(req.OldPassword) {
		resp = &types.BaseResponse{
			Code:    400,
			Message: "旧密码错误",
		}
		return resp, nil
	}
	user.Password = util.Md5(req.NewPassword)
	user.UpdatedAt = time.Now()

	l.svcCtx.UserModel.Update(l.ctx, user)
	resp = &types.BaseResponse{
		Code:    200,
		Message: "修改密码成功",
	}
	return resp, nil
}
