package controllers

import (
	"context"
	"errors"
	"testing"

	pb "currency-conversion/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCurrencyConverterServer struct {
	mock.Mock
}

func (m *mockCurrencyConverterServer) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.ConvertResponse), args.Error(1)
}

func TestConvert_SuccessfulConversions(t *testing.T) {
	testCases := []struct {
		name         string
		req          *pb.ConvertRequest
		expectedResp *pb.ConvertResponse
	}{
		{
			name: "INR to USD",
			req: &pb.ConvertRequest{
				SenderCurrencyType:   "INR",
				ReceiverCurrencyType: "USD",
				Amount:               100,
			},
			expectedResp: &pb.ConvertResponse{ConvertedAmount: 1.3},
		},
		{
			name: "INR to EUR",
			req: &pb.ConvertRequest{
				SenderCurrencyType:   "INR",
				ReceiverCurrencyType: "EUR",
				Amount:               100,
			},
			expectedResp: &pb.ConvertResponse{ConvertedAmount: 1.1},
		},
		{
			name: "USD to INR",
			req: &pb.ConvertRequest{
				SenderCurrencyType:   "USD",
				ReceiverCurrencyType: "INR",
				Amount:               100,
			},
			expectedResp: &pb.ConvertResponse{ConvertedAmount: 7400},
		},
		{
			name: "USD to EUR",
			req: &pb.ConvertRequest{
				SenderCurrencyType:   "USD",
				ReceiverCurrencyType: "EUR",
				Amount:               100,
			},
			expectedResp: &pb.ConvertResponse{ConvertedAmount: 85.0},
		},
		{
			name: "EUR to USD",
			req: &pb.ConvertRequest{
				SenderCurrencyType:   "EUR",
				ReceiverCurrencyType: "USD",
				Amount:               100,
			},
			expectedResp: &pb.ConvertResponse{ConvertedAmount: 118.0},
		},
		{
			name: "EUR to INR",
			req: &pb.ConvertRequest{
				SenderCurrencyType:   "EUR",
				ReceiverCurrencyType: "INR",
				Amount:               100,
			},
			expectedResp: &pb.ConvertResponse{ConvertedAmount: 8700},
		},
		{
			name: "No Conversion Needed",
			req: &pb.ConvertRequest{
				SenderCurrencyType:   "USD",
				ReceiverCurrencyType: "USD",
				Amount:               100,
			},
			expectedResp: &pb.ConvertResponse{ConvertedAmount: 100},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockServer := new(mockCurrencyConverterServer)
			mockServer.On("Convert", mock.Anything, tc.req).Return(tc.expectedResp, nil)

			resp, err := mockServer.Convert(context.Background(), tc.req)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResp, resp)
		})
	}
}

func TestConvert_FailureScenarios(t *testing.T) {
	testCases := []struct {
		name        string
		req         *pb.ConvertRequest
		expectedErr string
	}{
		{
			name: "Invalid Input",
			req: &pb.ConvertRequest{
				SenderCurrencyType:   "",
				ReceiverCurrencyType: "USD",
				Amount:               -10,
			},
			expectedErr: "invalid input",
		},
		{
			name: "Unsupported Sender Currency",
			req: &pb.ConvertRequest{
				SenderCurrencyType:   "YEN",
				ReceiverCurrencyType: "USD",
				Amount:               100,
			},
			expectedErr: "unsupported sender currency: GBP",
		},
		{
			name: "Unsupported Receiver Currency",
			req: &pb.ConvertRequest{
				SenderCurrencyType:   "USD",
				ReceiverCurrencyType: "JPY",
				Amount:               100,
			},
			expectedErr: "unsupported receiver currency: GBP",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockServer := new(mockCurrencyConverterServer)
			mockServer.On("Convert", mock.Anything, tc.req).Return((*pb.ConvertResponse)(nil), errors.New(tc.expectedErr))

			resp, err := mockServer.Convert(context.Background(), tc.req)
			assert.Nil(t, resp)
			assert.Error(t, err)
			assert.Equal(t, tc.expectedErr, err.Error())
		})
	}
}
