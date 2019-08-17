package gin_exporter

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func exporterAppInfo(c *gin.Context) {
	GinAppInfo.With(prometheus.Labels{
		"group":      "default",                                //	分组
		"type":       "default",                                //	类型
		"system":     "default-api",                            //	系统
		"instance":   "default-api-instances-1",                //	实例
		"version":    "0.0.0",                                  //	版本
		"commit_id":  "00000000000000000000000000000000000000", //	commit_id
		"started_at": "2019-08-17 00:00:00",                    //
		"platform":   "windows...",                             //
	})
}
func exporterGlobalRequestCount(c *gin.Context) {
	GinRequestCount.With(prometheus.Labels{"method": c.Request.Method, "endpoint": c.FullPath()}).Inc()
}

func exporterGinRequestHistogram(c *gin.Context) {
	start := time.Now().UnixNano()

	c.Next()

	end := time.Now().UnixNano()
	use_time := end - start
	GinRequestSummary.With(
		prometheus.Labels{
			"method":   c.Request.Method,
			"endpoint": c.FullPath(),
		},
	).Observe(float64(use_time))
}
func defaultExporter(c *gin.Context) {
	exporterAppInfo(c)
	exporterGlobalRequestCount(c)
	exporterGinRequestHistogram(c)
}
