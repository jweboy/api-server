package user

import (
	"strconv"

	. "github.com/jweboy/api-server/handler"
	"github.com/jweboy/api-server/model"
	"github.com/jweboy/api-server/pkg/errno"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	// log.Infof("userId: %d", userId)

	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
