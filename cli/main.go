package main

import (
	"io/ioutil"
	"log"
	"runtime"

	"github.com/davecheney/profile"
	"github.com/xyproto/mcbanner"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	defer profile.Start(profile.CPUProfile).Stop()

	targetImageFilename := "../web/public/img/c1.png"
	pngbytes, err := ioutil.ReadFile(targetImageFilename)
	if err != nil {
		log.Fatalf("Could not read: %s\n", targetImageFilename)
	}
	mcbanner.FindBest(mcbanner.Likeness, pngbytes, "/tmp/best.png")
}
