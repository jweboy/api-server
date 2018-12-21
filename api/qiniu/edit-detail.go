package qiniu

import (
	"fmt"
	"log"
	"strconv"

	. "github.com/jweboy/api-server/api"
	"github.com/jweboy/api-server/model"
	"github.com/jweboy/api-server/pkg/errno"
	"github.com/jweboy/api-server/pkg/setting"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"

	"github.com/gin-gonic/gin"
)

type EditModel struct {
	fileName string `form:"name" binding:"required"`
	ID       int    `form:"id" binding:"required"`
}

func EditDetail(c *gin.Context) {
	// var editModel EditModel

	log.Println(2, c.PostForm("fileName"))
	// TODO: 校验需要抽取公共函数
	// if err := c.ShouldBindJSON(&editModel); err != nil {
	// 	log.Println("====== Bind By JSON ======")
	// 	log.Println(1, editModel.fileName)
	// }

	// FIXME: 目前的编辑模式指定目标空间和源空间相同
	// FIXME: 不支持跨机房空间
	id, _ := strconv.Atoi(c.PostForm("id"))

	destFileName := c.PostForm("name")

	// 从数据库获取对应的文件名称
	detail, err := model.FileDetail(id)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	idUint64, _ := strconv.ParseUint(c.PostForm("id"), 10, 64)

	// step1. 更新七牛云存储的文件信息
	srcBucket := detail.Bucket
	destBucket := srcBucket
	destKey := c.PostForm("name")
	srcKey := detail.Name

	// TODO:
	// 这里可以提取为公共函数
	mac := qbox.NewMac(
		setting.QiniuSetting.AccessKey,
		setting.QiniuSetting.SecretKey,
	)
	cfg := storage.Config{}
	bucketManager := storage.NewBucketManager(mac, &cfg)

	fmt.Println(srcBucket, destBucket, destKey, srcKey)
	fmt.Println(detail.Bucket)

	putErr := bucketManager.Move(srcBucket, srcKey, destBucket, destKey, false)
	if putErr != nil {
		fmt.Println(putErr)
		SendResponse(c, errno.ErrListBucketError, nil)
		return
	}

	// step2. 更新数据库文件细信息
	file := model.FileModel{
		Name: destFileName,
		Id:   idUint64,
	}

	if err := file.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// step3. 返回请求体
	SendResponse(c, nil, detail)

}
