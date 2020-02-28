package rexfile

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	textkBlockVersion = 1
)

// Text datastructure
type Text struct {
	ID       uint64
	Red      float32
	Green    float32
	Blue     float32
	Alpha    float32
	Position mgl32.Vec3
	FontSize float32
	Text     string
}

// GetSize returns the estimated size of the block in bytes
func (block *Text) GetSize() int {
	return rexDataBlockHeaderSize + 8*4 + 2 + len(block.Text)
}

// ReadText reads a REX text
func ReadText(r io.Reader, hdr DataBlockHeader) (*Text, error) {
	var text Text

	if err := binary.Read(r, binary.LittleEndian, &text.Red); err != nil {
		return nil, fmt.Errorf("Reading failed: %v ", err)
	}

	if err := binary.Read(r, binary.LittleEndian, &text.Green); err != nil {
		return nil, fmt.Errorf("Reading failed: %v ", err)
	}

	if err := binary.Read(r, binary.LittleEndian, &text.Blue); err != nil {
		return nil, fmt.Errorf("Reading failed: %v ", err)
	}

	if err := binary.Read(r, binary.LittleEndian, &text.Alpha); err != nil {
		return nil, fmt.Errorf("Reading failed: %v ", err)
	}

	if err := binary.Read(r, binary.LittleEndian, &text.Position); err != nil {
		return nil, fmt.Errorf("Reading position failed: %v", err)
	}

	if err := binary.Read(r, binary.LittleEndian, &text.FontSize); err != nil {
		return nil, fmt.Errorf("Reading failed: %v ", err)
	}

	var textLen uint16
	if err := binary.Read(r, binary.LittleEndian, &textLen); err != nil {
		return nil, fmt.Errorf("Reading failed: %v ", err)
	}

	textArray := make([]byte, textLen)
	if err := binary.Read(r, binary.LittleEndian, &textArray); err != nil {
		return nil, fmt.Errorf("Reading text string failed: %v", err)
	}
	text.Text = string(textArray)

	return &text, nil
}

// Write writes the track to the given writer
func (block *Text) Write(w io.Writer) error {
	// return if nothing needs to be written
	if len(block.Text) == 0 {
		return nil
	}

	err := WriteDataBlockHeader(w, DataBlockHeader{
		Type:    typeText,
		Version: trackBlockVersion,
		Size:    uint32(block.GetSize() - rexDataBlockHeaderSize),
		ID:      block.ID,
	})
	if err != nil {
		return err
	}

	var data = []interface{}{
		block.Red,
		block.Green,
		block.Blue,
		block.Alpha,
		block.Position,
		block.FontSize,
		uint16(len(block.Text)),
		[]byte(block.Text),
	}
	for _, v := range data {
		err := binary.Write(w, binary.LittleEndian, v)
		if err != nil {
			return err
		}
	}

	return nil
}
