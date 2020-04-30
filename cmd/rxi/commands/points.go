package commands

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

// PointsCommand displays a points block
var PointsCommand = &cli.Command{
	Name:   "points",
	Usage:  "Display a points data block",
	Action: PointsAction,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "id",
			Usage: "Points block ID",
		},
	},
}

// PointsAction calculates the bounding box
func PointsAction(ctx *cli.Context) error {

	rexHeader, rexContent, err := OpenRexFileFromContext(ctx)
	if err != nil {
		return err
	}

	fmt.Println(rexHeader)

	if len(rexContent.PointLists) < 1 {
		color.Red.Println("No points content found")
	}

	id := uint64(ctx.Int("id"))

	for _, s := range rexContent.PointLists {
		if s.ID == id {
			for _, p := range s.Points {
				fmt.Println(p[0], p[1], p[2])
			}
		}
	}

	return nil
}
