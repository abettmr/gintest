package user

import (
	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"
	"cdisk/util"
	"context"
	"log"
	"strings"

	// "github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.DataResponse[types.LoginResp], err error) {
	// todo: add your logic here and delete this line
	// 1. 校验参数
	// model.withSession()
	resp = &types.DataResponse[types.LoginResp]{
		BaseResponse: types.BaseResponse{
			Code:    200,
			Message: "登录成功",
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

	// 连接数据库
	user, err := l.svcCtx.UserModel.FindOneName(l.ctx, req.Name)
	if err != nil && err != sqlx.ErrNotFound {
		log.Printf("Error: %v", err)
		resp.Code = 400
		resp.Message = "服务器内部错误"
		// log.Fatalln(resp.Message)
		return resp, err
	}
	if err != nil {
		resp.Code = 400
		resp.Message = "用户不存在"
		// log.Fatalln(resp.Message)
		return resp, err
	}
	if user.Password != util.Md5(req.Password) {
		resp.Code = 400
		resp.Message = "密码错误"
		// log.Fatalln(resp.Message)
		return resp, err
	}

	token, err := util.BuildTokens(
		util.TokenOptions{
			AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
			AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
			Fields: map[string]interface{}{
				"userId": user.Identity,
			},
		})
	if err != nil {
		resp.Code = 500
		resp.Message = "生成token失败"
		// log.Fatalln(resp.Message)
	}
	resp.Data.Token = token.AccessToken
	resp.Data.UserBasicInfo = types.UserBasicInfo{
		Identity: user.Identity,
		Name:     user.Name,
		Email:    user.Email,
	}
	log.Println("登录成功，token:", token.AccessToken)
	return resp, err

	// 2. 查询用户
	// 3. 校验密码
	// 4. 生成token
	// 5. 返回结果
	// 6. 记录登录日志
	// 7. 更新用户登录时间
	// 8. 更新用户登录ip
	// 9. 更新用户登录设备
	// 10. 更新用户登录状态
	// 11. 更新用户登录次数
}
