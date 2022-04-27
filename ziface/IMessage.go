package ziface

/*
	将请求消息封装到一个 Message 中，定义抽象消息接口
*/

type IMessage interface {
	GetDataLen() uint32
	GetMsgId() uint32
	GetData() []byte

	SetMsgId(uint32)
	SetData([]byte)
	SetDataLen(uint32)
}
