package span

import (
	"collector/db"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"time"
)

// 获取分页列表
func ListPage(query Query) (Page, error) {
	var startTime, endTime int64
	if query.Day == "" {
		now := time.Now()
		endTime = now.Unix() * 1000
		startTimeStr := now.Format("2006-01-02")
		startTimeStr = startTimeStr + " 00:00:00"
		dayStart, err := time.Parse("2006-01-02 15:04:05", startTimeStr)
		if err != nil {
			return Page{}, err
		}
		startTime = dayStart.Unix() * 1000
	} else {
		date, err := time.Parse("2006-01-02 15:04:05", query.Day+" 00:00:00")
		if err != nil {
			return Page{}, err
		}
		startTime = date.Unix() * 1000
		date, err = time.Parse("2006-01-02 15:04:05", query.Day+" 23:59:59")
		if err != nil {
			return Page{}, err
		}
		endTime = date.Unix() * 1000
	}
	query.startTime = startTime
	query.endTime = endTime
	session := wrapSessionQuery(query)
	count, err := session.Count()
	if err != nil {
		return Page{}, err
	}
	var list []EasySpanDO
	err = wrapSessionQuery(query).Limit(query.Size, (query.PageNo-1)*query.Size).Find(&list)
	if err != nil {
		return Page{}, err
	}
	pageRes := Page{
		Total:       count,
		List:        list,
		CurrentPage: query.PageNo,
	}
	return pageRes, nil
}

func wrapSessionQuery(query Query) *xorm.Session {
	session := db.GetEngine().Table("easy_span")
	session = session.Where("span_kind = ?", "server")
	if query.startTime > 0 && query.endTime > 0 {
		session.And(" start_time >= ? and start_time < ?", query.startTime, query.endTime)
	}
	if query.Error > -1 {
		session = session.And("error = ?", query.Error)
	}
	if query.OperationKey != "" {
		session = session.And("operation_name like ?", "%"+query.OperationKey+"%")
	}
	if query.ApplicationGroup != "" {
		session = session.And("application_group = ? ", query.ApplicationGroup)
	}
	if query.Application != "" {
		session = session.And("application = ?", query.Application)
	}
	session = session.Desc("start_time")
	return session
}

func TraceTree(traceId string) (TreeNode, error) {
	var list []EasySpanDO
	fmt.Println(traceId)
	err := db.GetEngine().Table("easy_span").Where("trace_id=?", traceId).Asc("start_time").Find(&list)
	if err != nil {
		return TreeNode{}, err
	}
	var span EasySpanDO
	childMap := make(map[string][]EasySpanDO)
	for _, it := range list {
		if it.ParentId == "-1" {
			span = it
			continue
		}
		v, ok := childMap[it.ParentId]
		if !ok {
			v = make([]EasySpanDO, 0)
		}
		v = append(v, it)
		childMap[it.ParentId] = v
	}
	//转换成树形结构
	treeNode := toTree(span, childMap, 0)
	return treeNode, nil
}

func toTree(node EasySpanDO, childMap map[string][]EasySpanDO, deepth int) TreeNode {
	treeNode := parseTreeNode(node)
	treeNode.Deepth = deepth
	children, ok := childMap[node.SpanId]
	if !ok || len(children) <= 0 {
		return treeNode
	}
	treeNodeChildren := make([]TreeNode, 0, len(children))
	for _, easySpan := range children {
		childTreeNode := toTree(easySpan, childMap, deepth+1)
		childTreeNode.ParentStartTime = treeNode.StartTime
		childTreeNode.ParentElapsedTime = treeNode.ElapsedTime
		treeNodeChildren = append(treeNodeChildren, childTreeNode)
	}
	treeNode.Children = treeNodeChildren
	return treeNode
}

var emptySpanNode = make([]TreeNode, 0)

func parseTreeNode(it EasySpanDO) TreeNode {
	tree := TreeNode{
		ElapsedTime:  it.FinishTime - it.StartTime,
		Children:     emptySpanNode,
		ShowChildren: true,
	}
	tree.OperationName = it.OperationName
	tree.StartTime = it.StartTime
	tree.Application = it.Application
	tree.SpanKind = it.SpanKind
	tree.Id = it.Id
	tree.ApplicationGroup = it.ApplicationGroup
	tree.SpanId = it.SpanId
	tree.Error = it.Error
	tree.Component = it.Component
	return tree
}

func Get(id int64) (EasySpanDO, error) {
	var list []EasySpanDO
	session := db.GetEngine().Table("easy_span")
	session = session.Where("id=?", id)
	err := session.Find(&list)
	if err != nil {
		fmt.Println("1111")
		return EasySpanDO{}, err
	}
	if len(list) <= 0 {
		fmt.Println("2222")
		return EasySpanDO{}, errors.New("记录不存在")
	}
	return list[0], nil

}