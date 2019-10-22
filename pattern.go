package mcbanner

import (
	"errors"
	"fmt"
	"github.com/xyproto/tinysvg"
	"math"
	"strconv"
)

const (
	maxPatterns = 6

	// Same order as minecraft.gamepedia.com/Banner
	PatternLowerThird         = iota // Base Fess Banner
	PatternUpperThird                // Chief Fess Banner
	PatternLeftThird                 // Pale Dexter Banner
	PatternRightThird                // Pale Sinister Banner
	PatternCenterThird               // Pale Banner
	PatternHorizontalLine            // Fess Banner
	PatternDiagonal1                 // Bend Banner
	PatternDiagonal2                 // Bend Sinister Banner
	patterStripes                    // Paly Banner
	PatternDiaCross                  // Saltire Banner
	PatternCross                     // Cross Banner
	PatternUpperLeftTriangle         // Per Bend Sinister Banner
	PatternUpperRightTriangle        // Per Bend Banner
	PatternLowerLeftTriangle         // Per Bend Inverted Banner
	PatternLowerRightTriangle        // Per Bend Sinister Inverted Banner
	PatternLeftHalf                  // Per Pale Banner
	PatternRightHalf                 // Per Pale Inverted Banner
	PatternUpperHalf                 // Per Fess Banner
	PatternLowerHalf                 // Per Fess Inverted Banner
	PatternLowerLeftSquare           // Base Dexter Canton Banner
	PatternLowerRightSquare          // Base Sinister Canton Banner
	PatternUpperLeftSquare           // Chief Dexter Canton Banner
	PatternUpperRightSqaure          // Chief Sinister Canton Banner
	PatternLowerTriangle             // Chevron Banner
	PatternUpperTriangle             // Inverted Chevron Banner
	PatternLowerWaves                // Base Indented Banner
	PatternUpperWaves                // Chief Indented Banner
	PatternCircle                    // Roundel Banner
	PatternDiamond                   // Lozenge Banner
	PatternBorder                    // Bordure Banner
	PatternWaveBorder                // Black/Dyed Borduer Indented Banner
	PatternBricks                    // Black/Dyed Field Masoned Banner
	PatternGradientDown              // Gradient Banner
	PatternGradientUp                // Base Gradient Banner
	PatternCreeper                   // Black/Dyed Creeper Charge Banner
	PatternSkull                     // Black/Dyed Skull Charge Banner
	PatternFlower                    // Black/Dyed Flower Charge Banner
	PatternLogo                      // Black/Dyed Mojang Charge Banner

	// Custom Patterns
	PatternFull // Background

	ColorWhite = iota
	ColorOrange
	ColorMagenta
	ColorLightBlue
	ColorYellow
	ColorLime
	ColorPink
	ColorGray
	ColorLightGray
	ColorCyan
	ColorPurple
	ColorBlue
	ColorBrown
	ColorGreen
	ColorRed
	ColorBlack

	// Custom Colors
	ColorBrightWhite
)

