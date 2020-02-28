package commands

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

// TextCommand displays a text block
var TextCommand = &cli.Command{
	Name:   "text",
	Usage:  "Display a text data block",
	Action: TextAction,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "id",
			Usage: "Text block ID",
		},
	},
}

// TextAction calculates the bounding box
func TextAction(ctx *cli.Context) error {

	rexHeader, rexContent, err := OpenRexFileFromContext(ctx)
	if err != nil {
		return err
	}

	fmt.Println(rexHeader)

	if len(rexContent.Texts) < 1 {
		color.Red.Println("No text content found")
	}

	id := uint64(ctx.Int("id"))

	for _, text := range rexContent.Texts {
		if text.ID == id {
			fmt.Printf("RGB %5.2f %5.2f %5.2f Alpha %5.2f\n", text.Red, text.Green, text.Blue, text.Alpha)
			fmt.Printf("Position %5.2f %5.2f %5.2f\n", text.Position[0], text.Position[1], text.Position[2])
			fmt.Printf("FontSize %5.2f\n", text.FontSize)
			fmt.Printf("Text %s\n\n", text.Text)
		}
	}

	return nil
}
