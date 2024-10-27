package proto

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockCurrencyConversionServer is a mock implementation of the CurrencyConverterServer interface
type MockCurrencyConversionServer struct {
	mock.Mock
}

func (m *MockCurrencyConversionServer) Convert(ctx context.Context, req *ConvertRequest) (*ConvertResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*ConvertResponse), args.Error(1)
}

func (m *MockCurrencyConversionServer) mustEmbedUnimplementedCurrencyConversionServer() {}
