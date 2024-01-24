## make shop_web 启动 商城web服务
## make shop_api 启动 商城api服务
## make run      启动 casdoor 服务单点登录服务
shop_web:
	yarn  --cwd ./app/shop/web start
	
shop_api:
	go run ./app/shop/api/main.go

run:down
	docker-compose up -d

down:
	docker-compose down


	