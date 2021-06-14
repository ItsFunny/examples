/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/2 10:24 上午
# @File : main.go
# @Description :
# @Attention :
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"strconv"
	"sync"
	"time"
)

const (
	metricsPath = "/metrics"
	faviconPath = "/favicon.ico"
)

var (
	// httpHistogram prometheus 模型
	httpHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   "http_server",
		Subsystem:   "",
		Name:        "requests_seconds",
		Help:        "Histogram of response latency (seconds) of http handlers.",
		ConstLabels: nil,
		Buckets:     nil,
	}, []string{"method", "code", "uri"})
)

// init 初始化prometheus模型
func init() {
	prometheus.MustRegister(httpHistogram)
}

// handlerPath 定义采样路由struct
type handlerPath struct {
	sync.Map
}

// get 获取path
func (hp *handlerPath) get(handler string) string {
	v, ok := hp.Load(handler)
	if !ok {
		return ""
	}
	return v.(string)
}

// set 保存path到sync.Map
func (hp *handlerPath) set(ri gin.RouteInfo) {
	hp.Store(ri.Handler, ri.Path)
}

// GinPrometheus gin调用Prometheus的struct
type GinPrometheus struct {
	engine  *gin.Engine
	ignored map[string]bool
	pathMap *handlerPath
	updated bool
}

type Option func(*GinPrometheus)

// Ignore 添加忽略的路径
func Ignore(path ...string) Option {
	return func(gp *GinPrometheus) {
		for _, p := range path {
			gp.ignored[p] = true
		}
	}
}

// New new gin prometheus
func New(e *gin.Engine, options ...Option) *GinPrometheus {
	if e == nil {
		return nil
	}

	gp := &GinPrometheus{
		engine: e,
		ignored: map[string]bool{
			metricsPath: true,
			faviconPath: true,
		},
		pathMap: &handlerPath{},
	}

	for _, o := range options {
		o(gp)
	}
	return gp
}

// updatePath 更新path
func (gp *GinPrometheus) updatePath() {
	gp.updated = true
	for _, ri := range gp.engine.Routes() {
		gp.pathMap.set(ri)
	}
}

// Middleware set gin middleware
func (gp *GinPrometheus) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !gp.updated {
			gp.updatePath()
		}
		// 过滤请求
		if gp.ignored[c.Request.URL.String()] {
			c.Next()
			return
		}

		start := time.Now()
		c.Next()

		httpHistogram.WithLabelValues(
			c.Request.Method,
			strconv.Itoa(c.Writer.Status()),
			gp.pathMap.get(c.HandlerName()),
		).Observe(time.Since(start).Seconds())
	}
}

func main() {
	// r:=gin.New()
	// ginPrometheus := New(r)
	// r.Use(ginPrometheus.Middleware())
	// // metrics采样
	// r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// r.Run(":10000")
	ExamplePusher_Push()
}

var ()

func ExamplePusher_Push() {
	opts := prometheus.HistogramOpts{
		Namespace: "DEMO",
		Subsystem: "state",
		Name:      "block_verify",
		Help:      "校验接收到的区块的时间花费",
	}
	histogram := prometheus.NewHistogram(opts)
	// completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
	// 	Name: "db_backup_last_completion_timestamp_seconds",
	// 	Help: "The timestamp of the last successful completion of a DB backup.",
	// })
	histogram.Observe(1)
	// completionTime.SetToCurrentTime()
	if err := push.New("http://localhost:9091", "demo").
		Collector(histogram).
		Grouping("db", "customers").

		Push(); err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}
