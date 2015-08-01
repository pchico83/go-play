package client

import (
	"errors"
)

type ErrorPingFactory struct{}

func (f ErrorPingFactory) Create(ip string, cert string, key string, ca string) (Client, error) {
	return ErrorPingClient{}, nil
}

type ErrorPingClient struct{}

func (f ErrorPingClient) Ping() error {
	return errors.New("Error to connect")
}
