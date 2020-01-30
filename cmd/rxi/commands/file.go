package commands

import (
	"bufio"
	"os"

	"github.com/gookit/color"
	"github.com/roboticeyes/gorexfile/encoding/rex"
)

// OpenRexFile tries to open up a REX file and performs all the error handling
func OpenRexFile(rexFile string) (*rex.Header, *rex.File, error) {

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
