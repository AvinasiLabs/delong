package ws

import (
	"delong/internal/models"
	"delong/pkg/bizcode"
	"delong/pkg/responser"
)

type Notifier struct {
	hub *Hub
}

func NewNotifier(hub *Hub) *Notifier {
	return &Notifier{hub: hub}
}

func (n *Notifier) Hub() *Hub {
	return n.hub
}

// func (n *Notifier) PushStatus(txHash string, status string) error {
// 	err := n.hub.Notify(txHash, map[string]any{
// 		"type":    "tx_status",
// 		"status":  status,
// 		"tx_hash": txHash,
// 	})
// 	if err == nil {
// 		n.hub.Remove(txHash)
// 	}
// 	return err
// }

func (n *Notifier) PushError(txHash string, code bizcode.Code) {
	respmsg := responser.Response{
		Code: code,
	}
	n.hub.Notify(txHash, respmsg)
}

func (n *Notifier) PushTxResult(txHash string, txResult *models.BlockchainTransaction) {
	respmsg := responser.Response{
		Code: bizcode.SUCCESS,
		Data: txResult,
	}
	n.hub.Notify(txHash, respmsg)
}
