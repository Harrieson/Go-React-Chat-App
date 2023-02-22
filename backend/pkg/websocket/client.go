package websocket

import (
	"go-reactt-chat/pkg/websocket"
	"sync"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
	mu   sync.Mutex
}

func (c *Client) Read() {

}
