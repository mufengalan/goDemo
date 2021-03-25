package controllers

import (
	"github.com/gin-gonic/gin"
	"head_server/app/model"
	"net/http"
)

func DoctorList(c *gin.Context) {

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

	appId := request.AppId
	headAppId := model.GetDoctorList(appId)
	c.JSON(http.StatusOK, gin.H{
		"head_app_id":      headAppId,
		"shop_name":        request.ShopName,
		"shop_status":      request.ShopName,
		"shop_auth_status": request.ShopAuthStatus,
		"page_num":         request.PageNum,
		"page_size":        request.PageSize,
	})
	return
}
