package user_api

import (
	"BlogServer/common/res"
	"BlogServer/global"
	"BlogServer/models"
	"github.com/gin-gonic/gin"
)

type UserBaseInfoResponse struct {
	UserID       uint   `json:"userId"`
	CodeAge      int    `json:"codeAge"`
	Avatar       string `json:"avatar"`
	Nickname     string `json:"nickname"`
	LookCount    int    `json:"lookCount"`
	ArticleCount int    `json:"articleCount"`
	FansCount    int    `json:"fansCount"`
	FollowCount  int    `json:"followCount"`
	Place        string `json:"place"` // ip 归属地
}

func (UserApi) UserBaseInfoView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, cr.ID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}

	data := UserBaseInfoResponse{
		UserID:       user.ID,
		CodeAge:      user.CodeAge(),
		Avatar:       user.Avatar,
		Nickname:     user.Nickname,
		LookCount:    1,
		ArticleCount: 1, // TODO:把文章做完回来做
		FansCount:    1, // TODO:把好友关系做完回来做
		FollowCount:  1,
		Place:        user.Addr,
	}

	res.OkWithData(data, c)
}
