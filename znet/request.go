package znet

import "github.com/chaggle/zinx-study/ziface"

type Request struct {
	//已经和客户端建立好的链接
	conn ziface.IConnection

	//客户端请求的数据
	msg ziface.IMessage
}

// GetConnection 获取请求连接信息
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// GetData 获取请求消息的信息
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// GetMsgID 获取请求消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
