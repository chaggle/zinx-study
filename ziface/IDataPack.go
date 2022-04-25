package ziface

/*
	封包、拆包模块
	直接面向 TCP 连接中的数据流，进行 TCP 粘包处理
*/

type IDataPack interface {
	// GetHeadLen 获取包的头部长度的方法
	GetHeadLen() uint32
	// Pack 封包方法
	Pack(msg IMessage) ([]byte, error)
	// Unpack 封包方法
	Unpack([]byte) (IMessage, error)
}
