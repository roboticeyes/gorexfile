// This binary can be used to work with a REXfile on the command prompt.
// Get additional information by using the --help argument.
// The binary can be compiled either by using go build or by the "make" command.
package main

import (
	"os"

	"github.com/roboticeyes/gorexfile/cmd/rxi/commands"
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
	"github.com/urfave/cli/v2"
)

var (
	rexHeader  *rexfile.Header
	rexContent *rexfile.File
	// Version string from ldflags
	Version string
	// Build string from ldflags
	Build string
)

func main() {

	app := cli.NewApp()
	app.Name = "rxi"
	app.Usage = "REXfile inspector"
	app.Version = Version + " - build " + Build
	app.Description = `
	This tool displays information about a given REX file. See all the available parameters
	by using the rxi --help command.

	Example: rxi test.rex
	`
	app.Copyright = "(c) 2020 Robotic Eyes GmbH"
	app.EnableBashCompletion = true

	app.Action = func(c *cli.Context) error {
		return commands.InfoAction(c)
	}

	app.Commands = []*cli.Command{
		commands.BboxCommand,
		commands.ImageCommand,
		commands.InfoCommand,
		commands.LineSetCommand,
		commands.PointsCommand,
		commands.MeshCommand,
		commands.ScaleCommand,
		commands.TextCommand,
		commands.TrackCommand,
		commands.TranslateCommand,
		commands.ExportCommand,
		commands.MirrorCommand,
		commands.DensityCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		// Currently ignore errors here
	}
}
