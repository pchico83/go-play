package stack

import (
	"errors"
	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func Exists(name string) bool {
	filePath := DefinitionFilePath(name)
	_, err := os.Stat(filePath)
	return err == nil
}

func Read(name string) (Stack, error) {
	s := Stack{Name: name}
	definitionFilePath := DefinitionFilePath(name)
	data, err := ioutil.ReadFile(definitionFilePath)
	if err != nil {
		return s, err
	}
	err = yaml.Unmarshal(data, &s.Definition)
	if err != nil {
		return s, err
	}
	s.SetDefaults()
	return s, nil
}

func List() ([]Stack, error) {
	fileFolder := FileFolder()
	result := make([]Stack, 0)
	files, err := ioutil.ReadDir(fileFolder)
	if err != nil {
		return result, err
	}
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".yml") {
			stack, err := Read(f.Name())
			if err != nil {
				return result, err
			}
			result = append(result, stack)
		}
	}
	return result, nil
}

func Write(name string, sourceFilePath string) error {
	data, err := ioutil.ReadFile(sourceFilePath)
	if err != nil {
		return err
	}

	if strings.Contains(string(data), "\t") {
		return errors.New("'\t' characters are not allowed in YAML files")
	}

	definitionFilePath := DefinitionFilePath(name)
	err = ioutil.WriteFile(definitionFilePath, data, 0644)
	if err != nil {
		return err
	}

	stateFilePath := StateFilePath(name)
	return ioutil.WriteFile(stateFilePath, []byte(""), 0644)
}

func Delete(name string) error {
	definitionFilePath := DefinitionFilePath(name)
	return os.Remove(definitionFilePath)
}

func FileFolder() (path string) {
	var fileFolder string
	if os.Getenv("ELORA_STORAGE_FOLDER") != "" {
		fileFolder = filepath.Join(os.Getenv("ELORA_STORAGE_FOLDER"), "stacks")
	} else {
		user, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		fileFolder = filepath.Join(user.HomeDir, "elora/stacks")
	}
	if err := os.MkdirAll(fileFolder, 0744); err != nil {
		log.Fatal(err)
	}
	return fileFolder
}

func DefinitionFilePath(name string) (path string) {
	fileFolder := FileFolder()
	definitionFileName := name + ".yml"
	return filepath.Join(fileFolder, definitionFileName)
}

func StateFilePath(name string) (path string) {
	fileFolder := FileFolder()
	return filepath.Join(fileFolder, name)
}
