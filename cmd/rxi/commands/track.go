package commands

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

// TrackCommand displays a track block
var TrackCommand = &cli.Command{
	Name:   "track",
	Usage:  "Display a track data block",
	Action: TrackAction,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "id",
			Usage: "Track block ID",
		},
	},
}

// TrackAction calculates the bounding box
func TrackAction(ctx *cli.Context) error {

	rexHeader, rexContent, err := OpenRexFileFromContext(ctx)
	if err != nil {
		return err
	}

	fmt.Println(rexHeader)

	if len(rexContent.Tracks) < 1 {
		color.Red.Println("No track content found")
	}

	id := uint64(ctx.Int("id"))

	for _, track := range rexContent.Tracks {
		if track.ID == id {
			fmt.Printf("Number of points %d\n", track.NrOfPoints)
			for _, elem := range track.Points {
				fmt.Printf("v %5.2f %5.2f %5.2f\n", elem.Point[0], elem.Point[1], elem.Point[2])
				fmt.Printf("n %5.2f %5.2f %5.2f\n", elem.NormalVec[0], elem.NormalVec[1], elem.NormalVec[2])
				fmt.Printf("c %5.2f\n\n", elem.Confidence)
			}
		}
	}

	return nil
}
