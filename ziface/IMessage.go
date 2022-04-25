package ziface

/*
	将请求消息封装到一个 Message 中，定义抽象消息接口
*/

type IMessage interface {
	// GetDataLen 获取消息长度
	GetDataLen() uint32
	// GetMsgId 获取消息ID
	GetMsgId() uint32
	// GetData 获取消息内容
	GetData() []byte
	// SetMsgId 设置消息ID
	SetMsgId(uint32)
	// SetData 设置消息内容
	SetData([]byte)
	// SetDataLen 设置消息长度
	SetDataLen(uint32)
}
