package stack

import (
	"github.com/pchico83/elora/node"
	"io/ioutil"
)

type Stack struct {
	Name       string
	Operation  string
	Definition map[string]ServiceDefinition
	Services   map[string]ServiceState
}

type ServiceDefinition struct {
	Image        string   `json:"image"`
	Command      string   `json:"command"`
	Links        []string `json:"links"`
	Ports        []string `json:"ports"`
	Volumes      []string `json:"volumes"`
	Volumes_from []string `json:"volumes_from"`
	Environment  []string `json:"environment"`
	Entrypoint   string   `json:"entrypoint"`
	Mem_limit    string   `json:"mem_limit"`
	Privileged   bool     `json:"privileged"`
	Restart      string   `json:"restart"`
	Strategy     string   `json:"strategy"`
	Scale        int      `json:"scale"`
	Tags         []string `json:"tags"`
}

type ServiceState struct {
	Operation  string
	Containers map[string]Container
}

type Container struct {
	Ip        string
	Node      string
	Status    string
	Operation string
	DockerId  string
	Ports     []string
}

type StackInterface interface{}

func (s *Stack) Export() (string, error) {
	definitionFilePath := DefinitionFilePath(s.Name)
	data, err := ioutil.ReadFile(definitionFilePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *Stack) SetDefaults() {
	for name, service := range s.Definition {
		if service.Restart == "" {
			service.Restart = "no"
		}
		if service.Strategy == "" {
			service.Strategy = "balance"
		}
		if service.Scale == 0 {
			service.Scale = 1
		}
		s.Definition[name] = service
	}
}

type NodeLoad struct {
	Containers int
	Ports      []string
}

func (s *Stack) SetBestNodes() error {
	nodes, err := node.List()
	if err != nil {
		return err
	}

	nodeLoads := make(map[string]NodeLoad)
	for _, nodeName := range nodes {
		nodeLoads[nodeName] = NodeLoad{Containers: 0, Ports: make([]string, 0)}
	}

	stacks, err := List()
	if err != nil {
		return err
	}

	for _, stack := range stacks {
		for serviceName, _ := range stack.Definition {
			for _, container := range stack.Services[serviceName].Containers {
				if container.Node != "" { //&& container.Node in nodeLoads {
					load := nodeLoads[container.Node]
					load.Containers += 1
					for _, p := range container.Ports {
						load.Ports = append(load.Ports, p)
					}
					nodeLoads[container.Node] = load
				}
			}
		}
	}

	return nil
}
