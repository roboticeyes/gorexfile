package rexfile

import (
	"github.com/go-gl/mathgl/mgl32"
)

// NewQuader returns a new quader with given size
func NewQuader(id, matID uint64, sizeX, sizeY, sizeZ float32) (Mesh, Material) {

	mesh := Mesh{
		ID:         id,
		Name:       "Quader",
		Coords:     getQuaderCoords(sizeX, sizeY, sizeZ),
		Triangles:  getQuaderTriangles(),
		MaterialID: matID,
	}

	mat := NewMaterial(matID)
	mat.KdRgb = mgl32.Vec3{0.9, 0.7, 0.1}

	return mesh, mat
}

func getQuaderCoords(sizeX, sizeY, sizeZ float32) []mgl32.Vec3 {

	return []mgl32.Vec3{
		{sizeX * 0.0, sizeY * 0.0, sizeZ * 1.0},
		{sizeX * 1.0, sizeY * 0.0, sizeZ * 1.0},
		{sizeX * 1.0, sizeY * 1.0, sizeZ * 1.0},
		{sizeX * 0.0, sizeY * 1.0, sizeZ * 1.0},
		{sizeX * 0.0, sizeY * 1.0, sizeZ * 1.0},
		{sizeX * 1.0, sizeY * 1.0, sizeZ * 1.0},
		{sizeX * 1.0, sizeY * 1.0, sizeZ * 0.0},
		{sizeX * 0.0, sizeY * 1.0, sizeZ * 0.0},
		{sizeX * 1.0, sizeY * 0.0, sizeZ * 0.0},
		{sizeX * 0.0, sizeY * 0.0, sizeZ * 0.0},
		{sizeX * 0.0, sizeY * 1.0, sizeZ * 0.0},
		{sizeX * 1.0, sizeY * 1.0, sizeZ * 0.0},
		{sizeX * 0.0, sizeY * 0.0, sizeZ * 0.0},
		{sizeX * 1.0, sizeY * 0.0, sizeZ * 0.0},
		{sizeX * 1.0, sizeY * 0.0, sizeZ * 1.0},
		{sizeX * 0.0, sizeY * 0.0, sizeZ * 1.0},
		{sizeX * 0.0, sizeY * 0.0, sizeZ * 0.0},
		{sizeX * 0.0, sizeY * 0.0, sizeZ * 1.0},
		{sizeX * 0.0, sizeY * 1.0, sizeZ * 1.0},
		{sizeX * 0.0, sizeY * 1.0, sizeZ * 0.0},
		{sizeX * 1.0, sizeY * 0.0, sizeZ * 1.0},
		{sizeX * 1.0, sizeY * 0.0, sizeZ * 0.0},
		{sizeX * 1.0, sizeY * 1.0, sizeZ * 0.0},
		{sizeX * 1.0, sizeY * 1.0, sizeZ * 1.0},
	}
}

func getQuaderTriangles() []Triangle {

	return []Triangle{
		{V0: 0, V1: 1, V2: 2},
		{V0: 2, V1: 3, V2: 0},
		{V0: 4, V1: 5, V2: 6},
		{V0: 6, V1: 7, V2: 4},
		{V0: 8, V1: 9, V2: 10},
		{V0: 10, V1: 11, V2: 8},
		{V0: 12, V1: 13, V2: 14},
		{V0: 14, V1: 15, V2: 12},
		{V0: 16, V1: 17, V2: 18},
		{V0: 18, V1: 19, V2: 16},
		{V0: 20, V1: 21, V2: 22},
		{V0: 22, V1: 23, V2: 20},
	}
}
