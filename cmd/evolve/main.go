package main

import (
	"flag"
	"io/ioutil"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/xyproto/mcbanner"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {

	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	targetImageFilename := "../mcweb/public/img/c1.png"
	pngbytes, err := ioutil.ReadFile(targetImageFilename)
	if err != nil {
		log.Fatalf("Could not read: %s\n", targetImageFilename)
	}
	mcbanner.FindBest(mcbanner.Likeness, pngbytes, "/tmp/best.png")

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
