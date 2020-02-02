package commands

import (
	"bytes"
	"fmt"
	"os"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/gookit/color"
	"github.com/roboticeyes/gorexfile/encoding/rex"
	"github.com/urfave/cli/v2"
)

// TranslateCommand moves the geometry according to the given parameters
var TranslateCommand = &cli.Command{
	Name:   "translate",
	Usage:  "Translates the whole REXfile",
	Action: TranslateAction,
	Flags: []cli.Flag{
		&cli.Float64Flag{
			Name:  "x",
			Usage: "x translation in world space [m]",
		},
		&cli.Float64Flag{
			Name:  "y",
			Usage: "y translation in world space [m]",
		},
		&cli.Float64Flag{
			Name:  "z",
			Usage: "z translation in world space (up) [m]",
		},
	},
}

// TranslateAction calculates the bounding box
func TranslateAction(ctx *cli.Context) error {

	output := ctx.Args().Get(1)

	if output == "" {
		color.Red.Println("Please specify an output file as second parameter")
		return fmt.Errorf("No output file specified")
	}

	rexHeader, rexContent, err := OpenRexFileFromContext(ctx)
	if err != nil {
		return err
	}

	// Convert into world space
	x := float32(ctx.Float64("x"))
	y := float32(ctx.Float64("z"))
	z := float32(ctx.Float64("y"))

	fmt.Println(rexHeader)

	translate := mgl32.Vec3{x, y, z}

	if len(rexContent.Meshes) > 0 {
		for i := 0; i < len(rexContent.Meshes); i++ {
			for j := 0; j < len(rexContent.Meshes[i].Coords); j++ {
				rexContent.Meshes[i].Coords[j] = rexContent.Meshes[i].Coords[j].Add(translate)
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
	e := rex.NewEncoder(&buf)
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
