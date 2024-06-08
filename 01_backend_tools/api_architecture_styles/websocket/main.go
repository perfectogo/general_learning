package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {

	router := gin.Default()

	router.GET("/ws", func(c *gin.Context) {

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		for {
			conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
			time.Sleep(time.Second)
		}
	})

	router.GET("/ws-s", func(c *gin.Context) {

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		i := 0
		for {
			i++
			conn.WriteMessage(websocket.TextMessage, []byte("New message (#"+strconv.Itoa(i)+")"))
			time.Sleep(time.Second)
			t, d, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("error in here:", err)
			}

			fmt.Println(t, string(d))
		}
	})

	router.Run(":8090")
}


