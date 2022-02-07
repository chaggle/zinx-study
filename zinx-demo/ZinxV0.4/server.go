package main

import (
	"fmt"

	"github.com/chaggle/zinx-study/ziface"
	"github.com/chaggle/zinx-study/znet"
)

/*
	基于 zinx 框架开发的服务器端应用程序
	单元进行函数的测试
*/

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

//Test PreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("Call Back Before Ping error")
	}
}

//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping...\n"))
	if err != nil {
		fmt.Println("Call Back Ping...Ping...Ping... error")
	}
}

//Test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping...\n"))
	if err != nil {
		fmt.Println("Call Back After Ping... error")
	}
}

//Server 模块的测试函数
func main() {

	//1 创建一个server 句柄 s
	s := znet.NewServer()

	//2 给当前zinx框架加入自定义的router
	s.AddRouter(&PingRouter{})

	//3 启动服务器
	s.Serve()
}
