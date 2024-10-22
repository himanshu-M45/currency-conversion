package main

import (
	"context"
	pb "currency-conversion/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Conversion rates
var conversionRates = map[string]map[string]float64{
	"USD": {"EUR": 0.85, "INR": 74.0},
	"EUR": {"USD": 1.18, "INR": 87.0},
	"INR": {"USD": 0.013, "EUR": 0.011},
}

type server struct {
	pb.UnimplementedCurrencyConverterServer
}

func (s *server) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	senderCurrency := req.GetSenderCurrencyType()
	receiverCurrency := req.GetReceiverCurrencyType()
	amount := req.GetAmount()

	// Call the conversion logic
	convertedAmount, err := convertCurrency(senderCurrency, receiverCurrency, amount)
	if err != nil {
		return nil, err
	}

	return &pb.ConvertResponse{ConvertedAmount: convertedAmount}, nil
}

func convertCurrency(senderCurrency, receiverCurrency string, amount float64) (float64, error) {
	// Check if senderCurrency exists in the conversionRates
	if rates, ok := conversionRates[senderCurrency]; ok {
		if rate, ok := rates[receiverCurrency]; ok {
			return amount * rate, nil
		}
		return 0, fmt.Errorf("unsupported receiver currency: %s", receiverCurrency)
	}
	return 0, fmt.Errorf("unsupported sender currency: %s", senderCurrency)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCurrencyConverterServer(s, &server{})
	log.Println("Server started at port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
