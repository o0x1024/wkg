package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func InputToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")

		if token == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"code": 400,
				"msg":  "认证信息为空",
			})
			c.Abort()
			return
		}
		if !strings.Contains(token, "41f330686cbf9e21d9f092c68d032b03e24c7284") {
			c.JSON(http.StatusForbidden, gin.H{
				"code": 400,
				"msg":  "认证信息错误",
			})
			c.Abort()
			return
		}
	}
}
