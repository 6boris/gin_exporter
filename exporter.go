package gin_exporter

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
)

var (
	GinAppInfo = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gin_app_info",
			Help: "The Gin Run info",
		},
		[]string{
			"group",      //	分组
			"type",       //	类型
			"system",     //	系统
			"instance",   //	实例
			"version",    //	版本
			"commit_id",  //	commit_id
			"started_at", //
			"platform",   //
		},
	)

	GinRequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gin_request_total",
			Help: "Number of hello requests in total",
		},
		[]string{"method", "endpoint"},
	)

	GinRequestGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gin_request_total2",
			Help: "Number of hello requests in total",
		},
		[]string{"method", "endpoint"},
	)

	GinRequestHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "gin_request_total3",
			Help:    "Number of hello requests in total",
			Buckets: []float64{},
		},
		[]string{"method", "endpoint"},
	)

	GinRequestSummary = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "gin_request_time_millisecond",
			Help:       "Number of hello requests in total",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"method", "endpoint"},
	)
)

type GinRoutesInfo struct {
	Method  string
	Path    string
	Handler string
}

var routeInfo []GinRoutesInfo

//
func init() {
	//prometheus.MustRegister(GinRequestCount)
	//prometheus.MustRegister(GinRequestGauge)
	//prometheus.MustRegister(GinRequestHistogram)
	//prometheus.MustRegister(GinRequestSummary)
	//registerHandler()
}

func registerDefaultHandler() {
	if err := prometheus.Register(GinAppInfo); err != nil {
		log.Panicln(err.Error())
	}

	if err := prometheus.Register(GinRequestCount); err != nil {
		log.Panicln(err.Error())
	}

	if err := prometheus.Register(GinRequestGauge); err != nil {
		log.Panicln(err.Error())
	}

	if err := prometheus.Register(GinRequestHistogram); err != nil {
		log.Panicln(err.Error())
	}
	if err := prometheus.Register(GinRequestSummary); err != nil {
		log.Panicln(err.Error())
	}
}

func Default(app *gin.Engine) {
	app.Use(defaultExporter)
	registerDefaultHandler()
	app.GET("metrics", metricsHandler())
	generateRouteInfo(app)
}

//	Gin Metrics api handler
func metricsHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// 初始化 Router信息
func generateRouteInfo(app *gin.Engine) {
	context.Background()
	for _, v := range app.Routes() {
		routeInfo = append(routeInfo, GinRoutesInfo{
			Method:  v.Method,
			Path:    v.Path,
			Handler: v.Handler,
		})
	}
}
