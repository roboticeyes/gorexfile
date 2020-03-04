package translate

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/roboticeyes/gorexfile/encoding/rexfile"
)

const (
	dump = false
)

func TestObjRexTranslator(t *testing.T) {

	testFileName := filepath.Join("..", "examples", "models", "capsule", "capsule.obj")

	translator, err := NewObjToRexTranslator(testFileName, false)
	if err != nil {
		t.Fatalf("Cannot create translator: %v", err)
	}

	res, err := translator.Translate()
	if err != nil {
		t.Fatalf("Got error in OBJ translation: %v", err)
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

	if len(rexBuf.Bytes()) != 650098 {
		t.Fatal("REXfile size has unexpected size")
	}

	if dump {
		f, _ := os.Create("out.rex")
		f.Write(rexBuf.Bytes())
		defer f.Close()
	}
}
