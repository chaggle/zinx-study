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

//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	fmt.Println("recieve from client: MsgId = ", request.GetMsgID(), ", data = ", string(request.GetData()))
	err := request.GetConnection().SendMsg(200, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	//先读取客户端的数据，再回写Hello Zinx Router V0.6
	fmt.Println("recieve from client: MsgId = ", request.GetMsgID(), ", data = ", string(request.GetData()))
	err := request.GetConnection().SendMsg(201, []byte("Hello Zinx Router V0.6"))
	if err != nil {
		fmt.Println(err)
	}
}

//Server 模块的测试函数
func main() {

	//1 创建一个server 句柄 s
	s := znet.NewServer()

	//2 给当前zinx框架加入自定义的router
	s.AddRouter(200, &PingRouter{})
	s.AddRouter(201, &HelloZinxRouter{})

	//3 启动服务器
	s.Serve()
}
