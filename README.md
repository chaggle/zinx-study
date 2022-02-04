---
title: "Zinx框架的学习"
date: 2022-01-10T10:14:32+08:00
tag: ["Go", "zinx"]
categories: ["Go"]
draft: true
---

# zinx 框架

开始学习 Go 语言实现的 zinx 框架,项目地址为：[https://github.com/chaggle/zinx-study](https://github.com/chaggle/zinx-study)

> 使用 go mod 管理, 初始化为 go mod init github.com/chaggle/zinx-study
> 并部署代码到 github.com 以及使用 go get 同步到本地 Gopath 的 github 包下！

## V0.1 基础的 server 模块

> 方法
>
> 初始化服务器 -- NewServer(name string) ziface.IServer
> 启动服务器 -- Start()
> 停止服务器 -- Stop()
> 运行服务器 -- Serve()
>
> 属性
>
> 名称 -- name
> IP 版本 -- IPVersion
> 监听 IP -- IP
> 监听端口 -- Port

## V0.2 简单的链接封装和业务绑定

> 方法
>
> 启动链接 -- Start()
> 停止链接 -- Stop()
> 获取当前链接的 conn 对象 (套接字) -- GetTCPConnection() \*net.TCPConn
> 得到链接 ID -- GetConnID() uint32
> 得到客户端连接的地址和端口 -- RemoteAddr() net.TCPAddr
> 发送数据的方法 -- Send(data []byte) error
>
> 属性
>
> socket TCP 套接字 -- Conn \*net.TCPConn
> 链接的 ID -- ConnID uint32
> 当前链接状态 (是否已经关闭) -- isClosed bool
> 与当前链接所绑定的处理业务与方法 -- handlerAPI ziface.HandleFunc
> 等待退出的 channel 管道 -- ExitChan chan bool

## V0.3 基础的 router 模块

>     Request 请求封装
>     	将链接与数据绑定一起
>     		属性
>     			链接的句柄 -- GetConnection() IConnection
>     			请求数据 -- GetData() []byte
>     		方法
>     			得到链接-- func (r *Request) GetConnection() ziface.IConnection
>     			得到数据 -- func (r *Request) GetData() []byte
>     			新建一个 Request 请求
>     Router 模块
>     	抽象的 IRouter
>     		处理业务之前的方法
>     			PreHandle(request IRequest)	//处理conn业务之前的方法
>     		处理业务的主方法
>     			Handle(request IRequest)	//处理conn业务的主方法
>     		处理业务之后的方法
>     			PostHandle(request IRequest)	//处理conn业务之后的方法
>     	具体的 BaseRouter
>     		处理业务之前的方法
>     			func (br *BaseRouter) PreHandle(request ziface.IRequest)
>     		处理业务的主方法
>     			func (br *BaseRouter) Handle(request ziface.IRequest)
>     		处理业务之后的方法
>     			func (br *BaseRouter) PostHandle(request ziface.IRequest)
>     	zinx 集成 Router 模块
>     		IServer 增添路由功能 - AddRouter(router IRouter)
>     		Server 类增加 Router 成员 ---> 去掉之前的HandAPI
>     		Connection 类绑定一个 Router 成员
>     		在 Connection 调用已经注册过的 Router 处理业务
>
>     使用zinxV0.3版本开发服务器
>     	1、创建一个服务器句柄
>     	2、给当前的 zinx 框架加一个自定义的业务路由
>     	3、启动 server
>     	4、需要继承 BaseRouter 去实现三个接口的方法
>     当前版本只有一个路由能使用，目前只能使用一个路由模块，在加入路由模块会使上一个路由模块方法进行重写覆盖。
