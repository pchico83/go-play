package utils

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pchico83/elora/client"
	"io/ioutil"
	"os"
)

func FillString(value string, size int) string {
	for i := len(value); i < size; i++ {
		value = value + " "
	}
	return value
}

func SetUpTestCommon() {
	folder, err := ioutil.TempDir("", "elora-")
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("ELORA_STORAGE_FOLDER", folder)
	if err != nil {
		log.Fatal(err)
	}
	client.Factory = client.MockFactory{}
}

func TearDownTestCommon() {
	folder := os.Getenv("ELORA_STORAGE_FOLDER")
	os.RemoveAll(folder)
}
