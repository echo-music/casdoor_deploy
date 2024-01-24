package middlewares

import (
	"casdoor-deploy/app/shop/api/exception"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Catch() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		length := len(c.Errors)
		if length <= 0 {
			return
		}
		e := c.Errors[length-1]
		if e == nil {
			return
		}
		code := exception.Code(e)
		msg := e.Error()
		switch code {
		case 0, -1:
			code = exception.CodeServerError
			msg = exception.CodeText(exception.CodeServerError)
			// 打印堆栈信息x
			fmt.Printf("%+v", exception.RealErr(e))

		case exception.CodeServerError:
			msg = exception.CodeText(exception.CodeServerError)
		}

		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  msg,
		})

	}
}
