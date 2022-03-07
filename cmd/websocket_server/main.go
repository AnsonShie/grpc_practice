package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Request struct {
	Num int32 `json:"num"`
}

type Response struct {
	Sum int32 `json:"sum"`
}

func main() {
	g := gin.New()
	g.Use(gin.Recovery())
	g.GET("/echo", echo)

	srv := &http.Server{
		Addr:    ":8899",
		Handler: g,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Println("server start at :8899")

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v \n", err)
	}

	log.Println("Server shutdown.")
}

func echo(ctx *gin.Context) {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer func() {
		log.Println("disconnect !!")
		c.Close()
	}()
	var sum int32
	for {
		var req Request
		err := c.ReadJSON(&req)
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("receive: %v\n", req)
		sum += req.Num
		resp := Response{
			Sum: sum,
		}
		err = c.WriteJSON(resp)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
