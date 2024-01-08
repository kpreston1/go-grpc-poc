package main

import (
	"fmt"
	"net"
	"log"
	"context"
	"github.com/kpreston1/go-grpc-poc/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf: []byte(req.From),
		Docx: []byte(req.To),
	}, nil
}

func main() {
	fmt.Println("Hello, world!")

	listener, error := net.Listen("tcp", ":8080")
	if error != nil {
		log.Fatalf("cannot create listener: %s", error)
	}

	serviceRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serviceRegistrar, service)
	error = serviceRegistrar.Serve(listener)
	if error != nil {
		log.Fatalf("cannot serve listener: %s", error)
	}	
}
