package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
	"time"

	pb "currency-conversion/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
)

type mockServer struct {
	pb.UnimplementedCurrencyConverterServer
}

// Mock server implementation and test
func (s *mockServer) Convert(ctx context.Context, req *pb.ConvertRequest) (*pb.ConvertResponse, error) {
	// Mock response
	return &pb.ConvertResponse{
		ConvertedAmount: 85.0, // Mock conversion rate
	}, nil
}

func TestServerInitialization(t *testing.T) {
	// Create a mock listener
	lis, err := net.Listen("tcp", ":0") // :0 means a random available port
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	// Create a new gRPC server and register the mock server
	s := grpc.NewServer()
	pb.RegisterCurrencyConverterServer(s, &mockServer{})

	// Start the server in a goroutine
	go func() {
		if err := s.Serve(lis); err != nil {
			t.Errorf("failed to serve: %v", err)
			return
		}
	}()
	defer s.Stop()

	// Create a client connection to the server
	conn, err := grpc.DialContext(
		context.Background(),
		lis.Addr().String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewCurrencyConverterClient(conn)

	// Test the connection by making a simple request
	req := &pb.ConvertRequest{
		SenderCurrencyType:   "USD",
		ReceiverCurrencyType: "EUR",
		Amount:               100,
	}

	_, err = client.Convert(context.Background(), req)
	if err != nil {
		t.Fatalf("Convert failed: %v", err)
	}
}

// test currency conversion
const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterCurrencyConverterServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestServer_Convert(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Fatalf("Failed to close connection: %v", err)
		}
	}(conn)
	client := pb.NewCurrencyConverterClient(conn)

	testCases := []struct {
		senderCurrency   string
		receiverCurrency string
		amount           float64
		expected         float64
	}{
		{"INR", "USD", 100, 1.3},
		{"INR", "EUR", 100, 1.1},
		{"USD", "INR", 100, 7400},
		{"USD", "EUR", 100, 85},
		{"EUR", "USD", 100, 118},
		{"EUR", "INR", 100, 8700},
	}

	for _, tc := range testCases {
		req := &pb.ConvertRequest{
			SenderCurrencyType:   tc.senderCurrency,
			ReceiverCurrencyType: tc.receiverCurrency,
			Amount:               tc.amount,
		}

		resp, err := client.Convert(ctx, req)
		if err != nil {
			t.Fatalf("Convert failed: %v", err)
		}

		if resp.ConvertedAmount != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, resp.ConvertedAmount)
		}
	}
}
