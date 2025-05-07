package ws

type Notifier struct {
	hub *Hub
}

func NewNotifier(hub *Hub) *Notifier {
	return &Notifier{hub: hub}
}

func (n *Notifier) Hub() *Hub {
	return n.hub
}

func (n *Notifier) PushStatus(txHash string, status string) error {
	err := n.hub.Notify(txHash, map[string]any{
		"type":    "tx_status",
		"status":  status,
		"tx_hash": txHash,
	})
	if err == nil {
		n.hub.Remove(txHash)
	}
	return err
}
