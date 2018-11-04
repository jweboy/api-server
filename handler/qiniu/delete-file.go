package qiniu

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	. "github.com/jweboy/api-server/handler"
	"github.com/jweboy/api-server/model"
	"github.com/jweboy/api-server/pkg/errno"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	log "qiniupkg.com/x/log.v7"
)

// Query 删除文件Query请求参数
type Query struct {
	Bucket string `form:"bucket"`
	Name   string `form:"name"`
}

// DeleteFile 删除指定空间的文件
func DeleteFile(c *gin.Context) {
	fileID, err := strconv.Atoi(c.Param("id"))

	// ID字段转换错误
	if err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// 请求Query中的bucket字段不为空
	var query Query
	if c.ShouldBindQuery(&query) == nil {
		if query.Bucket == "" || query.Name == "" {
			err := fmt.Errorf("请求参数buckt与name不能为空")
			log.Errorf("Query error:", err)
			SendResponse(c, err, nil)
			return
		}
	}

	// 删除七牛云指定文件
	accessKey, secretKey := getKeys()
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	if err := bucketManager.Delete(query.Bucket, query.Name); err != nil {
		log.Errorf("Qiniu SDK error:", err)
		SendResponse(c, err, nil)
		return
	}

	// 删除数据库对应数据
	if err := model.DeleteFile(uint64(fileID)); err != nil {
		log.Error("Databse error：", err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// 返回最终结果
	SendResponse(c, nil, nil)
}
