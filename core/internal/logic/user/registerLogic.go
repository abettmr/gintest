package user

import (
	"context"
	"strings"
	"time"

	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"
	"cdisk/mysql/model"
	"cdisk/util"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.DataResponse[types.LoginResp], err error) {
	// todo: add your logic here and delete this line
	resp = &types.DataResponse[types.LoginResp]{
		BaseResponse: types.BaseResponse{
			Code:    200,
			Message: "注册成功",
		},
		Data: types.LoginResp{
			Token:         "",
			UserBasicInfo: types.UserBasicInfo{}, // Initialize with an empty UserBasicInfo struct
		},
	}
	//第一步：判断异常值
	if len(strings.TrimSpace(req.Name)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		resp.Code = 400
		resp.Message = "用户ID或密码为空"
		return resp, nil
	}
	//第二步：查询数据库
	user, err := l.svcCtx.UserModel.FindOneName(l.ctx, req.Name)
	if err != nil {
		if err != sqlx.ErrNotFound {
			// log.Fatalln(err)
			resp.Code = 500
			resp.Message = "注册失败，请稍后再试"
			return resp, err
		}
	}
	if user != nil {
		// 用户已存在
		resp.Code = 400
		resp.Message = "用户已存在"
		return resp, nil
	}
	//第三步：插入数据库
	user = &model.User{
		Name:      req.Name,
		Password:  util.Md5(req.Password),
		Identity:  uuid.New().String(),
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		resp.Code = 500
		resp.Message = "注册失败，请稍后再试"
		return resp, err
	}
	return resp, nil
}
