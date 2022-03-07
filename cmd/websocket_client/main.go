package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

type Request struct {
	Num int32 `json:"num"`
}

type Response struct {
	Sum int32 `json:"sum"`
}

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:8899/echo", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		default:
			rand.Seed(time.Now().UnixNano())
			req := Request{
				Num: rand.Int31n(1000) + 5,
			}
			err = c.WriteJSON(req)
			if err != nil {
				log.Println(err)
				return
			}
			var resp Response
			err = c.ReadJSON(&resp)
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("receive: %v\n", resp)
			time.Sleep(time.Millisecond * 500)
		}
	}
}
