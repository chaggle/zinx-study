package znet

import (
	"fmt"
	"net"
	"time"

	"github.com/chaggle/zinx-study/utils"
	"github.com/chaggle/zinx-study/ziface"
)

// Server 定义服务器模块
type Server struct {

	//服务器的名称
	Name string

	//服务器的IP版本
	IPVersion string

	//服务器监听的IP号
	IP string

	//服务器监听的端口号
	Port int

	//当前 Server 的消息管理模块，用于绑定 MsgId 和对应的处理方法
	msgHandler ziface.IMsgHandle

	ConnManager ziface.IConnManager
}

//============== 实现 ziface.IServer 里的接口 ========

// Start 启动服务器方法实现
func (s *Server) Start() {

	//打印输出
	fmt.Printf("[START] Server listenner at IP : %s, Port : %d, is starting\n",
		utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx] Version : %s, MaxConn : %d, MaxPackageSize : %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPackageSize)

	go func() {
		//0 先启动 worker 工作池机制
		go s.msgHandler.StartWorkerPool()

		//1 获取一个TCP的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err :", err)
			return
		}

		//2 监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
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
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error :", err)
				continue //因为在一个 goroutine 里面
			}

			//如果超出了最大的链接数，则拒绝启动链接
			if s.GetConnManager().Len() > utils.GlobalObject.MaxConn {
				_ = conn.Close()
				continue
			}
			// 将处理新连接的业务方法 和 conn 进行绑定， 得到我们的链接模块
			dealConn := NewConnection(s, conn, cid, s.msgHandler)
			cid++

			//启动当前的链接业务
			go dealConn.Start()
		}
	}()
}

// Stop 停止服务器方法实现
func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server, name :", s.Name)

	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.GetConnManager().ClearConn()
}

// Serve 运行服务器方法实现
func (s *Server) Serve() {
	s.Start()

	//阻塞,否则主 Goroutine 退出， listenner的go将会退出
	for {
		time.Sleep(10 * time.Second)
	}
	//select {}
}

// AddRouter 路由功能，给当前服务器注册一个路由方法，供客户端链接使用
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(msgID, router)
	fmt.Println("Add Router Success")
}

func (s *Server) GetConnManager() ziface.IConnManager {
	return s.ConnManager
}

// NewServer 初始化服务器方法实现
func NewServer() ziface.IServer {
	//先初始化全局配置文件
	utils.GlobalObject.Reload()

	s := &Server{
		Name:        utils.GlobalObject.Name,
		IPVersion:   "tcp4", // 还没有实现 tcp6 版本
		IP:          utils.GlobalObject.Host,
		Port:        utils.GlobalObject.TcpPort,
		msgHandler:  NewMsgHandle(), //MsgHandler 初始化
		ConnManager: NewConnManager(),
	}

	return s
}
