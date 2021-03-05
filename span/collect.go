package span

import (
	"collector/db"
	"collector/easytrace"
	"github.com/golang/protobuf/proto"
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
	var isErr int = 0
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
	_, err = db.GetEngine().Table("easy_span").Insert(&do)
	if err != nil {
		return
	}
}
