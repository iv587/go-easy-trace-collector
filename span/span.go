package span

type EasySpanDO struct {
	Id               int64  `xorm:"'id'" json:"id"`
	StartTime        int64  `xorm:"'start_time'" json:"startTime"`
	FinishTime       int64  `xorm:"'finish_time'" json:"finishTime"`
	OperationName    string `xorm:"'operation_name'" json:"operationName"`
	TraceId          string `xorm:"'trace_id'" json:"traceId"`
	SpanId           string `xorm:"'span_id'" json:"spanId"`
	ParentId         string `xorm:"'parent_id'" json:"parentId"`
	Error            int    `xorm:"'error'" json:"error"`
	SpanKind         string `xorm:"'span_kind'" json:"spanKind"`
	Application      string `xorm:"'application'" json:"application"`
	ApplicationGroup string `xorm:"'application_group'" json:"applicationGroup"`
	LogDatas         string `xorm:"'log_datas'" json:"logDatas"`
	Tags             string `xorm:"'tags'" json:"tags"`
	AppInstance      string `xorm:"'app_instance'" json:"appInstance"`
	Component        string `xorm:"'component'" json:"component"`
}

type Page struct {
	Total       int64        `json:"total"`
	List        []EasySpanDO `json:"list"`
	CurrentPage int          `json:"current_page"`
}

type TreeNode struct {
	EasySpanDO
	StartTimeText     string     `json:"startTimeText"`
	FinishTimeText    string     `json:"finishTimeText"`
	ElapsedTime       int64      `json:"elapsedTime"`
	ParentStartTime   int64      `json:"parentStartTime"`
	ParentElapsedTime int64      `json:"parentElapsedTime"`
	ShowChildren      bool       `json:"_showChildren"`
	Children          []TreeNode `json:"children"`
	Deepth            int        `json:"deepth"`
}
