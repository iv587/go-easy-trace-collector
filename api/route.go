package api

import (
	"collector/auth"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {

	router.Static("/assets", "static/assets")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")
	router.StaticFile("/", "./static/index.html")
	apiGroup := router.Group("/api")
	apiGroup.Any("/login", trace.Login)
	traceApiGroup := apiGroup.Group("/trace")
	traceApiGroup.Use(checkAuth)
	traceApiGroup.Any("/list", trace.list)
	traceApiGroup.Any("/tree", trace.tree)
	traceApiGroup.Any("/getSpanById", trace.getSpanById)
	traceApiGroup.Any("/getApp", trace.getApp)

	connectApiGroup := apiGroup.Group("/connect")
	connectApiGroup.Use(checkAuth)
	connectApiGroup.Any("/list", pConnect.list)
}

var checkAuth = func(context *gin.Context) {
	token := context.PostForm("token")
	ok, err := auth.Verify(token)
	if err != nil {
		notLogin(context, "用户信息校验失败")
		context.Abort()
		return
	}
	if !ok {
		notLogin(context, "用户登录失败")
		context.Abort()
		return
	} else {
		context.Next()
	}
}
