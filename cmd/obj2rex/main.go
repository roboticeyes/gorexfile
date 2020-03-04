package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

var (
	// Version string from ldflags
	Version string
	// Build string from ldflags
	Build string
)

func main() {

	app := cli.NewApp()
	app.Name = "obj2rex"
	app.Usage = "Convert an OBJ file to a REX file"
	app.Version = Version + " - build " + Build
	app.HideHelp = true
	app.Description = `
	This tool converts an OBJ file into a REX file. Please specify the OBJ file
	with the -i flag and the REX file with the -o parameter.

	Example: obj2rex -i test.obj -o test.rex
	`
	app.Copyright = "(c) 2020 Robotic Eyes GmbH"
	app.EnableBashCompletion = true

	app.Action = func(c *cli.Context) error {
		return convert(c)
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "input",
			Aliases:  []string{"i"},
			Required: true,
			Usage:    "Input OBJ file",
		},
		&cli.StringFlag{
			Name:     "output",
			Required: true,
			Aliases:  []string{"o"},
			Usage:    "Output REX file",
		},
		&cli.BoolFlag{
			Name:  "verbose",
			Value: false,
			Usage: "verbose output of conversion",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		// Currently ignore errors here
	}
}
