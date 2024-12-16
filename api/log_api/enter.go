package log_api

import (
	"BlogServer/comment"
	"BlogServer/comment/res"
	"BlogServer/models"
	"BlogServer/models/enum"
	"github.com/gin-gonic/gin"
)

type LogApi struct{}

type LogListRequest struct {
	comment.PageInfo
	LogType     enum.LogType      `form:"logType"` // 日志类型 1 2 3
	Level       enum.LogLevelType `form:"level"`   // 日志级别 1 2 3
	UserID      uint              `form:"userID"`
	IP          string            `form:"ip"`
	LoginStatus bool              `form:"loginStatus"` // 登录的状态
	ServiceName string            `form:"serviceName"`
}

type LogListResponse struct {
	models.LogModel
	UserNickname string `json:"userNickname"`
	UserAvatar   string `json:"userAvatar"`
}

func (LogApi) LogListView(c *gin.Context) {
	// 分页 查询 (精确查询、模糊匹配)
	var cr LogListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	list, count, err := comment.ListQuery(models.LogModel{
		LogType:     cr.LogType,
		Level:       cr.Level,
		UserID:      cr.UserID,
		IP:          cr.IP,
		LoginStatus: cr.LoginStatus,
		ServiceName: cr.ServiceName,
	}, comment.Options{
		PageInfo:     cr.PageInfo,
		Likes:        []string{"title"},
		Preloads:     []string{"UserModel"},
		Debug:        true,
		DefaultOrder: "created_at desc",
	})

	var _list = make([]LogListResponse, 0)
	for _, logModel := range list {
		_list = append(_list, LogListResponse{
			LogModel:     logModel,
			UserNickname: logModel.UserModel.Nickname,
			UserAvatar:   logModel.UserModel.Avatar,
		})
	}

	res.OkWithList(_list, int(count), c)
	return
}
