package qiniu

import (
	"github.com/jweboy/api-server/service"
	"github.com/jweboy/api-server/util"

	"github.com/gin-gonic/gin"
	. "github.com/jweboy/api-server/api"
	"github.com/jweboy/api-server/model"
	"github.com/jweboy/api-server/pkg/errno"
)

// ListResponse 请求返回的json结构
type ListResponse struct {
	Total uint64             `json:"total"`
	Data  []*model.FileModel `json:"data"`
}

// Pagination 分页请求体
type Pagination struct {
	Bucket string `form:"bucket" binding:"required"` // 必填项
	Page   int    `form:"page"`
	Size   int    `form:"size"`
}

// ListFile 获取指定空间的文件列表
// @Summary 获取指定空间的文件列表
// @Description 获取指定存储空间的文件列表，带分页。
// @Tags qiniu
// @Accept  json
// @Produce  json
// @Param   bucket   query    string   true  "镜像空间名"
// @Param   page     query    int     false  "页码" default(1)
// @Param   size     query    int     false  "页数" default(10)
// @Router /qiniu/file [get]
// @Success 200 {array} model.FileModel
func ListFile(c *gin.Context) {
	var pagination Pagination

	// 默认第 1 页， 10 条 / 页
	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")

	// 检查 query 中的 bucket 字段是否存在
	if c.ShouldBindQuery(&pagination) != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// TODO: 这里如果需要在查找返回的数据中增加新字段的需求，需要增加锁处理
	// 具体参考 https://github.com/lexkong/apiserver_demos/blob/master/demo07/service/service.go
	// sql查询具体分页数据
	files, count, err := service.ListFile(
		pagination.Bucket,
		util.StrToInt(page),
		util.StrToInt(size),
	)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// 返回新组装的数据结构
	SendResponse(c, nil, ListResponse{
		Data:  files,
		Total: count,
	})
}
