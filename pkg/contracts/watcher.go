package contracts

import (
	"context"
	"log"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func WatchEventLoop[T any](
	ctx context.Context,
	watchFunc func(*bind.WatchOpts, chan T) (ethereum.Subscription, error),
	handle func(T),
) {
	ch := make(chan T)

	opts := &bind.WatchOpts{Context: ctx}
	sub, err := watchFunc(opts, ch)
	if err != nil {
		log.Printf("Subscription failed: %v", err)
		return
	}

	go func() {
		defer sub.Unsubscribe()

		for {
			select {
			case <-ctx.Done():
				log.Println("context done, stop listening")
				return
			case evt := <-ch:
				handle(evt)
			case err := <-sub.Err():
				log.Printf("Subscription error: %v", err)
				return
			}
		}
	}()
}
