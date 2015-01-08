package mcbanner

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/xyproto/permissions2"
)

// Use rsvg to render svg (convert bytes from svg to png)
func Convert(imagebytes []byte, fromformat, toformat string) []byte {
	randomString := permissions.RandomHumanFriendlyString(8) // Collisions are rare and not critical, for this function
	path := "/tmp/"
	inputFilename := path + randomString + ".svg"
	outputFilename := path + randomString + ".png"

	// write the .svg file
	if err := ioutil.WriteFile(inputFilename, imagebytes, 0600); err != nil {
		panic(err)
	}
	// convert the .svg file to the output format (perhaps png)
	if err := exec.Command("rsvg-convert", inputFilename, "-b", "white", "-f", toformat, "-o", outputFilename).Run(); err != nil {
		panic(err)
	}
	// read the converted image
	b, err := ioutil.ReadFile(outputFilename)
	if err != nil {
		panic(err)
	}
	// remove both temporary files
	if err := os.Remove(inputFilename); err != nil {
		panic(err)
	}
	if err = os.Remove(outputFilename); err != nil {
		panic(err)
	}
	return b
}
