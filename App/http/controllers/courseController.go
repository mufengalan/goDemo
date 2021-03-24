package controllers

import (
	"github.com/gin-gonic/gin"
	"head_server/app/model"
	"net/http"
)

type Request struct {
	AppId          string `json:"app_id"`
	ShopName       string `json:"shop_name"`
	ShopStatus     int    `json:"shop_status"`
	ShopAuthStatus int    `json:"shop_auth_status"`
	PageNum        int    `json:"page_num"`
	PageSize       int    `json:"page_size"`
}

func ShopList(c *gin.Context) {
	var request Request
	c.BindJSON(&request)

	appId := request.AppId
	shopName := request.ShopName
	shopStatus := request.ShopName
	shopAuthStatus := request.ShopAuthStatus
	pageNum := request.PageNum
	pageSize := request.PageSize

	headAppId := model.GetShopList(appId)
	c.JSON(http.StatusOK, gin.H{
		"app_id":           appId,
		"shop_name":        shopName,
		"shop_status":      shopStatus,
		"shop_auth_status": shopAuthStatus,
		"page_num":         pageNum,
		"page_size":        pageSize,
		"head_app_id":      headAppId,
	})
	return
}
