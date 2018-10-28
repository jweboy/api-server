package qiniu

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/spf13/viper"
)

// Key 七牛云上传的相关秘钥
type Key struct {
	access string
	secret string
}

func getKeys() Key {
	keys := Key{}
	keys.access = viper.GetString("qiniu.accessKey")
	keys.secret = viper.GetString("qiniu.secretKey")

	return keys
}

func getToken(bucket string) string {
	// get mac
	mac := qbox.NewMac(getKeys().access, getKeys().secret)
	// get policy
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	// 2小时有效期
	putPolicy.Expires = 7200
	// generate upload token
	upToken := putPolicy.UploadToken(mac)

	return upToken
}

func getCfg() storage.Config {
	cfg := storage.Config{}

	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	return cfg
}
