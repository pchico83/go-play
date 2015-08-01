package node

import (
	"errors"
	"fmt"
	"github.com/pchico83/elora/client"
	"path/filepath"
	"time"
)

type Node struct {
	Name    string
	Url     string
	Cert    string
	Key     string
	Cluster string
	Tags    []string
}

func (n *Node) Client() (client.Client, error) {
	client, err := client.Factory.Create(
		n.Url,
		filepath.Join(n.Cert, "cert.pem"),
		filepath.Join(n.Cert, "key.pem"),
		filepath.Join(n.Cert, "ca.pem"))
	if err != nil {
		return client, err
	}

	c1 := make(chan error, 1)
	go func() {
		c1 <- client.Ping()
	}()

	select {
	case err = <-c1:
		return client, err
	case <-time.After(time.Second * 5):
		return client, errors.New("Cannot connect to host after 5s")
	}
}

func (n *Node) Execute(command string) (string, error) {
	fmt.Println(command)
	return "", nil
}

func (n *Node) RemoveContainers() error {
	fmt.Println(n.Name)
	return nil
}
