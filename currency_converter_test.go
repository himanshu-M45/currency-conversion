package currency

import (
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
