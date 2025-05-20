package oss

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const s3Endpoint = "https://s3.bitiful.net"
const s3Region = "cn-east-1"

// func GetS3Client() *s3.Client {
// 	cli, err := CreateS3Client(s3AccessKey, s3SecretKey, s3Endpoint, s3Region)
// 	if err != nil {
// 		panic("failed to getS3Client: " + err.Error())
// 	}
// 	return cli
// }

func CreateS3Client(accessKey, secretKey, endpoint, region string) (*s3.Client, error) {
	// 使用静态凭证提供程序
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		config.WithRegion(region),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	// 创建 S3 客户端
	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
		o.UsePathStyle = true
		o.Region = region
	}), nil
}

func GeneratePresignedURL(s3Client *s3.Client, bucket, key string, duration time.Duration) (string, error) {
	// 创建 Presign 客户端
	presignClient := s3.NewPresignClient(s3Client)

	// 生成预签名 URL
	req, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}, func(opts *s3.PresignOptions) {
		opts.Expires = duration
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return req.URL, nil
}
