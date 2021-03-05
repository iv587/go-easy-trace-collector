package connection

import (
	"sync"
)

var connectMap = struct {
	sync.RWMutex
	connections map[string]*Context
}{
	connections: make(map[string]*Context),
}

func registContext(ctx *Context) {
	connectMap.Lock()
	defer connectMap.Unlock()
	connectMap.connections[ctx.conn.RemoteAddr().String()] = ctx
}

func removeContext(ctx *Context) {
	connectMap.Lock()
	defer connectMap.Unlock()
	delete(connectMap.connections, ctx.conn.RemoteAddr().String())
}

func ParseApp() []AppInfo {

	app := make(map[string]map[string]int)
	contexts := cloneConnection()
	for _, v := range contexts {
		group := v.App.Group
		v1, ok := app[group]
		if !ok {
			v1 = make(map[string]int)
		}
		_, ok = v1[v.App.Name]
		if !ok {
			v1[v.App.Name] = 1
		}
		app[group] = v1
	}

	appInfoList := make([]AppInfo, 0, len(app))
	for groupName, v := range app {
		groupInfo := AppInfo{
			Label: groupName,
			Value: groupName,
		}
		children := make([]AppInfo, 0, len(v))
		for appName, _ := range v {
			appNameInfo := AppInfo{
				Label: appName,
				Value: appName,
			}
			children = append(children, appNameInfo)
		}
		groupInfo.Children = children
		appInfoList = append(appInfoList, groupInfo)
	}
	return appInfoList

}

func cloneConnection() map[string]Context {
	connectMap.RLock()
	connectMap.RUnlock()
	contexts := connectMap.connections
	cloneCtx := make(map[string]Context)
	for k, v := range contexts {
		clon := Context{
			time:      v.time,
			alive:     v.alive,
			ReadBytes: v.ReadBytes,
			App:       v.App,
		}
		cloneCtx[k] = clon
	}
	return cloneCtx
}
