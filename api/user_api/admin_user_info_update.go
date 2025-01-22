package user_api

import (
	"BlogServer/common/res"
	"BlogServer/global"
	"BlogServer/models"
	"BlogServer/models/enum"
	"BlogServer/utlis/mps"
	"github.com/gin-gonic/gin"
)

type AdminUserInfoUpdateRequest struct {
	UserID   uint           `json:"userID" binding:"required"`
	Nickname *string        `json:"nickname" s-u:"nickname"`
	Avatar   *string        `json:"avatar" s-u:"avatar"` // 头像
	Abstract *string        `json:"abstract" s-u:"abstract"`
	Role     *enum.RoleType `json:"role" s-u:"role"`
}

func (UserApi) AdminUserInfoUpdateView(c *gin.Context) {
	var cr AdminUserInfoUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	userMap := mps.Struct2Map(cr, "s-u")
	var user models.UserModel

	err = global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	err = global.DB.Model(&user).Updates(userMap).Error
	if err != nil {
		res.FailWithMsg("用户信息修改失败", c)
		return
	}
	res.OkWithMsg("用户信息修改成功", c)

}
