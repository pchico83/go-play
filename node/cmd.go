package node

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/pchico83/elora/cluster"
	"github.com/pchico83/elora/utils"
	"strings"
)

const (
	FILL_SIZE = 30
)

func CreateCmd(name string, url string, cert string, key string, clusterName string, tags []string) error {
	url = CompleteUrl(url)
	log.WithFields(log.Fields{
		"name":    name,
		"url":     url,
		"cert":    cert,
		"key":     key,
		"cluster": clusterName,
		"tags":    tags,
	}).Info("Creating node")

	exists := Exists(name)
	if exists {
		return fmt.Errorf("Node %s already exists", name)
	}

	if !cluster.Exists(clusterName) {
		if clusterName != "default" {
			return fmt.Errorf("Cluster %s does not exist", clusterName)
		}

		defaultCluster := cluster.Cluster{Name: clusterName, Consul: true, Weave: true, Cleanup: true}
		err := cluster.Write(defaultCluster)
		if err != nil {
			return err
		}
	}

	node := Node{Name: name, Url: url, Cert: cert, Key: key, Cluster: clusterName, Tags: tags}
	_, err := node.Client()
	if err != nil {
		return err
	}

	err = Write(node)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"name":    name,
		"url":     url,
		"cert":    cert,
		"key":     key,
		"cluster": clusterName,
		"tags":    tags,
	}).Info("Created node")
	return nil
}

func EnvCmd(name string, unset bool) error {
	if !Exists(name) {
		return fmt.Errorf("Node %s does not exist", name)
	}

	node, err := Read(name)
	if err != nil {
		return err
	}
	if unset {
		fmt.Printf("unset DOCKER_TLS_VERIFY DOCKER_CERT_PATH DOCKER_HOST\n")
	} else {
		fmt.Printf("export DOCKER_HOST=%s\nexport DOCKER_CERT_PATH=%s\nexport DOCKER_TLS_VERIFY=yes\n", node.Url, node.Cert)
	}
	return nil
}

func InspectCmd(name string) error {
	if !Exists(name) {
		return fmt.Errorf("Node %s does not exist", name)
	}

	node, err := Read(name)
	if err != nil {
		return err
	}

	_, err = node.Client()
	if err != nil {
		fmt.Printf(
			"%s%s%s%s%s\n",
			utils.FillString(name, FILL_SIZE),
			utils.FillString("offline", FILL_SIZE),
			utils.FillString(node.Url, FILL_SIZE),
			utils.FillString(node.Cluster, FILL_SIZE),
			node.Tags)
	} else {
		fmt.Printf(
			"%s%s%s%s%s\n",
			utils.FillString(name, FILL_SIZE),
			utils.FillString("online", FILL_SIZE),
			utils.FillString(node.Url, FILL_SIZE),
			utils.FillString(node.Cluster, FILL_SIZE),
			node.Tags)
	}
	return nil
}

func ListCmd() error {
	nodes, err := List()
	if err != nil {
		return err
	}
	if len(nodes) == 0 {
		fmt.Printf("There are not nodes\n")
		return nil
	}

	fmt.Printf(
		"%s%s%s%sTAGS\n",
		utils.FillString("NAME", FILL_SIZE),
		utils.FillString("STATE", FILL_SIZE),
		utils.FillString("URL", FILL_SIZE),
		utils.FillString("CLUSTER", FILL_SIZE))
	for _, node := range nodes {
		err = InspectCmd(node)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateCmd(name string, url string, cert string, key string, tags []string) error {
	if !Exists(name) {
		return fmt.Errorf("Node %s does not exist", name)
	}

	if url != "" {
		url = CompleteUrl(url)
	}

	log.WithFields(log.Fields{
		"name": name,
		"url":  url,
		"cert": cert,
		"key":  key,
		"tags": tags,
	}).Info("Updating node")

	node, err := Read(name)
	if err != nil {
		return err
	}

	if url != "" {
		node.Url = url
	}
	if cert != "" {
		node.Cert = cert
	}
	if key != "" {
		node.Key = key
	}
	if len(tags) > 0 {
		node.Tags = tags
	}

	_, err = node.Client()
	if err != nil {
		return err
	}

	err = Write(node)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"name": node.Name,
		"url":  node.Url,
		"cert": node.Cert,
		"key":  key,
		"tags": tags,
	}).Info("Updated node")
	return nil
}

func RemoveCmd(name string) error {
	if !Exists(name) {
		return fmt.Errorf("Node %s does not exist", name)
	}

	log.WithFields(log.Fields{
		"name": name,
	}).Info("Removing node")

	//TODO: move container to another node if possible
	err := Delete(name)
	if err != nil {
		return err
	}

	log.WithFields(log.Fields{
		"name": name,
	}).Info("Removed node")
	return nil
}

func CompleteUrl(url string) string {
	if !strings.Contains(url, "://") {
		url = "tcp://" + url
	}
	if !strings.Contains(url, ":2376") {
		url = url + ":2376"
	}
	return url
}

func GetIp(url string) string {
	index := strings.Index(url, "//")
	url = url[index+2:]
	index = strings.Index(url, ":")
	return url[:index]
}
