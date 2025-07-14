package middleware

import (
	"fmt"
	"gin-app-start/app/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IPAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := Context{Ctx: c}
		ipList := config.Conf.Server.IpWhitelist
		flag := false
		clientIp := c.ClientIP()
		for _, value := range ipList {
			if clientIp == value {
				flag = true
				break
			}
		}
		if !flag {
			ctx.Response(http.StatusUnauthorized, fmt.Sprintf("%s not in ipList", clientIp), nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
