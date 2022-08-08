package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-simple/pkg/logger"
	"go-simple/pkg/response"
	"go.uber.org/zap"
	"net"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 获取用户请求信息
				httpRequest, _ := httputil.DumpRequest(c.Request, true)
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						errStr := strings.ToLower(se.Error())
						if strings.Contains(errStr, "broken pipe") || strings.Contains(errStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				// 链接中断的情况
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
							zap.Time("time", time.Now()),
							zap.Any("error", err),
							zap.String("request", string(httpRequest)),
						)
					c.Error(err.(error))
					c.Abort()
					// 链接已断开
					return
				}
				logger.Error("recovery time panic",
					zap.Time("time", time.Now()),
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.Stack("stacktrace"),
				)

				response.Abort500(c)
			}
		}()
		c.Next()
	}
}