package server

import (
	"currency-conversion/proto"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
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

func TestServer_StartServerSuccessfully(t *testing.T) {
	lis := new(mockListener)
	lis.On("Accept").Return(nil, errors.New("mock accept error"))
	lis.On("Close").Return(nil)
	lis.On("Addr").Return(&net.TCPAddr{})

	s := grpc.NewServer()
	mockServer := new(proto.MockCurrencyConversionServer)
	proto.RegisterCurrencyConversionServer(s, mockServer)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	assert.NotNil(t, s)
}

func TestServer_FailedToListen(t *testing.T) {
	_, err := net.Listen("tcp", ":99999")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid port")
}

func TestServer_FailedToServe(t *testing.T) {
	lis := new(mockListener)
	lis.On("Accept").Return(nil, errors.New("mock accept error"))
	lis.On("Close").Return(nil)
	lis.On("Addr").Return(&net.TCPAddr{})

	s := grpc.NewServer()
	mockServer := new(proto.MockCurrencyConversionServer)
	proto.RegisterCurrencyConversionServer(s, mockServer)

	go func() {
		err := s.Serve(lis)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "mock accept error")
	}()
}
