package service

import (
	"log"

	"github.com/gorilla/websocket"
)
struct User {
	id int
}

var connections = make(map[*websocket.Conn]int)

func processMessage(conn *websocket.Conn, msg string) {
	conn.WriteMessage(websocket.TextMessage, []byte("hello"))
}

func open(conn *websocket.Conn) {
	//添加连接
	connections[conn] = 1

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
			return//直到出错才退出循环
		}
		// if err := conn.WriteMessage(messageType, data); err != nil {
		// 	log.Println(err)
		// 	return
		// }
		log.Println(string(data))
		processMessage(conn, string(data))
	}
}
