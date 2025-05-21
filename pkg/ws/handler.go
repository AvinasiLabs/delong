package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var allowedOrigins = map[string]bool{
	"http://localhost:3000":  true,
	"https://somedomain.com": true,
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// prod
		// origin := r.Header.Get("Origin")
		// return allowedOrigins[origin]

		return true // test
	},
}

func NewHandler(hub *Hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		taskID := c.Query("task_id") // wss://yourdomain/ws?task_id=0xtxabc...
		if taskID == "" {
			c.JSON(400, gin.H{"error": "task_id is required"})
			return
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("Failed to upgrade connection: %v", err)
			return
		}

		hub.Register(taskID, conn)

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				_ = conn.Close()
				hub.mu.Lock()
				delete(hub.conns, taskID)
				delete(hub.buffer, taskID)
				hub.mu.Unlock()
				return
			}
		}
	}
}
