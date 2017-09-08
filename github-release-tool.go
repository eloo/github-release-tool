package main

import (
	"github.com/eloo/github-release-tool/cmd"
	"github.com/urfave/cli"
	"os"
)

var version string

func main() {
	app := cli.NewApp()
	app.Version = version
	app.Commands = []cli.Command{
		cmd.CmdDownload,
	}

	app.Run(os.Args)
}
