package main

import (
	"os"

	"github.com/roboticeyes/gorexfile/loader/obj"
)

func main() {

	if len(os.Args) < 2 {
		panic("Missing file")
	}
	baseName := os.Args[1]
	// dec, err := obj.Decode("modell.obj", "modell.mtl")
	dec, err := obj.Decode(baseName+".obj", baseName+".mtl")
	if err != nil {
		panic(err)
	}

	writer, err := os.Create("test.rex")
	if err != nil {
		panic(err)
	}

	err = dec.WriteREX(writer)
	if err != nil {
		panic(err)
	}
}
