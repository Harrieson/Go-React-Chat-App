package main

import (
	"fmt"
	"go-reactt-chat/pkg/websocket"
	"net/http"
)

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func main() {
	fmt.Println("Real Chat")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
