package user_api

import (
	"BlogServer/common/res"
	"BlogServer/global"
	"BlogServer/models"
	"BlogServer/models/enum"
	"BlogServer/service/email_service"
	"BlogServer/utlis/email_store"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type SendEmailRequest struct {
	Type  int8   `json:"type" binding:"oneof=1 2"` // 1 注册 2 重置密码
	Email string `json:"email" binding:"required"`
}

type SendEmailResponse struct {
	EmailID string `json:"emailID"`
}

func (UserApi) SendEmailView(c *gin.Context) {
	var cr SendEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	if !global.Config.Site.Login.EmailLogin {
		res.FailWithMsg("站点未启用邮箱注册", c)
		return
	}
	code := base64Captcha.RandText(4, "0123456789") // 验证码文本
	id := base64Captcha.RandomId()                  // 验证码的唯一标识符
	switch cr.Type {
	case 1:
		// 查邮箱是否不存在
		var user models.UserModel
		err = global.DB.Take(&user, "email = ?", cr.Email).Error
		if err == nil {
			res.FailWithMsg("该邮箱已被使用", c)
			return
		}
		err = email_service.SendRegisterCode(cr.Email, code)
	case 2:
		var user models.UserModel
		err = global.DB.Take(&user, "email = ?", cr.Email).Error
		if err != nil {
			res.FailWithMsg("邮箱不存在", c)
			return
		}
		// 必须得是邮箱注册的
		if user.RegisterSource != enum.RegisterEmailSourceType {
			res.FailWithMsg("非邮箱注册用户不能重置密码", c)
			return
		}
		err = email_service.SendResetPwdCode(cr.Email, code)
	}
	if err != nil {
		logrus.Errorf("邮件发送失败：%s", err)
		res.FailWithMsg("邮件发送失败", c)
		return
	}

	global.EmailVerifyStore.Store(id, email_store.EmailStoreInfo{
		Email: cr.Email,
		Code:  code,
	})
	email_store.Set(id, cr.Email, code)
	res.OkWithData(SendEmailResponse{
		EmailID: id,
	}, c)
}
