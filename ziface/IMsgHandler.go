package ziface

/*
	消息管理抽象层
*/

type IMsgHandle interface {

	// DoMsgHandler 调度/执行对应 Router 消息处理方法，马上以非阻塞的方法处理消息
	DoMsgHandler(request IRequest)

	// AddRouter 为消息添加具体的业务逻辑
	AddRouter(msgId uint32, router IRouter)

	// StartWorkerPool 启动 Work 工作池
	StartWorkerPool()

	// SendMsgToTaskQueue 将消息传递给 TaskQueue，由 work 进行处理
	SendMsgToTaskQueue(request IRequest)
}
