package adder

import (
	"context"
	"fmt"
)

type AdderServer struct{}

func (s *AdderServer) Add(ctx context.Context, in *NumMessage) (*NumMessage, error) {
	fmt.Printf("[SERVER] %d\n", in.Value)

	return &NumMessage{
		Value: in.Value + 1, // 1を加算
	}, nil
}
