package repository

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"cdisk/core/internal/logic/repository"
	"cdisk/core/internal/svc"
	"cdisk/core/internal/types"
	"cdisk/mysql/model"
	"cdisk/util/oss"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var req types.FileUploadReq
		// if err := httpx.Parse(r, &req); err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// 	return
		// }
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		b := make([]byte, fileHeader.Size)
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp, err := svcCtx.RepoModel.FindHash(r.Context(), hash)
		if err != nil && err != sqlx.ErrNotFound {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		var resp *types.BaseResponse
		if rp != nil && err == nil {
			resp = &types.BaseResponse{
				Code:    200,
				Message: "文件已存在",
			}
		}
		if err == sqlx.ErrNotFound {
			// 文件不存在，执行上传逻辑
			// 这里可以调用上传逻辑，比如将文件保存到指定位置
			// 这里假设上传成功，返回成功响应
			// 你可以根据实际情况修改这个逻辑
			// 例如，保存文件到本地磁盘或云存储等
			putres, err := svcCtx.S3Client.PutObject(r.Context(),
				&s3.PutObjectInput{
					Bucket: &svcCtx.Config.Oss.Bucket,
					Key:    &fileHeader.Filename,
					Body:   file,
				})
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
			if putres == nil {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
			// req.FileName = fileHeader.Filename
			// req.FileSize = fileHeader.Size
			path, err := oss.GeneratePresignedURL(svcCtx.S3Client,
				svcCtx.Config.Oss.Bucket, fileHeader.Filename, 7*24*time.Hour)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
				return
			}
			log.Println(path)
			l := repository.NewUploadLogic(r.Context(), svcCtx)
			rp := &model.RepoPool{}
			rp.Identity = uuid.New().String()
			rp.Name = fileHeader.Filename
			rp.Ext = filepath.Ext(fileHeader.Filename)
			rp.Hash = hash
			rp.Size = fileHeader.Size
			rp.CreatedAt = time.Now()
			rp.UpdatedAt = time.Now()
			rp.Path = "/"
			resp, err = l.Upload(rp)
		}

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
