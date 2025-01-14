package user_api

import (
	"BlogServer/common"
	"BlogServer/common/res"
	"BlogServer/global"
	"BlogServer/models"
	"BlogServer/utlis/jwts"
	"github.com/gin-gonic/gin"
	"time"
)

type UserLoginListRequest struct {
	common.PageInfo
	UserID    uint   `form:"userID"`
	IP        string `form:"ip"`
	Addr      string `form:"addr"`
	StartTime string `form:"startTime"` // 起始时间的 年 月 日 时 分 秒格式
	EndTime   string `form:"endTime"`
	Type      int8   `form:"type" binding:"required,oneof=1 2"` // 1 用户调用只能查自己  2 管理员 能查全部
}

type UserLoginListResponse struct {
	models.UserLoginModel
	UserNickname string `json:"userNickname,,omitempty"`
	UserAvatar   string `json:"userAvatar,,omitempty"`
}

func (UserApi) UserLoginListView(c *gin.Context) {

	var cr UserLoginListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	claims := jwts.GetClaims(c)
	if cr.Type == 1 {
		cr.UserID = claims.UserID
	}

	var query = global.DB.Where("")
	if cr.StartTime != "" {
		_, err = time.Parse("2006-01-02 15:04:05", cr.StartTime)
		if err != nil {
			res.FailWithMsg("开始时间格式错误", c)
			return
		}
		query.Where("created_at >= ?", cr.StartTime)
	}

	if cr.EndTime != "" {
		_, err = time.Parse("2006-01-02 15:04:05", cr.EndTime)
		if err != nil {
			res.FailWithMsg("结束时间格式错误", c)
			return
		}
		query.Where("created_at <= ?", cr.EndTime)
	}

	var preloads []string
	if cr.Type == 2 {
		preloads = []string{"UserModel"}
	}

	_list, count, _ := common.ListQuery(models.UserLoginModel{
		UserID: cr.UserID,
		IP:     cr.IP,
		Addr:   cr.Addr,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Where:    query,
		Preloads: preloads,
	})

	var list = make([]UserLoginListResponse, 0)
	for _, model := range _list {
		list = append(list, UserLoginListResponse{
			UserLoginModel: model,
			UserNickname:   model.UserModel.Nickname,
			UserAvatar:     model.UserModel.Avatar,
		})
	}

	res.OkWithList(list, count, c)
}
