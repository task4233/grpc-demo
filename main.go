package main

import (
	"fmt"
	"net"
	"os"

	"github.com/task4233/grpc-demo/adder"
	"google.golang.org/grpc"
)

const PORT = ":8080"

func main() {
	s := grpc.NewServer()

	adderServer := &adder.AdderServer{}
	adder.RegisterAdderServiceServer(s, adderServer)

	conn, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to listen: %s", err.Error())
	}
	defer conn.Close()

	fmt.Printf("Server ready with %s\n", PORT)
	s.Serve(conn)
}
