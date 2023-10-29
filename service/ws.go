package service

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type User struct {
	id int
}

var connections = make(map[*websocket.Conn]bool)

func processMessage(msg string) {
	for conn := range connections {
		conn.WriteMessage(websocket.TextMessage, []byte("hello"))
	}
}

func open(conn *websocket.Conn) {
	//添加连接
	connections[conn] = true
}
func close(conn *websocket.Conn) {
	//删除连接
	delete(connections, conn)
}
func Loop(conn *websocket.Conn) {
	open(conn)
	defer close(conn)
	//循环读取消息
	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return //直到出错才退出循环
		}
		// if err := conn.WriteMessage(messageType, data); err != nil {
		// 	log.Println(err)
		// 	return
		// }
		fmt.Println(messageType)
		fmt.Println(messageType)
		// log.Println(string(data))
		processMessage(string(data))
	}
}
