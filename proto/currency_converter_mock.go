package proto

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// MockCurrencyConverterServer is a mock implementation of the CurrencyConverterServer interface
type MockCurrencyConverterServer struct {
	mock.Mock
}

func (m *MockCurrencyConverterServer) Convert(ctx context.Context, req *ConvertRequest) (*ConvertResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*ConvertResponse), args.Error(1)
}

func (m *MockCurrencyConverterServer) mustEmbedUnimplementedCurrencyConverterServer() {}
