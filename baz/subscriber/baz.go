package subscriber

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	baz "go-micro-echo/baz/proto/baz"
)

type Baz struct{}

func (e *Baz) Handle(ctx context.Context, msg *baz.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *baz.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
