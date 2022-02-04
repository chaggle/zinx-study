package znet

import "github.com/chaggle/zinx-study/ziface"

// 实现 router 时，先使用此基类，然后进行isx-
type BaseRouter struct {
}

//BaseRouter 方法为空，目的是为了使客户端自行实现相关的业务

//处理conn业务之前的方法
func (br *BaseRouter) PreHandle(request ziface.IRequest) {}

//处理conn业务的主方法
func (br *BaseRouter) Handle(request ziface.IRequest) {}

//处理conn业务之后的方法
func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
