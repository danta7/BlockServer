package image_api

import (
	"BlogServer/common/res"
	"BlogServer/global"
	"BlogServer/service/qiniu_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type QiNiuTokenResponse struct {
	Token  string `json:"token"`
	Key    string `json:"key"`
	Region string `json:"region"`
	Url    string `json:"url"`
	Size   int    `json:"size"`
}

func (ImageApi) QiNiuGenToken(c *gin.Context) {
	q := global.Config.QiNiu
	if !q.Enable {
		res.FailWithMsg("未启用七牛云配置", c)
		return
	}

	token, err := qiniu_service.GenToken()
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	uid := uuid.New().String()
	key := fmt.Sprintf("%s/%s.png", q.Prefix, uid)
	url := fmt.Sprintf("%s/%s", q.Uri, key)

	res.OkWithData(QiNiuTokenResponse{
		Token:  token,
		Key:    key,
		Region: q.Region,
		Url:    url,
		Size:   q.Size,
	}, c)
}
