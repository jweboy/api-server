package user

import (
	"strconv"

	. "github.com/jweboy/restful-api-server/handler"
	"github.com/jweboy/restful-api-server/model"
	"github.com/jweboy/restful-api-server/pkg/errno"
	"github.com/jweboy/restful-api-server/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Update(c *gin.Context) {
	log.Info("Update function is called.", lager.Data{
		"X-Request-Id": util.GetReqID(c),
	})

	userId, _ := strconv.Atoi(c.Param("id"))
	log.Infof("userId:%d", userId)

	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u.Id = uint64(userId)

	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
