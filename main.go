package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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

	// Same order as minecraft.gamepedia.com/Banner
	patternLowerThird         = iota // Base Fess Banner
	patternUpperThird                // Chief Fess Banner
	patternLeftThird                 // Pale Dexter Banner
	patternRightThird                // Pale Sinister Banner
	patternCenterThird               // Pale Banner
	patternHorizontalLine            // Fess Banner
	patternDiagonal1                 // Bend Banner
	patternDiagonal2                 // Bend Sinister Banner
	patterStripes                    // Paly Banner
	patternDiaCross                  // Saltire Banner
	patternCross                     // Cross Banner
	patternUpperLeftTriangle         // Per Bend Sinister Banner
	patternUpperRightTriangle        // Per Bend Banner
	patternLowerLeftTriangle         // Per Bend Inverted Banner
	patternLowerRightTriangle        // Per Bend Sinister Inverted Banner
	patternLeftHalf                  // Per Pale Banner
	patternRightHalf                 // Per Pale Inverted Banner
	patternUpperHalf                 // Per Fess Banner
	patternLowerHalf                 // Per Fess Inverted Banner
	patternLowerLeftSquare           // Base Dexter Canton Banner
	patternLowerRightSquare          // Base Sinister Canton Banner
	patternUpperLeftSquare           // Chief Dexter Canton Banner
	patternUpperRightSqaure          // Chief Sinister Canton Banner
	patternLowerTriangle             // Chevron Banner
	patternUpperTriangle             // Inverted Chevron Banner
	patternLowerWaves                // Base Indented Banner
	patternUpperWaves                // Chief Indented Banner
	patternCircle                    // Roundel Banner
	patternDiamond                   // Lozenge Banner
	patternBorder                    // Bordure Banner
	patternWaveBorder                // Black/Dyed Borduer Indented Banner
	patternBricks                    // Black/Dyed Field Masoned Banner
	patternGradientDown              // Gradient Banner
	patternGradientUp                // Base Gradient Banner
	patternCreeper                   // Black/Dyed Creeper Charge Banner
	patternSkull                     // Black/Dyed Skull Charge Banner
	patternFlower                    // Black/Dyed Flower Charge Banner
	patternLogo                      // Black/Dyed Mojang Charge Banner

	// Custom patterns
	patternFull // Background

	colorWhite = iota
	colorOrange
	colorMagenta
	colorLightBlue
	colorYellow
	colorLime
	colorPink
	colorGray
	colorLightGray
	colorCyan
	colorPurple
	colorBlue
	colorBrown
	colorGreen
	colorRed
	colorBlack

	// Custom colors
	colorBrightWhite
)

var (
	colors = map[int]string{
		colorWhite:       "#989898",
		colorOrange:      "#804d23",
		colorMagenta:     "#6b3181",
		colorLightBlue:   "#3e5b7f",
		colorYellow:      "#898923",
		colorLime:        "#4c7815",
		colorPink:        "#8f4c62",
		colorGray:        "#313131",
		colorLightGray:   "#5b5b5b",
		colorCyan:        "#314c5b",
		colorPurple:      "#4d2b6a",
		colorBlue:        "#23316b",
		colorBrown:       "#3e3123",
		colorGreen:       "#3e4c23",
		colorRed:         "#5c2323",
		colorBlack:       "#151515",
		colorBrightWhite: "#ffffff",
	}
)

type Pattern struct {
	pattern int
	color   int
}

type Banner struct {
	patterns []*Pattern
	curpat   int
}