var PatternDesc = map[int]string{
	PatternLowerThird:         "Base Fess Banner",
	PatternUpperThird:         "Chief Fess Banner",
	PatternLeftThird:          "Pale Dexter Banner",
	PatternRightThird:         "Pale Sinister Banner",
	PatternCenterThird:        "Pale Banner",
	PatternHorizontalLine:     "Fess Banner",
	PatternDiagonal1:          "Bend Banner",
	PatternDiagonal2:          "Bend Sinister Banner",
	patterStripes:             "Paly Banner",
	PatternDiaCross:           "Saltire Banner",
	PatternCross:              "Cross Banner",
	PatternUpperLeftTriangle:  "Per Bend Sinister Banner",
	PatternUpperRightTriangle: "Per Bend Banner",
	PatternLowerLeftTriangle:  "Per Bend Inverted Banner",
	PatternLowerRightTriangle: "Per Bend Sinister Inverted Banner",
	PatternLeftHalf:           "Per Pale Banner",
	PatternRightHalf:          "Per Pale Inverted Banner",
	PatternUpperHalf:          "Per Fess Banner",
	PatternLowerHalf:          "Per Fess Inverted Banner",
	PatternLowerLeftSquare:    "Base Dexter Canton Banner",
	PatternLowerRightSquare:   "Base Sinister Canton Banner",
	PatternUpperLeftSquare:    "Chief Dexter Canton Banner",
	PatternUpperRightSqaure:   "Chief Sinister Canton Banner",
	PatternLowerTriangle:      "Chevron Banner",
	PatternUpperTriangle:      "Inverted Chevron Banner",
	PatternLowerWaves:         "Base Indented Banner",
	PatternUpperWaves:         "Chief Indented Banner",
	PatternCircle:             "Roundel Banner",
	PatternDiamond:            "Lozenge Banner",
	PatternBorder:             "Bordure Banner",
	PatternWaveBorder:         "Borduer Indented Banner",
	PatternBricks:             "Field Masoned Banner",
	PatternGradientDown:       "Gradient Banner",
	PatternGradientUp:         "Base Gradient Banner",
	PatternCreeper:            "Creeper Charge Banner",
	PatternSkull:              "Skull Charge Banner",
	PatternFlower:             "Flower Charge Banner",
	PatternLogo:               "Mojang Charge Banner",
	PatternFull:               "Background",
}

type Pattern struct {
	pattern int
	color   int
}

func (pat *Pattern) Valid() error {
	if (pat.pattern < PatternLowerThird) || (pat.pattern > PatternFull) {
		return errors.New(fmt.Sprintf("Invalid Pattern ID: %d\n", pat.pattern))
	}
	if (pat.color < ColorWhite) || (pat.color > ColorBrightWhite) {
		return errors.New(fmt.Sprintf("Invalid Color ID: %d\n", pat.color))
	}
	return nil
}

func NewPattern(patId, Color int) (*Pattern, error) {
	pat := &Pattern{patId, Color}
	if err := pat.Valid(); err != nil {
		return nil, err
	}
	return pat, nil
}

