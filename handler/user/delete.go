package user

import (
	. "restful-api-server/handler"
	"restful-api-server/model"
	"restful-api-server/pkg/errno"
	"strconv"

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
