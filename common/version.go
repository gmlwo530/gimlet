package common

import "fmt"

type Version struct {
	// Major and minor version.
	Number float32

	// Increment this for bug releases
	PatchLevel int
}

func (v Version) String() string {

	if (v.PatchLevel > 0){
		return fmt.Sprintf("%.1f.%d", v.Number, v.PatchLevel)
	}
	
	return fmt.Sprintf("%.1f", v.Number)
}