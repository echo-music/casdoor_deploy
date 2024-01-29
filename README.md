---
layout: post
title: casdoor 部署指南
categories: go
description: casdoor 部署指南
keywords: go
---

casdoor 部署指南，快速接入单点登录
# 一、背景
公司有个商城 (shop)需要接入到单点登录，在github 上找来找去发现 casdoor 开源项目满足我的需求：
1. 商城后端不需要记录会话信息
2. 用户只需要到认证中心登录一次，便可访问多个公司系统，无需重复登录

官方的 docker 部署文档不是很详细。所以自己摸索操作了部署，接入的流程，输出了如下文档，供大家参考。
# 二、环境准备
1. 电脑：mac （linux 系统都可）
2. 前端：node，yarn ，npm等 (我的 node是 v16.10.0)
3. 后端：go （我的是 go1.21.3）

# 三、casdoor 部署

## 3.1、 拉取仓库
为了方便部署和学习，我把部署的casdoor 的docker-compose.yaml 文件和商城代码都放到了了该仓库下，方便部署，部署目录如下：
[![pFuDUYT.png](https://s11.ax1x.com/2024/01/29/pFuDUYT.png)](https://imgse.com/i/pFuDUYT)

拉取仓库代码，进入master 分支
```
git clone git@github.com:echo-music/casdoor_deploy.git
```

## 3.2、启动服务
进入项目，在项目根目录下，执行如下命令：
```
make run
```
[![pFeIxKK.png](https://s11.ax1x.com/2024/01/24/pFeIxKK.png)](https://imgse.com/i/pFeIxKK)


## 3.3、验证服务是否启动
访问[http://localhost:8000](http://localhost:8000) 网址
出现如下登录页面，说明服务起来了
[![pFeo8x0.png](https://s11.ax1x.com/2024/01/24/pFeo8x0.png)](https://imgse.com/i/pFeo8x0)

默认账号: admin/123 登陆进入管理页面
[![pFeoDR1.png](https://s11.ax1x.com/2024/01/24/pFeoDR1.png)](https://imgse.com/i/pFeoDR1)

如果服务出现错误，请运行以下命令查看错误原因
```
docker logs casdoor-casdoor-1 -f
```

## 3.4、创建组织
进入 [http://localhost:8000/organizations](http://localhost:8000/organizations) 页创建组织：

[![pFe7aC9.png](https://s11.ax1x.com/2024/01/24/pFe7aC9.png)](https://imgse.com/i/pFe7aC9)

## 3.5、创建证书
进入 [http://localhost:8000/certs](http://localhost:8000/certs) 页 创建证书

[![pFe7xK0.png](https://s11.ax1x.com/2024/01/24/pFe7xK0.png)](https://imgse.com/i/pFe7xK0)

## 3.6、创建应用
进入[http://localhost:8000/applications](http://localhost:8000/applications)页 创建应用
[![pFeHEx1.png](https://s11.ax1x.com/2024/01/24/pFeHEx1.png)](https://imgse.com/i/pFeHEx1)

[![pFeHlPH.png](https://s11.ax1x.com/2024/01/24/pFeHlPH.png)](https://imgse.com/i/pFeHlPH)

[![pFeHdIg.png](https://s11.ax1x.com/2024/01/24/pFeHdIg.png)](https://imgse.com/i/pFeHdIg)

[![pFebSSA.png](https://s11.ax1x.com/2024/01/24/pFebSSA.png)](https://imgse.com/i/pFebSSA)

然后保持退出
# 四、应用接入casdoor

## 4.1、后端服务配置
进入 app/shop/api/conf 目录下，将前面画红框的信息写入到 conf.toml 这个文件，内容如下：
```
Title = "商城"

[Casdoor]
Endpoint = "http://localhost:8000"
GrantType = "authorization_code"   
ClientId = "27a7d6599cbd78f0e666"  
ClientSecret = "b9abe2f32b74f4268e57061334a639154b08a4b2" 
Certificate = "./certs/token_jwt_key.pem"
OrganizationName = "echo_music" 
ApplicationName = "shop"
RedirectUri = "http://localhost:3000/signin"

```

将前面下载的证书放到 app/shop/api/ 这个目录

[![pFebBTO.png](https://s11.ax1x.com/2024/01/24/pFebBTO.png)](https://imgse.com/i/pFebBTO)

到此，后端服务配置完成。

## 4.2、前端应用配置
打开 app/shop/web/src/package/Conf/下的配置文件 index.js 配置如下内容
该地址为shop应用的地址
```
export const API_URL = 'http://localhost:8080';

```

到此，前后端配置都完成了

## 4.3、启动 shop应用前后端
分别在两个终端执行 如下命令启动shop前后端服务
```
make shop_api

make shop_web
```

启动 shop 后端服务
[![pFevu6J.png](https://s11.ax1x.com/2024/01/24/pFevu6J.png)](https://imgse.com/i/pFevu6J)


启动 shop 前端服务
[![pFev8k6.png](https://s11.ax1x.com/2024/01/24/pFev8k6.png)](https://imgse.com/i/pFev8k6)

## 4.4、测试单点登录
访问shop 前端[http://localhost:3000/](http://localhost:3000/)

[![pFevtpD.png](https://s11.ax1x.com/2024/01/24/pFevtpD.png)](https://imgse.com/i/pFevtpD)

然后点击登录，进入 casdoor 授权登陆页：

[![pFevU6H.png](https://s11.ax1x.com/2024/01/24/pFevU6H.png)](https://imgse.com/i/pFevU6H)

输入账号，密码进行登录

[![pFevwnA.png](https://s11.ax1x.com/2024/01/24/pFevwnA.png)](https://imgse.com/i/pFevwnA)


此时会跳转到 shop 首页

[![pFevstf.png](https://s11.ax1x.com/2024/01/24/pFevstf.png)](https://imgse.com/i/pFevstf)


## 五、参考资料
[https://casdoor.org/zh/docs/how-to-connect/sdk](https://casdoor.org/zh/docs/how-to-connect/sdk)
