package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clientsMutex sync.Mutex
var rooms = make(map[string]map[*websocket.Conn]bool)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/send/:room", func(c *gin.Context) {
		room := c.Param("room")
		message := c.PostForm("message")

		broadcastMessage(room, message, nil)
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "message sent",
		})
	})

	r.Run(":8080")
}

func broadcastMessage(room, msg string, sender *websocket.Conn) {

	for client := range rooms[room] {
		// 将消息发送给同一房间的其他客户端
		if client != sender {
			err := client.WriteJSON(gin.H{
				"from":    "server",
				"message": msg,
			})
			if err != nil {
				return
			}
		}
	}
}
