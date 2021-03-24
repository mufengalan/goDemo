package main

import (
	"github.com/gin-gonic/gin"
	"head_server/Config"
	"head_server/database"
	"head_server/routers"
)

func main() {
	r := gin.Default()
	//加载env配置
	Config.Load(".env")
	database.Load()
	routers.Load(r)
	r.Run()
}
