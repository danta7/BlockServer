package site_api

import (
	"BlogServer/comment/res"
	"github.com/gin-gonic/gin"
)

type SiteApi struct {
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	res.OkWithData("xxx", c)
	return
}

type SiteUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (SiteApi) SiteUpdateView(c *gin.Context) {
	//log := log_service.GetLog(c)
	//log.ShowRequest()
	//log.ShowRequestHeader()
	//log.ShowResponseHeader()
	//log.ShowResponse()
	//log.SetTitle("更新站点")
	//log.SetItemInfo("请求时间", time.Now())
	//log.SetImage("/xxx/xxx")
	//log.SetLink("gin链接", "https://gin-gonic.com/zh-cn/")
	//c.Header("xxx", "xxxxe")

	var cr SiteUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	//log.SetItemInfo("结构体", cr)
	//log.SetItemInfo("切片", []string{"a", "b"})
	//log.SetItemInfo("字符串", "你好")
	//log.SetItemInfo("数字", 123)

	//id := log.Save()
	//fmt.Println(1, id)
	//id = log.Save()
	//fmt.Println(2, id)

	res.OkWithMsg("更新成功", c)
	return
}
