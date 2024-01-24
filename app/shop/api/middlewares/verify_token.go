package middlewares

import (
	"net/http"
	"strings"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code": 1400,
				"msg":  "authHeader is empty",
			})
			return
		}

		token := strings.Split(authHeader, "Bearer ")

		if len(token) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 1401,
				"msg":  "token is not valid Bearer token",
			})
			return
		}

		claims, err := casdoorsdk.ParseJwtToken(token[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": 1401,
				"msg":  "ParseJwtToken() error",
			})
			return
		}

		c.Set("user", claims.User)
		// 处理请求
		c.Next()
	}
}
