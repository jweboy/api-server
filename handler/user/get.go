package user

import (
	. "api-server/handler"
	"api-server/model"
	"api-server/pkg/errno"

	"github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {
	username := c.Param("username")

	user, err := model.GetUser(username)

	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
