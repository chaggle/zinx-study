package ziface

import "net"

// 定义链接模块的抽象层
type IConnection interface {
	Start() //启动链接

	Stop() //停止链接

	GetTCPConnection() *net.TCPConn //获取链接绑定的 socket conn

	GetConnID() uint32 //获取当前链接模块的 ID

	RemoteAddr() net.Addr //获取远程客户端的 TCP状态 IP port

	Send(data []byte) error //发送数据，将数据发送给远程的客户端
}

//定义一个处理链接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
