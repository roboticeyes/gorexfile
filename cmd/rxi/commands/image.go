package commands

import (
	"encoding/binary"
	"os"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

// ImageCommand extract an image from the image block
var ImageCommand = &cli.Command{
	Name:   "image",
	Usage:  "Extracts an image, if no output file given, content will be dumped to stdout",
	Action: ImageAction,
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "id",
			Usage: "Image block ID",
		},
		&cli.StringFlag{
			Name:  "o",
			Usage: "output file",
		},
	},
}

// ImageAction calculates the bounding box
func ImageAction(ctx *cli.Context) error {

	output := ctx.String("o")

	_, rexContent, err := OpenRexFileFromContext(ctx)
	if err != nil {
		return err
	}

	if len(rexContent.Images) < 1 {
		color.Red.Println("No image content found")
	}

	id := uint64(ctx.Int("id"))

	var outputFile *os.File
	if output != "" {
		outputFile, err = os.Create(output)
		if err != nil {
			color.Red.Println("Cannot create output file: ", err)
			return err
		}
		defer outputFile.Close()
	}

	for _, img := range rexContent.Images {
		if img.ID == id {
			if outputFile == nil {
				binary.Write(os.Stdout, binary.LittleEndian, img.Data)
			} else {
				binary.Write(outputFile, binary.LittleEndian, img.Data)
			}
		}
	}

	return nil
}
