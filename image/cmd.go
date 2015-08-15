package image

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/pchico83/d2k8/utils"
)

const (
	FILL_SIZE = 30
)

func CreateCmd(name string, user string, pwd string, url string) error {
	log.WithFields(log.Fields{
		"name": name,
		"user": user,
		"pwd":  pwd,
		"url":  url,
	}).Info("Creating image")

	exists := Exists(name)
	if exists {
		return fmt.Errorf("Image %s already exists", name)
	}

	image := Image{Name: name, User: user, Pwd: pwd, Url: url}
	err := Write(image)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"name": name,
		"user": user,
		"pwd":  pwd,
		"url":  url,
	}).Info("Created image")
	return nil
}

func InspectCmd(name string) error {
	if !Exists(name) {
		return fmt.Errorf("Image %s does not exist", name)
	}

	image, err := Read(name)
	if err != nil {
		return err
	}

	fmt.Printf(
		"%s%s%s\n",
		utils.FillString(image.Name, FILL_SIZE),
		utils.FillString(image.User, FILL_SIZE),
		utils.FillString(image.Url, FILL_SIZE))
	return nil
}

func ListCmd() error {
	images, err := List()
	if err != nil {
		return err
	}
	if len(images) == 0 {
		fmt.Printf("There are not private images\n")
		return nil
	}

	fmt.Printf(
		"%s%sURL\n",
		utils.FillString("NAME", FILL_SIZE),
		utils.FillString("USER", FILL_SIZE))
	for _, image := range images {
		err = InspectCmd(image)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateCmd(name string, user string, pwd string, url string) error {
	if !Exists(name) {
		return fmt.Errorf("Image %s does not exist", name)
	}

	log.WithFields(log.Fields{
		"name": name,
		"user": user,
		"pwd":  pwd,
		"url":  url,
	}).Info("Updating image")

	image, err := Read(name)
	if err != nil {
		return err
	}

	if user != "" {
		image.User = user
	}
	if pwd != "" {
		image.Pwd = pwd
	}
	if url != "" {
		image.Url = url
	}

	err = Write(image)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"name": image.Name,
		"user": user,
		"pwd":  pwd,
		"url":  url,
	}).Info("Updated image")
	return nil
}

func RemoveCmd(name string) error {
	if !Exists(name) {
		return fmt.Errorf("Image %s does not exist", name)
	}

	log.WithFields(log.Fields{
		"name": name,
	}).Info("Removing image")

	err := Delete(name)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"name": name,
	}).Info("Removed image")
	return nil
}
