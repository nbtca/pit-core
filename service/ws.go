package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type User struct {
	Name string
	Conn *websocket.Conn
}

type Msg struct {
	SendUser string
	ReceUser string
	SendTime string
	Msg      string
	IsPublic bool
	IsRecv   bool
	IsInfo   bool
}

func (msg *Msg) ParseMessage(message []byte) error {
	fmt.Println(string(message))
	err := json.Unmarshal(message, msg)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func (msg *Msg) EncodeMessage() []byte {
	b, _ := json.Marshal(msg)
	return b
}

var msgs = 
var users = make(map[string]User)

func HandleConnection(conn *websocket.Conn, r *http.Request) {
	defer conn.Close()
	user := User{}
	data := r.FormValue("data")
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		fmt.Println(err)
		conn.WriteMessage(websocket.TextMessage, []byte("连接发生错误"))
		return
	}
	_, u := users[user.Name]
	if u {
		conn.WriteMessage(websocket.TextMessage, []byte("该名称已存在,不允许重复链接"))
		return
	}
	user.Conn = conn
	userOn(user)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("连接已关闭"))
			log.Println(conn.RemoteAddr().String(), "关闭连接", err)
			break
		}
		msg := Msg{}
		err = msg.ParseMessage(message)
		if err != nil {
			log.Println(err)
			break
		}
		if msg.IsPublic {
			sendPublicMessage(msg)
		} else {
			sendPublicMessage(msg)
			sendPrivateMessage(msg)
		}
	}
	userOff(user)
}

func sendPublicMessage(msg Msg) {
	for _, user := range users {
		if user.Conn != nil {
			err := user.Conn.WriteMessage(websocket.TextMessage, msg.EncodeMessage())
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func sendPrivateMessage(msg Msg) {
	for _, user := range users {
		if user.Conn != nil {
			err := user.Conn.WriteMessage(websocket.TextMessage, msg.EncodeMessage())
			if err != nil {
				log.Println(err)
			}
		}
	}

}

func userOn(user User) {
	users[user.Name] = user
	str := fmt.Sprintf("%s 加入聊天室，当前聊天室人数为 %d。", user.Name, len(users))
	fmt.Println(str)
	msg := Msg{
		SendUser: user.Name,
		SendTime: time.Now().Format("2006-01-02 15:04:05"), // 日期格式化为 yyyy-MM-dd HH:mm:ss 格式
		Msg:      str,
		IsPublic: true,
		IsRecv:   true,
		IsInfo:   true,
	}
	sendPublicMessage(msg)
}

func userOff(user User) {
	for _, v := range users {
		if v.Name == user.Name {
			delete(users, v.Name)
			break
		}
	}
	str := fmt.Sprintf("%s 离开了聊天室，当前聊天室人数为 %d。", user.Name, len(users))
	fmt.Println(str)
	// 发送下线消息给其他用户
	msg := Msg{
		SendUser: user.Name,
		SendTime: time.Now().Format("2006-01-02 15:04:05"),
		Msg:      str,
		IsPublic: true,
		IsRecv:   true,
		IsInfo:   true,
	}
	sendPublicMessage(msg)
}
