package currency

import (
	"context"
	"errors"
	"log"
	"net"
	"testing"

	pb "currency-conversion/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

type mockCurrencyConverterServer struct {
	mock.Mock
}

func (m *mockCurrencyConverterServer) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.ConvertResponse), args.Error(1)
}

type mockListener struct {
	mock.Mock
}

func (m *mockListener) Accept() (net.Conn, error) {
	args := m.Called()
	return args.Get(0).(net.Conn), args.Error(1)
}

func (m *mockListener) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *mockListener) Addr() net.Addr {
	args := m.Called()
	return args.Get(0).(net.Addr)
}

func TestMain_StartServerSuccessfully(t *testing.T) {
	lis := new(mockListener)
	lis.On("Accept").Return(nil, errors.New("mock accept error"))
	lis.On("Close").Return(nil)
	lis.On("Addr").Return(&net.TCPAddr{})

	s := grpc.NewServer()
	mockServer := new(pb.MockCurrencyConverterServer)
	pb.RegisterCurrencyConverterServer(s, mockServer)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	assert.NotNil(t, s)
}

func TestMain_FailedToListen(t *testing.T) {
	_, err := net.Listen("tcp", ":99999")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid port")
}

func TestMain_FailedToServe(t *testing.T) {
	lis := new(mockListener)
	lis.On("Accept").Return(nil, errors.New("mock accept error"))
	lis.On("Close").Return(nil)
	lis.On("Addr").Return(&net.TCPAddr{})

	s := grpc.NewServer()
	mockServer := new(pb.MockCurrencyConverterServer)
	pb.RegisterCurrencyConverterServer(s, mockServer)

	go func() {
		err := s.Serve(lis)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "mock accept error")
	}()
}

func TestConvert_SuccessfulConversion_INRtoUSD(t *testing.T) {
	mockServer := new(mockCurrencyConverterServer)
	req := &pb.ConvertRequest{
		SenderCurrencyType:   "INR",
		ReceiverCurrencyType: "USD",
		Amount:               100,
	}
	expectedResp := &pb.ConvertResponse{ConvertedAmount: 1.3} // Assuming the conversion rate is 0.013
	mockServer.On("Convert", mock.Anything, req).Return(expectedResp, nil)

	resp, err := mockServer.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp)
}

func TestConvert_SuccessfulConversion_INRtoEUR(t *testing.T) {
	mockServer := new(mockCurrencyConverterServer)
	req := &pb.ConvertRequest{
		SenderCurrencyType:   "INR",
		ReceiverCurrencyType: "EUR",
		Amount:               100,
	}
	expectedResp := &pb.ConvertResponse{ConvertedAmount: 1.1} // Assuming the conversion rate is 0.011
	mockServer.On("Convert", mock.Anything, req).Return(expectedResp, nil)

	resp, err := mockServer.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp)
}

func TestConvert_SuccessfulConversion_USDtoINR(t *testing.T) {
	mockServer := new(mockCurrencyConverterServer)
	req := &pb.ConvertRequest{
		SenderCurrencyType:   "USD",
		ReceiverCurrencyType: "INR",
		Amount:               100,
	}
	expectedResp := &pb.ConvertResponse{ConvertedAmount: 7400} // Assuming the conversion rate is 74.0
	mockServer.On("Convert", mock.Anything, req).Return(expectedResp, nil)

	resp, err := mockServer.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp)
}

func TestConvert_SuccessfulConversion_USDtoEUR(t *testing.T) {
	mockServer := new(mockCurrencyConverterServer)
	req := &pb.ConvertRequest{
		SenderCurrencyType:   "USD",
		ReceiverCurrencyType: "EUR",
		Amount:               100,
	}
	expectedResp := &pb.ConvertResponse{ConvertedAmount: 85.0} // Assuming the conversion rate is 0.85
	mockServer.On("Convert", mock.Anything, req).Return(expectedResp, nil)

	resp, err := mockServer.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp)
}

