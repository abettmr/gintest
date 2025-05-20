package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	// Mysql mysql config
	Mysql struct {
		DataSource string
	} `json:"mysql"`
	Auth struct {
		AccessSecret string
		AccessExpire int64
	} `json:"auth"`
	// Redis redis config
	Redis struct {
		Address string
		Port    int
	} `json:"redis"`

	Oss struct {
		Endpoint  string
		AccessKey string
		SecretKey string
		Bucket    string
		Region    string
		// Domain     string
	} `json:"oss"`
	// Jwt struct {
	// 	AccessSecret string
	// 	AccessExpire int64
	// } `json:"jwt"`
}
