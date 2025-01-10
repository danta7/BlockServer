package middleware

import (
	"BlogServer/common/res"
	"BlogServer/global"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
)

type CaptchaMiddlewareRequest struct {
	CaptchaID   string `json:"captchaId" binding:"required"`
	CaptchaCode string `json:"captchaCode" binding:"required"`
}

func CaptchaMiddleware(c *gin.Context) {
	if !global.Config.Site.Login.Captcha {
		return
	}
	// 在绑定之前拿一下请求体
	body, err := c.GetRawData()
	if err != nil {
		res.FailWithMsg("获取请求体错误", c)
		c.Abort()
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	var cr CaptchaMiddlewareRequest
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		logrus.Errorf("图形验证失败 %s", err)
		res.FailWithMsg("图形验证失败", c)
		c.Abort()
		return
	}
	if !global.CaptchaStore.Verify(cr.CaptchaID, cr.CaptchaCode, true) {
		res.FailWithMsg("验证码错误", c)
		c.Abort()
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
}
