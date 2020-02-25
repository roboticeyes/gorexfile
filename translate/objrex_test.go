package translate

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/roboticeyes/gorexfile/encoding/rexfile"
)

func TestObjRexTranslator(t *testing.T) {

	objReader := bytes.NewReader([]byte(objSample))
	mtlReader := bytes.NewReader([]byte(mtlSample))

	translator := NewObjToRexTranslator(objReader, mtlReader, false)

	res, err := translator.Translate()
	if err != nil {
		t.Fatalf("Got error in OBJ translation: %v\n", err)
	}

	if res.Header() == nil {
		t.Fatalf("Cannot get header\n")
	}

	var rexBuf bytes.Buffer
	e := rexfile.NewEncoder(&rexBuf)
	err = e.Encode(res)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

	f, _ := os.Create("out.rex")
	f.Write(rexBuf.Bytes())
	defer f.Close()
}

const objSample = `
mtllib untitled.mtl
o Cube
v 1.000000 1.000000 -1.000000
v 1.000000 -1.000000 -1.000000
v 1.000000 1.000000 1.000000
v 1.000000 -1.000000 1.000000
v -1.000000 1.000000 -1.000000
v -1.000000 -1.000000 -1.000000
v -1.000000 1.000000 1.000000
v -1.000000 -1.000000 1.000000
vt 0.875000 0.500000
vt 0.625000 0.750000
vt 0.625000 0.500000
vt 0.375000 1.000000
vt 0.375000 0.750000
vt 0.625000 0.000000
vt 0.375000 0.250000
vt 0.375000 0.000000
vt 0.375000 0.500000
vt 0.125000 0.750000
vt 0.125000 0.500000
vt 0.625000 0.250000
vt 0.875000 0.750000
vt 0.625000 1.000000
usemtl Material
s off
f 5/1 3/2 1/3
f 3/2 8/4 4/5
f 7/6 6/7 8/8
f 2/9 8/10 6/11
f 1/3 4/5 2/9
f 5/12 2/9 6/7
f 5/1 7/13 3/2
f 3/2 7/14 8/4
f 7/6 5/12 6/7
f 2/9 4/5 8/10
f 1/3 3/2 4/5
f 5/12 1/3 2/9
`

const mtlSample = `
# Material Count: 1

newmtl Material
Ns 323.999994
Ka 1.000000 1.000000 1.000000
Kd 0.800000 0.226274 0.393412
Ks 0.500000 0.500000 0.500000
Ke 0.000000 0.000000 0.000000
Ni 1.450000
d 1.000000
illum 2
`
