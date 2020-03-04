package translate

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/breiting/gwob"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gookit/color"
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
)

const (
	objDefaultName = "objfile"
)

// ObjToRex container
type ObjToRex struct {
	objReader io.Reader
	inputPath string
	verbose   bool
}

// NewObjToRexTranslator returns a new translator for converting OBJ files to REX files.
// It tries to open up the files and creates the according reader
func NewObjToRexTranslator(inputFileName string, verbose bool) (*ObjToRex, error) {

	var err error
	o := &ObjToRex{
		inputPath: filepath.Dir(inputFileName),
		verbose:   verbose,
	}

	// Open OBJ file
	o.objReader, err = os.Open(inputFileName)
	if err != nil {
		color.Red.Println("Cannot read input file:", err)
		return nil, err
	}

	return o, nil
}

// Translate implements the translator interface
func (o *ObjToRex) Translate() (rexfile.File, error) {

	var rexFile rexfile.File

	options := gwob.ObjParserOptions{
		LogStats: o.verbose,
		Logger: func(msg string) {
			fmt.Printf("Translator: %s\n", msg)
		},
	}

	obj, err := gwob.NewObjFromReader(objDefaultName, o.objReader, &options)
	if err != nil {
		return rexFile, err
	}

	// Try to open the MTL file from the same directory
	var mtl gwob.MaterialLib

	mtlFileName := filepath.Join(o.inputPath, obj.Mtllib)
	mtlReader, err := os.Open(mtlFileName)
	if err != nil {
		color.Cyan.Println("Cannot open MTL file", mtlFileName, " skipping material information.")
	} else {
		mtl, err = gwob.ReadMaterialLibFromReader(mtlReader, &options)
		if err != nil {
			color.Cyan.Println("Error reading MTL file: ", err)
		}
	}

	vtxMap := make(map[int]int)
	texMap := make(map[string]uint64)
	dataID := uint64(1)

	for _, g := range obj.Groups {

		materialID := dataID
		if mtlMat, exists := mtl.Lib[g.Usemtl]; exists {
			rexMat := rexfile.NewMaterial(materialID)
			rexMat.KdRgb = mgl32.Vec3{mtlMat.Kd[0], mtlMat.Kd[1], mtlMat.Kd[2]}

			// add texture
			if mtlMat.MapKd != "" {

				// check if already available as block
				imgBlockID, ok := texMap[mtlMat.MapKd]
				if !ok {
					// check type
					imgType := rexfile.Raw24
					if strings.ToLower(filepath.Ext(mtlMat.MapKd)) == ".png" {
						imgType = rexfile.Png
					} else if strings.ToLower(filepath.Ext(mtlMat.MapKd)) == ".jpg" {
						imgType = rexfile.Jpeg
					}

					imgReader, err := os.Open(filepath.Join(o.inputPath, mtlMat.MapKd))
					if err != nil {
						panic(err)
					}
					buf, err := ioutil.ReadAll(imgReader)
					if err != nil {
						panic(err)
					}
					dataID++
					imgBlockID = dataID

					img := rexfile.Image{
						ID:          imgBlockID,
						Compression: uint32(imgType),
						Data:        buf,
					}
					rexFile.Images = append(rexFile.Images, img)
					texMap[mtlMat.MapKd] = imgBlockID
				}

				rexMat.KdTextureID = imgBlockID
			}
			rexFile.Materials = append(rexFile.Materials, rexMat)
		} else {
			fmt.Println("Group ", g.Name, " has no material, taking default")
			rexFile.Materials = append(rexFile.Materials, rexfile.NewMaterial(materialID))
		}
		dataID++

		// Get geometry
		mesh := rexfile.Mesh{ID: dataID, MaterialID: materialID, Name: g.Name}
		var c int
		var triangle [3]uint32
		var u, v float32
		for idx := g.IndexBegin; idx < g.IndexBegin+g.IndexCount; idx++ {
			oldVertexIndex := obj.Indices[idx]

			stride := obj.StrideSize / 4
			p := stride + obj.StrideOffsetPosition/4
			x := obj.Coord[oldVertexIndex*p]
			y := obj.Coord[oldVertexIndex*p+1]
			z := obj.Coord[oldVertexIndex*p+2]

			if obj.TextCoordFound {
				u = obj.Coord[oldVertexIndex*p+3]
				v = obj.Coord[oldVertexIndex*p+4]
			}

			rexCoordinate := rexifyObjCoordinate(x, y, z)

			if i, exists := vtxMap[oldVertexIndex]; exists {
				triangle[c] = uint32(i)
			} else {
				mesh.Coords = append(mesh.Coords, rexCoordinate)
				if obj.TextCoordFound {
					mesh.TexCoords = append(mesh.TexCoords, mgl32.Vec2{u, v})
				}
				triangle[c] = uint32(len(mesh.Coords) - 1)
				vtxMap[oldVertexIndex] = len(mesh.Coords) - 1
			}

			// Currently only triangles are assumed
			if c%2 == 0 && c > 0 {
				c = 0
				mesh.Triangles = append(mesh.Triangles, rexfile.Triangle{
					V0: triangle[0],
					V1: triangle[1],
					V2: triangle[2]})
			} else {
				c++
			}
		}

		// check if at least one triangle is contained
		if len(mesh.Triangles) > 0 {
			rexFile.Meshes = append(rexFile.Meshes, mesh)
			dataID++
		}
	}
	return rexFile, nil
}

// rexify converts the OBJ coordinate system into the REX coordinate system
func rexifyObjCoordinate(x, y, z float32) mgl32.Vec3 {
	return mgl32.Vec3{x, z, -y}
}
