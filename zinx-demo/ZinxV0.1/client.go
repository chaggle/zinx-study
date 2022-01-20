package main

import (
	"fmt"
	"net"
	"time"
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
		//_ 代表成功写入的字节数，在此并不关心字节数能写入多少
		_, err := conn.Write([]byte("Hello zinx V0.1"))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}

		buf := make([]byte, 512) //512Byte
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read buf error ", err)
			return
		}

		fmt.Printf("server call back: %s, cnt = %d\n", buf, cnt)

		//对cpu进行阻塞
		time.Sleep(10 * time.Second) //每隔10s 发一下Hello zinx V0.1
	}
}
