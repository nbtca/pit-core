package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nbtca/pit-core/config"
	"github.com/nbtca/pit-core/service"
	"github.com/nbtca/pit-core/util"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	config := config.ReadConfig()
	ip := util.GetLocalIP()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		go service.HandleConnection(conn, r)
	})

	fmt.Println("Server starts at " + ip[len(ip)-1] + ":" + config.SystemConfig.Port + "...")
	err := http.ListenAndServe(":"+config.SystemConfig.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
