package router

import (
	"BlogServer/api"
	"BlogServer/middleware"
	"github.com/gin-gonic/gin"
)

func LogRouter(r *gin.RouterGroup) {
	app := api.App.LogApi
	r.GET("logs", middleware.AdminMiddleware, app.LogListView)
	r.GET("logs/:id", middleware.AdminMiddleware, app.LogReadView)
	r.DELETE("logs", middleware.AdminMiddleware, app.LogRemoveView)
}
