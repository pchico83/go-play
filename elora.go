package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/pchico83/elora/client"
	"os"

	"github.com/pchico83/elora/image"
	"github.com/pchico83/elora/node"
	"github.com/pchico83/elora/stack"
)

func main() {
	app := cli.NewApp()
	app.Name = "elora"
	app.Usage = "container orchestration tool (from docker-compose.yml to kubernetes API)"
	app.Version = "0.1.0"
	app.Author = "Pablo Chico de Guzman"
	app.Email = "pchico83@gmail.com"

	app.Before = func(c *cli.Context) error {
		client.Factory = client.DockerFactory{}
		log.SetOutput(os.Stderr)
		return nil
	}

	app.Commands = []cli.Command{
		node.Commands,
		image.Commands,
		stack.Commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
