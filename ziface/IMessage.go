package ziface

/*
	将请求消息封装到一个 Message 中，定义抽象消息接口
*/

type IMessage interface {
	//获取消息长度
	GetDataLen() uint32
	//获取消息ID
	GetMsgId() uint32
	//获取消息内容
	GetData() []byte
	//设置消息ID
	SetMsgId(uint32)
	//设置消息内容
	SetData([]byte)
	//设置消息长度
	SetDataLen(uint32)
}
