// Copyright 2020 Robotic Eyes. All rights reserved.

package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/roboticeyes/gorexfile/cmd/rxi/commands"
	"github.com/roboticeyes/gorexfile/encoding/rex"
	"github.com/urfave/cli/v2"
)

const (
	version = "v0.1"
)

var (
	rexHeader  *rex.Header
	rexContent *rex.File
	// Version string from ldflags
	Version string
	// Build string from ldflags
	Build string
)

// the help text that gets displayed when something goes wrong or when you run
// help
const helpText = `
rxi - show REX file infos

actions:

  rxi -v                    prints version
  rxi help                  print this help

  rxi "file.rex"            show all REX blocks
  rxi bbox "file.rex"       displays the bounding box of the rex file

  rxi img ID "file.rex"     extract the given image and dump it to stdout (pipe to a viewer, e.g. | feh -)
  rxi mesh ID "file.rex"    extract the mesh block and dump it to stdout
  rxi lines ID "file.rex"   extract the lineset block and dump it to stdout

  rxi scale <factor> "input.rex" "output.rex" scales all mesh vertices by the given factor (e.g. 0.001)
`

// help prints the help text to stdout
func help(exit int) {
	fmt.Println(helpText)
	os.Exit(exit)
}

func openRexFile(rexFile string) {
	file, err := os.Open(rexFile)
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(file)
	d := rex.NewDecoder(r)
	rexHeader, rexContent, err = d.Decode()
	if err != nil && err.Error() != "unexpected EOF" {
		panic(err)
	}
}

// dumps the image to stdout (you can pipe it to an image viewer)
func rexExtractImage(rexFile, idString string) {
	openRexFile(rexFile)
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		panic(err)
	}

	for _, img := range rexContent.Images {
		if img.ID == id {
			binary.Write(os.Stdout, binary.LittleEndian, img.Data)
		}
	}
}

// dumps the mesh data block
func rexShowMesh(rexFile, idString string) {
	openRexFile(rexFile)
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		panic(err)
	}

	for _, mesh := range rexContent.Meshes {
		if mesh.ID == id {
			fmt.Println(mesh)
		}
	}
}

// dumps the lineset data blocks
func rexShowLines(rexFile, idString string) {
	openRexFile(rexFile)
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		panic(err)
	}

	for _, lineset := range rexContent.LineSets {
		if lineset.ID == id {
			for _, p := range lineset.Points {
				fmt.Printf("v %5.2f %5.2f %5.2f\n", p[0], p[1], p[2])

			}
		}
	}
}

func rexTranslate(rexFile string, x, y, z float32, output string) {
	openRexFile(rexFile)
	fmt.Println(rexHeader)

	translate := mgl32.Vec3{x, y, z}

	if len(rexContent.Meshes) > 0 {
		for i := 0; i < len(rexContent.Meshes); i++ {
			for j := 0; j < len(rexContent.Meshes[i].Coords); j++ {
				rexContent.Meshes[i].Coords[j] = rexContent.Meshes[i].Coords[j].Add(translate)
			}
		}
	}

	// create new file
	f, err := os.Create(output)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var buf bytes.Buffer
	e := rex.NewEncoder(&buf)
	err = e.Encode(*rexContent)
	if err != nil {
		panic(err)
	}
	n, err := f.Write(buf.Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully written %d bytes to file %s\n", n, output)

}

func rexBbox(rexFile string) {
	openRexFile(rexFile)

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
}

func rexScaleVertices(factor float32, input, output string) {

	openRexFile(input)

	for _, m := range rexContent.Meshes {
		for i := 0; i < len(m.Coords); i++ {
			for j := 0; j < 3; j++ {
				m.Coords[i][j] = m.Coords[i][j] * factor
			}
		}
	}

	var buf bytes.Buffer
	e := rex.NewEncoder(&buf)
	err := e.Encode(*rexContent)
	if err != nil {
		panic(err)
	}

	f, _ := os.Create(output)
	f.Write(buf.Bytes())
	defer f.Close()
}

func main() {

	app := cli.NewApp()
	app.Name = "rxi"
	app.Usage = "REXfile inspector"
	app.Version = version
	app.Copyright = "(c) 2020 Robotic Eyes GmbH"
	app.EnableBashCompletion = true
	app.Flags = GlobalFlags

	app.Action = func(c *cli.Context) error {
		return commands.InfoAction(c)
	}

	app.Commands = []*cli.Command{
		commands.InfoCommand,
		commands.BboxCommand,
		commands.TranslateCommand,
		commands.ImageCommand,
		commands.MeshCommand,
		commands.ScaleCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		// Currently ignore errors here
	}
}
