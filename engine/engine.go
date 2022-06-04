package engine

import (
	"github.com/chris1678/go-run/config"
	"github.com/chris1678/go-run/engine/middleware"
	"github.com/chris1678/go-run/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func Engine() *gin.Engine {

	if config.ApplicationConfig.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	var engine = gin.New()
	//注册验证

	if config.ApplicationConfig.Port == 0 {
		logger.LogHelper.Fatal("port is error")
		os.Exit(1)
	}
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("authorization")
	engine.Use(cors.New(corsConfig))
	//404错误
	engine.NoRoute(middleware.NoFound)
	//500错误
	engine.Use(middleware.ErrorHttp)
	//限流
	engine.Use(middleware.Sentinel())

	return engine

}
