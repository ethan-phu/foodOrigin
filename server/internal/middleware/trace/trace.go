package trace

import (
	"context"
	"knowFood/tools/uuid"
	"knowFood/utils/constant"
	"knowFood/utils/log"

	"github.com/gin-gonic/gin"
)

// SetRequestId 用来设置和透传requestId
func SetRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成流水号
		requestId := uuid.GenUUID16()
		// 设置到context里面去
		c.Header("X-Request-Id", requestId)

		// 设置requestId到context中，便于后面调用链的透传
		c.Set(constant.RequestId, requestId)
		c.Next()
	}
}

// RequestId 获取requestId，返回的是一个中间件
func RequestId() log.Valuer {
	return func(c context.Context) any {
		if rid := c.Value(constant.RequestId); rid != nil {
			return rid.(string)
		}
		return ""
	}
}
