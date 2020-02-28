package main

import (
	"bytes"
	"os"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
)

func main() {
	fileName := "testText.rex"

	text := rexfile.Text{ID: 1, Text: "rudi", Red: 23, Green: 34, Blue: 22, Alpha: 3, FontSize: 12}
	text.Position = mgl32.Vec3{0.5, 1.0, 0.0}

	rexFile := rexfile.File{}
	rexFile.Texts = append(rexFile.Texts, text)

	var buf bytes.Buffer
	e := rexfile.NewEncoder(&buf)
	err := e.Encode(rexFile)
	if err != nil {
		panic(err)
	}

	f, _ := os.Create(fileName)
	f.Write(buf.Bytes())
	defer f.Close()
}
