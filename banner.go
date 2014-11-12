package mcbanner

import (
	"github.com/xyproto/onthefly"
	"log"
)

const (
	bannerW = 20
	bannerH = 40

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
)

type Banner struct {
	patterns []*Pattern
	curpat   int
}

func NewBanner() *Banner {
	return &Banner{}
}

func (b *Banner) AddPattern(p *Pattern) {
	if len(b.patterns) > maxPatterns {
		log.Fatalln("Too many patters for banner, max", maxPatterns)
	}
	b.patterns = append(b.patterns, p)
}

func (b *Banner) Draw(svg *onthefly.Tag) {
	// draw each pattern
	for _, p := range b.patterns {
		color, ok := colors[p.color]
		if !ok {
			log.Fatalln("Invalid color ID: ", p.color)
		}
		DrawPattern(svg, p.pattern, color)
	}
}

// Generate a new SVG Page for a banner
func (b *Banner) SVGpage() *onthefly.Page {
	if b == nil {
		log.Fatalln("Can't generate SVG for a *Banner that is nil!")
	}
	page, svg := onthefly.NewTinySVG(0, 0, bannerW, bannerH)
	desc := svg.AddNewTag("desc")
	desc.AddContent("A banner")
	b.Draw(svg)
	return page
}

func NewRandomBanner() (b *Banner, how []*Pattern) {
	// Generate new banner
	b = NewBanner()
	how = []*Pattern{}
	p, _ := NewPattern(patternFull, randomColor())
	b.AddPattern(p)
	how = append(how, p)

	// Up to 6 different patterns
	for i := 0; i < maxPatterns; i++ {
		p, _ = NewPattern(randomPattern(), randomColor())
		b.AddPattern(p)
		how = append(how, p)
	}

	return b, how
}
