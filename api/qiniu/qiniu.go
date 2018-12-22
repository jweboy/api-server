package qiniu

import (
	"github.com/jweboy/api-server/pkg/setting"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

func getMac() *qbox.Mac {
	// get mac
	mac := qbox.NewMac(
		setting.QiniuSetting.AccessKey,
		setting.QiniuSetting.SecretKey,
	)
	return mac
}

func getToken(bucket string) string {
	// get mac
	mac := getMac()

	// get policy
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
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

func GetBucketManager() *storage.BucketManager {
	mac := getMac()
	cfg := getCfg()

	bucketManger := storage.NewBucketManager(mac, &cfg)
	return bucketManger
}
