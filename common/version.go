package common

import "fmt"

// Version is libary version info struct
type Version struct {
	// Major and minor version.
	Number float32

	// Increment this for bug releases
	PatchLevel int
}

// CurrentVersion is current library version
var CurrentVersion = Version{
	Number:     0.0,
	PatchLevel: 1,
}

func (v Version) String() string {

	if v.PatchLevel > 0 {
		return fmt.Sprintf("%.1f.%d", v.Number, v.PatchLevel)
	}

	return fmt.Sprintf("%.1f", v.Number)
}
