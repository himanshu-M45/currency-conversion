package main

import (
	"context"
	"currency-conversion/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

// Conversion rates
var conversionRates map[string]float64

type server struct {
	proto.UnimplementedCurrencyConverterServer
}

func (s *server) Convert(_ context.Context, req *proto.ConvertRequest) (*proto.ConvertResponse, error) {
	senderCurrency := req.GetSenderCurrencyType()
	receiverCurrency := req.GetReceiverCurrencyType()
	amount := req.GetAmount()

	if amount <= 0 || senderCurrency == "" || receiverCurrency == "" {
		return nil, fmt.Errorf("invalid input")
	}
	if senderCurrency == receiverCurrency { // No conversion needed
		return &proto.ConvertResponse{ConvertedAmount: amount}, nil
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

	return &proto.ConvertResponse{ConvertedAmount: formattedAmount}, nil
}

func convertCurrency(senderCurrency, receiverCurrency string, amount float64) (float64, error) {
	senderRate, senderOk := conversionRates[senderCurrency]
	receiverRate, receiverOk := conversionRates[receiverCurrency]

	if !senderOk {
		return 0, fmt.Errorf("unsupported sender currency: %s", senderCurrency)
	}
	if !receiverOk {
		return 0, fmt.Errorf("unsupported receiver currency: %s", receiverCurrency)
	}

	amountInBaseCurrency := amount / senderRate
	return amountInBaseCurrency * receiverRate, nil
}

func init() {
	var err error
	conversionRates, err = LoadConversionRates("conversion_rates.xml")
	if err != nil {
		log.Fatalf("failed to load conversion rates: %v", err)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterCurrencyConverterServer(s, &server{})
	log.Println("Server started at port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
