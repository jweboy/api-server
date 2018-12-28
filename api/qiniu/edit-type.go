package qiniu

import (
	"github.com/jweboy/api-server/model"
	"github.com/jweboy/api-server/pkg/errno"
	"github.com/jweboy/api-server/util"

	"github.com/gin-gonic/gin"
	. "github.com/jweboy/api-server/api"
)

// EditTypeModel 更新文件类型的请求body
type EditTypeModel struct {
	ID   int    `form:"id" binding:"required"`
	Type string `form:"type" binding:"required"`
}

// ChangeMime 更新文件类型
// @Summary 更新文件类型
// @Description 更新文件类型
// @Tags qiniu
// @Accept  json
// @Produce  json
// @Param type	query	string	 true	"文件类型"
// @Param id	query	int	 true	"文件id"
// @Router /qiniu/file/changeMime  [put]
// @Success 200
func ChangeMime(c *gin.Context) {
	var m EditTypeModel

	if err := c.ShouldBind(&m); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	/* ============= 查询数据库中的文件详情 ============= */
	var id = m.ID

	d, err := model.FileDetail(id)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	/* ============= 更改七牛云对应的文件类型 ============= */
	var bucket = d.Bucket
	var fileName = d.Name
	var fileType = m.Type

	bucketManager := util.GetBucketManager()

	repErr := bucketManager.ChangeMime(bucket, fileName, fileType)
	if repErr != nil {
		SendResponse(c, errno.ErrQiniuCloud, nil)
		return
	}

	// TODO: 新建文件时候的type解决之后需要再这里多加一步更新type的入库操作

	SendResponse(c, nil, nil)

}
