package routers

import (
	"github.com/gin-gonic/gin"
	"head_server/app/http/controllers"
)

func Load(g *gin.Engine) {
	/****课程管理****/
	prefix := g.Group("/course_manage")
	{
		prefix.POST("/doctor_list", controllers.DoctorList) //医生列表
		prefix.POST("/shop_list", controllers.ShopList)     //店铺列表

	}
}
