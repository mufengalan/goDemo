package main

import (
	"github.com/gin-gonic/gin"
	"head_server/database"
	"head_server/routers"
)

func main() {
	r := gin.Default()
	database.Load()
	routers.Load(r)
	r.Run()
}
