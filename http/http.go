package http

import (
	"collector/api"
	"collector/config"
	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()
	api.Route(r)
	r.Run(config.Http.Addr)
}
