package translate

import (
	"fmt"
	"io"

	"github.com/breiting/gwob"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
)

const (
	objDefaultName = "objfile"
)

// ObjToRex container
type ObjToRex struct {
	objReader io.Reader
	mtlReader io.Reader
	verbose   bool
}

// NewObjToRexTranslator returns a new translator for converting OBJ files to REX files
func NewObjToRexTranslator(objReader, mtlReader io.Reader, verbose bool) *ObjToRex {

	return &ObjToRex{
		objReader: objReader,
		mtlReader: mtlReader,
		verbose:   verbose,
	}
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

	mtl, err := gwob.ReadMaterialLibFromReader(o.mtlReader, &options)
	if err != nil {
		return rexFile, err
	}

	vtxMap := make(map[int]int)
	dataID := uint64(1)

	for _, g := range obj.Groups {

		if mtlMat, exists := mtl.Lib[g.Usemtl]; exists {
			rexMat := rexfile.NewMaterial(dataID)
			rexMat.KdRgb = mgl32.Vec3{mtlMat.Kd[0], mtlMat.Kd[1], mtlMat.Kd[2]}
			rexFile.Materials = append(rexFile.Materials, rexMat)
		} else {
			fmt.Println("Group ", g.Name, " has no material, taking default")
			rexFile.Materials = append(rexFile.Materials, rexfile.NewMaterial(dataID))
		}
		dataID++

		// Get geometry
		mesh := rexfile.Mesh{ID: dataID, MaterialID: dataID - 1, Name: g.Name}
		var c int
		var triangle [3]uint32
		for idx := g.IndexBegin; idx < g.IndexBegin+g.IndexCount; idx++ {
			oldVertexIndex := obj.Indices[idx]

			x := obj.Coord[oldVertexIndex*obj.StrideSize/4]
			y := obj.Coord[oldVertexIndex*obj.StrideSize/4+1]
			z := obj.Coord[oldVertexIndex*obj.StrideSize/4+2]

			rexCoordinate := rexifyObjCoordinate(x, y, z)

			if i, exists := vtxMap[oldVertexIndex]; exists {
				triangle[c] = uint32(i)
			} else {
				mesh.Coords = append(mesh.Coords, rexCoordinate)
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
