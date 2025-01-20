package user_api

import (
	"BlogServer/common/res"
	"BlogServer/global"
	"BlogServer/utlis/jwts"
	"github.com/gin-gonic/gin"
)

func (UserApi) BindEmailView(c *gin.Context) {

	if !global.Config.Site.Login.EmailLogin {
		res.FailWithMsg("站点未启用邮箱注册", c)
		return
	}

	_email, _ := c.Get("email")
	email := _email.(string)

	user, err := jwts.GetClaims(c).GetUser()
	if err != nil {
		res.FailWithMsg("不存在的用户", c)
		return
	}

	// 更新
	global.DB.Model(&user).Update("email", email)
	res.OkWithMsg("邮箱绑定成功", c)

}
