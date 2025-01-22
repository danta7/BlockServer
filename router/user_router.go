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
	r.GET("user/detail", middleware.AuthMiddleware, app.UserDetailView)
	r.GET("user/login", middleware.AuthMiddleware, app.UserLoginListView)
	r.GET("user/base", app.UserBaseInfoView)
	r.PUT("user/password", middleware.AuthMiddleware, app.UpdatePasswordView)
	r.PUT("user/password/reset", middleware.EmailVerifyMiddle, app.ResetPasswordView)
	r.PUT("user/email/bind", middleware.AuthMiddleware, middleware.EmailVerifyMiddle, app.BindEmailView)
	r.PUT("user", middleware.AuthMiddleware, app.UserInfoUpdateView)
	r.PUT("user/admin", middleware.AdminMiddleware, app.AdminUserInfoUpdateView)
}
