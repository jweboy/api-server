package user

import (
	. "restful-api-server/handler"
	"restful-api-server/model"
	"restful-api-server/pkg/errno"
	"restful-api-server/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Create 新建用户
// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user/{username} [post]...
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{
		"X-Request-Id": util.GetReqID(c),
	})

	var r CreateRequest

	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// 校验请求体
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 密码加密
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// log.Info(r.Username)
	// log.Info(r.Password)

	// 插入数据
	if err := u.Create(); err != nil {
		log.Error("test databse:", err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// 返回请求
	rsp := CreateResponse{
		Username: r.Username,
	}

	SendResponse(c, nil, rsp)
}
