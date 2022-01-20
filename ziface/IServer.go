package ziface

//服务器接口定义
type IServer interface {
	Start() //启动服务器

	Stop() //停止服务器

	Serve() //运行服务器
}
