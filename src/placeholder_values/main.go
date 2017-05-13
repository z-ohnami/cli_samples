package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c", // conma !
			Usage: "Load configuration from `FILE`",
		},
	}

	app.Run(os.Args)
}
