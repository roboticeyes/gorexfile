package commands

import (
	"bytes"
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
	"github.com/urfave/cli/v2"
)

// ScaleCommand scales all meshes in the given file
var ScaleCommand = &cli.Command{
	Name:   "scale",
	Usage:  "Scales all mesh blocks with the given factor",
	Action: ScaleAction,
	Flags: []cli.Flag{
		&cli.Float64Flag{
			Name:  "factor",
			Value: 1.0,
			Usage: "Scaling factor (1.0 means no scaling)",
		},
	},
}

// ScaleAction calculates the bounding box
func ScaleAction(ctx *cli.Context) error {

	output := ctx.Args().Get(1)

	if output == "" {
		color.Red.Println("Please specify an output file as second parameter")
		return fmt.Errorf("No output file specified")
	}

	factor := float32(ctx.Float64("factor"))

	if factor == 0.0 {
		color.Red.Println("Scale must be greater that zero")
		return fmt.Errorf("Scale is zero")
	}

	_, rexContent, err := OpenRexFileFromContext(ctx)
	if err != nil {
		return err
	}

	for _, m := range rexContent.Meshes {
		for i := 0; i < len(m.Coords); i++ {
			for j := 0; j < 3; j++ {
				m.Coords[i][j] = m.Coords[i][j] * factor
			}
		}
	}

	// create new file
	f, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var buf bytes.Buffer
	e := rexfile.NewEncoder(&buf)
	err = e.Encode(*rexContent)
	if err != nil {
		panic(err)
	}
	n, err := f.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
	color.Green.Printf("Successfully written %d bytes to file %s\n", n, output)
	return nil
}
