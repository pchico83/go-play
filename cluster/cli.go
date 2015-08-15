package cluster

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var Commands = cli.Command{
	Name:  "cluster",
	Usage: "manage d2k8 clusters",
	Subcommands: []cli.Command{
		{
			Name:  "create",
			Usage: "create a cluster",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "skip-consul",
					Usage: "'true' not to install consul agent for DNS",
				},
				cli.BoolFlag{
					Name:  "skip-weave",
					Usage: "'true' not to install weave router for SDN",
				},
				cli.StringFlag{
					Name:  "weave-pwd",
					Usage: "Password for weave traffic encryption.",
				},
				cli.BoolFlag{
					Name:  "skip-cleanup",
					Usage: "'true' not to install 'cleanup' service (cleans unused volumes and images)",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) != 1 {
					log.Fatal("'create' command only needs the name of the cluster to be created as a positional argument")
				}
				err := CreateCmd(c.Args().First(), !c.Bool("skip-consul"), !c.Bool("skip-weave"), c.String("weave-pwd"), !c.Bool("skip-cleanup"))
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:      "inspect",
			ShortName: "get",
			Usage:     "get the information of a cluster",
			Action: func(c *cli.Context) {
				if len(c.Args()) != 1 {
					log.Fatal("'inspect' command only needs the name of the cluster to retrieve as a postional argument")
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
			Usage:     "list the d2k8 current clusters",
			Action: func(c *cli.Context) {
				err := ListCmd()
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:      "remove",
			ShortName: "rm",
			Usage:     "remove a cluster",
			Action: func(c *cli.Context) {
				if len(c.Args()) != 1 {
					log.Fatal("'remove' command only needs the name of the cluster to be removed as a positional argument")
				}
				err := RemoveCmd(c.Args().First())
				if err != nil {
					log.Fatal(err)
				}
			},
		},
	},
}
