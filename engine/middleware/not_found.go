package middleware

import (
	"github.com/chris1678/go-run/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoFound(c *gin.Context) {

	// 实现内部重定向
	c.AbortWithStatusJSON(http.StatusOK, response.NotFound())

}
