package span

import (
	"collector/db"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"strings"
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
		query.time = now
	} else {
		date1, err := time.Parse("2006-01-02 15:04:05", query.Day+" 00:00:00")
		query.time = date1
		if err != nil {
			return Page{}, err
		}
		startTime = date1.Unix() * 1000
		date, err := time.Parse("2006-01-02 15:04:05", query.Day+" 23:59:59")
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
	if len(list) <= 0 {
		list = make([]EasySpanDO, 0, 0)
	}
	pageRes := Page{
		Total:       count,
		List:        list,
		CurrentPage: query.PageNo,
	}
	return pageRes, nil
}

func wrapSessionQuery(query Query) *xorm.Session {
	session := db.GetEngine().Table(getEasySpanTableName(query.time))
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

func TraceTree(id int64, time time.Time) (TreeNode, error) {
	po, err := Get(id, time)
	if err != nil {
		return TreeNode{}, err
	}
	var list []EasySpanDO
	err = db.GetEngine().Table(getEasySpanTableName(time)).Where("trace_id=?", po.TraceId).Asc("start_time").Find(&list)
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
	timeFmt := "15:04:05.000"
	startSec := it.StartTime / 1000
	startNSec := (it.StartTime % 1000) * 1000000

	finishSec := it.FinishTime / 1000
	finishNSec := (it.FinishTime % 1000) * 1000000

	tree := TreeNode{
		ElapsedTime:    it.FinishTime - it.StartTime,
		Children:       emptySpanNode,
		ShowChildren:   true,
		StartTimeText:  time.Unix(startSec, startNSec).Local().Format(timeFmt),
		FinishTimeText: time.Unix(finishSec, finishNSec).Format(timeFmt),
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

func Get(id int64, time time.Time) (EasySpanDO, error) {
	var list []EasySpanDO
	session := db.GetEngine().Table(getEasySpanTableName(time))
	session = session.Where("id=?", id)
	err := session.Find(&list)
	if err != nil {
		return EasySpanDO{}, err
	}
	if len(list) <= 0 {
		return EasySpanDO{}, errors.New("记录不存在")
	}
	return list[0], nil

}

func getEasySpanTableName(time time.Time) string {
	dateStr := time.Format("20060102")
	tableName := fmt.Sprintf("easy_span_%s", dateStr)
	return tableName
}

func createTable() {
	log.Println("easy_span创建开始")
	defer log.Println("easy_span创建完成")
	tableName := getEasySpanTableName(time.Now().Add(24 * time.Hour))
	res, err := db.GetEngine().Exec("DROP table  if exists " + tableName)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(res.RowsAffected())

	sql := strings.Replace(tableCreateSql, "{TABLE_NAME}", tableName, -1)
	res, err = db.GetEngine().Exec(sql)
	if err != nil {
		log.Print(err)
		return
	}
	log.Println(res.RowsAffected())
}

func PreCreateTable() {
	d := time.Hour
	timer := time.NewTimer(d)
	defer timer.Stop()
	for {
		<-timer.C
		createTable()
		timer.Reset(d)
	}
}
