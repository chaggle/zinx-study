package znet

import (
	"fmt"
	"strconv"

	"github.com/chaggle/zinx-study/ziface"
)

/*
	消息处理模块的实现
*/

type MsgHandle struct {
	//存放每个 MsgID 所对应的处理方法
	Apis map[uint32]ziface.IRouter
}

//创建 MsgHandle 的方法
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

//调度/执行对应 Router 消息处理方法，马上以非阻塞的方法处理消息
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

//为消息添加具体的业务逻辑
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
