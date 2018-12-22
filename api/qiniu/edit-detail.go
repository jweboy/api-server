package qiniu

import (
	"fmt"

	. "github.com/jweboy/api-server/api"
	"github.com/jweboy/api-server/model"
	"github.com/jweboy/api-server/pkg/errno"
	"github.com/jweboy/api-server/util"

	"github.com/gin-gonic/gin"
)

// EditModel 编辑文件请求 body
type EditModel struct {
	FileName string `form:"name" binding:"required"`
	ID       int    `form:"id" binding:"required"`
}

// EditDetail 更新文件信息
func EditDetail(c *gin.Context) {
	var m EditModel

	if err := c.ShouldBind(&m); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	/* ============= 从数据库获取文件信息 ============= */
	var id = m.ID
	d, err := model.FileDetail(id)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	/* ============= 更新七牛云存储的文件信息 ============= */
	// FIXME: 目前的编辑模式指定目标空间和源空间相同，并且不支持跨机房空间
	var srcBucket = d.Bucket
	var srcKey = d.Name
	var destBucket = srcBucket
	var destKey = m.FileName

	bucketManager := util.GetBucketManager()

	putErr := bucketManager.Move(srcBucket, srcKey, destBucket, destKey, false)
	if putErr != nil {
		fmt.Println(putErr)
		SendResponse(c, errno.ErrQiniuCloud, nil)
		return
	}

	/* ============= 更新数据库中的文件信息 ============= */
	f := model.FileModel{
		Name: destKey,
		Id:   uint64(id),
	}

	if err := f.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)

}
