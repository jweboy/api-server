package qiniu

import (
	"fmt"

	"github.com/jweboy/api-server/pkg/setting"

	"github.com/gin-gonic/gin"
	. "github.com/jweboy/api-server/api"
	"github.com/jweboy/api-server/pkg/errno"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

// ListBucket 获取存储空间列表
// @Summary 获取存储空间列表
// @Description 获取所有的存储空间列表，无分页。
// @Tags qiniu
// @Accept  json
// @Produce  json
// @Router /qiniu/bucket [get]
// @Success 200 {object} handler.Response "{"code":0,"message":"ok", "data": []}"
func ListBucket(c *gin.Context) {
	// new qbox mac
	mac := qbox.NewMac(
		setting.QiniuSetting.AccessKey,
		setting.QiniuSetting.SecretKey,
	)

	// set default storage config => default http
	cfg := storage.Config{}

	// new bucketmanager
	bucketManger := storage.NewBucketManager(mac, &cfg)

	// FIXME: @param shared 默认为true，文档里说`true的时候一同列表被授权访问的空间`不太理解
	// 后期有需求再做更改 => https://github.com/qiniu/api.v7/blob/master/storage/bucket.go
	buckets, err := bucketManger.Buckets(true)

	if err != nil {
		fmt.Print(err, accessKey, secretKey)
		SendResponse(c, errno.ErrListBucketError, nil)
		return
	}

	SendResponse(c, nil, buckets)
}
