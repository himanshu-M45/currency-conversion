package server

import (
	"currency-conversion/internal/app/cc/controllers"
	"currency-conversion/internal/app/cc/helper"
	"currency-conversion/internal/app/cc/service"
	"currency-conversion/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

func RunServer() {
	// gRPC server setup
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterCurrencyConversionServer(grpcServer, &controllers.Server{})

	log.Printf("Server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func init() {
	var err error
	service.ConversionRates, err = helper.LoadConversionRates("conversion_rates.xml")
	fmt.Println("Successfully loaded conversion rates")
	if err != nil {
		log.Fatalf("failed to load conversion rates: %v", err)
	}
}
