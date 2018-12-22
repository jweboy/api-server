package qiniu

import (
	"fmt"
	"strconv"

	"github.com/jweboy/api-server/model"

	"github.com/gin-gonic/gin"
	. "github.com/jweboy/api-server/api"
	"github.com/jweboy/api-server/pkg/errno"
	"github.com/jweboy/api-server/pkg/setting"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

func ChangeMime(c *gin.Context) {
	mime := c.PostForm("mime")

	id, _ := strconv.Atoi(c.PostForm("id"))

	detail, err := model.FileDetail(id)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// TODO: 这里可以提取为公共函数
	mac := qbox.NewMac(
		setting.QiniuSetting.AccessKey,
		setting.QiniuSetting.SecretKey,
	)
	cfg := storage.Config{}
	bucketManager := storage.NewBucketManager(mac, &cfg)

	bucket := detail.Bucket
	fileName := detail.Name

	repErr := bucketManager.ChangeMime(bucket, fileName, mime)
	if err != nil {
		fmt.Println(repErr)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)

}
