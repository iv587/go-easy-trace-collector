package api

import "github.com/gin-gonic/gin"

func Route(router *gin.Engine) {

	router.Static("/assets", "static/assets")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.StaticFile("/", "./static/index.html")
	apiGroup := router.Group("/api")

	traceApiGroup := apiGroup.Group("/trace")
	traceApiGroup.Any("/list", trace.list)
	traceApiGroup.Any("/tree", trace.tree)
	traceApiGroup.Any("/getSpanById", trace.getSpanById)
	traceApiGroup.Any("/getApp", trace.getApp)

	connectApiGroup := apiGroup.Group("/connect")
	connectApiGroup.Any("/list", pConnect.list)
}
