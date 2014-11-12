package main

import (
	"errors"
	"fmt"
	"github.com/xyproto/onthefly"
	"math"
	"strconv"
)

const (
	maxPatterns = 6

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

var patternDesc = map[int]string{
	patternLowerThird:         "Base Fess Banner",
	patternUpperThird:         "Chief Fess Banner",
	patternLeftThird:          "Pale Dexter Banner",
	patternRightThird:         "Pale Sinister Banner",
	patternCenterThird:        "Pale Banner",
	patternHorizontalLine:     "Fess Banner",
	patternDiagonal1:          "Bend Banner",
	patternDiagonal2:          "Bend Sinister Banner",
	patterStripes:             "Paly Banner",
	patternDiaCross:           "Saltire Banner",
	patternCross:              "Cross Banner",
	patternUpperLeftTriangle:  "Per Bend Sinister Banner",
	patternUpperRightTriangle: "Per Bend Banner",
	patternLowerLeftTriangle:  "Per Bend Inverted Banner",
	patternLowerRightTriangle: "Per Bend Sinister Inverted Banner",
	patternLeftHalf:           "Per Pale Banner",
	patternRightHalf:          "Per Pale Inverted Banner",
	patternUpperHalf:          "Per Fess Banner",
	patternLowerHalf:          "Per Fess Inverted Banner",
	patternLowerLeftSquare:    "Base Dexter Canton Banner",
	patternLowerRightSquare:   "Base Sinister Canton Banner",
	patternUpperLeftSquare:    "Chief Dexter Canton Banner",
	patternUpperRightSqaure:   "Chief Sinister Canton Banner",
	patternLowerTriangle:      "Chevron Banner",
	patternUpperTriangle:      "Inverted Chevron Banner",
	patternLowerWaves:         "Base Indented Banner",
	patternUpperWaves:         "Chief Indented Banner",
	patternCircle:             "Roundel Banner",
	patternDiamond:            "Lozenge Banner",
	patternBorder:             "Bordure Banner",
	patternWaveBorder:         "Borduer Indented Banner",
	patternBricks:             "Field Masoned Banner",
	patternGradientDown:       "Gradient Banner",
	patternGradientUp:         "Base Gradient Banner",
	patternCreeper:            "Creeper Charge Banner",
	patternSkull:              "Skull Charge Banner",
	patternFlower:             "Flower Charge Banner",
	patternLogo:               "Mojang Charge Banner",
	patternFull:               "Background",
}

type Pattern struct {
	pattern int
	color   int
}

func NewPattern(pattern, color int) (*Pattern, error) {
	if (pattern < patternLowerThird) || (pattern > patternFull) {
		return nil, errors.New(fmt.Sprintf("Invalid pattern ID: %d\n", pattern))
	}
	return &Pattern{pattern, color}, nil
}

func DrawPattern(svg *onthefly.Tag, pattern int, color string) {
	switch pattern {
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
		hy := (tenthH / 2)
		boxx := fifthW + (tenthW / 2)
		boxy := halfH - sixthH

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
			svg.Circle(x, y, radius, color)
			// Draw circle made out of lines
			oldx = linex
			oldy = liney
			linex = int(math.Floor(math.Cos(r)*spacing*0.8+0.5)) + halfW
			liney = int(math.Floor(math.Sin(r)*spacing*0.8+0.5)) + halfH
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

func (p *Pattern) PatternString() string {
	desc, ok := patternDesc[p.pattern]
	if !ok {
		return "Unknown pattern"
	}
	return desc
}

func (p *Pattern) ColorString() string {
	switch p.color {
	case colorWhite:
		return "White"
	case colorOrange:
		return "Orange"
	case colorMagenta:
		return "Magenta"
	case colorLightBlue:
		return "LightBlue"
	case colorYellow:
		return "Yellow"
	case colorLime:
		return "Lime"
	case colorPink:
		return "Pink"
	case colorGray:
		return "Gray"
	case colorLightGray:
		return "LightGray"
	case colorCyan:
		return "Cyan"
	case colorPurple:
		return "Purple"
	case colorBlue:
		return "Blue"
	case colorBrown:
		return "Brown"
	case colorGreen:
		return "Green"
	case colorRed:
		return "Red"
	case colorBlack:
		return "Black"
	case colorBrightWhite:
		return "Custom color"
	default:
		return "Unknown color"
	}
}

func (p *Pattern) String() string {
	return p.ColorString() + " " + p.PatternString()
}
