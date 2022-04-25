package znet

import (
	"fmt"
	"strconv"

	"github.com/chaggle/zinx-study/utils"
	"github.com/chaggle/zinx-study/ziface"
)

/*
	消息处理模块的实现
*/

type MsgHandle struct {
	//存放每个 MsgID 所对应的处理方法
	Apis map[uint32]ziface.IRouter

	//业务 work 工作池的数量
	WorkPoolSize uint32

	//Work 负责取任务的消息队列
	TaskQueue []chan ziface.IRequest
}

// NewMsgHandle 创建 MsgHandle 的方法
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:         make(map[uint32]ziface.IRouter),
		WorkPoolSize: utils.GlobalObject.WorkPoolSize,
		TaskQueue:    make([]chan ziface.IRequest, utils.GlobalObject.WorkPoolSize),
	}
}

// DoMsgHandler 调度或执行对应 Router 消息处理方法，马上以非阻塞的方法处理消息
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	//1 从 Request 中找到 msgID
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID = ", request.GetMsgID(), "is NOT FOUND! Need Register!")
	}

	//2 根据 MsgID 调度对应的 router 业务即可
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)

}

// AddRouter 为消息添加具体的业务逻辑
func (mh *MsgHandle) AddRouter(msgID uint32, router ziface.IRouter) {
	//1 判断当前 msg 绑定的 API 处理方法是否存在
	if _, ok := mh.Apis[msgID]; ok {
		//id 已经注册
		panic("repeated api, msgID = " + strconv.Itoa(int(msgID)))
	}

	//2 添加 msg 与 API 的绑定关系
	mh.Apis[msgID] = router
	fmt.Println("Add api MsgId = ", msgID)
}

// StartOneWorker 启动一个 Worker 的工作流程
func (mh *MsgHandle) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	fmt.Println("WorkID = ", workerID, "is started !")
	//不断等待队列中的消息
	for {
		select {
		//一有消息就取出队列的Request，并执行绑定的业务方法
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

// StartWorkerPool 启动 Work 工作池
func (mh *MsgHandle) StartWorkerPool() {
	//遍历需要启动的 worker 的数量，依次进行启动！
	for i := 0; i < int(mh.WorkPoolSize); i++ {
		//一个 worker 启动
		//给当前的 worker 开辟空间
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)

		//启动当前的 worker，阻塞的等待对应的任务队列是否有消息传递进来！
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}

}

// SendMsgToTaskQueue 将消息传递给 TaskQueue，由 work 进行处理
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	//根据 ConnID 来分配当前的连接应该由哪个 worker 负责
	//轮询的平均分配法则

	//得到需要处理此条链接的workID
	workID := request.GetConnection().GetConnID() % mh.WorkPoolSize
	fmt.Println("Add ConnID = ", request.GetConnection().GetConnID(),
		"request msgID = ", request.GetMsgID(), "to workerID = ", workID)

	mh.TaskQueue[workID] <- request

}
