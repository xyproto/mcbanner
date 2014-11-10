package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"math"

	"github.com/codegangsta/negroni"
	"github.com/xyproto/onthefly"
	"github.com/unrolled/render"
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
	fifthW = bannerW / 5
	fifthH = bannerH / 5
	sixthW = bannerW / 6
	sixthH = bannerH / 6
	tenthW = bannerW / 10
	tenthH = bannerH / 10

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
			svg.Box(0, 0, thirdW, fullH, color)
		case patternRightThird: // Pale Sinister Banner
			svg.Box(fullW-thirdW, 0, thirdW, fullH, color)
		case patternCenterThird: // Pale Banner
			svg.Box(halfW-(thirdW/2), 0, thirdW, fullH, color)
		case patternHorizontalLine: // Fess Banner
			svg.Box(0, halfH-1, fullW, thirdW/2, color)
		case patternDiagonal1: // Bend Banner
			svg.Line(0, 0, fullW, fullH, thirdW, color)
		case patternDiagonal2: // Bend Sinister Banner
			svg.Line(fullW, 0, 0, fullH, thirdW, color)
		case patterStripes: // Paly Banner
			for i := 0.15; i < 1.0; i += 0.25 {
				x := int(float64(fullW) * i)
				svg.Line(x, 0, x, fullH, sixthW, color)
			}
		case patternDiaCross: // Saltire Banner
			svg.Line(0, 0, fullW, fullH, thirdW, color)
			svg.Line(fullW, 0, 0, fullH, thirdW, color)
		case patternCross: // Cross Banner
			svg.Line(0, halfH, fullW, halfH, sixthW, color)
			svg.Line(halfW, 0, halfW, fullH, sixthW, color)
		case patternUpperLeftTriangle: // Per Bend Sinister Banner
			svg.Triangle(0, 0, fullW, 0, 0, fullH, color)
		case patternUpperRightTriangle: // Per Bend Banner
			svg.Triangle(0, 0, fullW, 0, fullW, fullH, color)
		case patternLowerLeftTriangle: // Per Bend Inverted Banner
			svg.Triangle(0, fullH, 0, 0, fullW, fullH, color)
		case patternLowerRightTriangle: // Per Bend Sinister Inverted Banner
			svg.Triangle(0, fullH, fullW, 0, fullW, fullH, color)
		case patternLeftHalf: // Per Pale Banner
			svg.Box(0, 0, halfW, fullH, color)
		case patternRightHalf: // Per Pale Inverted Banner
			svg.Box(halfW, 0, halfW, fullH, color)
		case patternUpperHalf: // Per Fess Banner
			svg.Box(0, 0, fullW, halfH, color)
		case patternLowerHalf: // Per Fess Inverted Banner
			svg.Box(0, halfH, fullW, halfH, color)
		case patternLowerLeftSquare: // Base Dexter Canton Banner
			svg.Box(0, fullH-thirdH, halfW, thirdH, color)
		case patternLowerRightSquare: // Base Sinister Canton Banner
			svg.Box(halfW, fullH-thirdH, halfW, thirdH, color)
		case patternUpperLeftSquare: // Chief Dexter Canton Banner
			svg.Box(0, 0, halfW, thirdH, color)
		case patternUpperRightSqaure: // Chief Sinister Canton Banner
			svg.Box(halfW, 0, halfW, thirdH, color)
		case patternLowerTriangle: // Chevron Banner
			svg.Triangle(0, fullH, fullW, fullH, halfW, halfH, color)
		case patternUpperTriangle: // Inverted Chevron Banner
			svg.Triangle(0, 0, fullW, 0, halfW, halfH, color)
		case patternLowerWaves: // Base Indented Banner
			svg.Ellipse(sixthW, fullH, fifthW, thirdW, color)
			svg.Ellipse(halfW, fullH, fifthW, thirdW, color)
			svg.Ellipse(fullW-sixthW, fullH, fifthW, thirdW, color)
		case patternUpperWaves: // Chief Indented Banner
			svg.Ellipse(sixthW, 0, fifthW, thirdW, color)
			svg.Ellipse(halfW, 0, fifthW, thirdW, color)
			svg.Ellipse(fullW-sixthW, 0, fifthW, thirdW, color)
		case patternCircle: // Roundel Banner
			svg.Circle(halfW, halfH, thirdW, color)
		case patternDiamond: // Lozenge Banner
			svg.Poly4(halfW, fifthH,
				fullW-fifthW, halfH,
				halfW, fullH-fifthH,
				fifthW, halfH,
				color)
		case patternBorder: // Bordure Banner
			svg.Line(0, 0, fullW, 0, sixthW, color)
			svg.Line(fullW, 0, fullW, fullH, sixthW, color)
			svg.Line(0, fullH, fullW, fullH, sixthW, color)
			svg.Line(0, 0, 0, fullH, sixthW, color)
		case patternWaveBorder: // Black/Dyed Borduer Indented Banner
			br := halfH             // radius of circles in corners
			ofs1 := fifthH          // offset away from the center, corner circles
			brx := sixthH           // radius x of circles along the sides
			bry := fullH / 7        // radius y of circles along the sides
			ofs2 := sixthH / 3      // offset away from the center, side circles
			bh := fullH/4 + fullH/8 // placement of circles along the sides
			// corners
			svg.Ellipse(0-ofs1, 0-ofs1, br, br, color)
			svg.Ellipse(0-ofs1, fullH+ofs1, br, br, color)
			svg.Ellipse(fullW+ofs1, 0-ofs1, br, br, color)
			svg.Ellipse(fullW+ofs1, fullH+ofs1, br, br, color)
			// sides
			svg.Ellipse(0-ofs2, bh, brx, bry, color)
			svg.Ellipse(0-ofs2, fullH-bh, brx, bry, color)
			svg.Ellipse(fullW+ofs2, bh, brx, bry, color)
			svg.Ellipse(fullW+ofs2, fullH-bh, brx, bry, color)
		case patternBricks: // Black/Dyed Field Masoned Banner
			brickW := fullW / 4
			brickH := fullH / 13
			ofs := 0
			for y := 0; y < fullH; y += brickH + 1 {
				for x := ofs; x < fullW; x += brickW + 1 {
					svg.Box(x, y, brickW, brickH, color)
				}
				// Alternate the x offset for every row of bricks
				if ofs == 0 {
					ofs = -brickH
				} else {
					ofs = 0
				}
			}
		case patternGradientDown: // Gradient Banner
			ur, _ := strconv.ParseUint(color[1:3], 16, 0)
			ug, _ := strconv.ParseUint(color[3:5], 16, 0)
			ub, _ := strconv.ParseUint(color[5:7], 16, 0)
			r := int(ur)
			g := int(ug)
			b := int(ub)
			a := 1.0
			for y := 0; y < fullH; y++ {
				a = 1.0 - (float64(y) / float64(fullH))
				svg.Box(0, y, fullW, 1, onthefly.ColorStringAlpha(r, g, b, a))
			}
		case patternGradientUp: // Base Gradient Banner
			ur, _ := strconv.ParseUint(color[1:3], 16, 0)
			ug, _ := strconv.ParseUint(color[3:5], 16, 0)
			ub, _ := strconv.ParseUint(color[5:7], 16, 0)
			r := int(ur)
			g := int(ug)
			b := int(ub)
			a := 1.0
			for y := 0; y < fullH; y++ {
				a = float64(y) / float64(fullH)
				svg.Box(0, y, fullW, 1, onthefly.ColorStringAlpha(r, g, b, a))
			}
		case patternCreeper: // Black/Dyed Creeper Charge Banner
			// eyes
			svg.Box(sixthW-tenthW/2, thirdH-tenthH/2, halfW/2, halfW/2, color)
			svg.Box(fullW-(sixthW+halfW/2)+tenthW/2, thirdH-tenthH/2, halfW/2, halfW/2, color)
			// nose
			svg.Line(halfW, thirdH+tenthH, halfW, halfH+(halfW/3), halfW/2, color)
			// whiskers
			svg.Box(sixthW+sixthW, thirdH+fifthH, tenthW, halfW/2, color)
			svg.Box(fullW-(sixthW+sixthW)-tenthW, thirdH+fifthH, tenthW, halfW/2, color)
		case patternSkull: // Black/Dyed Skull Charge Banner
			hy := (tenthH/2)
			boxx := fifthW+(tenthW/2)
			boxy := halfH-sixthH

			// Top of the head
			svg.Box(boxx, boxy-hy, halfW, hy*2, color)

			// Side of face
			svg.Box(boxx, boxy+hy, hy, hy*2, color)
			svg.Box(boxx+(halfW-hy), boxy+hy, hy, hy*2, color)

			// Nose
			svg.Box(boxx+(halfW-hy)/2, boxy+hy, hy, hy/2, color)

			// Cheeks
			svg.Box(boxx+(halfW-hy)/2-hy, boxy+hy+hy/2, hy, hy/2, color)
			svg.Box(boxx+(halfW-hy)/2+hy, boxy+hy+hy/2, hy, hy/2, color)

			// Bottom of face
			svg.Box(boxx, boxy+hy*3-(hy/2), halfW, 1, color)

			// The cross
			ofs := thirdH // offset from bottom
			svg.Line(sixthW, fullH-(sixthH+ofs), fullW-sixthW, fullH-ofs, tenthW, color)
			svg.Line(fullW-sixthW, fullH-(sixthH+ofs), sixthW, fullH-ofs, tenthW, color)

			// Ends of cross
			svg.Circle(hy+hy/2, halfH+hy/2, hy/2, color)
			svg.Circle(fullW-(hy+hy/2), halfH+hy/2, hy/2, color)
			svg.Circle(hy+hy/2, halfH+hy*4-hy/2, hy/2, color)
			svg.Circle(fullW-(hy+hy/2), halfH+hy*4-hy/2, hy/2, color)

		case patternFlower: // Black/Dyed Flower Charge Banner
			svg.Circle(halfW, halfH, sixthW, color)
			numcircles := 10
			step := (math.Pi*2.0)/float64(numcircles)
			radius := fullW/20
			spacing := float64(thirdW)
			linet := tenthW/2 // line thickness
			jointt := tenthW/2 // line joint thickness (circles)
			var (
				linex, liney, oldx, oldy, firstx, firsty int
			)
			for r := 0.0; r < math.Pi*2.0; r += step {
				// Draw outer circles
				x := int(math.Floor(math.Cos(r)*spacing + 0.5)) + halfW
				y := int(math.Floor(math.Sin(r)*spacing + 0.5)) + halfH
				svg.Circle(x, y, radius, color)
				// Draw circle made out of lines
				oldx = linex
				oldy = liney
				linex = int(math.Floor(math.Cos(r)*spacing*0.8 + 0.5)) + halfW
				liney = int(math.Floor(math.Sin(r)*spacing*0.8 + 0.5)) + halfH
				if oldx == 0 {
					firstx = linex
					firsty = liney
				} else {
					svg.Line(oldx, oldy, linex, liney, linet, color)
					svg.Circle(linex, liney, jointt, color)
				}
			}
			svg.Line(linex, liney, firstx, firsty, linet, color)
			svg.Circle(firstx, firsty, jointt, color)
		case patternLogo: // Black/Dyed Mojang Charge Banner
			numsteps := 10
			step := (math.Pi*2.0)/float64(numsteps)
			spacing := float64(thirdW)
			linet := fifthW // line thickness
			var (
				linex, liney, oldx, oldy int
			)
			for r := math.Pi/2; r < (math.Pi*2.0)-(math.Pi/2); r += step {
				// Draw circle made out of lines
				oldx = linex
				oldy = liney
				linex = int(math.Floor(math.Cos(r)*spacing*0.8 + 0.5)) + halfW
				liney = int(math.Floor(math.Sin(r)*spacing*0.8 + 0.5)) + halfH
				if oldx == 0 {
					svg.Line(linex, liney, fullW-fifthW, liney, linet, color)
				} else {
					svg.Line(oldx, oldy, linex, liney, linet, color)
				}
			}
			svg.Line(linex, liney, fullW-fifthW, liney+fifthW, linet, color)
			svg.Line(fullW-thirdW, thirdH, (fullW-thirdW)+tenthW, thirdH+tenthW, tenthW, color)
		case patternFull:
			svg.Box(0, 0, fullW, fullH, color)
		}
	}
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
func patternGalleryPage(title string, svgurls []string) *onthefly.Page {
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

// Create a pattern gallery under /patterns
func patternGallery(mux *http.ServeMux, path string) {
	var (
		svgurls []string
		b       *Banner
	)
	for i := patternLowerThird; i <= patternLogo; i++ {

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
	page := patternGalleryPage("Pattern gallery", svgurls)
	// Publish the generated Page in a way that connects the HTML and CSS
	page.Publish(mux, path, "/css/banner.css", false)
}

// Set up the paths and handlers then start serving.
func main() {
	// Seed the random number generator
	seed()

	// Create a Negroni instance and a ServeMux instance
	n := negroni.Classic()
	r := render.New(render.Options{})
	mux := http.NewServeMux()

	// Main page
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		data := map[string]string{
			"title": "Banner Generator",
		}

		// Reload template
		r = render.New(render.Options{})

		// Render and return
		r.HTML(w, http.StatusOK, "index", data)
	})

	// Publish a pattern gallery
    patternGallery(mux, "/patterns")

	// Handler goes last
	n.UseHandler(mux)

	// Listen for requests at port 3000
	n.Run(":3000")
}
