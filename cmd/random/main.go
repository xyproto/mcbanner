package main

import (
	"github.com/xyproto/mcbanner"
	"os"
)

func main() {
	mcbanner.Seed()
	ban, _ := mcbanner.NewRandomBanner()
	os.Stdout.Write(ban.PNG())
}
