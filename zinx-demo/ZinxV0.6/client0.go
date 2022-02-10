package main

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/chaggle/zinx-study/znet"
)

func main() {
	fmt.Println("client start")

	time.Sleep(5 * time.Second)

	//1 直接链接远程服务器，得到一个conn链接
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		//发送封包消息
		dp := znet.NewDataPack()
		msg, _ := dp.Pack(znet.NewMsgPackage(200, []byte("Zinx V0.6 search!")))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("Write error err: ", err)
			return
		}

		//读取流中的 head 部分
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData) //此处不能用 := 因为_, err := conn.Write(binaryMsg) 已经使用
		if err != nil {
			fmt.Println("read head error")
			break
		}

		//将 headData 字节流，拆包到 msg 中
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack error: ", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			//msg 有 data 数据，需要再次读取 data 数据
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			//根据 datalen 从io中读取数据流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("Server unpack data error: ", err)
				return
			}
			fmt.Println("==> Recieve Msg: ID = ", msg.Id, "Datalen = ", msg.DataLen, "Data = ", string(msg.Data))
		}
		//对cpu进行阻塞
		time.Sleep(2 * time.Second)
	}
}
