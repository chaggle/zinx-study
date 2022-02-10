package ziface

/*
	消息管理抽象层
*/

type IMsgHandle interface {

	//调度/执行对应 Router 消息处理方法，马上以非阻塞的方法处理消息
	DoMsgHandler(request IRequest)

	//为消息添加具体的业务逻辑
	AddRouter(msgId uint32, router IRouter)
}
