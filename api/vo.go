package api

type GroupAppVo struct {
	Label    string       `json:"label"`
	Value    string       `json:"value"`
	Children []GroupAppVo `json:"children"`
}
