package api

type GroupAppVo struct {
	Label    string       `json:"label"`
	Value    string       `json:"value"`
	Children []GroupAppVo `json:"children"`
}

type ConnectInfoVo struct {

	// 地址
	Addr string `json:"addr"`

	// 链接建立时间
	CreateTime string `json:"createTime"`

	// 链接存活时间
	AliveTime string `json:"aliveTime"`

	// 应用组
	AppGroup string `json:"appGroup"`

	// 应用
	AppName string `json:"appName"`

	// 应用启动时间
	AppStartTime string `json:"appStartTime"`

	// 网络入口流量
	NetInput string `json:"netInput"`
}
