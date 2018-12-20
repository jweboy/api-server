package qiniu

import (
	"fmt"

	"github.com/gin-gonic/gin"
	. "github.com/jweboy/api-server/api"
	"github.com/jweboy/api-server/model"
	"github.com/jweboy/api-server/pkg/errno"
	"github.com/jweboy/api-server/pkg/setting"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

// DetailQuery 文件详情请求query
type DetailQuery struct {
	ID int `form:"id" binding:"required"`
}

// FileDetail 获取文件详情
func FileDetail(c *gin.Context) {
	var query DetailQuery

	// 检查query
	if c.BindQuery(&query) != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	id := query.ID
	detail, err := model.FileDetail(id)

	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	bucket := detail.Bucket
	key := detail.Name

	// TODO: 这里可以提取为公共函数
	mac := qbox.NewMac(
		setting.QiniuSetting.AccessKey,
		setting.QiniuSetting.SecretKey,
	)

	cfg := storage.Config{}

	bucketManager := storage.NewBucketManager(mac, &cfg)

	info, err := bucketManager.Stat(bucket, key)

	// TODO: 返回的中文文件的名称需要在getFileList的时候转为对应的中文名
	if err != nil {
		fmt.Println(err)
		SendResponse(c, errno.ErrListBucketError, nil)
		return
	}

	SendResponse(c, nil, info)
}
