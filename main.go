package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/xyproto/onthefly"
)

const (
	bannerW = 20
	bannerH = 40

	zoomFactor = 8 // 8x zoom

	fullW  = bannerW
	fullH  = bannerH
	halfW  = bannerW / 2
	halfH  = bannerH / 2
	thirdW = bannerW / 3
	thirdH = bannerH / 3

	maxX = fullW - 1
	maxY = fullH - 1

	maxPatterns = 99 // 6

	patternUpperHalf = iota
	patternLowerHalf
	patternUpperThird
	patternLowerThird
	patternUpperTriangle
	patternLowerTriangle
	patternCircle
	patternHorizontalLine
	patternVerticalLine
)

type Pattern struct {
	color string
	ptype int
}

type Banner struct {
	patterns []*Pattern
	curpat   int
}

func (b *Banner) Draw(svg *onthefly.Tag) {
	// generate the picture, then draw
	for _, pattern := range b.patterns {
		switch pattern.ptype {
		case patternUpperHalf:
			svg.Box(0, 0, fullW, halfH, pattern.color)
		case patternLowerHalf:
			svg.Box(0, halfH, fullW, halfH, pattern.color)
		case patternUpperThird:
			svg.Box(0, 0, fullW, thirdH, pattern.color)
		case patternLowerThird:
			svg.Box(0, fullH-thirdH, fullW, fullH-thirdH, pattern.color)
		case patternUpperTriangle:
			svg.Triangle(0, 0, fullW, 0, halfW, halfH, pattern.color)
		case patternLowerTriangle:
			svg.Triangle(0, fullH, fullW, fullH, halfW, halfH, pattern.color)
		case patternCircle:
			svg.Circle(halfW, halfH, thirdW, pattern.color)
		case patternHorizontalLine:
			svg.Box(0, halfH-1, fullW, 2, pattern.color)
		case patternVerticalLine:
			svg.Box(halfW-1, 0, 2, fullH, pattern.color)
		}
	}
	// For debugging
	//svg.Pixel(0, 0, 0, 255, 0)
	//svg.Pixel(maxX, 0, 0, 0, 255)
	//svg.Pixel(0, maxY, 255, 255, 0)
	//svg.Pixel(maxX, maxY, 0, 255, 255)
}

// Generate a new SVG Page for a banner
func (b *Banner) SvgPage() *onthefly.Page {
	page, svg := onthefly.NewTinySVG(0, 0, bannerW, bannerH)
	desc := svg.AddNewTag("desc")
	desc.AddContent("Hello SVG")
	b.Draw(svg)
	return page
}

func newBanner() *Banner {
	return &Banner{}
}

func newPattern(color string, pattern int) *Pattern {
	return &Pattern{color, pattern}
}

func (b *Banner) AddPattern(p *Pattern) {
	if len(b.patterns) > maxPatterns {
		log.Fatalln("Too many patters for banner, max", maxPatterns)
	}
	b.patterns = append(b.patterns, p)
}

// Generate a new onthefly Page (HTML5 and CSS combined)
func indexPage(svgurl string) *onthefly.Page {

	// Create a new HTML5 page, with CSS included
	page := onthefly.NewHTML5Page("Banner")

	page.AddContent("Banner")

	// Change the margin (em is default)
	page.SetMargin(4)

	// Change the font family
	page.SetFontFamily("serif") // or: sans-serif

	// Change the color scheme
	page.SetColor("black", "#d0d0d0")

	// Include the generated SVG image on the page
	body, _ := page.GetTag("body")

	// CSS attributes for the body tag
	body.AddStyle("font-size", "2em")

	// Paragraph
	p := body.AddNewTag("p")

	// CSS style
	p.AddStyle("margin-top", "2em")

	var (
		tag          *onthefly.Tag
		useObjectTag = true
	)
	if useObjectTag {
		// object tag
		tag = p.AddNewTag("object")
		// HTML attributes
		tag.AddAttrib("data", svgurl)
		tag.AddAttrib("type", "image/svg+xml")
	} else {
		// img tag
		tag = p.AddNewTag("img")
		// HTML attributes
		tag.AddAttrib("src", svgurl)
		tag.AddAttrib("alt", "Banner")
	}

	// CSS style
	w := strconv.Itoa(bannerW * zoomFactor)
	h := strconv.Itoa(bannerH * zoomFactor)
	tag.AddStyle("width", w+"px")
	tag.AddStyle("height", h+"px")
	tag.AddStyle("border", "8px solid black")

	return page
}

// Set up the paths and handlers then start serving.
func main() {
	log.Println("onthefly ", onthefly.Version)

	b := newBanner()
	p := newPattern("red", patternUpperHalf)
	b.AddPattern(p)
	p = newPattern("blue", patternLowerHalf)
	b.AddPattern(p)
	p = newPattern("white", patternUpperTriangle)
	b.AddPattern(p)
	p = newPattern("black", patternLowerTriangle)
	b.AddPattern(p)
	p = newPattern("yellow", patternCircle)
	b.AddPattern(p)
	p = newPattern("green", patternUpperThird)
	b.AddPattern(p)
	p = newPattern("purple", patternLowerThird)
	b.AddPattern(p)
	p = newPattern("orange", patternHorizontalLine)
	b.AddPattern(p)
	p = newPattern("red", patternVerticalLine)
	b.AddPattern(p)

	// Create a Negroni instance and a ServeMux instance
	n := negroni.Classic()
	mux := http.NewServeMux()

	// Publish the generated SVG as "/circles.svg"
	svgurl := "/circles.svg"
	mux.HandleFunc(svgurl, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "image/svg+xml")
		fmt.Fprint(w, b.SvgPage().String())
	})

	// Generate a Page that includes the svg image
	page := indexPage(svgurl)
	// Publish the generated Page in a way that connects the HTML and CSS
	page.Publish(mux, "/", "/style.css", false)

	// Handler goes last
	n.UseHandler(mux)

	// Listen for requests at port 3000
	n.Run(":3000")
}
