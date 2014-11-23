package mcbanner

import (
	"log"
	"math/rand"
	"time"
)

func Seed() {
	rand.Seed(time.Now().UnixNano())
}

func randomPattern() int {
	// PatternLowerThird is the first Pattern
	// PatternFull is the one after the last one we wish to return
	return PatternLowerThird + rand.Intn(PatternFull-PatternLowerThird)
}

func randomColor() int {
	// colorWhite is the first constant, colorBlack is the last one
	return ColorWhite + rand.Intn((ColorBlack-ColorWhite)+1)
}

func NewRandomPattern() Pattern {
	// error is only returned if the values are invalid
	// randomPattern and randomColor only returns valid values, so the error can be ignored
	pat, err := NewPattern(randomPattern(), randomColor())
	if err != nil {
		log.Println("ERROR WHEN GENERATING RANDOM PATTERN")
	}
	return *pat
}
