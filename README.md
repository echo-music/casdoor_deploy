# 一、背景

# 二、环境准备
电脑：mac
前端：node，yarn ，npm等
后端：go


# 二、casdoor 部署

## 2.1、 拉取仓库
拉取仓库代码，进入master 分支
```
git clone git@github.com:echo-music/casdoor_deploy.git
```

## 2.2、启动服务
进入项目，在项目跟目录下，执行如下命令：
```
    make run
```
[![pFeIxKK.png](https://s11.ax1x.com/2024/01/24/pFeIxKK.png)](https://imgse.com/i/pFeIxKK)


## 2.3、验证服务是否启动
访问[http://localhost:8000](http://localhost:8000) 网址
出现如下登录页面，说明服务起来了
[![pFeo8x0.png](https://s11.ax1x.com/2024/01/24/pFeo8x0.png)](https://imgse.com/i/pFeo8x0)

默认账号: admin/123 登陆进入管理页面
[![pFeoDR1.png](https://s11.ax1x.com/2024/01/24/pFeoDR1.png)](https://imgse.com/i/pFeoDR1)

如果服务出现错误，请运行以下命令查看错误原因
```
    docker logs casdoor-casdoor-1 -f
```

## 2.4、创建组织

## 2.5、创建证书


## 2.6、创建应用


## 三、应用接入casdoor

## 3.1、后端服务配置
进入 app/shop/api/conf 目录下，
```
Title = "商城"

[Casdoor]
Endpoint = "http://localhost:8000" ## casdoor 服务地址
GrantType = "authorization_code"   ## GrantType oauth2授权方式
ClientId = "27a7d6599cbd78f0e666"  ## 客户端ID， 备注：在 casdoor管理端 创建应用生成的
ClientSecret = "b9abe2f32b74f4268e57061334a639154b08a4b2" ## 客户端密码， 备注：在 casdoor管理端创建应用生成的
Certificate = "./certs/token_jwt_key.pem" ## 证书 备注：在 casdoor管理端 创建证书生成的
OrganizationName = "echo_music" ## 组织名称
ApplicationName = "shop" ## 应用名称
RedirectUri = "http://localhost:3000/signin" ## 跳转到应用的客户端地址

```




