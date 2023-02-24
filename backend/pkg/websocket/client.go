package websocket
import (
	"go-react-chat/pkg/websocket"
	"github.com/gorilla/websocket"
	"fmt"
	"log"
	"sync"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
	mu   sync.Mutex
}

type Message struct{
	Type int `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister -< c
		c.Conn.Close()
	}()
	for {
	messageType, p, err := c.Conn.ReadMessage()
	if err !=nil{
		log.Println(err)
		return
	}
	message := Message{Type:messageType, Body: string(p)}
	c.Pool.Broadcast <-message
	fmt.Printf("Message received:%+v\n", message)
	}
}