func DrawPattern(svg *tinysvg.Tag, Pattern int, Color string) {
	switch Pattern {
	case PatternLowerThird: // Base Fess Banner
		svg.Box(0, fullH-thirdH, fullW, fullH-thirdH, Color)
	case PatternUpperThird: // Chief Fess Banner
		svg.Box(0, 0, fullW, thirdH, Color)
	case PatternLeftThird: // Pale Dexter Banner
		svg.Box(0, 0, thirdW, fullH, Color)
	case PatternRightThird: // Pale Sinister Banner
		svg.Box(fullW-thirdW, 0, thirdW, fullH, Color)
	case PatternCenterThird: // Pale Banner
		svg.Box(halfW-(thirdW/2), 0, thirdW, fullH, Color)
	case PatternHorizontalLine: // Fess Banner
		svg.Box(0, halfH-1, fullW, thirdW/2, Color)
	case PatternDiagonal1: // Bend Banner
		svg.Line(0, 0, fullW, fullH, thirdW, Color)
	case PatternDiagonal2: // Bend Sinister Banner
		svg.Line(fullW, 0, 0, fullH, thirdW, Color)
	case patterStripes: // Paly Banner
		for i := 0.15; i < 1.0; i += 0.25 {
			x := int(float64(fullW) * i)
			svg.Line(x, 0, x, fullH, sixthW, Color)
		}
	case PatternDiaCross: // Saltire Banner
		svg.Line(0, 0, fullW, fullH, thirdW, Color)
		svg.Line(fullW, 0, 0, fullH, thirdW, Color)
	case PatternCross: // Cross Banner
		svg.Line(0, halfH, fullW, halfH, sixthW, Color)
		svg.Line(halfW, 0, halfW, fullH, sixthW, Color)
	case PatternUpperLeftTriangle: // Per Bend Sinister Banner
		svg.Triangle(0, 0, fullW, 0, 0, fullH, Color)
	case PatternUpperRightTriangle: // Per Bend Banner
		svg.Triangle(0, 0, fullW, 0, fullW, fullH, Color)
	case PatternLowerLeftTriangle: // Per Bend Inverted Banner
		svg.Triangle(0, fullH, 0, 0, fullW, fullH, Color)
	case PatternLowerRightTriangle: // Per Bend Sinister Inverted Banner
		svg.Triangle(0, fullH, fullW, 0, fullW, fullH, Color)
	case PatternLeftHalf: // Per Pale Banner
		svg.Box(0, 0, halfW, fullH, Color)
	case PatternRightHalf: // Per Pale Inverted Banner
		svg.Box(halfW, 0, halfW, fullH, Color)
	case PatternUpperHalf: // Per Fess Banner
		svg.Box(0, 0, fullW, halfH, Color)
	case PatternLowerHalf: // Per Fess Inverted Banner
		svg.Box(0, halfH, fullW, halfH, Color)
	case PatternLowerLeftSquare: // Base Dexter Canton Banner
		svg.Box(0, fullH-thirdH, halfW, thirdH, Color)
	case PatternLowerRightSquare: // Base Sinister Canton Banner
		svg.Box(halfW, fullH-thirdH, halfW, thirdH, Color)
	case PatternUpperLeftSquare: // Chief Dexter Canton Banner
		svg.Box(0, 0, halfW, thirdH, Color)
	case PatternUpperRightSqaure: // Chief Sinister Canton Banner
		svg.Box(halfW, 0, halfW, thirdH, Color)
	case PatternLowerTriangle: // Chevron Banner
		svg.Triangle(0, fullH, fullW, fullH, halfW, halfH, Color)
	case PatternUpperTriangle: // Inverted Chevron Banner
		svg.Triangle(0, 0, fullW, 0, halfW, halfH, Color)
	case PatternLowerWaves: // Base Indented Banner
		svg.Ellipse(sixthW, fullH, fifthW, thirdW, Color)
		svg.Ellipse(halfW, fullH, fifthW, thirdW, Color)
		svg.Ellipse(fullW-sixthW, fullH, fifthW, thirdW, Color)
	case PatternUpperWaves: // Chief Indented Banner
		svg.Ellipse(sixthW, 0, fifthW, thirdW, Color)
		svg.Ellipse(halfW, 0, fifthW, thirdW, Color)
		svg.Ellipse(fullW-sixthW, 0, fifthW, thirdW, Color)
	case PatternCircle: // Roundel Banner
		svg.Circle(halfW, halfH, thirdW, Color)
	case PatternDiamond: // Lozenge Banner
		svg.Poly4(halfW, fifthH,
			fullW-fifthW, halfH,
			halfW, fullH-fifthH,
			fifthW, halfH,
			Color)
	case PatternBorder: // Bordure Banner
		svg.Line(0, 0, fullW, 0, sixthW, Color)
		svg.Line(fullW, 0, fullW, fullH, sixthW, Color)
		svg.Line(0, fullH, fullW, fullH, sixthW, Color)
		svg.Line(0, 0, 0, fullH, sixthW, Color)
	case PatternWaveBorder: // Black/Dyed Borduer Indented Banner
		br := halfH             // radius of circles in corners
		ofs1 := fifthH          // offset away from the center, corner circles
		brx := sixthH           // radius x of circles along the sides
		bry := fullH / 7        // radius y of circles along the sides
		ofs2 := sixthH / 3      // offset away from the center, side circles
		bh := fullH/4 + fullH/8 // placement of circles along the sides
		// corners
		svg.Ellipse(0-ofs1, 0-ofs1, br, br, Color)
		svg.Ellipse(0-ofs1, fullH+ofs1, br, br, Color)
		svg.Ellipse(fullW+ofs1, 0-ofs1, br, br, Color)
		svg.Ellipse(fullW+ofs1, fullH+ofs1, br, br, Color)
		// sides
		svg.Ellipse(0-ofs2, bh, brx, bry, Color)
		svg.Ellipse(0-ofs2, fullH-bh, brx, bry, Color)
		svg.Ellipse(fullW+ofs2, bh, brx, bry, Color)
		svg.Ellipse(fullW+ofs2, fullH-bh, brx, bry, Color)
	case PatternBricks: // Black/Dyed Field Masoned Banner
		brickW := fullW / 4
		brickH := fullH / 13
		ofs := 0
		for y := 0; y < fullH; y += brickH + 1 {
			for x := ofs; x < fullW; x += brickW + 1 {
				svg.Box(x, y, brickW, brickH, Color)
			}
			// Alternate the x offset for every row of bricks
			if ofs == 0 {
				ofs = -brickH
			} else {
				ofs = 0
			}
		}
	case PatternGradientDown: // Gradient Banner
		ur, _ := strconv.ParseUint(Color[1:3], 16, 0)
		ug, _ := strconv.ParseUint(Color[3:5], 16, 0)
		ub, _ := strconv.ParseUint(Color[5:7], 16, 0)
		r := int(ur)
		g := int(ug)
		b := int(ub)
		a := 1.0
		for y := 0; y < fullH; y++ {
			a = 1.0 - (float64(y) / float64(fullH))
			svg.Box(0, y, fullW, 1, string(tinysvg.ColorBytesAlpha(r, g, b, a)))
		}
	case PatternGradientUp: // Base Gradient Banner
		ur, _ := strconv.ParseUint(Color[1:3], 16, 0)
		ug, _ := strconv.ParseUint(Color[3:5], 16, 0)
		ub, _ := strconv.ParseUint(Color[5:7], 16, 0)
		r := int(ur)
		g := int(ug)
		b := int(ub)
		a := 1.0
		for y := 0; y < fullH; y++ {
			a = float64(y) / float64(fullH)
			svg.Box(0, y, fullW, 1, string(tinysvg.ColorBytesAlpha(r, g, b, a)))
		}
	case PatternCreeper: // Black/Dyed Creeper Charge Banner
		// eyes
		svg.Box(sixthW-tenthW/2, thirdH-tenthH/2, halfW/2, halfW/2, Color)
		svg.Box(fullW-(sixthW+halfW/2)+tenthW/2, thirdH-tenthH/2, halfW/2, halfW/2, Color)
		// nose
		svg.Line(halfW, thirdH+tenthH, halfW, halfH+(halfW/3), halfW/2, Color)
		// whiskers
		svg.Box(sixthW+sixthW, thirdH+fifthH, tenthW, halfW/2, Color)
		svg.Box(fullW-(sixthW+sixthW)-tenthW, thirdH+fifthH, tenthW, halfW/2, Color)
	case PatternSkull: // Black/Dyed Skull Charge Banner
		hy := (tenthH / 2)
		boxx := fifthW + (tenthW / 2)
		boxy := halfH - sixthH

		// Top of the head
		svg.Box(boxx, boxy-hy, halfW, hy*2, Color)

		// Side of face
		svg.Box(boxx, boxy+hy, hy, hy*2, Color)
		svg.Box(boxx+(halfW-hy), boxy+hy, hy, hy*2, Color)

		// Nose
		svg.Box(boxx+(halfW-hy)/2, boxy+hy, hy, hy/2, Color)

		// Cheeks
		svg.Box(boxx+(halfW-hy)/2-hy, boxy+hy+hy/2, hy, hy/2, Color)
		svg.Box(boxx+(halfW-hy)/2+hy, boxy+hy+hy/2, hy, hy/2, Color)

		// Bottom of face
		svg.Box(boxx, boxy+hy*3-(hy/2), halfW, 1, Color)

		// The cross
		ofs := thirdH // offset from bottom
		svg.Line(sixthW, fullH-(sixthH+ofs), fullW-sixthW, fullH-ofs, tenthW, Color)
		svg.Line(fullW-sixthW, fullH-(sixthH+ofs), sixthW, fullH-ofs, tenthW, Color)

		// Ends of cross
		svg.Circle(hy+hy/2, halfH+hy/2, hy/2, Color)
		svg.Circle(fullW-(hy+hy/2), halfH+hy/2, hy/2, Color)
		svg.Circle(hy+hy/2, halfH+hy*4-hy/2, hy/2, Color)
		svg.Circle(fullW-(hy+hy/2), halfH+hy*4-hy/2, hy/2, Color)

	case PatternFlower: // Black/Dyed Flower Charge Banner
		svg.Circle(halfW, halfH, sixthW, Color)
		numcircles := 10
		step := (math.Pi * 2.0) / float64(numcircles)
		radius := fullW / 20
		spacing := float64(thirdW)
		linet := tenthW / 2  // line thickness
		jointt := tenthW / 2 // line joint thickness (circles)
		var (
			linex, liney, oldx, oldy, firstx, firsty int
		)
		for r := 0.0; r < math.Pi*2.0; r += step {
			// Draw outer circles
			x := int(math.Floor(math.Cos(r)*spacing+0.5)) + halfW
			y := int(math.Floor(math.Sin(r)*spacing+0.5)) + halfH
			svg.Circle(x, y, radius, Color)
			// Draw circle made out of lines
			oldx = linex
			oldy = liney
			linex = int(math.Floor(math.Cos(r)*spacing*0.8+0.5)) + halfW
			liney = int(math.Floor(math.Sin(r)*spacing*0.8+0.5)) + halfH
			if oldx == 0 {
				firstx = linex
				firsty = liney
			} else {
				svg.Line(oldx, oldy, linex, liney, linet, Color)
				svg.Circle(linex, liney, jointt, Color)
			}
		}
		svg.Line(linex, liney, firstx, firsty, linet, Color)
		svg.Circle(firstx, firsty, jointt, Color)
	case PatternLogo: // Black/Dyed Mojang Charge Banner
		numsteps := 10
		step := (math.Pi * 2.0) / float64(numsteps)
		spacing := float64(thirdW)
		linet := fifthW // line thickness
		var (
			linex, liney, oldx, oldy int
		)
		for r := math.Pi / 2; r < (math.Pi*2.0)-(math.Pi/2); r += step {
			// Draw circle made out of lines
			oldx = linex
			oldy = liney
			linex = int(math.Floor(math.Cos(r)*spacing*0.8+0.5)) + halfW
			liney = int(math.Floor(math.Sin(r)*spacing*0.8+0.5)) + halfH
			if oldx == 0 {
				svg.Line(linex, liney, fullW-fifthW, liney, linet, Color)
			} else {
				svg.Line(oldx, oldy, linex, liney, linet, Color)
			}
		}
		svg.Line(linex, liney, fullW-fifthW, liney+fifthW, linet, Color)
		svg.Line(fullW-thirdW, thirdH, (fullW-thirdW)+tenthW, thirdH+tenthW, tenthW, Color)
	case PatternFull:
		svg.Box(0, 0, fullW, fullH, Color)
	}
}

func (pat *Pattern) PatternString() string {
	desc, ok := PatternDesc[pat.pattern]
	if !ok {
		return "Unknown Pattern"
	}
	return desc
}

// TODO: Switch the switch with a lookup like in PatternString
func (pat *Pattern) ColorString() string {
	switch pat.color {
	case ColorWhite:
		return "White"
	case ColorOrange:
		return "Orange"
	case ColorMagenta:
		return "Magenta"
	case ColorLightBlue:
		return "LightBlue"
	case ColorYellow:
		return "Yellow"
	case ColorLime:
		return "Lime"
	case ColorPink:
		return "Pink"
	case ColorGray:
		return "Gray"
	case ColorLightGray:
		return "LightGray"
	case ColorCyan:
		return "Cyan"
	case ColorPurple:
		return "Purple"
	case ColorBlue:
		return "Blue"
	case ColorBrown:
		return "Brown"
	case ColorGreen:
		return "Green"
	case ColorRed:
		return "Red"
	case ColorBlack:
		return "Black"
	case ColorBrightWhite:
		return "Custom color"
	default:
		return "Unknown color"
	}
}

func (pat *Pattern) String() string {
	return pat.ColorString() + " " + pat.PatternString()
}
