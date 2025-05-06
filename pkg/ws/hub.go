package ws

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	mu     sync.RWMutex
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
	h.conns[taskID] = conn

	if msg, ok := h.buffer[taskID]; ok {
		conn.WriteJSON(msg)
		delete(h.buffer, taskID)
	}
}

func (h *Hub) Remove(taskID string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.conns, taskID)
}

func (h *Hub) Notify(taskID string, payload any) error {
	h.mu.RLock()
	defer h.mu.RUnlock()
	conn, ok := h.conns[taskID]
	if ok {
		return conn.WriteJSON(payload)
	} else {
		h.buffer[taskID] = payload
		return nil
	}
}
