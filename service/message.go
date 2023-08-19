package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sync"
	"tiktok/pjdata"
)

// 线程安全映射类型，可以保证修改map时的并发安全
var chatConnMap = sync.Map{}

func RunMessageServer() {

	//创建tcp监听器
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("Run message sever failed: %v\n", err)
		return
	}

	for {

		//conn为tcp连接对象
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Accept conn failed: %v\n", err)
			continue
		}

		//成功接受连接，调用函数process，利用协程并发处理多个连接
		go process(conn)
	}

}

func process(conn net.Conn) {
	defer conn.Close()

	var buf [256]byte
	for {

		//读取内容到整个buf数组。n为字节数
		n, err := conn.Read(buf[:])

		//字节数为零：1.文件读完了  2.发生了读取错误
		if n == 0 {
			if err == io.EOF {
				break
			}
			fmt.Printf("Read message failed: %v\n", err)
			continue
		}

		//将接收到的json数据转换为结构体MessageSendEvent
		event := pjdata.MessageSendEvent{}
		_ = json.Unmarshal(buf[:n], &event)
		fmt.Printf("Receive Message：%+v\n", event)

		//发送消息方（如果无消息内容仅仅代表处于连接状态）
		fromChatKey := fmt.Sprintf("%d_%d", event.UserId, event.ToUserId)

		//len为0，表示无消息内容
		if len(event.MsgContent) == 0 {
			chatConnMap.Store(fromChatKey, conn)
			continue
		}

		//查看接收消息方是否处于连接状态
		toChatKey := fmt.Sprintf("%d_%d", event.ToUserId, event.UserId)
		writeConn, exist := chatConnMap.Load(toChatKey)

		//exist表示map中是否存在toChatKey的值
		if !exist {
			fmt.Printf("User %d offline\n", event.ToUserId)
			continue
		}

		//发送消息
		pushEvent := pjdata.MessagePushEvent{
			FromUserId: event.UserId,
			MsgContent: event.MsgContent,
		}

		//将结构体转换为json格式
		pushData, _ := json.Marshal(pushEvent)

		//此处为类型断言，因为sync.Map的值是interface{}，其可以表示任意的值
		_, err = writeConn.(net.Conn).Write(pushData)
		if err != nil {
			fmt.Printf("Push message failed: %v\n", err)
		}
	}
}
