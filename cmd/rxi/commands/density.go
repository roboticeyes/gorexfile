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
	Usage:  "Reduces density of pointLists to a specified grid size, an absolute amount or by percentage",
	Action: DensityActions,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "resolution",
			Value: false,
			Usage: "set the minimum distance between points in meters, this option yields the best results " +
				"because it evens out the density across the entire point cloud",
			Aliases: []string{"res"},
		},
		&cli.BoolFlag{
			Name:    "percent",
			Value:   false,
			Usage:   "reduction in percent every pointList will be reduced by",
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

	if !ctx.Bool("resolution") && !ctx.Bool("percent") && !ctx.Bool("absolute") {
		color.Red.Println("Please specify: resolution, percent or absolute")
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

	if ctx.Bool("resolution") {
		ReducePointListDensityVoxelBased(ctx, rexContent)
	} else {
		ReducePointListDensityNaive(ctx, rexContent)
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

type GridLocation struct {
	x int
	y int
	z int
}

type GridEntry struct {
	location mgl32.Vec3
	color    mgl32.Vec3
}

// ReducePointListDensityVoxelBased evens out and thins out pointLists
func ReducePointListDensityVoxelBased(ctx *cli.Context, rexContent *rexfile.File) {
	for i := 0; i < len(rexContent.PointLists); i++ {
		originalPointListLength := len(rexContent.PointLists[i].Points)
		voxelCellSize := float32(ctx.Float64("val"))

		//sort points into voxel construct
		voxelGrid := make(map[GridLocation][]GridEntry)

		for j := 0; j < originalPointListLength; j++ {
			gridLocation := GetGridLocationOfVec3(rexContent.PointLists[i].Points[j], voxelCellSize)
			gridEntry := GridEntry{rexContent.PointLists[i].Points[j], rexContent.PointLists[i].Colors[j]}

			if voxelGrid[gridLocation] != nil {
				voxelGrid[gridLocation] = append(voxelGrid[gridLocation], gridEntry)
			} else {
				voxelGrid[gridLocation] = []GridEntry{gridEntry}
			}
		}

		//process voxel averages
		averagedPoints := make([]mgl32.Vec3, len(voxelGrid))
		averagedColors := make([]mgl32.Vec3, len(voxelGrid))
		iter := 0

		for gridLocation, pointsInGridCell := range voxelGrid {
			summedLocation := mgl32.Vec3{}
			summedColor := mgl32.Vec3{}

			for j := 0; j < len(pointsInGridCell); j++ {
				summedLocation = mgl32.Vec3.Add(pointsInGridCell[j].location, summedLocation)
				summedColor = mgl32.Vec3.Add(pointsInGridCell[j].color, summedColor)
			}

			//translate voxel grid to real-world coords
			averagedLocation := mgl32.Vec3{float32(gridLocation.x) * voxelCellSize, float32(gridLocation.y) * voxelCellSize, float32(gridLocation.z) * voxelCellSize}

			//use this instead for avg location, but grid looks nicer
			//averagedLocation := mgl32.Vec3.Mul(summedLocation, float32(1)/float32(len(pointsInGridCell)))

			averagedColor := mgl32.Vec3.Mul(summedColor, float32(1)/float32(len(pointsInGridCell)))

			averagedPoints[iter] = averagedLocation
			averagedColors[iter] = averagedColor

			iter++
		}

		rexContent.PointLists[i].Points = averagedPoints
		rexContent.PointLists[i].Colors = averagedColors
	}
}

func GetGridLocationOfVec3(vec3 mgl32.Vec3, cellSize float32) GridLocation {
	return GridLocation{
		int(vec3[0] / cellSize),
		int(vec3[1] / cellSize),
		int(vec3[2] / cellSize),
	}
}

// ReducePointListDensityNaive heavily depends on pointList's spatial distribution to work correctly. works fine for laser scans to achieve desired counts/percentages.
func ReducePointListDensityNaive(ctx *cli.Context, rexContent *rexfile.File) {
	for i := 0; i < len(rexContent.PointLists); i++ {
		originalPointListLength := len(rexContent.PointLists[i].Points)
		reducedPointListLength := GetNewPointArraySize(originalPointListLength, ctx.Float64("val"), ctx.Bool("percent"))

		if originalPointListLength <= reducedPointListLength {
			color.Red.Println("Skipped pointList already smaller or equal to the desired size. PointListID:", rexContent.PointLists[i].ID)
			continue
		}

		tempListPoints := make([]mgl32.Vec3, reducedPointListLength)
		tempListColors := make([]mgl32.Vec3, reducedPointListLength)

		for j := 0; j < reducedPointListLength; j++ {
			adjustedIndex := int((float32(j) / float32(reducedPointListLength)) * float32(originalPointListLength))
			tempListPoints[j] = rexContent.PointLists[i].Points[adjustedIndex]
			tempListColors[j] = rexContent.PointLists[i].Colors[adjustedIndex]
		}

		rexContent.PointLists[i].Points = tempListPoints
		rexContent.PointLists[i].Colors = tempListColors
	}
}

func GetNewPointArraySize(pointListSize int, reduction float64, isPercentage bool) int {
	if isPercentage {
		return int(math.Ceil(float64(pointListSize) * reduction / 100))
	} else {
		return int(reduction)
	}
}
