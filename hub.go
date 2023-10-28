package main

import "encoding/json"

var h = hub{
	clients:  make(map[*connection]bool),
	user:     make(chan *connection),
	data:     make(chan []byte),
	received: make(chan *connection),
}

type hub struct {
	clients  map[*connection]bool
	data     chan []byte
	received chan *connection
	user     chan *connection
}

func (h *hub) run() {
	for {
		select {
		case cc := <-h.received:
			h.clients[cc] = true
			cc.data.Ip = cc.ws.RemoteAddr().String()
			cc.data.Type = "handshake"
			cc.data.UserList = user_list
			data_b, _ := json.Marshal(cc.data)
			cc.sc <- data_b
		case c := <-h.user:
			if _, ok := h.clients[c]; ok {
				delete(h.clients, c)
				close(c.sc)
			}
		case data := <-h.data:
			for c := range h.clients {
				select {
				case c.sc <- data:
				default:
					delete(h.clients, c)
					close(c.sc)
				}
			}
		}
	}
}