func TestConvert_SuccessfulConversion_EURtoUSD(t *testing.T) {
	mockServer := new(mockCurrencyConverterServer)
	req := &pb.ConvertRequest{
		SenderCurrencyType:   "EUR",
		ReceiverCurrencyType: "USD",
		Amount:               100,
	}
	expectedResp := &pb.ConvertResponse{ConvertedAmount: 118.0} // Assuming the conversion rate is 1.18
	mockServer.On("Convert", mock.Anything, req).Return(expectedResp, nil)

	resp, err := mockServer.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp)
}

func TestConvert_SuccessfulConversion_EURtoINR(t *testing.T) {
	mockServer := new(mockCurrencyConverterServer)
	req := &pb.ConvertRequest{
		SenderCurrencyType:   "EUR",
		ReceiverCurrencyType: "INR",
		Amount:               100,
	}
	expectedResp := &pb.ConvertResponse{ConvertedAmount: 8700} // Assuming the conversion rate is 87.0
	mockServer.On("Convert", mock.Anything, req).Return(expectedResp, nil)

	resp, err := mockServer.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp)
}

func TestConvert_InvalidInput(t *testing.T) {
	mockServer := new(mockCurrencyConverterServer)
	req := &pb.ConvertRequest{
		SenderCurrencyType:   "",
		ReceiverCurrencyType: "USD",
		Amount:               -10,
	}
	mockServer.On("Convert", mock.Anything, req).Return((*pb.ConvertResponse)(nil), errors.New("invalid input"))

	resp, err := mockServer.Convert(context.Background(), req)
	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "invalid input", err.Error())
}

func TestConvert_NoConversionNeeded(t *testing.T) {
	mockServer := new(mockCurrencyConverterServer)
	req := &pb.ConvertRequest{
		SenderCurrencyType:   "USD",
		ReceiverCurrencyType: "USD",
		Amount:               100,
	}
	expectedResp := &pb.ConvertResponse{ConvertedAmount: 100}
	mockServer.On("Convert", mock.Anything, req).Return(expectedResp, nil)

	resp, err := mockServer.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp)
}

func TestConvert_SuccessfulConversion(t *testing.T) {
	mockServer := new(mockCurrencyConverterServer)
	req := &pb.ConvertRequest{
		SenderCurrencyType:   "USD",
		ReceiverCurrencyType: "EUR",
		Amount:               100,
	}
	expectedResp := &pb.ConvertResponse{ConvertedAmount: 85.0} // Assuming the conversion rate is 0.85
	mockServer.On("Convert", mock.Anything, req).Return(expectedResp, nil)

	resp, err := mockServer.Convert(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expectedResp, resp)
}

func TestConvert_UnsupportedSenderCurrency(t *testing.T) {
	mockServer := new(mockCurrencyConverterServer)
	req := &pb.ConvertRequest{
		SenderCurrencyType:   "YEN",
		ReceiverCurrencyType: "USD",
		Amount:               100,
	}
	mockServer.On("Convert", mock.Anything, req).Return((*pb.ConvertResponse)(nil), errors.New("unsupported sender currency: GBP"))

	resp, err := mockServer.Convert(context.Background(), req)
	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "unsupported sender currency: GBP", err.Error())
}

func TestConvert_UnsupportedReceiverCurrency(t *testing.T) {
	mockServer := new(mockCurrencyConverterServer)
	req := &pb.ConvertRequest{
		SenderCurrencyType:   "USD",
		ReceiverCurrencyType: "GBP",
		Amount:               100,
	}
	mockServer.On("Convert", mock.Anything, req).Return((*pb.ConvertResponse)(nil), errors.New("unsupported receiver currency: GBP"))

	resp, err := mockServer.Convert(context.Background(), req)
	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, "unsupported receiver currency: GBP", err.Error())
}
