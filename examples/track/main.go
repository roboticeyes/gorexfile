package main

import (
	"bytes"
	"os"
	"time"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
)

func main() {
	fileName := "testTrack.rex"

	track := rexfile.Track{ID: 0}
	track.Timestamp = time.Now().Unix()

	elem := rexfile.TrackElement{Point: mgl32.Vec3{0.0, 0.0, 0.0}, NormalVec: mgl32.Vec3{1.0, 1.0, 1.0}, Confidence: 1.0}
	track.Points = append(track.Points, elem)
	elem = rexfile.TrackElement{Point: mgl32.Vec3{1.0, 2.0, 4.0}, NormalVec: mgl32.Vec3{1.1, 1.0, 1.0}, Confidence: 1.0}
	track.Points = append(track.Points, elem)
	elem = rexfile.TrackElement{Point: mgl32.Vec3{0.7, 0.2, 2.0}, NormalVec: mgl32.Vec3{1.2, 1.0, 1.0}, Confidence: 1.0}
	track.Points = append(track.Points, elem)
	elem = rexfile.TrackElement{Point: mgl32.Vec3{0.6, 0.5, 0.4}, NormalVec: mgl32.Vec3{1.3, 1.0, 1.0}, Confidence: 1.0}
	track.Points = append(track.Points, elem)
	track.NrOfPoints = uint32(len(track.Points))

	rexFile := rexfile.File{}
	rexFile.Tracks = append(rexFile.Tracks, track)

	var buf bytes.Buffer
	e := rexfile.NewEncoder(&buf)
	err := e.Encode(rexFile)
	if err != nil {
		panic(err)
	}

	f, _ := os.Create(fileName)
	f.Write(buf.Bytes())
	defer f.Close()
}
