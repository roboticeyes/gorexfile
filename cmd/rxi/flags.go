package main

import (
	"github.com/urfave/cli/v2"
)

// GlobalFlags are global CLI flags
var GlobalFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:  "verbose",
		Usage: "set output to verbose",
	},
}
