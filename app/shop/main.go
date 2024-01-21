package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")                                                     // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization") //你想放行的header也可以在后面自行添加
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")                                   //我自己只使用 get post 所以只放行它
		// c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func main() {
	router := gin.Default()
	router.Use(Cors())

	router.GET("/callback", callbackHandle)
	router.GET("/order/list", listHandle)

	log.Fatal(router.Run(":8080"))
}

type CallbackResp struct {
	AccesToken   string `json:"access_token"`
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

func callbackHandle(c *gin.Context) {
	client := resty.New()

	var req = make(map[string]string, 5)
	req["grant_type"] = "authorization_code"
	req["client_id"] = "72c162ebfb9a7f597be9"
	req["client_secret"] = "cad87299a9f59861010dec3e91d99512b7b210a5"
	req["code"] = c.Query("code")

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		Post("http://localhost:8000/api/login/oauth/access_token")
	if err != nil {
		panic("hello")
	}
	var res CallbackResp
	fmt.Println(resp.String())

	json.Unmarshal([]byte(resp.String()), &res)

	// 处理 GET 请求
	c.JSON(http.StatusOK, gin.H{
		"code":  c.Query("code"),
		"state": c.Query("state"),
		"resp":  res,
	})
}
func listHandle(c *gin.Context) {
	// 处理 GET 请求
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":       401,
		"redict_url": "http://localhost:8000/login/oauth/authorize?client_id=72c162ebfb9a7f597be9&response_type=code&redirect_uri=http://localhost:3000/callback&scope=read&state=casdoor",
	})
}
