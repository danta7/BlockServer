package router

import (
	"BlogServer/global"
	"BlogServer/middleware"
	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(global.Config.System.GinMode)
	r := gin.Default()

	// 配置静态路由
	r.Static("/uploads", "uploads")

	nr := r.Group("/api")
	nr.Use(middleware.LogMiddleWare)

	SiteRouter(nr)
	LogRouter(nr)

	addr := global.Config.System.Addr()
	err := r.Run(addr)
	if err != nil {
		return
	}

}
