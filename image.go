package mcbanner

import (
	"bytes"
	"image/color"
	"image/png"
	"log"
)

// The difference between two color values, as float64 (0..1)
func Colordiff(a, b uint32) float64 {
	af := float64(a) / (256.0 * 256.0)
	bf := float64(b) / (256.0 * 256.0)
	res := af - bf
	if res < 0 {
		return -res
	}
	return res
}

func Valuediff(a, b float64) float64 {
	res := a - b
	if res < 0 {
		return -res
	}
	return res
}

// Lightness
func Value(r, g, b uint32) float64 {
	rf := float64(r) / (256.0 * 256.0)
	gf := float64(g) / (256.0 * 256.0)
	bf := float64(b) / (256.0 * 256.0)
	return (rf + gf + bf) / 3.0
}

// Find the distance between two colors, 0..1
func Distance(c1, c2 color.Color) float64 {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()
	//v1 := Value(r1, g1, b1)
	//v2 := Value(r2, g2, b2)
	//return (Colordiff(r1, r2) + Colordiff(g1, g2) + Colordiff(b1, b2) + Valuediff(v1, v2)) * 0.25
	return (Colordiff(r1, r2) + Colordiff(g1, g2) + Colordiff(b1, b2)) * 0.33
}

// Find how visually similar two images are, from 0..1
// Does not take the human vision into account, only rgb values
func Likeness(png1, png2 []byte) float64 {
	buf1 := bytes.NewBuffer(png1)
	buf2 := bytes.NewBuffer(png2)

	i1, err := png.Decode(buf1)
	if err != nil {
		log.Fatalln(err.Error())
	}
	i2, err := png.Decode(buf2)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if i1.Bounds() != i2.Bounds() {
		log.Fatalf("Can only compare images of same size! Got %v and %v.\n", i1.Bounds(), i2.Bounds())
	}
	counter := 0
	sum := 0.0
	for y := i1.Bounds().Min.Y; y < i1.Bounds().Max.Y; y++ {
		for x := i1.Bounds().Min.X; x < i1.Bounds().Max.X; x++ {
			//log.Printf("d %.3f\n", Distance(i1.At(x, y), i2.At(x, y)))
			sum += Distance(i1.At(x, y), i2.At(x, y))
			counter++
		}
	}
	//log.Printf("Difference %.3f\n", res)
	return sum / float64(counter)
}

// Compare a banner with a png file, return the likeness as float64
func Compare(b *Banner, pngbytes []byte) float64 {
	b_pngbytes := b.PNG()
	return Likeness(b_pngbytes, pngbytes)
}
