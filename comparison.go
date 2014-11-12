package main

import "github.com/gographics/imagick/imagick"

// Render SVG to PNG
func Render(svgxml string) []byte {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	err := mw.SetVectorGraphics(svgxml)
	err := mw.SetImageFormat("png")
	return mw.GetImageBlob()
}
