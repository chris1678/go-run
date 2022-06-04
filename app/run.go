package app

import (
	"context"
	"fmt"
	"github.com/chris1678/go-run/api"
	"github.com/chris1678/go-run/cache"
	"github.com/chris1678/go-run/cert"
	"github.com/chris1678/go-run/config"
	"github.com/chris1678/go-run/crons"
	"github.com/chris1678/go-run/logger"
	"github.com/chris1678/go-run/orm"
	"github.com/chris1678/go-run/utils"
	"github.com/chris1678/go-run/utils/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//初始化所有
func Initialize(conf string) {
	//初始化配置和
	config.Initialize(conf)
	//日志初始化
	logger.Initialize()
	//初始化数据库
	orm.Initialize()
	//redis缓存初始化
	cache.Initialize()
	//定时任务初始化
	crons.Initialize()
	//验证码初始化
	captcha.Initialize()
	//	队列初始化
	//验证器中文注册
	api.TransInit()
	//token证书初始化
	cert.Initialize()
}
func Run(handles *gin.Engine) *http.Server {

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d",
			config.ApplicationConfig.Port,
		),
		ReadTimeout:    time.Duration(config.ApplicationConfig.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.ApplicationConfig.WriterTimeout) * time.Second,
		MaxHeaderBytes: config.ApplicationConfig.MaxHeader << 20, //1M
		Handler:        handles,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.LogHelper.Fatal("listen: ", err)
			os.Exit(1)
		}
	}()

	logger.LogHelper.Infof("-  Local:   http://localhost:%d/", config.ApplicationConfig.Port)

	return srv

}
func Close(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.LogHelper.Infof("%s Shutdown Server ...", time.Now().Format(utils.TimeFormat))

	if err := srv.Shutdown(ctx); err != nil {
		logger.LogHelper.Fatal("Server Shutdown:", err)
	}
	logger.LogHelper.Info("Server exiting")
}
