package middleware

import (
	"github.com/alibaba/sentinel-golang/core/system"
	"github.com/alibaba/sentinel-golang/logging"
	sentinelPlugin "github.com/alibaba/sentinel-golang/pkg/adapters/gin"
	"github.com/chris1678/go-run/logger"
	"github.com/chris1678/go-run/response"
	"github.com/chris1678/go-run/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Sentinel 限流
func Sentinel() gin.HandlerFunc {
	//重置限流这个日志的级别，不然显示info出来
	logging.ResetGlobalLoggerLevel(logging.WarnLevel)
	if _, err := system.LoadRules([]*system.Rule{
		{
			MetricType:   system.InboundQPS,
			TriggerCount: 10,
			Strategy:     system.BBR,
		},
	}); err != nil {
		logger.LogHelper.Fatalf("Unexpected error: %+v", err)
	}
	return sentinelPlugin.SentinelMiddleware(
		//根据IP限流
		sentinelPlugin.WithResourceExtractor(func(c *gin.Context) string {
			return utils.GetClientIP(c)
		}),
		//风险提示
		sentinelPlugin.WithBlockFallback(func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusOK, response.AuthError("too many request; the quota used up"))
		}))

}
