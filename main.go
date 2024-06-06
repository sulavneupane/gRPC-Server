package main

import (
	"context"
	"github.com/sulavneupane/gRPC-Server/invoicer"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (server myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	myInvoicerService := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegistrar, myInvoicerService)
	err = serverRegistrar.Serve(listener)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
