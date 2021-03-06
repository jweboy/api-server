package qiniu

import (
	"github.com/gin-gonic/gin"
	. "github.com/jweboy/api-server/api"
	"github.com/jweboy/api-server/pkg/errno"
	"github.com/jweboy/api-server/util"
)

// ListBucket 获取存储空间列表
// @Summary 获取存储空间列表
// @Description 获取所有的存储空间列表，无分页。
// @Tags qiniu
// @Accept  json
// @Produce  json
// TODO: 数组需要定义
// @Router /qiniu/bucket [get]
// @Success 200
func ListBucket(c *gin.Context) {
	bucketManger := util.GetBucketManager()

	// FIXME: @param shared 默认为true，文档里说`true的时候一同列表被授权访问的空间`不太理解
	// 后期有需求再做更改 => https://github.com/qiniu/api.v7/blob/master/storage/bucket.go
	buckets, err := bucketManger.Buckets(true)

	if err != nil {
		SendResponse(c, errno.ErrListBucketError, nil)
		return
	}

	SendResponse(c, nil, buckets)
}
