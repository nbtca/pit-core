package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nbtca/pit-core/service"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	//设置路由
	http.HandleFunc("/", handler)
	//监听端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//处理消息
	go service.Loop(conn)
}
