package router

import (
	"BlogServer/global"
	"BlogServer/middleware"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	nr := r.Group("/api")
	nr.Use(middleware.LogMiddleWare)
	SiteRouter(nr)

	addr := global.Config.System.Addr()
	r.Run(addr)

}
