package captcha_api

import (
	"BlogServer/common/res"
	"BlogServer/global"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
	"image/color"
)

type CaptchaApi struct {
}

type CaptchaResponse struct {
	CaptchaId string `json:"captchaID"`
	Captcha   string `json:"captcha"`
}

func (CaptchaApi) CaptchaView(c *gin.Context) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      1,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
	}

	driverString = captchaConfig
	//将driverString中指定的字体文件转换为驱动程序所需的字体格式，并将结果赋值给driver变量。这个步骤是为了将字体文件转换为正确的格式，以便在生成验证码时使用正确的字体。
	driver = driverString.ConvertFonts()
	//使用driver和stores参数创建一个新的验证码实例，并将其赋值给captcha变量。这里的stores参数表示验证码存储器，用于存储和验证验证码。
	captcha := base64Captcha.NewCaptcha(driver, global.CaptchaStore)
	//调用captcha实例的Generate方法生成验证码。lid是生成的验证码的唯一标识符，lb64s是生成的验证码图片的Base64编码字符串，lerr是生成过程中的任何错误。
	lid, lb64s, _, lerr := captcha.Generate()
	if lerr != nil {
		res.FailWithMsg("图片验证码生成失败", c)
		logrus.Error(lerr)
		return
	}
	res.OkWithData(CaptchaResponse{
		CaptchaId: lid,
		Captcha:   lb64s,
	}, c)
}
