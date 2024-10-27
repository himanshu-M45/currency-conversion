package controllers

import (
	"context"
	"currency-conversion/internal/app/cc/service"
	"currency-conversion/proto"
	"fmt"
)

type Server struct {
	proto.UnimplementedCurrencyConversionServer
}

func (s *Server) Convert(_ context.Context, req *proto.ConvertRequest) (*proto.ConvertResponse, error) {
	senderCurrency := req.GetSenderCurrencyType()
	receiverCurrency := req.GetReceiverCurrencyType()
	amount := req.GetAmount()

	convertedAmount, err := service.ConvertCurrency(senderCurrency, receiverCurrency, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to convert currency: %v", err)
	}

	return &proto.ConvertResponse{ConvertedAmount: convertedAmount}, nil
}
