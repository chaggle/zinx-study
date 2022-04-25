package ziface

import "net"

// IConnection 定义链接模块的抽象层
type IConnection interface {
	// Start 启动链接
	Start()
	// Stop 停止链接
	Stop()
	// GetTCPConnection 获取链接绑定的 socket conn
	GetTCPConnection() *net.TCPConn
	// GetConnID 获取当前链接模块的ID
	GetConnID() uint32
	// RemoteAddr
	//获取远程客户端的 TCP 状态 IP port
	RemoteAddr() net.Addr
	// SendMsg 发送数据，将数据发送给远程的客户端
	SendMsg(msgId uint32, data []byte) error
}

// HandleFunc 定义一个处理链接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
