package utils

import (
	"encoding/json"
	"io/ioutil"

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
	TcpServer ziface.IServer //zinx 全局的 Server 对象
	Host      string         //当前服务器主机监听IP
	TcpPort   int            //当前服务器主机监听端口
	Name      string         //当前服务器名称

	/*
		zinx 配置
	*/
	Version        string //当前 zinx 版本号
	MaxConn        int    //当前服务器允许的最大链接数
	MaxPackageSize uint32 //当前zinx框架数据包的最大值
}

/*
	定义全局的对象
*/

var GlobalObject *GlobalObj

/*
	从 zinx.json 加载配置文件
*/
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	//将 json 数据解析到 struct 中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err) //没必要执行下去，直接 panic 返回即可
	}
}

func init() {
	//配置文件未加载时候的一些默认值
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "V0.4",
		TcpPort:        7777,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	GlobalObject.Reload()
}
