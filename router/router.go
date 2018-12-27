package router

import (
	"net/http"

	_ "github.com/jweboy/api-server/docs"
	// "github.com/jweboy/api-server/handler/user"
	"github.com/jweboy/api-server/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jweboy/api-server/api/qiniu"
	"github.com/jweboy/api-server/api/sd"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Load 加载中间件、API、公用函数
func Load() *gin.Engine {
	route := gin.New()

	// 注入logger
	route.Use(gin.Logger())

	// 处理请求可能因为程序bug或者其他异常情况导致程序挂了
	// 调用gin.Recovery()恢复 API 服务器从而不影响下一次请求的调用
	// 如果程序出现panic错误，会将请求定义为500错误返回
	route.Use(gin.Recovery())

	// 强制浏览器不使用缓存
	route.Use(middleware.NoCache)

	// 浏览器跨域 OPTIONS 请求设置
	route.Use(middleware.Options)

	// 安全设置
	route.Use(middleware.Secure)

	// 404 Handler
	route.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route")
	})

	// Swagger docs
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// The health check handlers
	svcd := route.Group("/api/v1/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	q := route.Group("/api/v1/qiniu")
	{
		q.GET("/file", qiniu.ListFile)
		q.POST("/file", qiniu.UploadFile)
		q.DELETE("/file", qiniu.DeleteFile)
		q.GET("/file/detail", qiniu.FileDetail)
		q.PUT("/file/edit", qiniu.EditDetail)
		q.PUT("/file/changeMime", qiniu.ChangeMime)
		q.GET("/bucket", qiniu.ListBucket)
	}

	// The user
	// u := g.Group("/v1/user")
	// {
	// 	u.POST("/:username", user.Create)
	// 	u.GET("/:username", user.GET)
	// 	u.DELETE("/:id", user.Delete)
	// 	u.PUT("/:id", user.Update)
	// }

	return route
}
