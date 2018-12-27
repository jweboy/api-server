package qiniu

import (
	"github.com/gin-gonic/gin"
	. "github.com/jweboy/api-server/api"
	"github.com/jweboy/api-server/model"
	"github.com/jweboy/api-server/pkg/errno"
	"github.com/jweboy/api-server/util"
)

// DetailQuery 文件详情请求体
type DetailQuery struct {
	ID int `form:"id" binding:"required"`
}

// FileDetail 获取文件详情
// @Summary 获取文件详情
// @Description 获取文件详情
// @Tags qiniu
// @Accept  json
// @Produce  json
// @Param id	query	int	 true	"文件id"
// @Router /qiniu/file/detail  [get]
// TODO: nil需要定义
// @Success 200 {object} nil
func FileDetail(c *gin.Context) {
	var query DetailQuery

	// 检查 query 中的 ID 字段是否存在
	if c.ShouldBindQuery(&query) != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// 查询文件详情
	d, err := model.FileDetail(query.ID)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, model.FileModel{
		Id:        d.Id,
		CreatedAt: d.CreatedAt,
		Name:      util.DecodeStr(d.Name), // 反序列化文件名
		Key:       d.Key,
		Bucket:    d.Bucket,
		Size:      d.Size,
	})
}
