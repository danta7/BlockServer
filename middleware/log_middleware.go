package middleware

import (
	"BlogServer/service/log_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseWriter 响应体读取 自己定义方法，增强write方法，通过中间变量存响应体
type ResponseWriter struct {
	gin.ResponseWriter
	Body []byte
	Head http.Header
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	w.Body = append(w.Body, data...)
	return w.ResponseWriter.Write(data)
}

func (w *ResponseWriter) Header() http.Header {
	return w.Head
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
	c.Writer = res
	c.Next()
	// 响应中间件
	log.SetResponse(res.Body)
	log.SetResponseHeader(res.Head)
	log.MiddleSave()
}
