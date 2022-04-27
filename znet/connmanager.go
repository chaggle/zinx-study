package znet

import (
	"errors"
	"fmt"
	"github.com/chaggle/zinx-study/ziface"
	"sync"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection
	connLock    sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection, 0),
	}
}

func (cm *ConnManager) Add(conn ziface.IConnection) {
	//添加操作加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	cm.connections[conn.GetConnID()] = conn
	fmt.Println("connection add to ConnManager successfully: conn num = ", cm.Len())
}
func (cm *ConnManager) Remove(conn ziface.IConnection) {
	//移除操作加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	delete(cm.connections, conn.GetConnID())

	fmt.Println("connection Remove ConnID=", conn.GetConnID(), " successfully: conn num = ", cm.Len())
}
func (cm *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	//读取操作加读锁
	cm.connLock.RLock()
	defer cm.connLock.RLock()

	if conn, ok := cm.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}
func (cm *ConnManager) Len() int {
	return len(cm.connections)
}
func (cm *ConnManager) ClearConn() {
	//保护共享资源 map 加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	for connID, conn := range cm.connections {
		//清空所有的conn链接，然后删除
		conn.Stop()

		delete(cm.connections, connID)
	}

	fmt.Println("Clear All Connections successfully: conn num = ", cm.Len())
}
