package main

import (
	"net/http"
	"runtime"

	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
)

const (
	version_string = "MC Banner Generator 2.0.1"
)

// Set up the paths and handlers then start serving.
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Create a Negroni instance and a ServeMux instance
	n := negroni.Classic()
	r := render.New(render.Options{})
	mux := http.NewServeMux()

	// Publish the main page
	mainPage(mux, "/", r)

	// Publish the pattern gallery page
	patternGallery(mux, "/patterns")

	// Publish the random banner page
	randomBanner(mux, "/random", r)

	// Publish the svg->pixel vs png->pixel comparison page
	comparison(mux, "/comparison", r)

	// Handler goes last
	n.UseHandler(mux)

	// Listen for requests
	n.Run(":3020")
}
