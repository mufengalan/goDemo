package routers

import (
	"github.com/gin-gonic/gin"
	"head_server/app/http/controllers"
)

func Load(g *gin.Engine) {
	/****课程管理****/
	prefix := g.Group("/course_manage")
	{
		prefix.POST("/shop_list", controllers.DoctorList)
	}
}
