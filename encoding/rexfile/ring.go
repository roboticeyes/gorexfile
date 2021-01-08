package rexfile

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

// NewRing returns a new ring with specified radius (meters)
func NewRing(id, matID uint64, radius, height float32, nrOfSegments int, color mgl32.Vec3, doubleSided bool) (Mesh, Material) {
	mesh := Mesh{
		ID:         id,
		Name:       "Ring",
		Coords:     getRingCoords(radius, height, nrOfSegments),
		Triangles:  getRingTriangles(nrOfSegments, doubleSided),
		MaterialID: matID,
	}

	mat := NewMaterial(matID)
	mat.KdRgb = color

	return mesh, mat
}

func getRingCoords(radius, height float32, nrOfSegments int) []mgl32.Vec3 {
	var coords []mgl32.Vec3

	alpha := 2 * math.Pi / float64(nrOfSegments)

	angle := 0.0
	for i := 0; i < nrOfSegments; i++ {
		angle = float64(i) * alpha
		x := radius * float32(math.Cos(angle))
		y := radius * float32(math.Sin(angle))
		coords = append(coords, mgl32.Vec3{x, 0, -y})
		coords = append(coords, mgl32.Vec3{x, height, -y})
	}
	return coords
}

func getRingTriangles(nrOfSegments int, doubleSided bool) []Triangle {
	var triangles []Triangle
	nrOfSegmentsFloat := float64(nrOfSegments * 2)

	for i := 0; i < nrOfSegments*2-1; i = i + 2 {
		triangle1 := Triangle{V0: uint32(i), V1: uint32(math.Mod(float64(i+2), nrOfSegmentsFloat)), V2: uint32(math.Mod(float64(i+3), nrOfSegmentsFloat))}
		triangles = append(triangles, triangle1)

		if doubleSided {
			triangles = append(triangles, Triangle{V0: triangle1.V0, V1: triangle1.V2, V2: triangle1.V1})
		}

		triangle2 := Triangle{V0: uint32(i), V1: uint32(math.Mod(float64(i+3), nrOfSegmentsFloat)), V2: uint32(math.Mod(float64(i+1), nrOfSegmentsFloat))}
		triangles = append(triangles, triangle2)

		if doubleSided {
			triangles = append(triangles, Triangle{V0: triangle2.V0, V1: triangle2.V2, V2: triangle2.V1})
		}
	}

	return triangles
}
