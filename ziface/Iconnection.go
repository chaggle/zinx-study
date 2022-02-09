package ziface

import "net"

// 定义链接模块的抽象层
type IConnection interface {
	//启动链接
	Start()
	//停止链接
	Stop()
	//获取链接绑定的 socket conn
	GetTCPConnection() *net.TCPConn
	//获取当前链接模块的ID
	GetConnID() uint32
	//获取远程客户端的 TCP 状态 IP port
	RemoteAddr() net.Addr
	//发送数据，将数据发送给远程的客户端
	SendMsg(msgId uint32, data []byte) error
}

//定义一个处理链接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
