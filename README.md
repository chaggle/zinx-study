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
