package commands

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
	"github.com/roboticeyes/gorexfile/translate"
	"github.com/urfave/cli/v2"
)

var supportedFileTypes []string

func init() {

	supportedFileTypes = []string{
		"obj",
	}
}

func isTypeSupported(t string) bool {

	for _, v := range supportedFileTypes {
		if v == t {
			return true
		}
	}
	return false
}

// ImportCommand displays file information
var ImportCommand = &cli.Command{
	Name:   "import",
	Usage:  "Imports a geometry file and stores it to the given REX file",
	Action: ImportAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "type",
			Aliases: []string{"t"},
			Usage:   "specifies the type of input (e.g. obj)",
		},
		&cli.BoolFlag{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "lists the supported input formats types",
		},
		&cli.BoolFlag{
			Name:  "verbose",
			Usage: "verbose output of conversion",
		},
		&cli.StringSliceFlag{
			Name:    "input",
			Aliases: []string{"i"},
			Usage:   "input file, can be used multiple times for more than one file",
		},
	},
}

// ImportAction is the default action, therefore it is set to public
func ImportAction(ctx *cli.Context) error {

	if ctx.Bool("list") == true {
		color.Green.Println("Supported file type(s):")
		for _, t := range supportedFileTypes {
			fmt.Println("\t-", t)
		}
		return nil
	}

	fileType := ctx.String("type")
	if fileType == "" {
		color.Red.Println("Please specify a file type")
		return nil
	}
	if !isTypeSupported(fileType) {
		color.Red.Printf("Type %s is not supported\n", fileType)
		return nil
	}

	// a slice of input files
	input := ctx.StringSlice("input")

	var result rexfile.File
	var err error

	switch fileType {
	case "obj":
		result, err = importObj(input, ctx.Bool("verbose"))
	}

	// Write REXfile
	var rexBuf bytes.Buffer
	e := rexfile.NewEncoder(&rexBuf)
	err = e.Encode(result)
	if err != nil {
		color.Red.Println("Error during encoding:", err)
		return err
	}

	rexFileName := ctx.Args().First()
	f, _ := os.Create(rexFileName)
	f.Write(rexBuf.Bytes())
	defer f.Close()
	color.Green.Println("Successfully written", rexFileName)
	return nil
}

func importObj(input []string, verbose bool) (rexfile.File, error) {

	var objReader, mtlReader io.Reader
	var err error
	for _, v := range input {
		if strings.ToLower(filepath.Ext(v)) == ".obj" {
			objReader, err = os.Open(v)
			if err != nil {
				return rexfile.File{}, err
			}
		} else if strings.ToLower(filepath.Ext(v)) == ".mtl" {
			mtlReader, err = os.Open(v)
			if err != nil {
				return rexfile.File{}, err
			}
		}
	}

	translator := translate.NewObjToRexTranslator(objReader, mtlReader, verbose)
	return translator.Translate()
}
