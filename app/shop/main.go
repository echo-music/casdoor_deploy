package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-gonic/gin"
)

func init() {
	content, err := ioutil.ReadFile("./certs/token_jwt_key.pem")
	if err != nil {
		panic(err)
	}
	var certificate = string(content)

	casdoorsdk.InitConfig("http://localhost:8000", "72c162ebfb9a7f597be9", "cad87299a9f59861010dec3e91d99512b7b210a5", certificate, "organization_ep5s2s", "shop")
}

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

func main() {
	router := gin.Default()
	router.Use(Cors())
	router.GET("/api/signin", signinHandler)
	router.GET("/api/authorize", authorizeHandle)
	router.GET("/api/userinfo", VerifyToken(), userinfoHandler)
	router.GET("/api/goods", VerifyToken(), goodsHandle)

	log.Fatal(router.Run(":8080"))
}

// 根据code 获取令牌 token
func signinHandler(c *gin.Context) {
	// client := resty.New()

	// var req = make(map[string]string, 5)
	// req["grant_type"] = "authorization_code"
	// req["client_id"] = "72c162ebfb9a7f597be9"
	// req["client_secret"] = "cad87299a9f59861010dec3e91d99512b7b210a5"
	// req["code"] = c.Query("code")

	// resp, err := client.R().
	// 	SetHeader("Content-Type", "application/json").
	// 	SetBody(req).
	// 	Post("http://localhost:8000/api/login/oauth/access_token")
	// if err != nil {
	// 	panic("hello")
	// }
	// var res CallbackResp

	// json.Unmarshal([]byte(resp.String()), &res)
	token, err := casdoorsdk.GetOAuthToken(c.Query("code"), c.Query("state"))
	if err != nil {
		panic(err)
	}

	claims, err := casdoorsdk.ParseJwtToken(token.AccessToken)
	if err != nil {
		panic(err)
	}

	claims.AccessToken = token.AccessToken

	c.JSON(http.StatusOK, gin.H{
		"code":  1200,
		"token": claims.AccessToken,
	})
}

// 根据令牌 token 获取用户信息
func userinfoHandler(c *gin.Context) {
	user, ok := c.Get("user")
	fmt.Println(user, ok)

	c.JSON(http.StatusOK, gin.H{
		"code":  1200,
		"token": user,
	})
}
func authorizeHandle(c *gin.Context) {
	// 处理 GET 请求
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":       1401,
		"redict_url": "http://localhost:8000/login/oauth/authorize?client_id=72c162ebfb9a7f597be9&response_type=code&redirect_uri=http://localhost:3000/api/signin&scope=read&state=casdoor",
	})
}

func goodsHandle(c *gin.Context) {
	// 处理 GET 请求
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": 1200,
		"data": map[string]interface{}{
			"uuid":  "xsdcdsc-cdscds-xcdscdsc",
			"name":  "苹果",
			"prize": 18,
		},
	})
}
