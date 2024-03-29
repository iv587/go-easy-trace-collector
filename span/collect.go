package span

import (
	"collector/db"
	"collector/easytrace"
	"github.com/golang/protobuf/proto"
	"time"
)

func buildSpan(bSpan []byte) (easytrace.EasySpan, error) {
	var easySpan easytrace.EasySpan
	err := proto.Unmarshal(bSpan, &easySpan)
	if err != nil {
		return easySpan, err
	}
	return easySpan, nil
}

func Process(body []byte) {
	span, err := buildSpan(body)
	var isErr = 0
	if span.Error {
		isErr = 1
	}
	if err != nil {
		return
	}
	do := EasySpanDO{
		StartTime:        span.StartTime,
		FinishTime:       span.FinishTime,
		OperationName:    span.OperationName,
		TraceId:          span.TraceId,
		SpanId:           span.SpanId,
		ParentId:         span.ParentId,
		Error:            isErr,
		SpanKind:         span.SpanKind,
		Application:      span.Application,
		ApplicationGroup: span.ApplicationGroup,
		LogDatas:         span.LogDatas,
		Tags:             span.Tags,
		AppInstance:      span.AppInstance,
		Component:        span.Component,
	}
	time1 := time.Unix(do.StartTime/1000, 0)
	_, err = db.GetEngine().Table(getEasySpanTableName(time1)).Insert(&do)
	if err != nil {
		return
	}
}
