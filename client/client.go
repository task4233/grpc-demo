package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/task4233/grpc-demo/adder"
	"google.golang.org/grpc"
)

const PORT = "127.0.0.1:8080"

func NewConn(host string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if host != "" {
		opts = append(opts, grpc.WithAuthority(host))
	}

	// 今回はinsercureな設定で進めます
	opts = append(opts, grpc.WithInsecure())
	return grpc.Dial(host, opts...)
}

func AddRequest(conn *grpc.ClientConn, req *adder.NumMessage) (*adder.NumMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client := adder.NewAdderServiceClient(conn)
	return client.Add(ctx, req)
}

func main() {
	conn, err := NewConn(PORT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed NewConn: %s", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	num := &adder.NumMessage{
		Value: 57,
	}
	fmt.Printf("[CLIENT] %d\n", num.Value)

	gotNum, err := AddRequest(conn, num)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed AddRequest: %s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("[CLIENT] %d\n", gotNum.Value)
}
