package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/roboticeyes/gorexfile/encoding/rex"
	"github.com/urfave/cli/v2"
)

// OpenRexFileFromContext tries to open up a REX file and performs all the error handling
func OpenRexFileFromContext(ctx *cli.Context) (*rex.Header, *rex.File, error) {

	rexFile := ctx.Args().First()

	if rexFile == "" {
		color.Red.Println("Please specify a REX file ...")
		return nil, nil, fmt.Errorf("REXfile is missing")
	}

	file, err := os.Open(rexFile)
	if err != nil {
		color.Red.Println("Cannot open file ", rexFile)
		return nil, nil, err
	}
	r := bufio.NewReader(file)
	d := rex.NewDecoder(r)
	rexHeader, rexContent, err := d.Decode()
	if err != nil && err.Error() != "unexpected EOF" {
		color.Red.Println("Error during decoding occurs", err)
	}
	return rexHeader, rexContent, err
}
