package api

import (
	"collector/connection"
	"collector/span"
	"github.com/gin-gonic/gin"
	"strconv"
)

var trace = new(Trace)

type Trace struct {
}

func (t *Trace) list(c *gin.Context) {
	var query span.Query
	query.Error = -1
	err := c.ShouldBind(&query)
	if err != nil {
		panic(err)
	}
	page, err := span.ListPage(query)
	if err != nil {
		panic(err)
	}
	succ(c, "", page)
}

func (t *Trace) tree(c *gin.Context) {
	traceId := c.PostForm("traceId")
	node, err := span.TraceTree(traceId)
	if err != nil {
		panic(err)
	}
	succ(c, "", node)
}

func (t *Trace) getSpanById(c *gin.Context) {
	idStr := c.PostForm("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		panic(err)
	}
	span, err := span.Get(id)
	if err != nil {
		panic(err)
	}
	succ(c, "", span)
}

func (t *Trace) getApp(c *gin.Context) {
	list := connection.ParseApp()
	succ(c, "", list)
}
