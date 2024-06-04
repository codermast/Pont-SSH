package service

import "context"

type Connection struct {
	ctx context.Context
}

// NewConnection 创建一个 Connection 对象
func NewConnection() *Connection {
	return &Connection{}
}

// Startup 启动方法，用来注册一些资源
func (a *Connection) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Shutdown 销毁方法，用户释放一些资源
func (a *Connection) Shutdown() {
	a.ctx = nil
}
