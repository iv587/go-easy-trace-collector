package client

import (
	"sync"
)

var clientMap = struct {
	sync.RWMutex
	connections map[string]*Context
}{
	connections: make(map[string]*Context),
}

func registContext(ctx *Context) {
	clientMap.Lock()
	defer clientMap.Unlock()
	clientMap.connections[ctx.conn.RemoteAddr().String()] = ctx
}

func removeContext(ctx *Context) {
	clientMap.Lock()
	defer clientMap.Unlock()
	delete(clientMap.connections, ctx.conn.RemoteAddr().String())
}

func GroupApp() map[string]map[string]int {

	app := make(map[string]map[string]int)
	contexts := cloneConnection()
	for _, v := range contexts {
		group := v.AppInfo.Group
		v1, ok := app[group]
		if !ok {
			v1 = make(map[string]int)
		}
		_, ok = v1[v.AppInfo.Name]
		if !ok {
			v1[v.AppInfo.Name] = 1
		}
		app[group] = v1
	}
	return app
}

func ConnectList() map[string]Context {
	return cloneConnection()
}

func cloneConnection() map[string]Context {
	clientMap.RLock()
	clientMap.RUnlock()
	contexts := clientMap.connections
	cloneCtx := make(map[string]Context)
	for k, v := range contexts {
		clon := Context{
			Time:      v.Time,
			Alive:     v.Alive,
			ReadBytes: v.ReadBytes,
			AppInfo:   v.AppInfo,
		}
		cloneCtx[k] = clon
	}
	return cloneCtx
}
