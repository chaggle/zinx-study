---
title: "Zinx框架的学习"
date: 2022-01-10T10:14:32+08:00
lastMod: 2022-02-10T16:47:32+08:00
tag: ["Go", "zinx"]
categories: ["Go"]
---

# zinx 框架

开始学习 Go 语言实现的 zinx 框架,项目地址为：[https://github.com/chaggle/zinx-study](https://github.com/chaggle/zinx-study)

> 使用 go mod 管理, 初始化为 go mod init github.com/chaggle/zinx-study
> 并部署代码到 github.com 以及使用 go get 同步到本地 Gopath 的 github 包下！

## V0.1 基础的 server 模块

> 方法
>
> 初始化服务器 -- NewServer(name string) ziface.IServer
>
> 启动服务器 -- Start()
>
> 停止服务器 -- Stop()
>
> 运行服务器 -- Serve()
>
> 属性
>
> 名称 -- name
>
> IP 版本 -- IPVersion
>
> 监听 IP -- IP
>
> 监听端口 -- Port

## V0.2 简单的链接封装和业务绑定

> 方法
>
> 启动链接 -- Start()
>
> 停止链接 -- Stop()
>
> 获取当前链接的 conn 对象 (套接字) -- GetTCPConnection() \*net.TCPConn
>
> 得到链接 ID -- GetConnID() uint32
>
> 得到客户端连接的地址和端口 -- RemoteAddr() net.TCPAddr
>
> 发送数据的方法 -- Send(data []byte) error
>
> 属性
> socket TCP 套接字 -- Conn \*net.TCPConn
>
> 链接的 ID -- ConnID uint32
>
> 当前链接状态 (是否已经关闭) -- isClosed bool
>
> 与当前链接所绑定的处理业务与方法 -- handlerAPI ziface.HandleFunc
>
> 等待退出的 channel 管道 -- ExitChan chan bool

## V0.3 基础的 router 模块

> Request 请求封装
>
> 将链接与数据绑定一起
>
> 属性
> 链接的句柄 -- GetConnection() IConnection
>
> 请求数据 -- GetData() []byte
>
> 方法
> 得到链接-- func (r \*Request) GetConnection() ziface.IConnection
>
> 得到数据 -- func (r \*Request) GetData() []byte
>
> 新建一个 Request 请求
>
> Router 模块
>
> 抽象的 IRouter
>
> 处理业务之前的方法 PreHandle(request IRequest) //处理 conn 业务之前的方法
>
> 处理业务的主方法 Handle(request IRequest) //处理 conn 业务的主方法
>
> 处理业务之后的方法 PostHandle(request IRequest) //处理 conn 业务之后的方法
>
> 具体的 BaseRouter
>
> 处理业务之前的方法 func (br \*BaseRouter) PreHandle(request ziface.IRequest)
>
> 处理业务的主方法 func (br \*BaseRouter) Handle(request ziface.IRequest)
>
> 处理业务之后的方法 func (br \*BaseRouter) PostHandle(request ziface.IRequest)
>
> zinx 集成 Router 模块
>
> IServer 增添路由功能 - AddRouter(router IRouter)
>
> Server 类增加 Router 成员 ---> 去掉之前的 HandAPI
>
> Connection 类绑定一个 Router 成员
>
> 在 Connection 调用已经注册过的 Router 处理业务
>
> 使用 zinxV0.3 版本开发服务器
>
> 1、创建一个服务器句柄
>
> 2、给当前的 zinx 框架加一个自定义的业务路由
>
> 3、启动 server
>
> 4、需要继承 BaseRouter 去实现三个接口的方法
>
> 当前版本只有一个路由能使用，目前只能使用一个路由模块，在加入路由模块会使上一个路由模块方法进行重写覆盖。

## V0.4 全局配置模块

> 路径 ：服务器项目主地址/conf/zinx.json (用户进行填写)
>
> 创建一个 zinx 的全局配置模块 utils/globalobj.go
>
> 初始化后读取用户配置的 zinx.json ---> globalobj 对象中对应的 zinx 服务器句柄代码进行参数替换
>
> 提供一个 GlobalObject 对象 --- var GlobalObject \*GlobalObj
>
> 使用 zinx0.4 版本进行开发

## V0.5 消息封装

> 定义消息的结构 Message
>
> 属性
> 消息的 ID
>
> 消息的长度
>
> 消息的内容
>
> 方法
> SetMsgId(uint32) //设置消息 ID
>
> SetData([]byte) //设置消息内容
>
> SetDataLen(uint32) //设置消息长度
>
> GetDataLen() uint32 //获取消息长度
>
> GetMsgId() uint32 //获取消息 ID
>
> GetData() []byte //获取消息内容
>
> 定义解决 TCP 粘包问题的封包拆包的模块
>
> 针对 Message 进行 TLV 格式的封装 -- func (dp \*DataPack) Pack(msg ziface.IMessage) ([]byte, error)
>
> 写 Message 的长度
>
> 写 Message 的 ID
>
> 写 Message 的内容
>
> 针对 Message 进行 TLV 格式的拆包 -- func (dp \*DataPack) Unpack(binaryData []byte) (ziface.IMessage, error)
>
> 先读取固定长度的 head ---> 消息的长度和消息的类型
>
> 再根据消息内容的长度，再进行一次读写，从 conn 中读取消息的内容
>
> 将消息封装机制集成到 Zinx 框架中
>
> 将 Message 添加到 Request 属性字段
>
> 修改连接读取数据的机制，将之前的单纯读取 byte 改成拆包形式，读取按照 TLV 形式进行读取
>
> 给链接提供一个发包的机制：将发送的消息打包，再发送
>
> 使用 zinxV0.5 开发

## V0.6 多路由模式

> 消息管理模块(支持多路由 API 调度管理)
> 属性
> 集合 - 消息 ID 与对应 router 的关系 - map Apis map[uint32]ziface.IRouter
>
> 方法
>
> ​ 根据 MsgId 来索引调度路由方法 func (mh \*MsgHandle) DoMsgHandler(request ziface.IRequest)
>
> ​ 添加路由方法到 map 集合中 func (mh \*MsgHandle) AddRouter(msgID uint32, router ziface.IRouter)
>
> 消息管理模块集成到 Zinx 框架中
> 将 server 模块里面的 Router 属性变为 MsgHandle 属性
> 将 server 模块中的 AddRouter 修改调用 MsgHandler 的 AddRouter
> 将 connection 模块中的 Router 属性替换为 MsgHandle 属性
> 将 connection 模块中的 Router 业务调度 Router 的业务改为调度 MsgHandle 调度， 并修改 StartRead 方法
>
> 使用 Zinx V0.6 版本开发
