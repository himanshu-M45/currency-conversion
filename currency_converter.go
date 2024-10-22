package main

import (
	"context"
	pb "currency-conversion/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
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

	if amount <= 0 || senderCurrency == "" || receiverCurrency == "" {
		return nil, fmt.Errorf("invalid input")
	}
	if senderCurrency == receiverCurrency { // No conversion needed
		return &pb.ConvertResponse{ConvertedAmount: amount}, nil
	}

	// Call the conversion logic
	convertedAmount, err := convertCurrency(senderCurrency, receiverCurrency, amount)
	if err != nil {
		return nil, err
	}

	// Format the converted amount to two decimal places
	formattedAmount, err := strconv.ParseFloat(fmt.Sprintf("%.2f", convertedAmount), 64)
	if err != nil {
		return nil, fmt.Errorf("failed to format converted amount: %v", err)
	}

	return &pb.ConvertResponse{ConvertedAmount: formattedAmount}, nil
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
