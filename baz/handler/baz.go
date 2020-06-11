package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	baz "go-micro-echo/baz/proto/baz"
)

type Baz struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Baz) Call(ctx context.Context, req *baz.Request, rsp *baz.Response) error {
	log.Info("Received Baz.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Baz) Stream(ctx context.Context, req *baz.StreamingRequest, stream baz.Baz_StreamStream) error {
	log.Infof("Received Baz.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&baz.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Baz) PingPong(ctx context.Context, stream baz.Baz_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&baz.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
