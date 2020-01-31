package commands

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/urfave/cli/v2"
)

// BboxCommand displays bounding box of the given REXfile
var BboxCommand = &cli.Command{
	Name:   "bbox",
	Usage:  "Calculate the bounding box",
	Action: BboxAction,
}

// BboxAction calculates the bounding box
func BboxAction(ctx *cli.Context) error {

	rexHeader, rexContent, err := OpenRexFileFromContext(ctx)
	if err != nil {
		return err
	}

	fmt.Println(rexHeader)

	bbmin := mgl32.Vec3{mgl32.MaxValue, mgl32.MaxValue, mgl32.MaxValue}
	bbmax := mgl32.Vec3{mgl32.MinValue, mgl32.MinValue, mgl32.MinValue}

	if len(rexContent.Meshes) > 0 {
		for _, mesh := range rexContent.Meshes {
			for _, c := range mesh.Coords {
				for i := 0; i < 3; i++ {
					if c[i] < bbmin[i] {
						bbmin[i] = c[i]
					}
					if c[i] > bbmax[i] {
						bbmax[i] = c[i]
					}
				}
			}
		}
	}
	fmt.Println("BoundingBox MIN: ", bbmin)
	fmt.Println("BoundingBox MAX: ", bbmax)
	return nil
}
