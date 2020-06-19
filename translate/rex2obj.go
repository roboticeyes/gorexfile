package translate

import (
	"bufio"
	"fmt"
	"io"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
)

func RexToObj(header *rexfile.Header, content *rexfile.File, output io.Writer) error {

	w := bufio.NewWriter(output)

	w.WriteString("# Converted by rxi, (c) Robotic Eyes\n")
	w.WriteString(fmt.Sprintf("# REX meshes: %d\n", len(content.Meshes)))

	vtxCounter := uint32(1)
	vtxOffset := make(map[int]uint32)

	// write vertices
	for i := 0; i < len(content.Meshes); i++ {
		vtxOffset[i] = vtxCounter
		for _, v := range content.Meshes[i].Coords {
			w.WriteString(fmt.Sprintf("v %f %f %f\n", v.X(), v.Y(), v.Z()))
		}
		vtxCounter += uint32(len(content.Meshes[i].Coords))
	}

	// write faces
	for i := 0; i < len(content.Meshes); i++ {
		for _, v := range content.Meshes[i].Triangles {
			w.WriteString(fmt.Sprintf("f %d %d %d\n", vtxOffset[i]+v.V0, vtxOffset[i]+v.V1, vtxOffset[i]+v.V2))
		}
	}

	w.Flush()

	return nil
}

// unrexify converts the REX coordinate system into the project coordinate system
func unrexify(x, y, z float32) mgl32.Vec3 {
	return mgl32.Vec3{x, -z, y}
}
