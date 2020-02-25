package translate

import (
	"github.com/roboticeyes/gorexfile/encoding/rexfile"
)

// Translator interface is a generic interface which converts anything to a rex file
type Translator interface {
	Translate() (rexfile.File, error)
}
