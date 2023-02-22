package api

import (
	"collector/auth"
	"collector/client"
	"collector/config"
	"collector/span"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var trace = new(Trace)

type Trace struct {
}

func (t *Trace) Login(c *gin.Context) {
	userName := c.PostForm("userName")
	password := c.PostForm("password")
	if userName == config.User.Name && password == config.User.Password {
		token, err := auth.Token(userName, password)
		if err != nil {
			error(c, err.Error())
			return
		}
		succ(c, "", gin.H{
			"token": token,
		})
	}
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
	idStr := c.PostForm("id")
	startTimeStr := c.PostForm("startTime")
	startTime, err := strconv.ParseInt(startTimeStr, 10, 64)
	if err != nil {
		panic(err)
	}
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		panic(err)
	}
	node, err := span.TraceTree(id, time.Unix(startTime/1000, 0))
	if err != nil {
		panic(err)
	}
	succ(c, "", node)
}

func (t *Trace) getSpanById(c *gin.Context) {
	idStr := c.PostForm("id")
	startTimeStr := c.PostForm("startTime")
	startTime, err := strconv.ParseInt(startTimeStr, 10, 64)
	if err != nil {
		panic(err)
	}
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		panic(err)
	}
	span, err := span.Get(id, time.Unix(startTime/1000, 0))
	if err != nil {
		panic(err)
	}
	succ(c, "", span)
}

func (t *Trace) getApp(c *gin.Context) {
	groupApp := client.GroupApp()
	groupAppList := make([]GroupAppVo, 0, len(groupApp))
	for groupName, v := range groupApp {
		groupInfo := GroupAppVo{
			Label: groupName,
			Value: groupName,
		}
		children := make([]GroupAppVo, 0, len(v))
		for appName, _ := range v {
			appNameInfo := GroupAppVo{
				Label: appName,
				Value: appName,
			}
			children = append(children, appNameInfo)
		}
		groupInfo.Children = children
		groupAppList = append(groupAppList, groupInfo)
	}
	succ(c, "", groupAppList)
}
