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
	bannerW    = 2
	bannerH    = 2
	zoomFactor = 20 // 20x zoom

	halfW = bannerW / 2
	halfH = bannerH / 2
	fullW = bannerW - 1
	fullH = bannerH - 1

	maxPatterns = 6

	patternTopHalf    = 0
	patternBottomHalf = 1
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
		if pattern.ptype == patternTopHalf {
			svg.Box(0, 0, fullW, halfH, pattern.color)
		}
	}
	svg.Pixel(0, 0, 0, 255, 0)
	svg.Pixel(fullW, 0, 0, 0, 255)
	svg.Pixel(0, fullH, 255, 255, 0)
	svg.Pixel(fullW, fullH, 0, 255, 255)
	// x, y, w, h, color
	//svg.Pixel(x, y, r, g, b)
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
	if len(b.patterns) > 6 {
		log.Fatalln("Too many patters for banner, max 6")
	}
	b.patterns = append(b.patterns, p)
}

// Generate a new onthefly Page (HTML5 and CSS combined)
func indexPage(svgurl string) *onthefly.Page {

	// Create a new HTML5 page, with CSS included
	page := onthefly.NewHTML5Page("Banner")

	page.AddContent("blabla")

	// Change the margin (em is default)
	page.SetMargin(4)

	// Change the font family
	page.SetFontFamily("serif") // or: sans-serif

	// Change the color scheme
	page.SetColor("black", "#d0d0d0")

	// Include the generated SVG image on the page
	body, err := page.GetTag("body")
	if err == nil {
		// CSS attributes for the body tag
		body.AddStyle("font-size", "2em")

		// Paragraph
		p := body.AddNewTag("p")

		// CSS style
		p.AddStyle("margin-top", "2em")

		objectOrImg := "object"

		if objectOrImg == "object" {
			// object tag
			tag := p.AddNewTag("object")
			// HTML attributes
			tag.AddAttrib("data", svgurl)
			tag.AddAttrib("type", "image/svg+xml")
		} else {
			// img tag
			tag := p.AddNewTag("img")
			// HTML attributes
			tag.AddAttrib("src", svgurl)
			tag.AddAttrib("alt", "Banner")
		}

		// CSS style
		w := strconv.Itoa(bannerW * zoomFactor)
		h := strconv.Itoa(bannerH * zoomFactor)
		object.AddStyle("width", w + "px")
		object.AddStyle("height", h + "px")
		object.AddStyle("border", "8px solid black")
	}

	return page
}

// Set up the paths and handlers then start serving.
func main() {
	log.Println("onthefly ", onthefly.Version)

	b := newBanner()
	p := newPattern("red", patternTopHalf)
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
