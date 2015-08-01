package image

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	"path/filepath"
)

var Commands = cli.Command{
	Name:  "image",
	Usage: "manage elora images",
	Subcommands: []cli.Command{
		{
			Name:  "create",
			Usage: "create a private image",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "user, u",
					Usage: "Specify the username of the image",
				},
				cli.StringFlag{
					Name:  "pwd, p",
					Usage: "Specify the password to access the image",
				},
				cli.StringFlag{
					Name:  "url",
					Usage: "Specify the url where the image is available",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) != 1 {
					log.Fatal("'create' command only needs the name of the image to be created as a positional argument")
				}
				if c.String("user") == "" {
					log.Fatal("flag 'user' is mandatory")
				}
				if c.String("pwd") == "" {
					log.Fatal("flag 'pwd' is mandatory")
				}
				if c.String("url") == "" {
					log.Fatal("flag 'url' is mandatory")
				}
				if !filepath.IsAbs(c.String("cert")) {
					log.Fatal("'cert' must be an absolute path")
				}
				err := CreateCmd(c.Args().First(), c.String("user"), c.String("pwd"), c.String("url"))
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:      "inspect",
			ShortName: "get",
			Usage:     "get the information of a private image",
			Action: func(c *cli.Context) {
				if len(c.Args()) != 1 {
					log.Fatal("'inspect' command only needs the name of the image to retrieve as a postional argument")
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
			Usage:     "list the elora current private images",
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
			Usage:     "update a private image",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "user, u",
					Usage: "Specify the username of the image",
				},
				cli.StringFlag{
					Name:  "pwd, p",
					Usage: "Specify the password to access the image",
				},
				cli.StringFlag{
					Name:  "url",
					Usage: "Specify the url where the image is available",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) != 1 {
					log.Fatal("'update' command only needs the name of the image to be updated as a positional argument")
				}
				err := UpdateCmd(c.Args().First(), c.String("user"), c.String("pwd"), c.String("url"))
				if err != nil {
					log.Fatal(err)
				}
			},
		},
		{
			Name:      "remove",
			ShortName: "rm",
			Usage:     "remove a private image",
			Action: func(c *cli.Context) {
				if len(c.Args()) != 1 {
					log.Fatal("'remove' command only needs the name of the image to be removed as a positional argument")
				}
				err := RemoveCmd(c.Args().First())
				if err != nil {
					log.Fatal(err)
				}
			},
		},
	},
}
