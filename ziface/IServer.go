package ziface

//服务器接口定义
type IServer interface {

	//启动服务器
	Start()

	//停止服务器
	Stop()

	//运行服务器
	Serve()
}
