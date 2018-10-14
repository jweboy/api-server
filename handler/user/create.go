package user

import (
	. "github.com/jweboy/restful-api-server/handler"
	"github.com/jweboy/restful-api-server/model"
	"github.com/jweboy/restful-api-server/pkg/errno"
	"github.com/jweboy/restful-api-server/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Create 新建用户
// @Summary 创建用户
// @Description 新增用户入库
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "用户名1-32个字符,密码4-128个字符，都必填"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"ok","data":{"username":"Jack"}}"
// @Router /user/{username} [post]
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
