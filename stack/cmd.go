package stack

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/pchico83/d2k8/utils"
	"gopkg.in/yaml.v2"
)

const (
	FILL_SIZE = 30
)

func CreateCmd(name string, filePath string) error {
	log.WithFields(log.Fields{
		"name": name,
		"file": filePath,
	}).Info("Creating stack")

	if Exists(name) {
		return fmt.Errorf("Stack %s already exists", name)
	}

	err := Validate(name, filePath)
	if err != nil {
		return err
	}

	err = Write(name, filePath)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"name": name,
		"file": filePath,
	}).Info("Created stack")
	return nil
}

func InspectCmd(name string) error {
	if !Exists(name) {
		return fmt.Errorf("Stack %s does not exist", name)
	}

	stack, err := Read(name)
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(&stack)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", data)
	return nil
}

func ExportCmd(name string) error {
	if !Exists(name) {
		return fmt.Errorf("Stack %s does not exist", name)
	}

	stack, err := Read(name)
	if err != nil {
		return err
	}
	data, err := stack.Export()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", data)
	return nil
}

func ListCmd() error {
	stacks, err := List()
	if err != nil {
		return err
	}
	if len(stacks) == 0 {
		fmt.Printf("There are not stacks\n")
		return nil
	}

	fmt.Printf("%sLAST OPERATION\n", utils.FillString("NAME", FILL_SIZE))
	for _, stack := range stacks {
		fmt.Printf("%s\n", utils.FillString(stack.Name, FILL_SIZE))
	}
	return nil
}

func UpdateCmd(name string, filePath string) error {
	if !Exists(name) {
		return fmt.Errorf("Stack %s does not exist", name)
	}

	log.WithFields(log.Fields{
		"name": name,
		"file": filePath,
	}).Info("Updating stack")

	err := Validate(name, filePath)
	if err != nil {
		return err
	}

	err = Write(name, filePath)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"name": name,
		"file": filePath,
	}).Info("Updated stack")
	return nil
}

func RemoveCmd(name string) error {
	if !Exists(name) {
		return fmt.Errorf("Stack %s does not exist", name)
	}

	log.WithFields(log.Fields{
		"name": name,
	}).Info("Removing stack")

	err := Delete(name)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"name": name,
	}).Info("Removed stack")
	return nil
}
