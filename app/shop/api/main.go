package main

import (
	"casdoor-deploy/app/shop/api/exception"
	"casdoor-deploy/app/shop/api/middlewares"
	"casdoor-deploy/app/shop/api/types"
	"log"
	"net/http"
	"net/url"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	code := c.Query("code")
	state := c.Query("state")

	token, err := casdoorsdk.GetOAuthToken(code, state)
	if err != nil {
		c.Error(exception.WithStack(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": exception.CodeOK,
		"user": token,
	})
}

// 根据令牌 token 获取用户信息
func userinfoHandler(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"code":  exception.CodeOK,
		"token": user,
	})
}
func authorizeHandle(c *gin.Context) {
	baseURL := types.Cfg.Casdoor.Endpoint + "/login/oauth/authorize"
	params := url.Values{}
	params.Set("client_id", types.Cfg.Casdoor.ClientId)
	params.Set("response_type", "code")
	params.Set("redirect_uri", types.Cfg.Casdoor.RedirectUri)
	params.Set("scope", "read")
	params.Set("state", uuid.New().String())

	u, err := url.Parse(baseURL)
	if err != nil {
		c.Error(exception.New("授权地址解析失败"))
		return
	}

	u.RawQuery = params.Encode()
	// 处理 GET 请求
	c.JSON(http.StatusOK, gin.H{
		"code":       exception.CodeOK,
		"signin_url": u.String(),
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
