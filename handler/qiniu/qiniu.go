package qiniu

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/spf13/viper"
)

func getKeys() (string, string) {
	accessKey := viper.GetString("qiniu.accessKey")
	secretKey := viper.GetString("qiniu.secretKey")
	return accessKey, secretKey
}

func getToken(bucket string) string {
	accessKey, secretKey := getKeys()
	// get mac
	mac := qbox.NewMac(accessKey, secretKey)
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
