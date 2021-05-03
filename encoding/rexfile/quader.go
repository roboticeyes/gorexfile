package rexfile

import (
	"github.com/go-gl/mathgl/mgl32"
)

// NewQuader returns a new cube with size (meters)
func NewQuader(id, matID uint64, sizeX, sizeY, sizeZ float32) (Mesh, Material) {

	// default geometry is 2 meters
	scaleX := sizeX / 2.0
	scaleY := sizeY / 2.0
	scaleZ := sizeZ / 2.0

	mesh := Mesh{
		ID:         id,
		Name:       "Quader",
		Coords:     getCoords(scaleX, scaleY, scaleZ),
		Triangles:  getTriangles(),
		MaterialID: matID,
	}

	mat := NewMaterial(matID)
	mat.KdRgb = mgl32.Vec3{0.9, 0.7, 0.1}

	return mesh, mat
}

func getCoords(scaleX, scaleY, scaleZ float32) []mgl32.Vec3 {

	return []mgl32.Vec3{
		mgl32.Vec3{scaleX * -1.0, scaleY * -1.0, scaleZ * 1.0},
		mgl32.Vec3{scaleX * 1.0, scaleY * -1.0, scaleZ * 1.0},
		mgl32.Vec3{scaleX * 1.0, scaleY * 1.0, scaleZ * 1.0},
		mgl32.Vec3{scaleX * -1.0, scaleY * 1.0, scaleZ * 1.0},
		mgl32.Vec3{scaleX * -1.0, scaleY * 1.0, scaleZ * 1.0},
		mgl32.Vec3{scaleX* 1.0, scaleY * 1.0, scaleZ * 1.0},
		mgl32.Vec3{scaleX * 1.0, scaleY * 1.0, scaleZ * -1.0},
		mgl32.Vec3{scaleX * -1.0, scaleY * 1.0, scaleZ * -1.0},
		mgl32.Vec3{scaleX * 1.0, scaleY * -1.0, scaleZ * -1.0},
		mgl32.Vec3{scaleX * -1.0, scaleY * -1.0, scaleZ * -1.0},
		mgl32.Vec3{scaleX * -1.0, scaleY * 1.0, scaleZ * -1.0},
		mgl32.Vec3{scaleX* 1.0, scaleY * 1.0, scalez * -1.0},
		mgl32.Vec3{scaleX * -1.0, scaleY * -1.0, scaleZ * -1.0},
		mgl32.Vec3{scaleX * 1.0, scaleY * -1.0, scaleZ * -1.0},
		mgl32.Vec3{scaleX * 1.0, scaleY * -1.0, scaleZ * 1.0},
		mgl32.Vec3{scaleX * -1.0, scaleY * -1.0, scaleZ * 1.0},
		mgl32.Vec3{scaleX * -1.0, scaleY * -1.0, scaleZ * -1.0},
		mgl32.Vec3{scaleX * -1.0, scaleY * -1.0, scaleZ * 1.0},
		mgl32.Vec3{scaleX * -1.0, scaleY * 1.0, scaleZ * 1.0},
		mgl32.Vec3{scaleX * -1.0, scaleY * 1.0, scaleZ * -1.0},
		mgl32.Vec3{scaleX * 1.0, scaleY * -1.0, scaleZ * 1.0},
		mgl32.Vec3{scaleX * 1.0, scaleY * -1.0, scaleZ * -1.0},
		mgl32.Vec3{scaleX * 1.0, scaleY * 1.0, scaleZ * -1.0},
		mgl32.Vec3{scaleX * 1.0, scaleY * 1.0, scaleZ * 1.0},
	}
}

func getTriangles() []Triangle {

	return []Triangle{
		Triangle{V0: 0, V1: 1, V2: 2},
		Triangle{V0: 2, V1: 3, V2: 0},
		Triangle{V0: 4, V1: 5, V2: 6},
		Triangle{V0: 6, V1: 7, V2: 4},
		Triangle{V0: 8, V1: 9, V2: 10},
		Triangle{V0: 10, V1: 11, V2: 8},
		Triangle{V0: 12, V1: 13, V2: 14},
		Triangle{V0: 14, V1: 15, V2: 12},
		Triangle{V0: 16, V1: 17, V2: 18},
		Triangle{V0: 18, V1: 19, V2: 16},
		Triangle{V0: 20, V1: 21, V2: 22},
		Triangle{V0: 22, V1: 23, V2: 20},
	}
}
