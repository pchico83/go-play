package client

import (
	docker "github.com/fsouza/go-dockerclient"
)

var Factory ClientFactory

type ClientFactory interface {
	Create(ip string, cert string, key string, ca string) (Client, error)
}

type Client interface {
	Ping() error
}

type DockerFactory struct {
}

func (f DockerFactory) Create(ip string, cert string, key string, ca string) (Client, error) {
	return docker.NewTLSClient(ip, cert, key, ca)
}
