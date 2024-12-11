package middleware

import (
	"bytes"
	"fmt"
	"io"
	"knowFood/utils/log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// ReqLogger 记录每次请求的请求信息和响应信息
func ReqLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前
		t := time.Now()
		// 获取请求路径
		reqPath := c.Request.URL.Path
		// 获取请求方法
		method := c.Request.Method
		// 获得客户端ip
		ip := c.ClientIP()
		// 获得原始查询参数
		rawQuery := c.Request.URL.RawQuery
		// JSON和FORM表单打印请求的Body, 其他内容类型，比如文件上传不打印
		var requestBody string
		contentType := c.GetHeader("Content-Type")
		/*
			关键步骤：使用 bytes.NewBuffer(requestBody) 创建一个新的 bytes.Buffer，它包含了读取出来的请求体数据。然后，使用 io.NopCloser 包装这个 bytes.Buffer。
			io.NopCloser 是一个实现了 io.Closer 接口但不做任何实际关闭操作的类型。这很重要，因为 http.Request.Body 需要是一个 io.ReadCloser 接口，而 bytes.Buffer 本身只实现了 io.Reader 接口。

			最后，将包装后的 bytes.Buffer 赋值给 c.Request.Body。这样，你就替换了原始的 Request.Body，使其可以被再次读取，而原始数据不会丢失。

			执行 c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody)) 这一步是为了确保请求体可以被多次读取，这是处理 Web 请求时常见的一种模式，特别是在使用中间件或需要记录原始请求数据的情况下。
		*/
		if contentType != "" &&
			(strings.HasPrefix(contentType, "application/json") ||
				strings.HasPrefix(contentType, "application/x-www-form-urlencoded")) {
			requestBody, err := io.ReadAll(c.Request.Body)
			if err != nil {
				requestBody = []byte{}
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		log.WithCtx(c).Info(fmt.Sprintf("host:%s %s %s start", ip, method, reqPath), "query", rawQuery, "body", requestBody)

		c.Next()
		// 请求后
		latency := time.Since(t).Microseconds()
		log.WithCtx(c).Info(fmt.Sprintf("host:%s %s %s end", ip, method, reqPath), "cost/us", latency)
	}
}
