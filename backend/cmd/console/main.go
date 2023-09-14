package main

import (
	"fmt"
	"pixel-art-tools/backend/pkg/imageinfo"
	"pixel-art-tools/backend/pkg/imageloader"
)

func main() {
	imagePath := "/home/neidz/Projects/pixel-art-tools/backend/assets/crewmate.png"

	image, err := imageloader.LoadImage(imagePath)
	if err != nil {
		fmt.Println("Error changing working directory:", err)
		return
	}

	metadata := imageinfo.GenerateMetadata(image)

	fmt.Println("Dimensions: ", metadata.ImageDimensions.X, metadata.ImageDimensions.Y)
}
