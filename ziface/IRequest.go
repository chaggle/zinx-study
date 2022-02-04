package ziface

/*
	IRequest 接口
	实际上把客户端请求的链接信息和请求的数据包，封装到了一个 Request 中
*/

type IRequest interface {
	GetConnection() IConnection //得到当前的链接

	GetData() []byte //得到请求的数据
}
