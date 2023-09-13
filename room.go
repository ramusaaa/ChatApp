package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type room struct {
	clients map[*client]bool

	//Odaya girişler
	join chan *client

	//Odadan çıkışlar
	leave chan *client

	forward chan []byte
}

func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool)}
}

func (r *room) run() {

	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.receive)

		case msg := <-r.forward:
			for client := range r.clients {
				client.receive <- msg
			}

		}
	}

}

const (
	socketBuffersize  = 1024
	messageBuffersize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBuffersize, WriteBufferSize: socketBuffersize}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)

	if err != nil {
		log.Fatal("Servehttp:", err)
		return
	}
	client := &client{
		socket:  socket,
		receive: make(chan []byte, messageBuffersize),
		room:    r,
	}
	r.join <- client
	defer func() { r.leave <- client }()

	go client.write()
	client.read()
}
