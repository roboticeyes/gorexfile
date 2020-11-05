package translate

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
)

func RexToObj(header *rexfile.Header, content *rexfile.File, writer *bufio.Writer) error {

	writer.WriteString(fmt.Sprintf("# REX meshes: %d\n", len(content.Meshes)))

	vtxCounter := uint32(1)
	vtxOffset := make(map[int]uint32)

	// write vertices
	for i := 0; i < len(content.Meshes); i++ {
		vtxOffset[i] = vtxCounter
		for _, v := range content.Meshes[i].Coords {
			writer.WriteString(fmt.Sprintf("v %f %f %f\n", v.X(), v.Y(), v.Z()))
		}
		for _, v := range content.Meshes[i].TexCoords {
			writer.WriteString(fmt.Sprintf("vt %f %f\n", v.X(), v.Y()))
		}
		vtxCounter += uint32(len(content.Meshes[i].Coords))
	}

	// write faces
	for i := 0; i < len(content.Meshes); i++ {
		for _, v := range content.Meshes[i].Triangles {
			if len(content.Meshes[i].TexCoords) > 0 {
				writer.WriteString(fmt.Sprintf("f %d/%d %d/%d %d/%d\n",
					vtxOffset[i]+v.V0, vtxOffset[i]+v.V0,
					vtxOffset[i]+v.V1, vtxOffset[i]+v.V1,
					vtxOffset[i]+v.V2, vtxOffset[i]+v.V2))
			} else {
				writer.WriteString(fmt.Sprintf("f %d %d %d\n", vtxOffset[i]+v.V0, vtxOffset[i]+v.V1, vtxOffset[i]+v.V2))
			}
		}
	}

	return nil
}

// RexToWavefront is a high-level function which generates OBJ/MTL/texture files based
// on the REXfile content.
// dir ... directory for the output, needs to exist
// name ... the name of the files being generation (e.g. name="sphere" would output sphere.obj and sphere.mtl)
func RexToWavefront(header *rexfile.Header, content *rexfile.File, path, name string) error {

	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("Target directory does not exist")
	}
	if !info.IsDir() {
		return fmt.Errorf("Given path is not a directory")
	}

	// obj file
	obj, err := os.Create(filepath.Join(path, name+".obj"))
	if err != nil {
		return err
	}
	defer obj.Close()

	objWriter := bufio.NewWriter(obj)
	objWriter.WriteString("# Converted by gorexfile, (c) Robotic Eyes\n")
	objWriter.WriteString("mtllib " + name + ".mtl\n")
	objWriter.WriteString("usemtl default\n")

	err = RexToObj(header, content, objWriter)
	if err != nil {
		return err
	}

	err = objWriter.Flush()
	if err != nil {
		return err
	}

	// mtl file
	mtl, err := os.Create(filepath.Join(path, name+".mtl"))
	if err != nil {
		return err
	}
	defer mtl.Close()

	// TODO refactor
	mtlWriter := bufio.NewWriter(mtl)
	mtlWriter.WriteString("newmtl default\n")
	mtlWriter.WriteString("Ka 1 1 1\n")
	mtlWriter.WriteString("Kd 1 1 1\n")
	mtlWriter.WriteString("Ks 1 1 1\n")
	mtlWriter.WriteString("illum 2\n")
	mtlWriter.WriteString("Ns 1.4\n")
	mtlWriter.WriteString("map_Ka texture.jpg\n")
	mtlWriter.WriteString("map_Kd texture.jpg\n")
	mtlWriter.WriteString("map_Ks texture.jpg\n")

	// image
	for _, img := range content.Images {
		if img.Compression == rexfile.Jpeg {
			f, err := os.Create(filepath.Join(path, "texture.jpg"))
			if err != nil {
				return err
			}
			defer f.Close()
			binary.Write(f, binary.LittleEndian, img.Data)
		}
	}

	return mtlWriter.Flush()
}

// unrexify converts the REX coordinate system into the project coordinate system
func unrexify(x, y, z float32) mgl32.Vec3 {
	return mgl32.Vec3{x, -z, y}
}
