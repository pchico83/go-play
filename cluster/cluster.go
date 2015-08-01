package cluster

import (
	"fmt"
)

type Cluster struct {
	Name     string
	Consul   bool
	Weave    bool
	WeavePwd string
	Cleanup  bool
}

func (c *Cluster) RemoveNodes() error {
	fmt.Println(c.Name)
	return nil
}
