package qiniu

import (
	"github.com/gin-gonic/gin"
	. "github.com/jweboy/api-server/api"
	"github.com/jweboy/api-server/model"
	"github.com/jweboy/api-server/pkg/errno"
	"github.com/jweboy/api-server/util"
	log "qiniupkg.com/x/log.v7"
)

// DeleteQuery 删除文件 Query 请求参数
type DeleteQuery struct {
	ID int `form:"id" binding:"required"` // 必填项
}

// DeleteFile 删除指定空间的文件
// @Summary 删除指定空间的文件
// @Description 删除指定空间的文件
// @Tags qiniu
// @Accept  json
// @Produce  json
// @Param id	query	int	 true	"文件id"
// @Router /qiniu/file  [delete]
// @Success 200
func DeleteFile(c *gin.Context) {
	var deleteQuery DeleteQuery

	// 检查 query 中的 id 字段是否存在
	if c.ShouldBindQuery(&deleteQuery) != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	/* ============= 获取数据库中对应的文件信息 ============= */
	var id = deleteQuery.ID

	detail, err := model.FileDetail(id)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	/* ============= 删除七牛云指定文件 ============= */
	var bucket = detail.Bucket
	var fileName = detail.Name

	bucketManager := util.GetBucketManager()

	if err := bucketManager.Delete(bucket, fileName); err != nil {
		// log.Errorf("Qiniu SDK error:", err)
		SendResponse(c, errno.ErrFileDelete, nil)
		return
	}

	/* ============= 删除数据库对应文件 ============= */
	var fileID = uint64(id)
	if err := model.DeleteFile(fileID); err != nil {
		log.Error("Databse error：", err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
