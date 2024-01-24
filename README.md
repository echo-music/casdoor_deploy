# 一、背景

# 二、环境准备
1. 电脑：mac
2. 前端：node，yarn ，npm等
3. 后端：go

# 三、casdoor 部署

## 3.1、 拉取仓库
拉取仓库代码，进入master 分支
```
git clone git@github.com:echo-music/casdoor_deploy.git
```

## 3.2、启动服务
进入项目，在项目跟目录下，执行如下命令：
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
## 四、应用接入casdoor

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




