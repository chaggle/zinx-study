package ziface

//服务器接口定义
type IServer interface {
	// Start 启动服务器
	Start()
	// Stop 停止服务器
	Stop()
	// Serve 运行服务器
	Serve()
	// AddRouter 路由功能，给当前服务器注册一个路由方法，供客户端链接使用
	AddRouter(msgId uint32, router IRouter)
}
