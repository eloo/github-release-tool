package main

import (
	"fmt"
	"os"

	"github.com/eloo/github-release-tool/cmd"
	"github.com/eloo/github-release-tool/log"
)

var version string

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	log.DefaultCallerDepth = 3
	log.ShowDepth = true
}
