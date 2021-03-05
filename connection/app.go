package connection

type App struct {
	Name  string
	Group string
}

type AppInfo struct {
	Label    string    `json:"label"`
	Value    string    `json:"value"`
	Children []AppInfo `json:"children"`
}
