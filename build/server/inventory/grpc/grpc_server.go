package grpc

import (
	"net"

	"google.golang.org/grpc"
)

func Start() error {
	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	return grpcServer.Serve(l)
}
