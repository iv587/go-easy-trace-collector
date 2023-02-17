package span

import "time"

type Query struct {
	Application      string `form:"application"`
	ApplicationGroup string `form:"applicationGroup"`
	SpanKind         string `form:"spanKind"`
	OperationName    string `form:"operationName"`
	Day              string `form:"day"`
	PageNo           int    `form:"pageNo"`
	Size             int    `form:"size"`
	Error            int    `form:"error"`
	OperationKey     string `form:"operationKey"`
	startTime        int64
	endTime          int64
	time             time.Time
}
