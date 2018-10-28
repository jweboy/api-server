package qiniu

import (
	"github.com/gin-gonic/gin"
	. "github.com/jweboy/api-server/handler"
	"github.com/jweboy/api-server/pkg/errno"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

// ListBucket 获取空间列表
func ListBucket(c *gin.Context) {
	// get keys
	accessKey := getKeys().access
	secretKey := getKeys().secret

	// new qbox mac
	mac := qbox.NewMac(accessKey, secretKey)

	// set default storage config => default http
	cfg := storage.Config{}

	// new bucketmanager
	bucketManger := storage.NewBucketManager(mac, &cfg)

	// @param shared 默认为true，文档里说`true的时候一同列表被授权访问的空间`不太理解
	// 后期有需求再做更改 => https://github.com/qiniu/api.v7/blob/master/storage/bucket.go
	buckets, err := bucketManger.Buckets(true)

	if err != nil {
		SendResponse(c, errno.ErrListBucketError, nil)
		return
	}

	SendResponse(c, nil, buckets)
}
