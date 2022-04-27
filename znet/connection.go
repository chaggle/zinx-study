package znet

import (
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/chaggle/zinx-study/ziface"
)

type Connection struct {

	//当前连接归属于哪个 Server
	TCPServer ziface.IServer

	Conn   *net.TCPConn
	ConnID uint32

	//当前链接的状态
	isClosed bool

	//消息管理 MsgID 和对应处理方法的消息管理模块
	msgHandler ziface.IMsgHandle

	//告知当前链接已经退出的 channel
	ExitBuffChan chan bool

	//无缓冲的管道，用于读、写两个 Goroutine 之间的通信
	msgChan chan []byte
}

// NewConnection 初始化链接模块
func NewConnection(server ziface.IServer, conn *net.TCPConn, connid uint32, msgHandler ziface.IMsgHandle) *Connection {
	c := &Connection{
		TCPServer:    server,
		Conn:         conn,
		ConnID:       connid,
		isClosed:     false,
		msgHandler:   msgHandler,
		ExitBuffChan: make(chan bool, 1),
		msgChan:      make(chan []byte), //msgChan 初始化
	}

	c.TCPServer.GetConnManager().Add(c)
	return c
}

// StartRead 链接的读业务的方法
func (c *Connection) StartRead() {
	fmt.Println("[Read Goroutine is running ...]")
	defer fmt.Println("connID = ", c.ConnID, "Reader is exit, remoter addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//创建拆包解包对象
		dp := NewDataPack()

		//读取客户端的 Msg Head 二进制流的 8 个字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error: ", err)
			c.ExitBuffChan <- true
			continue
		}

		//拆包，得到 msgID 和 datalen 放在 msg 中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error: ", err)
			c.ExitBuffChan <- true
			continue
		}

		//根据得到的 datalen 读取 data，放在 msg.Data 中
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen()) //上面有 var 定义了全局，此处是传值，并非赋值！
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error", err)
				c.ExitBuffChan <- true
				continue
			}
		}
		msg.SetData(data)

		//得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			msg:  msg,
		}

		//从绑定好的消息和对应的消息处理方法中执行对应的 Handle 方法
		go c.msgHandler.DoMsgHandler(&req)
	}

}

// StartWrite 创建写业务的 Goroutine
func (c *Connection) StartWrite() {
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println("connID = ", c.ConnID, "Writer is exit, remoter addr is ", c.RemoteAddr().String())
	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send Data error: ", err)
				return
			}
		case <-c.ExitBuffChan:
			//conn 已经关闭
			return
		}
	}
}

// Start 启动链接
func (c *Connection) Start() {
	fmt.Println("Conn start ... ConnID = ", c.ConnID)

	//启动用户从客户端读取数据的 go
	go c.StartRead()

	//启动用户从客户端写回客户端数据的 go
	go c.StartWrite()

	for {
		<-c.ExitBuffChan
		//得到退出消息，不在阻塞
		return
	}
}

// Stop 停止链接
func (c *Connection) Stop() {
	fmt.Println("Conn stop ... ConnID = ", c.ConnID)

	//如果当前链接已经关闭
	if c.isClosed {
		return
	}

	c.isClosed = true

	//关闭 socket 链接
	_ = c.Conn.Close()

	//关闭Writer Goroutine
	c.ExitBuffChan <- true

	//将链接从 连接管理器中删除
	c.TCPServer.GetConnManager().Remove(c)

	// channel 资源回收
	close(c.ExitBuffChan)
	close(c.msgChan)
}

// GetTCPConnection 获取链接绑定的 socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取当前链接模块的 ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取远程客户端的 TCP状态 IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// SendMsg 提供一个 SendMsg 方法，将我们要发送给客户端的数据，先进行封包，在进行发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("connection closed when send Msg")
	}

	//将 data 封包， 并发送
	dp := NewDataPack()
	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("pack error msg")
	}

	//客户端写回
	c.msgChan <- binaryMsg

	return nil
}
