package container

import (
	"fmt"
	"github.com/pchico83/elora/node"
	"github.com/pchico83/elora/stack"
)

type DockerManager struct{}

func (d DockerManager) Run(stack stack.Stack, service string, container string) (string, error) {
	nodeName := stack.Services[service].Containers[container].Node

	if nodeName == "" {
		return "", fmt.Errorf("Container %s does not have node", container)
	}

	if !node.Exists(nodeName) {
		return "", fmt.Errorf("Node %s does not exist", nodeName)
	}

	_, err := node.Read(nodeName)
	if err != nil {
		return "", err
	}

	command := "docker run -d --name " + container
	command += " --dns=[consulIP] --restart " + stack.Definition[service].Restart
	if stack.Definition[service].Mem_limit != "" {
		command += " -m=" + stack.Definition[service].Mem_limit
	}
	if stack.Definition[service].Privileged {
		command += " --privileged=true"
	}

	if stack.Definition[service].Entrypoint != "" {
		command += " --entrypoint=" + stack.Definition[service].Entrypoint
	}

	command += stack.Definition[service].Image

	if stack.Definition[service].Command != "" {
		command += " " + stack.Definition[service].Command
	}

	return command, nil
}

func (d DockerManager) Stop(stack stack.Stack, service string, container string) error {
	dockerId := stack.Services[service].Containers[container].DockerId
	nodeName := stack.Services[service].Containers[container].Node

	if nodeName == "" {
		return fmt.Errorf("Container %s does not have node", container)
	}

	if !node.Exists(nodeName) {
		return fmt.Errorf("Node %s does not exist", nodeName)
	}

	n, err := node.Read(nodeName)
	if err != nil {
		return err
	}

	var command string
	if dockerId != "" {
		command = "docker stop " + dockerId
	} else {
		command = "docker stop " + container
	}

	_, err = n.Execute(command)
	return err
}

func (d DockerManager) Start(stack stack.Stack, service string, container string) error {
	dockerId := stack.Services[service].Containers[container].DockerId
	nodeName := stack.Services[service].Containers[container].Node

	if nodeName == "" {
		return fmt.Errorf("Container %s does not have node", container)
	}

	if !node.Exists(nodeName) {
		return fmt.Errorf("Node %s does not exist", nodeName)
	}

	n, err := node.Read(nodeName)
	if err != nil {
		return err
	}

	var command string
	if dockerId != "" {
		command = "docker start " + dockerId
	} else {
		command = "docker start " + container
	}

	_, err = n.Execute(command)
	return err
}

func (d DockerManager) Remove(stack stack.Stack, service string, container string, clean bool) error {
	dockerId := stack.Services[service].Containers[container].DockerId
	nodeName := stack.Services[service].Containers[container].Node

	if nodeName == "" {
		return fmt.Errorf("Container %s does not have node", container)
	}

	if !node.Exists(nodeName) {
		return fmt.Errorf("Node %s does not exist", nodeName)
	}

	n, err := node.Read(nodeName)
	if err != nil {
		return err
	}

	var command string
	if dockerId != "" {
		if clean {
			command = "docker rm -fv " + dockerId
		} else {
			command = "docker rm -f " + dockerId
		}
	} else {
		if clean {
			command = "docker rm -fv " + container
		} else {
			command = "docker rm -f " + container
		}
	}

	_, err = n.Execute(command)
	return err
}
