package node

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"path/filepath"
	"strings"
)

var Commands = cli.Command{
	Name:  "node",
	Usage: "manage elora nodes",
	Subcommands: []cli.Command{
		{
			Name:  "create",
			Usage: "create a node",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "url",
					Usage:  "Specify the url of the node",
					EnvVar: "DOCKER_HOST",
				},
				cli.StringFlag{
					Name:   "cert",
					Usage:  "Specify the docker certificate path",
					EnvVar: "DOCKER_CERT_PATH",
				},
				cli.StringFlag{
					Name:  "key, k",
					Usage: "Specify the SSH certificate path",
				},
				cli.StringFlag{
					Name:  "cluster, c",
					Usage: "Specify the node cluster",
				},
				cli.StringFlag{
					Name:  "tags, t",
					Usage: "Specify the node tags separated by commas",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) != 1 {
					log.Fatal("'create' command only needs the name of the node to be created as a positional argument")
				}
				if c.String("url") == "" {
					log.Fatal("flag 'url' is mandatory, or set the variable 'DOCKER_HOST'")
				}
				if !filepath.IsAbs(c.String("cert")) {
					log.Fatal("'cert' must be an absolute path")
				}
				key := c.String("key")
				if key == "" {
					if strings.Contains(c.String("cert"), "/.docker/machine/machines/") {
						key = filepath.Join(c.String("cert"), "id_rsa")
					} else {
						log.Fatal("flag 'key' is mandatory, or set the variable 'DOCKER_CERT_PATH' of a machine created with docker-machine")
					}
				}
				if !filepath.IsAbs(key) {
					log.Fatal("'key' must be an absolute path")
				}
				cluster := c.String("cluster")
				if cluster == "" {
					cluster = "default"
				}
				tags := strings.Split(c.String("tags"), ",")
				err := CreateCmd(c.Args().First(), c.String("url"), c.String("cert"), key, cluster, tags)
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:  "env",
			Usage: "show docker env variables. Configure the docker client to connect to this node by executing '$(elora node env test)'.",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "unset, u",
					Usage: "show unset commands",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) != 1 {
					log.Fatal("'env' command only needs the name of the node to retrieve as a postional argument")
				}
				err := EnvCmd(c.Args().First(), c.Bool("unset"))
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:      "inspect",
			ShortName: "get",
			Usage:     "get the information of a node",
			Action: func(c *cli.Context) {
				if len(c.Args()) != 1 {
					log.Fatal("'inspect' command only needs the name of the node to retrieve as a postional argument")
				}
				err := InspectCmd(c.Args().First())
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:      "list",
			ShortName: "ls",
			Usage:     "list the elora current nodes",
			Action: func(c *cli.Context) {
				err := ListCmd()
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:      "update",
			ShortName: "up",
			Usage:     "update a node",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "url",
					Usage: "Specify the url of the node",
				},
				cli.StringFlag{
					Name:  "cert, c",
					Usage: "Specify the docker certificate path",
				},
				cli.StringFlag{
					Name:  "key, k",
					Usage: "Specify the SSH certificate path",
				},
				cli.StringFlag{
					Name:  "tags, t",
					Usage: "Specify the host tags separated by commas",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) != 1 {
					log.Fatal("'update' command only needs the name of the node to be updated as a positional argument")
				}
				key := c.String("key")
				if key == "" {
					if strings.Contains(c.String("cert"), "/.docker/machine/machines/") {
						key = filepath.Join(c.String("cert"), "id_rsa")
					}
				}
				if !filepath.IsAbs(key) {
					log.Fatal("'key' must be an absolute path")
				}
				tags := strings.Split(c.String("tags"), ",")
				err := UpdateCmd(c.Args().First(), c.String("url"), c.String("cert"), key, tags)
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:      "remove",
			ShortName: "rm",
			Usage:     "remove a node",
			Action: func(c *cli.Context) {
				if len(c.Args()) != 1 {
					log.Fatal("'remove' command only needs the name of the node to be removed as a positional argument")
				}
				err := RemoveCmd(c.Args().First())
				if err != nil {
					log.Fatal(err)
				}
			},
		},
	},
}
