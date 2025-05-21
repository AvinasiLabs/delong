package ws

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	mu     sync.Mutex
	conns  map[string]*websocket.Conn // taskID -> conn
	buffer map[string]any             // taskID -> pending register req (if not yet registered)
}

func NewHub() *Hub {
	return &Hub{
		conns:  make(map[string]*websocket.Conn),
		buffer: make(map[string]any),
	}
}

func (h *Hub) Register(taskID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if msg, ok := h.buffer[taskID]; ok {
		conn.WriteJSON(msg)
		delete(h.buffer, taskID)

		err := conn.Close()
		if err != nil {
			log.Printf("Failed to close connection: %v", err)
		}
	} else {
		h.conns[taskID] = conn
	}
}

func (h *Hub) Notify(taskID string, payload any) {
	h.mu.Lock()
	defer h.mu.Unlock()
	conn, ok := h.conns[taskID]
	if ok {
		err := conn.WriteJSON(payload)
		if err != nil {
			log.Printf("Failed to write msg: %v", err)
		}
		err = conn.Close()
		if err != nil {
			log.Printf("Failed to close connection: %v", err)
		}
		delete(h.conns, taskID)
	} else {
		h.buffer[taskID] = payload
	}
}
