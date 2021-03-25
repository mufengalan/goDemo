package controllers

import (
	"github.com/gin-gonic/gin"
	"head_server/app/model"
	"net/http"
)

func DoctorList(c *gin.Context) {

	type Request struct {
		ID int `json:"id"`
	}
	var request Request
	c.BindJSON(&request)

	id := request.ID
	name := model.GetDoctorList(id)
	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
	return
}

func ShopList(c *gin.Context) {
	type Request struct {
		AppId          string `json:"app_id"`
		ShopName       string `json:"shop_name"`
		ShopStatus     int    `json:"shop_status"`
		ShopAuthStatus int    `json:"shop_auth_status"`
		PageNum        int    `json:"page_num"`
		PageSize       int    `json:"page_size"`
	}
	var request Request
	c.BindJSON(&request)
	return
}
