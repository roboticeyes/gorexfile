package commands

import (
	"fmt"

	"github.com/roboticeyes/gorexfile/encoding/rexfile"
	"github.com/urfave/cli/v2"
)

// InfoCommand displays file information
var InfoCommand = &cli.Command{
	Name:   "info",
	Usage:  "Show file information",
	Action: InfoAction,
}

// InfoAction is the default action, therefore it is set to public
func InfoAction(ctx *cli.Context) error {

	rexHeader, rexContent, err := OpenRexFileFromContext(ctx)
	if err != nil {
		return err
	}

	fmt.Println(rexHeader)

	// Meshes
	if len(rexContent.Meshes) > 0 {
		fmt.Printf("Meshes (%d)\n", len(rexContent.Meshes))
		fmt.Printf("%10s %8s %8s %12s %s\n", "ID", "#Vtx", "#Tri", "Material", "Name")
		for _, mesh := range rexContent.Meshes {
			fmt.Printf("%10d %8d %8d %12d %s\n", mesh.ID, len(mesh.Coords), len(mesh.Triangles), mesh.MaterialID, mesh.Name)
		}
	}
	// Materials
	if len(rexContent.Materials) > 0 {
		fmt.Printf("Materials (%d)\n", len(rexContent.Materials))
		fmt.Printf("%10s %17s %16s %16s %5s %5s %s\n", "ID", "Ambient", "Diffuse", "Specular", "Ns", "Opacity", "TextureID (ADS)")
		for _, mat := range rexContent.Materials {
			texA, texD, texS := int(mat.KaTextureID), int(mat.KdTextureID), int(mat.KsTextureID)
			if mat.KaTextureID == rexfile.NotSpecified {
				texA = -1
			}
			if mat.KdTextureID == rexfile.NotSpecified {
				texD = -1
			}
			if mat.KsTextureID == rexfile.NotSpecified {
				texS = -1
			}
			fmt.Printf("%10d, [%.2f,%.2f,%.2f] [%.2f,%.2f,%.2f] [%.2f,%.2f,%.2f] %5.1f %7.2f [%d,%d,%d]\n", mat.ID,
				mat.KaRgb.X(), mat.KaRgb.Y(), mat.KaRgb.Z(),
				mat.KdRgb.X(), mat.KdRgb.Y(), mat.KdRgb.Z(),
				mat.KsRgb.X(), mat.KsRgb.Y(), mat.KsRgb.Z(),
				mat.Ns, mat.Alpha,
				texA, texD, texS)
		}
	}
	// Images
	if len(rexContent.Images) > 0 {
		fmt.Printf("Images (%d)\n", len(rexContent.Images))
		fmt.Printf("%10s %8s %12s\n", "ID", "Compression", "Bytes")
		for _, img := range rexContent.Images {
			compression := "raw"
			if img.Compression == 1 {
				compression = "jpg"
			} else if img.Compression == 2 {
				compression = "png"
			}
			fmt.Printf("%10d %11s %12d\n", img.ID, compression, len(img.Data))
		}
	}

	// PointList
	if len(rexContent.PointLists) > 0 {
		fmt.Printf("PointLists (%d)\n", len(rexContent.PointLists))
		fmt.Printf("%10s %8s %8s\n", "ID", "#Vtx", "#Col")
		for _, pl := range rexContent.PointLists {
			fmt.Printf("%10d %8d %8d\n", pl.ID, len(pl.Points), len(pl.Colors))
		}
	}

	// LineSet
	if len(rexContent.LineSets) > 0 {
		fmt.Printf("LineSets (%d)\n", len(rexContent.LineSets))
		fmt.Printf("%10s %8s %8s\n", "ID", "#Vtx", "#Col")
		for _, pl := range rexContent.LineSets {
			fmt.Printf("%10d %8d %8d\n", pl.ID, len(pl.Points), len(pl.Colors))
		}
	}

	// SceneNodes
	if len(rexContent.SceneNodes) > 0 {
		fmt.Printf("SceneNodes (%d)\n", len(rexContent.SceneNodes))
		fmt.Printf("%10s %14s %21s %28s %21s %s\n", "ID", "GeometryID", "Translation", "Rotation", "Scale", "Name")
		for _, pl := range rexContent.SceneNodes {

			fmt.Printf("%10d %14d [%+.2f, %+.2f, %+.2f] [%+.2f, %+.2f, %+.2f, %+.2f] [%+.2f, %+.2f, %+.2f] %s\n", pl.ID, pl.GeometryID,
				pl.Translation.X(), pl.Translation.Y(), pl.Translation.Z(),
				pl.Rotation.X(), pl.Rotation.Y(), pl.Rotation.Z(), pl.Rotation.W(),
				pl.Scale.X(), pl.Scale.Y(), pl.Scale.Z(), pl.Name)
		}
	}

	if rexContent.UnknownBlocks > 0 {
		fmt.Printf("Unknown blocks (%d)\n", rexContent.UnknownBlocks)
	}

	return nil
}
