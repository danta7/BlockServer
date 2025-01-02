package router

import (
	"BlogServer/api"
	"BlogServer/middleware"
	"github.com/gin-gonic/gin"
)

func ImageRouter(r *gin.RouterGroup) {
	app := api.App.ImageApi
	r.POST("images", middleware.AuthMiddleware, app.ImageUploadView)
	r.POST("images/qiniu", middleware.AuthMiddleware, app.QiNiuGenToken)
	r.GET("images", middleware.AuthMiddleware, app.ImageListView)
	r.DELETE("images", middleware.AdminMiddleware, app.ImageRemoveView)
}
