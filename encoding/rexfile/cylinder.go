package rexfile

import (
	"github.com/go-gl/mathgl/mgl32"
)

// NewCylinder returns a new cylinder with radius and height (meters)
func NewCylinder(id, matID uint64, radius float32, height float32) (Mesh, Material) {
	mesh := Mesh{
		ID:         id,
		Name:       "Cylinder",
		Coords:     getCylinderCoords(radius, height),
		Triangles:  getCylinderTriangles(),
		MaterialID: matID,
	}

	mat := NewMaterial(matID)
	mat.KdRgb = mgl32.Vec3{0.9, 0.7, 0.1}

	return mesh, mat
}

//values extracted from generated cylinder
func getCylinderCoords(radius float32, height float32) []mgl32.Vec3 {
	return []mgl32.Vec3{
		{radius * (+0.000000e+000), height * 0.000000, radius * (+1.000000e+000)},
		{radius * (+3.826834e-001), height * 0.000000, radius * (+9.238795e-001)},
		{radius * (+7.071068e-001), height * 0.000000, radius * (+7.071068e-001)},
		{radius * (+9.238795e-001), height * 0.000000, radius * (+3.826834e-001)},
		{radius * (+1.000000e+000), height * 0.000000, radius * (+6.123234e-017)},
		{radius * (+9.238795e-001), height * 0.000000, radius * (-3.826834e-001)},
		{radius * (+7.071068e-001), height * 0.000000, radius * (-7.071068e-001)},
		{radius * (+3.826834e-001), height * 0.000000, radius * (-9.238795e-001)},
		{radius * (+1.224647e-016), height * 0.000000, radius * (-1.000000e+000)},
		{radius * (-3.826834e-001), height * 0.000000, radius * (-9.238795e-001)},
		{radius * (-7.071068e-001), height * 0.000000, radius * (-7.071068e-001)},
		{radius * (-9.238795e-001), height * 0.000000, radius * (-3.826834e-001)},
		{radius * (-1.000000e+000), height * 0.000000, radius * (-1.836970e-016)},
		{radius * (-9.238795e-001), height * 0.000000, radius * (+3.826834e-001)},
		{radius * (-7.071068e-001), height * 0.000000, radius * (+7.071068e-001)},
		{radius * (-3.826834e-001), height * 0.000000, radius * (+9.238795e-001)},
		{radius * (+0.000000e+000), height * 1.000000, radius * (+1.000000e+000)},
		{radius * (+3.826834e-001), height * 1.000000, radius * (+9.238795e-001)},
		{radius * (+7.071068e-001), height * 1.000000, radius * (+7.071068e-001)},
		{radius * (+9.238795e-001), height * 1.000000, radius * (+3.826834e-001)},
		{radius * (+1.000000e+000), height * 1.000000, radius * (+6.123234e-017)},
		{radius * (+9.238795e-001), height * 1.000000, radius * (-3.826834e-001)},
		{radius * (+7.071068e-001), height * 1.000000, radius * (-7.071068e-001)},
		{radius * (+3.826834e-001), height * 1.000000, radius * (-9.238795e-001)},
		{radius * (+1.224647e-016), height * 1.000000, radius * (-1.000000e+000)},
		{radius * (-3.826834e-001), height * 1.000000, radius * (-9.238795e-001)},
		{radius * (-7.071068e-001), height * 1.000000, radius * (-7.071068e-001)},
		{radius * (-9.238795e-001), height * 1.000000, radius * (-3.826834e-001)},
		{radius * (-1.000000e+000), height * 1.000000, radius * (-1.836970e-016)},
		{radius * (-9.238795e-001), height * 1.000000, radius * (+3.826834e-001)},
		{radius * (-7.071068e-001), height * 1.000000, radius * (+7.071068e-001)},
		{radius * (-3.826834e-001), height * 1.000000, radius * (+9.238795e-001)},
	}
}

func getCylinderTriangles() []Triangle {
	return []Triangle{
		{16, 0, 1},
		{1, 17, 16},
		{17, 1, 2},
		{2, 18, 17},
		{18, 2, 3},
		{3, 19, 18},
		{19, 3, 4},
		{4, 20, 19},
		{20, 4, 5},
		{5, 21, 20},
		{21, 5, 6},
		{6, 22, 21},
		{22, 6, 7},
		{7, 23, 22},
		{23, 7, 8},
		{8, 24, 23},
		{24, 8, 9},
		{9, 25, 24},
		{25, 9, 10},
		{10, 26, 25},
		{26, 10, 11},
		{11, 27, 26},
		{27, 11, 12},
		{12, 28, 27},
		{28, 12, 13},
		{13, 29, 28},
		{29, 13, 14},
		{14, 30, 29},
		{30, 14, 15},
		{15, 31, 30},
		{31, 15, 0},
		{0, 16, 31},
		{2, 1, 0},
		{16, 17, 18},
		{4, 3, 2},
		{18, 19, 20},
		{6, 5, 4},
		{20, 21, 22},
		{8, 7, 6},
		{22, 23, 24},
		{10, 9, 8},
		{24, 25, 26},
		{12, 11, 10},
		{26, 27, 28},
		{14, 13, 12},
		{28, 29, 30},
		{0, 15, 14},
		{30, 31, 16},
		{4, 2, 0},
		{16, 18, 20},
		{8, 6, 4},
		{20, 22, 24},
		{12, 10, 8},
		{24, 26, 28},
		{0, 14, 12},
		{28, 30, 16},
		{8, 4, 0},
		{16, 20, 24},
		{0, 12, 8},
		{24, 28, 16},
	}
}
