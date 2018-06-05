package api

import (
	"mockapi/config"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	registerRouters(router)
}

func registerRouters(router *gin.Engine) {
	mockApiRouterRegister(router)
	router.Use(cors.Default())
}

func Run() {
	port := config.App.Api.Port
	address := config.App.Api.Address
	debug := config.App.Api.Debug
	if debug {
		gin.SetMode(gin.DebugMode)
	}else{
		gin.SetMode(gin.ReleaseMode)
	}
	router.Run(fmt.Sprintf("%s:%s", address, port))
}
