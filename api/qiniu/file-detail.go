package qiniu

import (
	"github.com/gin-gonic/gin"
	. "github.com/jweboy/api-server/api"
	"github.com/jweboy/api-server/model"
	"github.com/jweboy/api-server/pkg/errno"
)

// DetailQuery 文件详情请求体
type DetailQuery struct {
	ID int `form:"id" binding:"required"`
}

// FileDetail 获取文件详情
func FileDetail(c *gin.Context) {
	var query DetailQuery

	// 检查query
	if c.ShouldBindQuery(&query) != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	detail, err := model.FileDetail(query.ID)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// TODO: 需要对返回的数据中的Name字段做转义
	SendResponse(c, nil, detail)
}
