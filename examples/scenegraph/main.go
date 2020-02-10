package main

import (
	"encoding/json"
	"os"

	// "bytes"
	// "math"
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
)

func main() {

	root := rexfile.NewSceneNode(1, 0, "root")
	s10 := rexfile.NewSceneNode(2, 10, "mesh 10")
	s11 := rexfile.NewSceneNode(3, 11, "mesh 11")
	s12 := rexfile.NewSceneNode(4, 12, "mesh 12")
	t1 := rexfile.NewSceneNode(5, 0, "position 1")
	t2 := rexfile.NewSceneNode(6, 0, "position 2")
	t2.Translation[1] -= 5

	t1.Children = append(t1.Children, []rexfile.SceneNode{s10, s11, s12}...)
	t2.Children = append(t2.Children, []rexfile.SceneNode{s10, s11, s12}...)
	root.Children = append(root.Children, []rexfile.SceneNode{t1, t2}...)

	sg := rexfile.SceneGraph{
		GUID: "74de49f3-88c5-40bd-8121-dd866059f7b3",
		Name: "Scenegraph Sample",
		Root: root,
	}

	e := json.NewEncoder(os.Stdout)
	e.Encode(sg)

	// cube, mat := rexfile.NewCube(1, 2, 1)
	//
	// rexFile := rexfile.File{}
	// rexFile.Meshes = append(rexFile.Meshes, cube)
	// rexFile.Materials = append(rexFile.Materials, mat)
	//
	// id := 3
	// for x := -10; x < 10; x++ {
	// 	for y := -10; y < 10; y++ {
	// 		for z := -10; z < 10; z++ {
	// 			rexFile.rexfile.SceneNodes = append(rexFile.SceneNodes, getSceneNode(uint64(id), 1, float32(x), float32(y), float32(z), 0.5))
	// 			id++
	// 		}
	// 	}
	// }
	// // rexFile.rexfile.SceneNodes = append(rexFile.SceneNodes, getSceneNode(4, 1, 4.5, 0, 0, 0.5))
	//
	// var buf bytes.Buffer
	// e := rexfile.NewEncoder(&buf)
	// err := e.Encode(rexFile)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// f, _ := os.Create(fileName)
	// f.Write(buf.Bytes())
	// defer f.Close()

}
