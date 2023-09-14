package patternmatcher

import (
	"errors"
	"image/color"
)

func AreColorsSimilar(color1 color.Color, color2 color.Color, args ...int) (bool, error) {
	threshold := 1

	if len(args) > 0 {
		threshold = args[0]
	}

	if threshold < 0 {
		return false, errors.New("threshold cannot be negative")
	}

	rgbaColor1 := color.RGBAModel.Convert(color1).(color.RGBA)
	rgbaColor2 := color.RGBAModel.Convert(color2).(color.RGBA)

	deltaR := int(rgbaColor1.R) - int(rgbaColor2.R)
	deltaG := int(rgbaColor1.G) - int(rgbaColor2.G)
	deltaB := int(rgbaColor1.B) - int(rgbaColor2.B)

	return absInt(deltaR) <= threshold &&
		absInt(deltaG) <= threshold &&
		absInt(deltaB) <= threshold, nil
}

func absInt(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}
