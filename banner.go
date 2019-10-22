package mcbanner

import (
	"github.com/xyproto/tinysvg"
	"log"
)

const (
	BannerW = 20
	BannerH = 40

	fullW  = BannerW
	fullH  = BannerH
	halfW  = BannerW / 2
	halfH  = BannerH / 2
	thirdW = BannerW / 3
	thirdH = BannerH / 3
	fifthW = BannerW / 5
	fifthH = BannerH / 5
	sixthW = BannerW / 6
	sixthH = BannerH / 6
	tenthW = BannerW / 10
	tenthH = BannerH / 10

	maxX = fullW - 1
	maxY = fullH - 1
)

type Banner struct {
	Patterns []*Pattern
	curpat   int
}

func NewBanner() *Banner {
	return &Banner{}
}

func (b *Banner) AddPattern(p *Pattern) {
	if len(b.Patterns) > maxPatterns {
		log.Fatalln("Too many patters for banner, max", maxPatterns)
	}
	b.Patterns = append(b.Patterns, p)
}

func (b *Banner) Draw(svg *tinysvg.Tag) {
	// draw each Pattern
	for _, p := range b.Patterns {
		color, ok := colors[p.color]
		if !ok {
			log.Fatalln("Invalid color ID: ", p.color)
		}
		DrawPattern(svg, p.pattern, color)
	}
}

// Generate a new SVG image for the banner
func (b *Banner) Image() *tinysvg.Document {
	if b == nil {
		log.Fatalln("Can't generate SVG for a *Banner that is nil!")
	}
	document, svg := tinysvg.NewTinySVG(BannerW, BannerH)
	svg.Describe("A banner")
	b.Draw(svg)
	return document
}

func (b *Banner) SVG() []byte {
	return b.Image().Bytes()
}

func (b *Banner) PNG() []byte {
	return Convert(b.SVG(), "svg", "png")
}

func NewRandomBanner() (b *Banner, how []*Pattern) {
	// Generate new banner
	b = NewBanner()
	how = []*Pattern{}
	p, _ := NewPattern(PatternFull, randomColor())
	b.AddPattern(p)
	how = append(how, p)

	// Up to 6 different Patterns
	for i := 0; i < maxPatterns; i++ {
		p, _ = NewPattern(randomPattern(), randomColor())
		b.AddPattern(p)
		how = append(how, p)
	}

	return b, how
}
