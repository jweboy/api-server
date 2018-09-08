package router

import (
	"net/http"
	"restful-api-server/handler/sd"
	"restful-api-server/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load 加载中间件、API、公用函数
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares

	// 处理请求可能因为程序bug或者其他异常情况导致程序挂了
	// 调用gin.Recovery()恢复 API 服务器从而不影响下一次请求的调用
	g.Use(gin.Recovery())

	// 强制浏览器不使用缓存
	g.Use(middleware.NoCache)

	// 浏览器跨域 OPTIONS 请求设置
	g.Use(middleware.Options)

	// 安全设置
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404 Handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route")
	})

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
