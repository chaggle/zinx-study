package znet

import (
	"fmt"
	"net"

	"github.com/chaggle/zinx-study/ziface"
)

/*
	链接模块
*/

type Connection struct {
	Conn *net.TCPConn //当前链接的 socket TCP套接字

	ConnID uint32 //链接的ID

	isClosed bool //当前链接的状态

	handlerAPI ziface.HandleFunc //当前处理链接业务模块的API

	ExitChan chan bool //告知当前链接已经退出的 channel
}

//初始化链接模块
func NewConnection(conn *net.TCPConn, connid uint32, callback_api ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connid,
		handlerAPI: callback_api,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
	}

	return c
}

//链接的读业务的方法
func (c *Connection) StartRead() {
	fmt.Println("Read Goroutine is running ...")
	defer fmt.Println("connID = ", c.ConnID, "Reader is exit, remoter addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端的数据到 buf 中
		buf := make([]byte, 512)

		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("receive buf err ", err)
			continue
		}

		//调用当前链接所绑定的HandleAPI
		if err := c.handlerAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("connID ", c.ConnID, "handle is error ", err)
			break
		}
	}

}

//启动链接
func (c *Connection) Start() {
	fmt.Println("Conn start ... ConnID = ", c.ConnID)

	//启动从当前链接读取数据的 go
	go c.StartRead()

}

//停止链接
func (c *Connection) Stop() {
	fmt.Println("Conn stop ... ConnID = ", c.ConnID)

	//如果当前链接已经关闭
	if c.isClosed == true {
		return
	}

	c.isClosed = true

	//关闭 socket 链接
	c.Conn.Close()

	// channel 资源回收
	close(c.ExitChan)
}

//获取链接绑定的 socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取当前链接模块的 ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取远程客户端的 TCP状态 IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//发送数据，将数据发送给远程的客户端
func (c *Connection) Send(data []byte) error {
	return nil
}
