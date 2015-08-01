package cluster

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/pchico83/elora/utils"
)

const (
	FILL_SIZE = 30
)

func CreateCmd(name string, consul bool, weave bool, weavePwd string, cleanup bool) error {
	log.WithFields(log.Fields{
		"name":      name,
		"consul":    consul,
		"weave":     weave,
		"weave-pwd": weavePwd,
		"cleanup":   cleanup,
	}).Info("Creating cluster")

	exists := Exists(name)
	if exists {
		return fmt.Errorf("Cluster %s already exists", name)
	}

	cluster := Cluster{Name: name, Consul: consul, Weave: weave, WeavePwd: weavePwd, Cleanup: cleanup}
	err := Write(cluster)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"name":      name,
		"consul":    consul,
		"weave":     weave,
		"weave-pwd": weavePwd,
		"cleanup":   cleanup,
	}).Info("Created cluster")
	return nil
}

func InspectCmd(name string) error {
	if !Exists(name) {
		return fmt.Errorf("Cluster %s does not exist", name)
	}

	//	cluster, err := Read(name)
	//	if err != nil {
	//		return err
	//	}

	//	hasWeavePwd := cluster.WeavePwd != ""
	//	fmt.Printf(
	//		"%s%b%s%s%s\n",
	//		utils.FillString(name, FILL_SIZE),
	//		utils.FillString(cluster.Consul, FILL_SIZE),
	//		utils.FillString(cluster.Weave, FILL_SIZE),
	//		utils.FillString(hasWeavePwd, FILL_SIZE),
	//		cluster.Cleanup)
	return nil
}

func ListCmd() error {
	clusters, err := List()
	if err != nil {
		return err
	}
	if len(clusters) == 0 {
		fmt.Printf("There are not clusters\n")
		return nil
	}

	fmt.Printf(
		"%s%s%sCLEANUP\n",
		utils.FillString("NAME", FILL_SIZE),
		utils.FillString("CONSUL", FILL_SIZE),
		utils.FillString("WEAVE", FILL_SIZE),
		utils.FillString("WEAVE-PWD", FILL_SIZE))
	for _, cluster := range clusters {
		err = InspectCmd(cluster)
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveCmd(name string) error {
	if !Exists(name) {
		return fmt.Errorf("Cluster %s does not exist", name)
	}

	log.WithFields(log.Fields{
		"name": name,
	}).Info("Removing cluster")

	err := Delete(name)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"name": name,
	}).Info("Removed cluster")
	return nil
}
