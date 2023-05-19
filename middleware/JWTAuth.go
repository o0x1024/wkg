package middleware

import (
	"net/http"
	Gjwt "wkg/pkg/libs/jwt"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			// zap.S().Errorf("请求未携带token，无权限访问")
			c.JSON(http.StatusForbidden, gin.H{
				"code": 400,
				"msg":  "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		//Glog.Print("get token: ", token)

		if !Gjwt.CheckToken(token) {
			c.JSON(http.StatusForbidden, gin.H{
				"status": -1,
				"msg":    "授权已过期",
			})
			c.Abort()
			return

		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		//c.Set("claims", claims)
	}
}
