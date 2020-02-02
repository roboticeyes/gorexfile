// Copyright 2020 Robotic Eyes. All rights reserved.

package main

import (
	"os"

	"github.com/roboticeyes/gorexfile/cmd/rxi/commands"
	"github.com/roboticeyes/gorexfile/encoding/rex"
	"github.com/urfave/cli/v2"
)

const (
	version = "v0.2"
)

var (
	rexHeader  *rex.Header
	rexContent *rex.File
	// Version string from ldflags
	Version string
	// Build string from ldflags
	Build string
)

func main() {

	app := cli.NewApp()
	app.Name = "rxi"
	app.Usage = "REXfile inspector"
	app.Version = version
	app.Copyright = "(c) 2020 Robotic Eyes GmbH"
	app.EnableBashCompletion = true

	app.Action = func(c *cli.Context) error {
		return commands.InfoAction(c)
	}

	app.Commands = []*cli.Command{
		commands.InfoCommand,
		commands.BboxCommand,
		commands.TranslateCommand,
		commands.ImageCommand,
		commands.MeshCommand,
		commands.ScaleCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		// Currently ignore errors here
	}
}
