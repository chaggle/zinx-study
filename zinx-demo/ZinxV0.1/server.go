package main

import "github.com/chaggle/zinx-study/znet"

//Server 模块的测试函数
func main() {

	//1 创建一个server 句柄 s
	s := znet.NewServer("[zinx V0.1]")

	//2 启动服务器
	s.Serve()
}
