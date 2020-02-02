package commands

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

// MeshCommand displays a mesh block
var MeshCommand = &cli.Command{
	Name:   "mesh",
	Usage:  "Display a mesh data block",
	Action: MeshAction,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "id",
			Usage: "Mesh block ID",
		},
	},
}

// MeshAction calculates the bounding box
func MeshAction(ctx *cli.Context) error {

	rexHeader, rexContent, err := OpenRexFileFromContext(ctx)
	if err != nil {
		return err
	}

	fmt.Println(rexHeader)

	if len(rexContent.Meshes) < 1 {
		color.Red.Println("No mesh content found")
	}

	id := uint64(ctx.Int("id"))

	for _, mesh := range rexContent.Meshes {
		if mesh.ID == id {
			fmt.Println(mesh)
		}
	}

	return nil
}
