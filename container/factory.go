package container

import (
	"github.com/pchico83/d2k8/stack"
)

type Manager interface {
	Run(stack stack.Stack, service string, container string) (string, error)
	Stop(stack stack.Stack, service string, container string) error
	Start(stack stack.Stack, service string, container string) error
	Remove(stack stack.Stack, service string, container string, clean bool) error
}