func (b *Banner) Draw(svg *onthefly.Tag) {
	// generate the picture, then draw
	for _, p := range b.patterns {
		color, ok := colors[p.color]
		if !ok {
			log.Fatalln("Invalid color ID: ", p.color)
		}
		//log.Println("color:", color)
		switch p.pattern {
		case patternLowerThird: // Base Fess Banner
			svg.Box(0, fullH-thirdH, fullW, fullH-thirdH, color)
		case patternUpperThird: // Chief Fess Banner
			svg.Box(0, 0, fullW, thirdH, color)
		case patternLeftThird: // Pale Dexter Banner
		case patternRightThird: // Pale Sinister Banner
		case patternCenterThird: // Pale Banner
			svg.Box(halfW-(thirdW/2), 0, thirdW, fullH, color)
		case patternHorizontalLine: // Fess Banner
			svg.Box(0, halfH-1, fullW, 2, color)
		case patternDiagonal1: // Bend Banner
		case patternDiagonal2: // Bend Sinister Banner
		case patterStripes: // Paly Banner
		case patternDiaCross: // Saltire Banner
		case patternCross: // Cross Banner
		case patternUpperLeftTriangle: // Per Bend Sinister Banner
		case patternUpperRightTriangle: // Per Bend Banner
		case patternLowerLeftTriangle: // Per Bend Inverted Banner
		case patternLowerRightTriangle: // Per Bend Sinister Inverted Banner
		case patternLeftHalf: // Per Pale Banner
		case patternRightHalf: // Per Pale Inverted Banner
		case patternUpperHalf: // Per Fess Banner
			svg.Box(0, 0, fullW, halfH, color)
		case patternLowerHalf: // Per Fess Inverted Banner
			svg.Box(0, halfH, fullW, halfH, color)
		case patternLowerLeftSquare: // Base Dexter Canton Banner
		case patternLowerRightSquare: // Base Sinister Canton Banner
		case patternUpperLeftSquare: // Chief Dexter Canton Banner
		case patternUpperRightSqaure: // Chief Sinister Canton Banner
		case patternLowerTriangle: // Chevron Banner
			svg.Triangle(0, fullH, fullW, fullH, halfW, halfH, color)
		case patternUpperTriangle: // Inverted Chevron Banner
			svg.Triangle(0, 0, fullW, 0, halfW, halfH, color)
		case patternLowerWaves: // Base Indented Banner
		case patternUpperWaves: // Chief Indented Banner
		case patternCircle: // Roundel Banner
			svg.Circle(halfW, halfH, thirdW, color)
		case patternDiamond: // Lozenge Banner
		case patternBorder: // Bordure Banner
		case patternWaveBorder: // Black/Dyed Borduer Indented Banner
		case patternBricks: // Black/Dyed Field Masoned Banner
		case patternGradientDown: // Gradient Banner
		case patternGradientUp: // Base Gradient Banner
		case patternCreeper: // Black/Dyed Creeper Charge Banner
		case patternSkull: // Black/Dyed Skull Charge Banner
		case patternFlower: // Black/Dyed Flower Charge Banner
		case patternLogo: // Black/Dyed Mojang Charge Banner
		case patternFull:
			svg.Box(0, 0, fullW, fullH, color)
		}
	}
	// For debugging
	//svg.Pixel(0, 0, 0, 255, 0)
	//svg.Pixel(maxX, 0, 0, 0, 255)
	//svg.Pixel(0, maxY, 255, 255, 0)
	//svg.Pixel(maxX, maxY, 0, 255, 255)
}

// Generate a new SVG Page for a banner
func (b *Banner) SVGpage() *onthefly.Page {
	page, svg := onthefly.NewTinySVG(0, 0, bannerW, bannerH)
	desc := svg.AddNewTag("desc")
	desc.AddContent("A banner")
	b.Draw(svg)
	return page
}

func NewBanner() *Banner {
	return &Banner{}
}

func NewPattern(pattern, color int) *Pattern {
	if (pattern < patternLowerThird) || (pattern > patternFull) {
		log.Fatalln("Invalid pattern ID: ", pattern)
	}
	return &Pattern{pattern, color}
}

func (b *Banner) AddPattern(p *Pattern) {
	if len(b.patterns) > maxPatterns {
		log.Fatalln("Too many patters for banner, max", maxPatterns)
	}
	b.patterns = append(b.patterns, p)
}

// Generate a new onthefly Page (HTML5 and CSS combined)
func mainPage(svgurls []string) *onthefly.Page {

	title := "Banners"

	// Create a new HTML5 page, with CSS included
	page := onthefly.NewHTML5Page(title)
	page.AddContent(title)

	// Change the margin (em is default)
	page.SetMargin(4)

	// Change the font family
	page.SetFontFamily("sans-serif")

	// Change the color scheme
	page.SetColor("black", "#a0a0a0")

	// Include the generated SVG image on the page
	body, _ := page.GetTag("body")

	// CSS attributes for the body tag
	body.AddStyle("font-size", "2em")

	// Paragraph
	p := body.AddNewTag("p")

	// CSS style
	p.AddStyle("margin-top", "2em")

	// Add images
	var (
		tag          *onthefly.Tag
		useObjectTag = false
	)
	for _, svgurl := range svgurls {
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
	}

	// CSS style
	w := strconv.Itoa(bannerW * zoomFactor)
	h := strconv.Itoa(bannerH * zoomFactor)
	tag.AddStyle("width", w+"px")
	tag.AddStyle("height", h+"px")
	tag.AddStyle("border", "8px solid black")

	return page
}

func randomColor() int {
	return colorWhite + rand.Intn((colorBlack-colorWhite)+1)
}

func seed() {
	rand.Seed(time.Now().UnixNano())
}

// Set up the paths and handlers then start serving.
func main() {
	seed()

	// Create a Negroni instance and a ServeMux instance
	n := negroni.Classic()
	mux := http.NewServeMux()

	var (
		svgurls []string
		b       *Banner
	)
	for i := patternLowerThird; i < patternLogo; i++ {

		b = NewBanner()
		b.AddPattern(NewPattern(patternFull, colorBrightWhite))
		b.AddPattern(NewPattern(i, colorRed))

		svgString := b.SVGpage().String()

		// Publish the generated SVG as "/img/banner_NNN.svg"
		svgurl := fmt.Sprintf("/img/banner_%d.svg", i)
		mux.HandleFunc(svgurl, func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "image/svg+xml")
			fmt.Fprint(w, svgString)
		})

		svgurls = append(svgurls, svgurl)
	}

	// Generate a Page that includes the svg images
	page := mainPage(svgurls)
	// Publish the generated Page in a way that connects the HTML and CSS
	page.Publish(mux, "/", "/css/banner.css", false)

	// Handler goes last
	n.UseHandler(mux)

	// Listen for requests at port 3000
	n.Run(":3000")
}
