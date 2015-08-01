package client

type MockFactory struct{}

func (f MockFactory) Create(ip string, cert string, key string, ca string) (Client, error) {
	return MockClient{}, nil
}

type MockClient struct{}

func (f MockClient) Ping() error {
	return nil
}
