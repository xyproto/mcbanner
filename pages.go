package main

import (
	"fmt"
	"github.com/unrolled/render"
	"github.com/xyproto/onthefly"
	"net/http"
	"strconv"
)

const (
	zoomFactor = 8 // 8x zoom when displaying banners
)

// TODO: This works as long as there are not too many users. Fix.
var b *Banner

// Generate a new onthefly Page (HTML5 and CSS combined)
func patternGalleryPage(title string, svgurls, captions []string) *onthefly.Page {
	var specialPageZoomFactor = 0.7

	// Create a new HTML5 page, with CSS included
	page := onthefly.NewHTML5Page(title)
	page.AddContent(title)

	// Change the margin (em is default)
	page.SetMargin(3)

	// Change the font family
	page.SetFontFamily("sans-serif")

	// Change the color scheme
	page.SetColor("black", "#a0a0a0")

	// Include the generated SVG image on the page
	body, _ := page.GetTag("body")

	// CSS attributes for the body tag
	body.AddStyle("font-size", "2em")

	// Div
	div := body.AddNewTag("div")
	div.AddStyle("font-size", "0.5em")

	// CSS style
	div.AddStyle("margin-top", "2em")

	// Add images
	var (
		tag, figure  *onthefly.Tag
		useObjectTag = false
	)
	for i, svgurl := range svgurls {
		figure = div.AddNewTag("figure")
		if useObjectTag {
			// object tag
			tag = figure.AddNewTag("object")
			// HTML attributes
			tag.AddAttrib("data", svgurl)
			tag.AddAttrib("type", "image/svg+xml")
		} else {
			// img tag
			tag = figure.AddNewTag("img")
			// HTML attributes
			tag.AddAttrib("src", svgurl)
			tag.AddAttrib("alt", "Banner")
		}
		cap := figure.AddNewTag("figcaption")
		cap.AddContent(captions[i])
	}

	page.AddStyle(`
	  body { margin: 2em auto; width: 60%; padding: 1em; height: 100%; }
	  figure { float: left; display; block; width: 12%; }
	  figure img { display: block: width: 90%; margin: 0 auto; padding: 0; }
	  figure figcaption { display: block; margin 0.5em auto; width: 90%; padding: 0.3em; font: italic small Arial, sans-serif; text-align: left; }}
`)

	// CSS style
	w := strconv.Itoa(int(bannerW * zoomFactor * specialPageZoomFactor))
	h := strconv.Itoa(int(bannerH * zoomFactor * specialPageZoomFactor))
	tag.AddStyle("width", w+"px")
	tag.AddStyle("height", h+"px")
	tag.AddStyle("border", "4px solid black")

	return page
}

// Create a pattern gallery under the given path
func patternGallery(mux *http.ServeMux, path string) {
	var (
		svgurls, descs []string
		b              *Banner
	)
	for i := patternLowerThird; i <= patternLogo; i++ {

		b = NewBanner()
		p, _ := NewPattern(patternFull, colorBrightWhite)
		b.AddPattern(p)
		p, _ = NewPattern(i, colorRed)
		b.AddPattern(p)

		svgString := b.SVGpage().String()

		// Publish the generated SVG as "/img/banner_NNN.svg"
		svgurl := fmt.Sprintf("/img/banner_%d.svg", i)
		mux.HandleFunc(svgurl, func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "image/svg+xml")
			fmt.Fprint(w, svgString)
		})

		svgurls = append(svgurls, svgurl)
		descs = append(descs, p.PatternString())
	}

	// Generate a Page that includes the svg images
	page := patternGalleryPage("Pattern gallery", svgurls, descs)
	// Publish the generated Page in a way that connects the HTML and CSS
	page.Publish(mux, path, "/css/banner.css", false)
}

func mainPage(mux *http.ServeMux, path string, r *render.Render) {
	mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		data := map[string]string{
			"title": version_string,
		}

		// Reload template
		r = render.New(render.Options{})

		// Render and return
		r.HTML(w, http.StatusOK, "index", data)
	})
}

func randomBanner(mux *http.ServeMux, path string, r *render.Render) {
	seed()

	mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		bw := strconv.Itoa(bannerW*zoomFactor) + "px"
		bh := strconv.Itoa(bannerH*zoomFactor) + "px"

		// Reload template
		r = render.New(render.Options{})

		// The recipe
		how := []string{}

		// Generate new banner
		b = NewBanner()
		p, _ := NewPattern(patternFull, randomColor())
		b.AddPattern(p)
		how = append(how, p.String())

		// Up to 6 different patterns
		for i := 0; i < maxPatterns; i++ {
			p, _ = NewPattern(randomPattern(), randomColor())
			b.AddPattern(p)
			how = append(how, p.String())
		}

		data := map[string]interface{}{
			"title":        "Random banner",
			"bannerWidth":  bw,
			"bannerHeight": bh,
			"how":          how,
			"howTitle":     "How to make it yourself",
		}

		// Render and return
		r.HTML(w, http.StatusOK, "random", data)
	})

	mux.HandleFunc("/img/random.svg", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "image/svg+xml")
		fmt.Fprint(w, b.SVGpage().String())
	})

}
