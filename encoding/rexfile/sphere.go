package rexfile

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

// NewSphere returns a new sphere (unit meters)
func NewSphere(id, matID uint64, radius float64, widthSegments, heightSegments int, color mgl32.Vec3) (Mesh, Material) {

	phiStart, phiLength, thetaStart, thetaLength := 0.0, math.Pi*2, 0.0, math.Pi

	thetaEnd := thetaStart + thetaLength
	vertexCount := (widthSegments + 1) * (heightSegments + 1)

	positions := make([]mgl32.Vec3, vertexCount)
	triangles := make([]Triangle, 0)

	index := 0
	vertices := make([][]uint32, 0)

	for y := 0; y <= heightSegments; y++ {
		verticesRow := make([]uint32, 0)
		v := float64(y) / float64(heightSegments)
		for x := 0; x <= widthSegments; x++ {
			u := float64(x) / float64(widthSegments)
			px := -radius * math.Cos(phiStart+u*phiLength) * math.Sin(thetaStart+v*thetaLength)
			py := radius * math.Cos(thetaStart+v*thetaLength)
			pz := radius * math.Sin(phiStart+u*phiLength) * math.Sin(thetaStart+v*thetaLength)

			positions[index] = mgl32.Vec3{float32(px), float32(py), float32(pz)}
			verticesRow = append(verticesRow, uint32(index))
			index++
		}
		vertices = append(vertices, verticesRow)
	}

	for y := 0; y < heightSegments; y++ {
		for x := 0; x < widthSegments; x++ {
			v1 := vertices[y][x+1]
			v2 := vertices[y][x]
			v3 := vertices[y+1][x]
			v4 := vertices[y+1][x+1]
			if y != 0 || thetaStart > 0 {
				triangles = append(triangles, Triangle{v1, v2, v4})
			}
			if y != heightSegments-1 || thetaEnd < math.Pi {
				triangles = append(triangles, Triangle{v2, v3, v4})
			}
		}
	}

	mesh := Mesh{
		ID:         id,
		Name:       "Sphere",
		Coords:     positions,
		Triangles:  triangles,
		MaterialID: matID,
	}

	mat := NewMaterial(matID)
	mat.KdRgb = color
	mat.Alpha = 0.8

	return mesh, mat
}
