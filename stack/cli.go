package stack

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var Commands = cli.Command{
	Name:  "stack",
	Usage: "manage stacks",
	Subcommands: []cli.Command{
		{
			Name:  "create",
			Usage: "create a stack",
			Action: func(c *cli.Context) {
				if len(c.Args()) != 2 {
					log.Fatal("'create' command accepts 2 arguments, the name of the stack and the path to the YAML elora template")
				}
				err := CreateCmd(c.Args()[0], c.Args()[1])
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:      "inspect",
			ShortName: "get",
			Usage:     "get the information of a stack",
			Action: func(c *cli.Context) {
				if len(c.Args()) != 1 {
					log.Fatal("'get' command only accepts one argument, the name of the stack")
				}
				err := InspectCmd(c.Args().First())
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:  "export",
			Usage: "get the elora YAML template of a stack",
			Action: func(c *cli.Context) {
				if len(c.Args()) != 1 {
					log.Fatal("'export' command only accepts one argument, the name of the stack")
				}
				err := ExportCmd(c.Args().First())
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:      "list",
			ShortName: "ls",
			Usage:     "list the elora current stacks",
			Action: func(c *cli.Context) {
				if len(c.Args()) != 0 {
					log.Fatal("'list' command does not have arguments")
				}
				err := ListCmd()
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:      "update",
			ShortName: "up",
			Usage:     "update a stack",
			Action: func(c *cli.Context) {
				if len(c.Args()) != 2 {
					log.Fatal("'update' command accepts 2 arguments, the name of the stack and the path to its new YAML elora template")
				}
				err := UpdateCmd(c.Args()[0], c.Args()[1])
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:      "remove",
			ShortName: "rm",
			Usage:     "remove a stack",
			Action: func(c *cli.Context) {
				if len(c.Args()) != 1 {
					log.Fatal("'remove' command only accepts one argument, the name of stack to remove")
				}
				err := RemoveCmd(c.Args().First())
				if err != nil {
					log.Fatal(err)
				}
			},
		},
	},
}
