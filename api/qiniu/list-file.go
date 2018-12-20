package qiniu

import (
	"github.com/gin-gonic/gin"
	. "github.com/jweboy/api-server/api"
	"github.com/jweboy/api-server/model"
	"github.com/jweboy/api-server/pkg/errno"
)

// ListResponse 重新组织的请求返回json
type ListResponse struct {
	Total uint64             `json:"total"`
	Data  []*model.FileModel `json:"data"`
}

// Pagination 请求体Query的类型等定义
type Pagination struct {
	Bucket string `form:"bucket" binding:"required"`
	Page   int    `form:"page" binding:"required"`
	Size   int    `form:"size" binding:"required"`
}

// ListFile 获取指定空间的文件列表
// @Summary 获取指定空间的文件列表
// @Description 获取指定存储空间的文件列表，带分页。
// @Tags qiniu
// @Accept  json
// @Produce  json
// @Param   bucket     query    string     true        "镜像空间名"
// @Param   page     query    int     true        "页码"
// @Param   size     query    int     true        "页数"
// @Router /qiniu/file [get]
// @Success 200 {object} handler.Response "{"code":0,"message":"ok", "data": []}"
func ListFile(c *gin.Context) {
	var pagination Pagination

	// 检查Query完整
	if c.BindQuery(&pagination) != nil {
		// TODO: 边界值处理 page size 都为0的时候 会有warning => Wanted to override status code 400 with 200
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// TODO: 这里如果需要在查找返回的数据中增加新字段的需求，需要增加锁处理，具体参考 https://github.com/lexkong/apiserver_demos/blob/master/demo07/service/service.go
	// sql查询具体分页数据
	files, count, err := model.ListFile(pagination.Bucket, pagination.Page, pagination.Size)
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
