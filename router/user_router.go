package router

import (
	"BlogServer/api"
	"BlogServer/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	app := api.App.UserApi
	r.POST("user/send_email", middleware.CaptchaMiddleware, app.SendEmailView)
	r.POST("user/email", middleware.EmailVerifyMiddle, app.RegisterEmailView)
	r.POST("user/login", middleware.CaptchaMiddleware, app.PwdLoginApi)
}
