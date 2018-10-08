package user

import (
	. "restful-api-server/handler"
	"restful-api-server/pkg/errno"
	"restful-api-server/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{
		"X-Request-Id": util.GetReqID(c),
	})

	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
}
