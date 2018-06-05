package api

import (
	"github.com/gin-gonic/gin"
	"mockapi/config"
)

func mockApiRouterRegister(router *gin.Engine) {
	router.GET("/bp/info", bpInfoHandler)
	router.GET("/upgrade", upgradeHandler)
	router.GET("/network", networkHandler)
}

func bpInfoHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	info := config.App.BpInfo
	c.Data(200,"application/json; charset=utf-8", []byte(info))
}

func upgradeHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	info := config.App.Upgrade
	c.Data(200,"application/json; charset=utf-8", []byte(info))
}

func networkHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	info := config.App.Network
	c.Data(200,"application/json; charset=utf-8", []byte(info))
}

