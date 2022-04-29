package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/chaggle/zinx-study/ziface"
)

/*
	存储一切有关于 zinx 框架的全局参数，以供其他模块进行使用
	有些参数也是能通过zinx.json由用户进行配置的
*/
type GlobalObj struct {
	/*
		server 配置
	*/

	//zinx 全局的 Server 对象
	TcpServer ziface.IServer

	//当前服务器主机监听IP
	Host string

	//当前服务器主机监听端口
	TcpPort int

	//当前服务器名称
	Name string

	/*
		zinx 配置
	*/

	//当前 zinx 版本号
	Version string

	//当前服务器允许的最大链接数
	MaxConn int

	//当前zinx框架数据包的最大值
	MaxPackageSize uint32

	//业务工作 Work 池的对应Goroutine数量
	WorkPoolSize uint32

	//业务工作 Work 池的对应负责的消息队列的最大存储数量
	MaxWorkerTaskLen uint32

	//业务工作 buffChannel 最大的存储数
	MaxMsgChanLen uint32
}

/*
	定义全局的对象
*/

var GlobalObject *GlobalObj

/*
	从 zinx.json 加载配置文件
    优化
    1、可以进一步进行配置的热加载，实现动态修改配置文件
    2、配置保存全局变量，为了模拟业务场景，存放服务器的配置数据，是否能存到
       redis 这样的内存缓存数据库中，虽然有点感觉是浪费，但是也能考虑。
*/
func (g *GlobalObj) Reload() {

	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	//将 json 数据解析到 struct 中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		/*panic(err)*/ //解析配置文件都失败了，当然没必要执行下去。
		log.Fatal("解析配置文件出错！")
	}
}

func init() {
	//配置文件未加载时候的一些默认值
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerApp",
		Version:          "V0.4",
		TcpPort:          7777,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkPoolSize:     10,
		MaxWorkerTaskLen: 1024,
		MaxMsgChanLen:    1024,
	}

	//V0.5版本进行注释掉！
	GlobalObject.Reload()
}
