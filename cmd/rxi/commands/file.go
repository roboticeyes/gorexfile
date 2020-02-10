package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
	"github.com/urfave/cli/v2"
)

// OpenRexFileFromContext tries to open up a REX file and performs all the error handling
func OpenRexFileFromContext(ctx *cli.Context) (*rexfile.Header, *rexfile.File, error) {

	rexFile := ctx.Args().First()

	if rexFile == "" {
		color.Red.Println("\nPlease specify a REX file ...\n")
		fmt.Println("For more information please run rxi --help or rxi <command> --help")
		return nil, nil, fmt.Errorf("REXfile is missing")
	}

	file, err := os.Open(rexFile)
	if err != nil {
		color.Red.Println("Cannot open file ", rexFile)
		return nil, nil, err
	}
	r := bufio.NewReader(file)
	d := rexfile.NewDecoder(r)
	rexHeader, rexContent, err := d.Decode()
	if err != nil && err.Error() != "unexpected EOF" {
		color.Red.Println("Error during decoding occurs", err)
	}
	return rexHeader, rexContent, err
}
