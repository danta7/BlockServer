package middleware

import (
	"BlogServer/service/log_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseWriter 增强Gin的ResponseWriter,增加Body和Head字段，用于捕获响应体和响应头
type ResponseWriter struct {
	gin.ResponseWriter
	Body []byte      // 存储响应体
	Head http.Header // 存储响应头
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	w.Body = append(w.Body, data...)    // 将响应体数据追加到Body中
	return w.ResponseWriter.Write(data) // 再调用 gin.ResponseWriter 的Write方法
}

func (w *ResponseWriter) Header() http.Header {
	return w.Head // 返回调用的响应头
}

func LogMiddleWare(c *gin.Context) {

	log := log_service.NewActionLogByGin(c)
	// 请求中间件
	log.SetRequest(c)
	// 把这份请求中间件设置好的log对象传递给gin.context,通过gin.context获取日志对象，保证是一个日志流
	c.Set("log", log)

	res := &ResponseWriter{
		ResponseWriter: c.Writer,
		Head:           make(http.Header),
	}

	c.Writer = res // 将gin默认的c.writer替换为自己创建的res，所有通过 c.Writer 写入的响应数据都将首先进入到自定义的 ResponseWriter 中
	c.Next()
	// 响应中间件
	log.SetResponse(res.Body)
	log.SetResponseHeader(res.Head)
	log.MiddleSave()
}
