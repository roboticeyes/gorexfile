package rexfile

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
	"sort"
)

// NewCylinder returns a new cylinder with radius and height (meters)
func NewCylinder(id, matID uint64, radius float32, height float32) (Mesh, Material) {
	const numberOfSegments = 16

	mesh := Mesh{
		ID:         id,
		Name:       "Cylinder",
		Coords:     getCylinderCoords(radius, height, numberOfSegments),
		Triangles:  getCylinderTriangles(numberOfSegments),
		MaterialID: matID,
	}

	mat := NewMaterial(matID)
	mat.KdRgb = mgl32.Vec3{0.9, 0.7, 0.1}

	return mesh, mat
}

func getCylinderCoords(radius float32, height float32, numberOfSegments int) []mgl32.Vec3 {
	coords := make([]mgl32.Vec3, numberOfSegments*2)
	for i := 0; i < numberOfSegments; i++ {
		angle := (2 * math.Pi) / float64(numberOfSegments)

		coords[i] = mgl32.Vec3{
			radius * float32(math.Sin(float64(i)*angle)),
			0,
			radius * float32(math.Cos(float64(i)*angle)),
		}

		coords[i+numberOfSegments] = mgl32.Vec3{
			coords[i].X(),
			height,
			coords[i].Z(),
		}
	}

	return coords
}

func getCylinderTriangles(numberOfSegments int) []Triangle {
	baseShapeVertexRange := numberOfSegments
	wallTriangles := make([]Triangle, baseShapeVertexRange*2)

	ti := 0
	for i := 0; i < baseShapeVertexRange; i++ {
		a := i + baseShapeVertexRange
		b := i
		c := i + 1
		d := i + 1 + baseShapeVertexRange

		if d > numberOfSegments*2-1 {
			d = baseShapeVertexRange
			c = 0
		}

		wallTriangles[ti].V0 = uint32(a)
		wallTriangles[ti].V1 = uint32(b)
		wallTriangles[ti].V2 = uint32(c)
		ti++
		wallTriangles[ti].V0 = uint32(c)
		wallTriangles[ti].V1 = uint32(d)
		wallTriangles[ti].V2 = uint32(a)
		ti++
	}

	capTriangles := make([]Triangle, baseShapeVertexRange*2-4)

	currentValidVertices := make([]int, baseShapeVertexRange)
	for i := range currentValidVertices {
		currentValidVertices[i] = i
	}

	ti = 0
	for len(currentValidVertices) > 2 {
		sort.Ints(currentValidVertices)
		var nextValidVertices []int
		for i := 0; i < len(currentValidVertices)-1; i += 2 {

			aIndex := 0
			if i+2 < len(currentValidVertices) {
				aIndex = i + 2
			}

			a := currentValidVertices[aIndex]
			b := currentValidVertices[i+1]
			c := currentValidVertices[i]

			//add bottom tri
			capTriangles[ti].V0 = uint32(a)
			capTriangles[ti].V1 = uint32(b)
			capTriangles[ti].V2 = uint32(c)
			ti++

			//add top tri
			capTriangles[ti].V0 = uint32(c + baseShapeVertexRange)
			capTriangles[ti].V1 = uint32(b + baseShapeVertexRange)
			capTriangles[ti].V2 = uint32(a + baseShapeVertexRange)
			ti++

			//keep list of unique remaining vertices
			if !contains(nextValidVertices, a) {
				nextValidVertices = append(nextValidVertices, a)
			}
			if !contains(nextValidVertices, c) {
				nextValidVertices = append(nextValidVertices, c)
			}
		}

		currentValidVertices = nextValidVertices
	}

	return append(wallTriangles, capTriangles...)
}

func contains(arr []int, val int) bool {
	for _, av := range arr {
		if av == val {
			return true
		}
	}
	return false
}
