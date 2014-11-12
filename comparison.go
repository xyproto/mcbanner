package mcbanner

import "github.com/gographics/imagick/imagick"

// Render SVG to PNG (SLOW AND DOES NOT WORK) :(
func Render(svgxml string) []byte {
	// Initialize / destroy
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	dw := imagick.NewDrawingWand()
	defer dw.Destroy()

	// Render the SVG with the DrawingWand onto the MagickWand
	dw.SetVectorGraphics(svgxml)
	mw.DrawImage(dw)

	// Return the bytes as PNG
	mw.SetImageFormat("png")
	return mw.GetImageBlob()
}
