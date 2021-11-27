package grpc

import (
	_interface "bookstore/internal/app/domain/inventory/books/module/interface"
	"bookstore/internal/app/domain/server"
	"net"

	"google.golang.org/grpc"
)

func Start(bookModule _interface.BookModuleInterface) error {
	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	server.RegisterOrdersServiceServer(grpcServer, bookModule)
	return grpcServer.Serve(l)
}
