package ziface

/*
	路由抽象接口(接口都是 interface 类型，注意别写成 struct 类)
	路由藜麦数据都是IRequest
*/

type IRouter interface {
	PreHandle(request IRequest) //处理conn业务之前的方法

	Handle(request IRequest) //处理conn业务的主方法

	PostHandle(request IRequest) //处理conn业务之后的方法

}
