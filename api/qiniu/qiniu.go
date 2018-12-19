package qiniu

import (
	"github.com/jweboy/api-server/pkg/setting"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

var (
	accessKey = setting.QiniuSetting.AccessKey
	secretKey = setting.QiniuSetting.SecretKey
	AccessKey = setting.QiniuSetting.AccessKey
	// secretKey = setting.QiniuSetting.SecretKey
)

func getKeys() (string, string) {
	return accessKey, secretKey
}

func getToken(bucket string) string {
	// get mac
	mac := qbox.NewMac(
		setting.QiniuSetting.AccessKey,
		setting.QiniuSetting.SecretKey,
	)
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
