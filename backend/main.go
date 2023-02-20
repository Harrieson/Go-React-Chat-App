package main

import (
	"fmt"
	"net/http"
)

func setupRoutes() {
	pool := websocket.NewPool()
}

func main() {
	fmt.Println("Real Chat")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
