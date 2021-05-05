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
	bbmax := mgl32.Vec3{-mgl32.MaxValue, -mgl32.MaxValue, -mgl32.MaxValue}

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
	if len(rexContent.PointLists) > 0 {
		for _, pl := range rexContent.PointLists {
			for _, c := range pl.Points {
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
	fmt.Println("BoundingBox in worldspace (z is up)")
	fmt.Printf("\tmin: %9.2f %9.2f %9.2f\n", bbmin.X(), bbmin.Z(), bbmin.Y())
	fmt.Printf("\tmax: %9.2f %9.2f %9.2f\n", bbmax.X(), bbmax.Z(), bbmax.Y())
	return nil
}
