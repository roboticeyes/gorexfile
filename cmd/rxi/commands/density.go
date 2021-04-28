package commands

import (
	"bytes"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gookit/color"
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
	"github.com/urfave/cli/v2"
	"math"
	"os"
)

// DensityCommand reduces density of pointLists
var DensityCommand = &cli.Command{
	Name:   "density",
	Usage:  "Reduces density of pointLists to an absolute amount or by percentage",
	Action: DensityActions,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "percent",
			Value:   false,
			Usage:   "reduction in percent to which every pointList will be reduced",
			Aliases: []string{"pct"},
		},
		&cli.BoolFlag{
			Name:    "absolute",
			Value:   false,
			Usage:   "amount to which every pointList will be reduced",
			Aliases: []string{"abs"},
		},
		&cli.Float64Flag{
			Name:    "value",
			Usage:   "reduction amount",
			Aliases: []string{"val"},
		},
	},
}

// MirrorActions reduces density of pointLists
func DensityActions(ctx *cli.Context) error {

	output := ctx.Args().Get(1)

	if ctx.Bool("percent") && ctx.Bool("absolute") {
		color.Red.Println("Percent and absolute are mutually exclusive")
		return fmt.Errorf("Mutally exclusive arguments given")
	}

	if !ctx.Bool("percent") && !ctx.Bool("absolute") {
		color.Red.Println("Please specify either percent or absolute argument")
		return fmt.Errorf("No density arguments given")
	}

	if output == "" {
		color.Red.Println("Please specify an output file as second parameter")
		return fmt.Errorf("No output file specified")
	}

	_, rexContent, err := OpenRexFileFromContext(ctx)
	if err != nil {
		return err
	}

	if len(rexContent.PointLists) == 0 {
		color.Red.Println("REX file must have at least one PointList. Density function only affects PointLists")
		return fmt.Errorf("File contains no PointLists")
	}

	for i := 0; i < len(rexContent.PointLists); i++ {
		originalPointListLength := len(rexContent.PointLists[i].Points)
		reducedPointListLength := GetNewPointListSize(originalPointListLength, ctx.Float64("val"), ctx.Bool("percent"))

		if originalPointListLength <= reducedPointListLength {
			color.Red.Println("Skipped pointList already smaller or equal to the desired size. PointListID:", rexContent.PointLists[i].ID)
			continue
		}

		tempListPoints := make([]mgl32.Vec3, reducedPointListLength)
		tempListColors := make([]mgl32.Vec3, reducedPointListLength)

		for j := 0; j < reducedPointListLength; j++ {
			adjustedIndex := int((float64(j) / float64(reducedPointListLength)) * float64(originalPointListLength))
			tempListPoints[j] = rexContent.PointLists[i].Points[adjustedIndex]
			tempListColors[j] = rexContent.PointLists[i].Colors[adjustedIndex]
		}

		rexContent.PointLists[i].Points = tempListPoints
		rexContent.PointLists[i].Colors = tempListColors
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

func GetNewPointListSize(pointListSize int, reduction float64, isPercentage bool) int {
	if isPercentage {
		return int(math.Ceil(float64(pointListSize) * reduction / 100))
	} else {
		return int(reduction)
	}
}
