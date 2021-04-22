package commands

import (
	"bytes"
	"fmt"
	"github.com/gookit/color"
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
	"github.com/urfave/cli/v2"
	"os"
)

// MirrorCommand mirrors points and meshes
var MirrorCommand = &cli.Command{
	Name:   "mirror",
	Usage:  "Mirrors meshes and points in global space",
	Action: MirrorActions,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "x",
			Value: false,
			Usage: "mirror along output x-axis",
		},
		&cli.BoolFlag{
			Name:  "y",
			Value: false,
			Usage: "mirror along output y-axis",
		},
		&cli.BoolFlag{
			Name:  "z",
			Value: false,
			Usage: "mirror along output z-axis",
		},
	},
}

// MirrorActions mirrors points and meshes along the specified axis
func MirrorActions(ctx *cli.Context) error {

	output := ctx.Args().Get(1)

	if output == "" {
		color.Red.Println("Please specify an output file as second parameter")
		return fmt.Errorf("No output file specified")
	}

	_, rexContent, err := OpenRexFileFromContext(ctx)
	if err != nil {
		return err
	}

	mirrorVector := [3]float32{
		GetValFromMirrorArg(ctx.Bool("x")),
		GetValFromMirrorArg(ctx.Bool("y")),
		GetValFromMirrorArg(ctx.Bool("z")),
	}

	for _, m := range rexContent.Meshes {
		for i := 0; i < len(m.Coords); i++ {
			for j := 0; j < 3; j++ {
				m.Coords[i][j] *= mirrorVector[j]
			}
		}
	}

	for _, s := range rexContent.PointLists {
		for i := 0; i < len(s.Points); i++ {
			for j := 0; j < 3; j++ {
				s.Points[i][j] *= mirrorVector[j]
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

func GetValFromMirrorArg(isFlipped bool) float32 {
	if isFlipped {
		return -1
	} else {
		return 1
	}
}
