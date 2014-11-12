package mcbanner

import (
	"math/rand"
	"time"
)

func seed() {
	rand.Seed(time.Now().UnixNano())
}

func randomPattern() int {
	// patternLowerThird is the first pattern
	// patternFull is the one after the last one we wish to return
	return patternLowerThird + rand.Intn(patternFull-patternLowerThird)
}

func randomColor() int {
	// colorWhite is the first constant, colorBlack is the last one
	return colorWhite + rand.Intn((colorBlack-colorWhite)+1)
}
