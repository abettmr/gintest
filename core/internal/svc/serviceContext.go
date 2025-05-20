package svc

import (
	"cdisk/core/internal/config"
	"cdisk/core/internal/middleware"
	"cdisk/mysql/model"
	"cdisk/util/oss"
	"os"

	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config        config.Config
	Auth          rest.Middleware
	UserModel     model.UserModel
	RepoModel     model.RepoPoolModel
	ShareModel    model.UserShareModel
	UserRepoModel model.UserRepoModel

	S3Client *s3.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// test db connection
	file, err := os.OpenFile("bucket.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)
	db := sqlx.NewMysql(c.Mysql.DataSource)
	res, err := db.Exec("SELECT 1")
	log.Println("sql test:", res)
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	s3client, err := oss.CreateS3Client(c.Oss.AccessKey, c.Oss.SecretKey, c.Oss.Endpoint, c.Oss.Region)
	if err != nil {
		panic("failed to create oss session: " + err.Error())
	}
	return &ServiceContext{
		Config:        c,
		Auth:          middleware.NewAuthMiddleware(c.Auth.AccessSecret).Handle,
		UserModel:     model.NewUserModel(db),
		RepoModel:     model.NewRepoPoolModel(db),
		ShareModel:    model.NewUserShareModel(db),
		UserRepoModel: model.NewUserRepoModel(db),
		S3Client:      s3client,
	}
}
