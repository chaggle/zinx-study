package main

import (
	"fmt"
	"io"
	"net"

	"github.com/chaggle/zinx-study/znet"
)

/*
	pack封包、unpack拆包的单元测试
*/

func main() {
	/*
		模拟的服务器
	*/

	//1 创建SocketTCP
	listenner, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen error: ", err)
		return
	}
	//创建一个 go 承载负责从客户端处理事务
	//2 从客户端读取数据，拆包处理
	for {
		conn, err := listenner.Accept()
		if err != nil {
			fmt.Println("server accept error: ", err)
		}

		go func(conn net.Conn) {
			//处理客户端的请求
			//----> 拆包过程 <----
			//定义拆包对象
			dp := znet.NewDataPack()
			for {
				//1 先从 conn 中把 head 读出来
				headData := make([]byte, dp.GetHeadLen())
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					fmt.Println("read head error ")
					break
				}

				msgHead, err := dp.Unpack(headData)
				if err != nil {
					fmt.Println("server unpack error ", err)
					return
				}
				//2 第二次读，通过 head 中 datalen，在读 data 数据
				if msgHead.GetDataLen() > 0 {
					// msg 有数据的，需要第二次进行 data 数据的读取
					msg := msgHead.(*znet.Message)
					msg.Data = make([]byte, msg.GetDataLen())

					//根据datalen的长度再次从IO流中读取data数据
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("server unpack error ", err)
						return
					}

					//完整的一个消息已经读取完毕
					fmt.Println("----> Recieve MsgID: ", msg.Id, "dataLen = ", msg.DataLen, "data = ", string(msg.Data))
				}
			}

		}(conn)
	}
}
