package translate

import (
	"bufio"
	"fmt"
	"os"

	"github.com/roboticeyes/gorexfile/encoding/rexfile"
)

const (
	item = "e39718-b33b-4d2f-b30e-493264d55de3"
)

var (
	rexFile = "./" + item + ".rex"
	outDir  = "./" + item
)

func ExampleRexToWavefront() {

	file, err := os.Open(rexFile)
	if err != nil {
		fmt.Println("Cannot open input file", rexFile)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	d := rexfile.NewDecoder(r)
	rexHeader, rexContent, err := d.Decode()
	if err != nil && err.Error() != "unexpected EOF" {
		fmt.Println(err)
	}

	err = RexToWavefront(rexHeader, rexContent, outDir, item)
	if err != nil {
		fmt.Println(err)
	}
}
