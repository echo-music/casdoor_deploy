package main

import (
	"casdoor-deploy/app/shop/api/exception"
	"casdoor-deploy/app/shop/api/middlewares"
	"casdoor-deploy/app/shop/api/types"
	"fmt"
	"log"
	"net/http"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func init() {
	//app config 初始化
	types.Init()
	// casdoor sdk confg 初始化
	casdoorsdk.InitConfig(
		types.Cfg.Casdoor.Endpoint,
		types.Cfg.Casdoor.ClientId,
		types.Cfg.Casdoor.ClientSecret,
		types.Cfg.Casdoor.Certificate,
		types.Cfg.Casdoor.OrganizationName,
		types.Cfg.Casdoor.ApplicationName,
	)
}

func main() {
	router := gin.Default()
	router.Use(middlewares.Cors(), middlewares.Catch())
	router.GET("/api/signin", signinHandler)
	router.GET("/api/authorize", authorizeHandle)
	router.GET("/api/userinfo", middlewares.VerifyToken(), userinfoHandler)
	router.GET("/api/goods", middlewares.VerifyToken(), goodsHandle)

	log.Fatal(router.Run(":8080"))
}

// 根据code 获取令牌 token
func signinHandler(c *gin.Context) {

	client := resty.New()
	var req = make(map[string]string, 5)
	req["grant_type"] = types.Cfg.Casdoor.GrantType
	req["client_id"] = types.Cfg.Casdoor.ClientId
	req["client_secret"] = types.Cfg.Casdoor.ClientSecret
	req["code"] = c.Query("code")

	var result = types.SigninResp{}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		SetResult(&result).
		Post("http://localhost:8000/api/login/oauth/access_token")
	if err != nil {
		c.Error(exception.WithStack(err))
		return
	}
	if !resp.IsSuccess() {
		c.Error(exception.New("获取令牌失败"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   exception.CodeOK,
		"result": result,
	})
}

// 根据令牌 token 获取用户信息
func userinfoHandler(c *gin.Context) {
	user, ok := c.Get("user")
	fmt.Println(user, ok)

	c.JSON(http.StatusOK, gin.H{
		"code":  exception.CodeOK,
		"token": user,
	})
}
func authorizeHandle(c *gin.Context) {
	// 处理 GET 请求
	c.JSON(http.StatusOK, gin.H{
		"code":       exception.CodeOK,
		"signin_url": "http://localhost:8000/login/oauth/authorize?client_id=72c162ebfb9a7f597be9&response_type=code&redirect_uri=http://localhost:3000/api/signin&scope=read&state=casdoor",
	})
}

func goodsHandle(c *gin.Context) {
	// 处理 GET 请求
	c.JSON(http.StatusOK, gin.H{
		"code": exception.CodeOK,
		"data": map[string]interface{}{
			"uuid":  "xsdcdsc-cdscds-xcdscdsc",
			"name":  "苹果",
			"prize": 18,
		},
	})
}
