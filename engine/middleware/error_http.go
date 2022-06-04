package middleware

import (
	"github.com/chris1678/go-run/config"
	"github.com/chris1678/go-run/logger"
	"github.com/chris1678/go-run/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

// ErrorHttp  统一500错误处理函数
func ErrorHttp(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logger.LogHelper.Error(r)
			d := config.ApplicationConfig
			if d.Mode != "prod" {
				debug.PrintStack()
			}
			c.AbortWithStatusJSON(http.StatusOK, response.ServerError())
		}
	}()
	c.Next()
}
