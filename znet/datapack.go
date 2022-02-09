package znet

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/chaggle/zinx-study/utils"
	"github.com/chaggle/zinx-study/ziface"
)

//封包、拆包的具体模块
type DataPack struct {
}

//拆包封包的实例的一个初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

//获取包的头部长度的方法
func (dp *DataPack) GetHeadLen() uint32 {
	//DataLen uint32(4 Byte) + Id uint32(4 Byte)
	return 8
}

//封包方法
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//创建一个存放 byte 字节流的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//将datalen写入dataBuff 中
	//二进制写法
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	//将MsgId 写入dataBuff 中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//将data数据写入dataBuff 中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

//拆包方法 (只需要将包的Head读出来)，再根据head里面的data长度再进行一次读取
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	//创建一个从输入二进制数据的 ioReader
	dataBuff := bytes.NewReader(binaryData)

	//只解压 head 的信息，得到 dataLen 跟 msgId
	msg := &Message{}

	//读datalen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	//读msgID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//判断dataLen的长度超出我们允许处理的最大包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too Large msg data recieved")
	}

	//再通过 head 的长度对conn读取一次数据。此处后续回顾发现概念较为模糊！需要进一步加深相应的理解
	return msg, nil
}
