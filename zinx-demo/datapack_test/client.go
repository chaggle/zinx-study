package main

import (
	"fmt"
	"net"

	"github.com/chaggle/zinx-study/znet"
)

/*
	出现以下问题
	panic: open conf/zinx.json: The system cannot find the path specified.
	goroutine 1 [running]:
	github.com/chaggle/zinx-study/utils.(*GlobalObj).Reload(0x3f)
	E:/Gopath/src/github.com/chaggle/zinx-study/utils/globalobj.go:43 +0x6e
	github.com/chaggle/zinx-study/utils.init.0()
	E:/Gopath/src/github.com/chaggle/zinx-study/utils/globalobj.go:64 +0x98
	exit status 2
	可以选用v0.4版本的conf文件，也可以直接将globalobj.go文件中的最后一行reload()注释掉
*/
func main() {
	/*
		模拟客户端
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial error: ", err)
		return
	}

	//创建一个封包的对象
	dp := znet.NewDataPack()

	//模拟粘包的过程，封装两个 msg 一起发送
	//封装第一个msg1包
	msg1 := &znet.Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'z', 'i', 'n', 'x'},
	}

	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 error: ", err)
		return
	}

	//封装第二个msg2包
	msg2 := &znet.Message{
		Id:      1,
		DataLen: 7,
		Data:    []byte{'n', 'i', ',', 'h', 'a', 'o', '!'},
	}

	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 error: ", err)
		return
	}

	//两包粘一起
	sendData1 = append(sendData1, sendData2...)

	//一起发送
	conn.Write(sendData1)

	//客户端阻塞
	select {}
}
