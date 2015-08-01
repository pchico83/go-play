package node

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

func Exists(name string) bool {
	filePath := FilePath(name)
	_, err := os.Stat(filePath)
	return err == nil
}

func Read(name string) (Node, error) {
	n := Node{}
	filePath := FilePath(name)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return n, err
	}

	return n, yaml.Unmarshal(data, &n)
}

func List() ([]string, error) {
	fileFolder := FileFolder()
	result := make([]string, 0)
	files, err := ioutil.ReadDir(fileFolder)
	if err != nil {
		return result, err
	}
	for _, f := range files {
		result = append(result, f.Name())
	}
	return result, nil
}

func Write(node Node) error {
	filePath := FilePath(node.Name)
	data, err := yaml.Marshal(&node)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, []byte(data), 0644)
}

func Delete(name string) error {
	filePath := FilePath(name)
	return os.Remove(filePath)
}

func FileFolder() string {
	var fileFolder string
	if os.Getenv("ELORA_STORAGE_FOLDER") != "" {
		fileFolder = filepath.Join(os.Getenv("ELORA_STORAGE_FOLDER"), "nodes")
	} else {
		user, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		fileFolder = filepath.Join(user.HomeDir, "elora/nodes")
	}
	if err := os.MkdirAll(fileFolder, 0744); err != nil {
		log.Fatal(err)
	}
	return fileFolder
}

func FilePath(name string) string {
	fileFolder := FileFolder()
	return filepath.Join(fileFolder, name)
}
