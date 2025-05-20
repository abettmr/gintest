// goctl 1.8.2

package handler

import (
	"net/http"

	repository "cdisk/core/internal/handler/repository"
	share "cdisk/core/internal/handler/share"
	user "cdisk/core/internal/handler/user"
	userRepo "cdisk/core/internal/handler/userRepo"
	"cdisk/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: repository.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/download",
					Handler: repository.DownloadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/folder/create",
					Handler: repository.CreateFolderHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: repository.ListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/move",
					Handler: repository.MoveHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/rename",
					Handler: repository.RenameHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/upload",
					Handler: repository.UploadHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/repo/v1"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/cancel",
					Handler: share.CancelShareHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: share.CreateShareHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/detail",
					Handler: share.GetShareDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/download",
					Handler: share.DownloadShareHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/my/list",
					Handler: share.ListMySharesHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/popular",
					Handler: share.GetPopularSharesHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/stats",
					Handler: share.GetShareStatsHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: share.UpdateShareHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/verify",
					Handler: share.VerifyShareHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/share/v1"),
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user/v1"),
	)
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/detail",
					Handler: user.DetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/password/change",
					Handler: user.ChangePasswordHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/user/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/create",
				Handler: userRepo.CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: userRepo.DeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/detail",
				Handler: userRepo.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: userRepo.ListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/move",
				Handler: userRepo.MoveHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user-repo/v1"),
	)
}
