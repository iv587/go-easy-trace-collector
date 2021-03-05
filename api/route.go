package api

import "github.com/gin-gonic/gin"

func Route(router *gin.Engine) {
	router.Static("/css", "static/css")
	router.Static("/js", "static/js")
	router.Static("/img", "static/img")
	router.Static("/fonts", "static/fonts")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.StaticFile("/", "./static/index.html")
	apiGroup := router.Group("/api")
	apiGroup.Any("trace/list", trace.list)
	apiGroup.Any("trace/tree", trace.tree)
	apiGroup.Any("trace/getSpanById", trace.getSpanById)
	apiGroup.Any("trace/getApp", trace.getApp)
}
