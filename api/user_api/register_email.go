package user_api

import (
	"BlogServer/common/res"
	"BlogServer/global"
	"BlogServer/models"
	"BlogServer/models/enum"
	"BlogServer/service/user_service"
	"BlogServer/utlis/jwts"
	"BlogServer/utlis/pwd"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type RegisterEmailRequest struct {
	EmailID   string `json:"emailID" binding:"required"`
	EmailCode string `json:"emailCode" binding:"required"`
	Pwd       string `json:"pwd" binding:"required"`
}

func (UserApi) RegisterEmailView(c *gin.Context) {
	var cr RegisterEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	if !global.Config.Site.Login.EmailLogin {
		res.FailWithMsg("站点未启用邮箱注册", c)
		return
	}

	// 创建用户
	uname := base64Captcha.RandText(5, "0123456789")

	_email, _ := c.Get("email")
	email := _email.(string)

	hashPwd, _ := pwd.GenerateFromPassword(cr.Pwd)
	var user = models.UserModel{
		Username:       fmt.Sprintf("b_%s", uname),
		Nickname:       "邮箱用户",
		Avatar:         "",
		RegisterSource: enum.RegisterEmailSourceType,
		Password:       hashPwd,
		Email:          email,
		Role:           enum.UserRole,
	}
	err = global.DB.Create(&user).Error
	if err != nil {
		res.FailWithMsg("邮箱注册失败", c)
		logrus.Errorf("创建用户失败%s", err)
		return
	}

	// 颁发token
	token, err := jwts.GetToken(jwts.Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
	})
	if err != nil {
		res.FailWithMsg("邮箱登录失败", c)
		return
	}

	// 记录登录日志
	user_service.NewUserService(user).UserLogin(c)
	res.OkWithData(token, c)
}
