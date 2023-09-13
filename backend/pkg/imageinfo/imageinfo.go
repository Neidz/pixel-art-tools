package imageinfo

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
)

type ImageDimensions struct {
	X int
	Y int
}

type Metadata struct {
	ImageDimensions ImageDimensions
}

func GenerateMetadata(img image.Image) Metadata {
	metadata := Metadata{ImageDimensions{X: img.Bounds().Dx(), Y: img.Bounds().Dy()}}

	return metadata
}
