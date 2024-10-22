package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	//pb "path/to/generated/proto"
	pb "github.com/himanshu-m45/currency-conversion/proto"
)

type server struct {
	pb.UnimplementedCurrencyConverterServer
}

func (s *server) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	senderCurrency := req.GetSenderCurrencyType()
	receiverCurrency := req.GetReceiverCurrencyType()
	amount := req.GetAmount()

	// Implement the conversion logic here
	convertedAmount := convertCurrency(senderCurrency, receiverCurrency, amount)

	return &pb.ConvertResponse{ConvertedAmount: convertedAmount}, nil
}

func convertCurrency(senderCurrency, receiverCurrency string, amount float64) float64 {
	// Implement your conversion logic here
	// This is a placeholder implementation
	return amount * 1.1
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCurrencyConverterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
