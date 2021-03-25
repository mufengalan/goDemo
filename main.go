package main

import (
	"github.com/gin-gonic/gin"
	"head_server/Config"
	"head_server/routers"
)

func main() {
	r := gin.Default()

	Config.Load(".env") //加载env配置

	routers.Load(r) //加载路由
	r.Run()
}
