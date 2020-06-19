package commands

import (
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/roboticeyes/gorexfile/translate"
	"github.com/urfave/cli/v2"
)

// ExportCommand moves the geometry according to the given parameters
var ExportCommand = &cli.Command{
	Name:   "export",
	Usage:  "Exports the REX mesh geometry into an OBJ file (no material!)",
	Action: ExportAction,
}

// ExportAction exports REXmesh geometries into an OBJ file
func ExportAction(ctx *cli.Context) error {

	output := ctx.Args().Get(1)

	if output == "" {
		color.Red.Println("Please specify an output file as second parameter")
		return fmt.Errorf("No output file specified")
	}

	rexHeader, rexContent, err := OpenRexFileFromContext(ctx)
	if err != nil {
		return err
	}

	fmt.Println(rexHeader)

	f, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	translate.RexToObj(rexHeader, rexContent, f)

	color.Green.Printf("Successfully exported to file %s\n", output)
	return nil
}
