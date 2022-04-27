package ziface

// IConnManager  conn 链接管理
type IConnManager interface {
	Add(conn IConnection)
	Remove(conn IConnection)
	Get(connID uint32) (IConnection, error)
	Len() int   //获取当前链接
	ClearConn() //清除所有的 Conn 链接
}
