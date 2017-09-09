package main

import (
	"github.com/eloo/github-release-tool/src/cmd"
	"github.com/urfave/cli"
	"os"

	"github.com/eloo/github-release-tool/src/log"
)

var version string

func main() {
	app := cli.NewApp()
	app.Version = version
	app.Commands = []cli.Command{
		cmd.Download,
	}
	app.Run(os.Args)
}

func init() {
	log.DefaultCallerDepth = 3
	log.ShowDepth = true
}
