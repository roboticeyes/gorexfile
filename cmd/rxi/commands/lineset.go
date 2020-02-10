package commands

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

// LineSetCommand displays a lineset block
var LineSetCommand = &cli.Command{
	Name:   "lines",
	Usage:  "Display a lineset data block",
	Action: LineSetAction,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "id",
			Usage: "LineSet block ID",
		},
	},
}

// LineSetAction calculates the bounding box
func LineSetAction(ctx *cli.Context) error {

	rexHeader, rexContent, err := OpenRexFileFromContext(ctx)
	if err != nil {
		return err
	}

	fmt.Println(rexHeader)

	if len(rexContent.LineSets) < 1 {
		color.Red.Println("No lineset content found")
	}

	id := uint64(ctx.Int("id"))

	for _, s := range rexContent.LineSets {
		if s.ID == id {
			for _, p := range s.Points {
				fmt.Println(p[0], p[1], p[2])
			}
		}
	}

	return nil
}
