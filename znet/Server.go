package znet

import (
	"fmt"
	"net"
	"time"

	"github.com/chaggle/zinx-study/ziface"
)

//定义服务器模块
type Server struct {
	//相应服务器处理业务需要接受的信息
	Name string //服务器的名称

	IPVersion string //服务器的IP版本

	IP string //服务器监听的IP号

	Port int //服务器监听的端口号

	Router ziface.IRouter //从当前的Server添加一个router，server注册的链接对应的处理业务
}

//============== 实现 ziface.IServer 里的接口 ========

//启动服务器方法实现
func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP : %s, Port : %d, is starting\n", s.IP, s.Port)
	go func() {
		//1 获取一个TCP的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err :", err)
			return
		}

		//2 监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IP, "error :", err)
			return
		}

		fmt.Println("start zinx server", s.Name, "success, now listenning")
		var cid uint32
		cid = 0
		//3 阻塞等待客户端的连接，处理（读写）客户端链接的业务
		//  启动 server 的网络链接业务

		for {
			//阻塞等待客户端的建立连接的请求，如果有客户端链接，阻塞返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error :", err)
				continue //因为在一个 goroutine 里面
			}

			// 将处理新连接的业务方法 和 conn 进行绑定， 得到我们的链接模块
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			//启动当前的链接业务
			go dealConn.Start()
		}
	}()
}

//停止服务器方法实现
func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server, name :", s.Name)
}

//运行服务器方法实现
func (s *Server) Serve() {
	s.Start()

	//阻塞,否则主 Goroutine 退出， listenner的go将会退出
	for {
		time.Sleep(10 * time.Second)
	}
	//select {}
}

//路由功能，给当前服务器注册一个路由方法，供客户端链接使用
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Success")
}

//初始化服务器方法实现
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "127.0.0.1",
		Port:      8999,
		Router:    nil,
	}

	return s
}
