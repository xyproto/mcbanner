package main

import (
	"github.com/davecheney/profile"
	"github.com/xyproto/mcbanner"
	"io/ioutil"
	"log"
)

func main() {
	defer profile.Start(profile.CPUProfile).Stop()

	targetImageFilename := "../web/public/img/c1.png"
	pngbytes, err := ioutil.ReadFile(targetImageFilename)
	if err != nil {
		log.Fatalf("Could not read: %s\n", targetImageFilename)
	}
	mcbanner.FindBest(mcbanner.Likeness, pngbytes)
}
