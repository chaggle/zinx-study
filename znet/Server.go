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
	Name      string //服务器的名称
	IPVersion string //服务器的IP版本
	IP        string //服务器监听的IP号
	Port      int    //服务器监听的端口号
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

		//3 阻塞等待客户端的连接，处理（读写）客户端链接的业务
		//  启动 server 的网络链接业务

		for {
			//阻塞等待客户端的建立连接的请求，如果有客户端链接，阻塞返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error :", err)
				continue //因为在一个 goroutine 里面
			}

			//TODO Server.Start() 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接

			//TODO Server.Start() 处理该新连接请求的业务方法， 此时应该有 handler 和 conn是绑定的

			//v0.1版本实现一个回显的业务即可
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("receive buf error", err)
						continue
					}

					//函数回显
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back err", err)
						continue
					}
				}
			}()
		}
	}()
}

//停止服务器方法实现
func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server, name :", s.Name)

	//TODO Server.Stop() 进行服务器关闭连接清理的相应的业务
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

//初始化服务器方法实现
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "127.0.0.1",
		Port:      8999,
	}

	return s
}
