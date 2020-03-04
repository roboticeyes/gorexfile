package main

import (
	"bytes"
	"os"

	"github.com/gookit/color"
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
	"github.com/roboticeyes/gorexfile/translate"
	"github.com/urfave/cli/v2"
)

func convert(ctx *cli.Context) error {

	inputFileName := ctx.String("input")
	outputFileName := ctx.String("output")
	verbose := ctx.Bool("verbose")

	translator, err := translate.NewObjToRexTranslator(inputFileName, verbose)
	if err != nil {
		color.Red.Println("Cannot create translator:", err)
		return err
	}

	rexFile, err := translator.Translate()
	if err != nil {
		color.Red.Println("Cannot perform translation:", err)
		return err
	}

	// Write REXfile
	var rexBuf bytes.Buffer
	e := rexfile.NewEncoder(&rexBuf)
	err = e.Encode(rexFile)
	if err != nil {
		color.Red.Println("Error during encoding:", err)
		return err
	}

	f, err := os.Create(outputFileName)
	if err != nil {
		color.Red.Println("Cannot write output file:", err)
		return err
	}
	_, err = f.Write(rexBuf.Bytes())
	if err != nil {
		color.Red.Println("Error during write operation:", err)
		return err
	}
	defer f.Close()
	color.Green.Println("Successfully written", outputFileName)
	return nil
}
