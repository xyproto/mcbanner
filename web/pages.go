package main

import (
	"fmt"
	"github.com/unrolled/render"
	"github.com/xyproto/mcbanner"
	"github.com/xyproto/onthefly"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	zoomFactor = 8 // 8x zoom when displaying banners
)

// TODO: This works as long as there are not too many users. Fix.
var gB *mcbanner.Banner
var gHow []string

func mainPage(mux *http.ServeMux, path string, r *render.Render) {
	mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		data := map[string]string{
			"title": version_string,
		}

		// TODO: Enable this if a development environment variable is on
		// Reload template (great for development)
		//r = render.New(render.Options{})

		// Render and return
		r.HTML(w, http.StatusOK, "index", data)
	})
}

func comparison(mux *http.ServeMux, path string, r *render.Render) {
	mcbanner.Seed()

	svgurl1 := "/img/a.svg"
	pngurl1 := "/img/a.png"
	pngurl2 := "/img/c1.png"
	pngbytes2, err := ioutil.ReadFile("public" + pngurl2)
	if err != nil {
		log.Fatalln("Could not read: " + "public" + pngurl2)
	}

	b, _ := mcbanner.NewRandomBanner()
	svgxml1 := b.SVG()
	pngbytes1 := b.PNG()
	likeness := mcbanner.Equality(pngbytes1, pngbytes2)

	mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		b, _ = mcbanner.NewRandomBanner()
		svgxml1 = b.SVG()
		pngbytes1 = b.PNG()
		likeness = mcbanner.Equality(pngbytes1, pngbytes2)

		bw := strconv.Itoa(mcbanner.BannerW) + "px"
		bh := strconv.Itoa(mcbanner.BannerH) + "px"

		data := map[string]string{
			"title":        "Comparison",
			"likeness":     fmt.Sprintf("%.0f%%", likeness*100.0),
			"bannerWidth":  bw,
			"bannerHeight": bh,
			"svg":          svgurl1,
			"png1":         pngurl1,
			"png2":         pngurl2,
		}

		// TODO: Enable this only if a development environment variable is on
		// Reload template (great for development)
		r = render.New(render.Options{})

		// Render and return
		r.HTML(w, http.StatusOK, "comparison", data)
	})
	mux.HandleFunc(svgurl1, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "image/svg+xml")
		fmt.Fprint(w, svgxml1)
	})
	//mux.HandleFunc(pngurl1, func(w http.ResponseWriter, req *http.Request) {
	//	w.Header().Add("Content-Type", "image/png")
	//	w.Write(pngbytes1)
	//})
	mux.HandleFunc(pngurl2, func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "image/png")
		w.Write(pngbytes2)
	})

}

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
	w := strconv.Itoa(int(mcbanner.BannerW * zoomFactor * specialPageZoomFactor))
	h := strconv.Itoa(int(mcbanner.BannerH * zoomFactor * specialPageZoomFactor))
	tag.AddStyle("width", w+"px")
	tag.AddStyle("height", h+"px")
	tag.AddStyle("border", "4px solid black")

	return page
}

// Create a pattern gallery under the given path
func patternGallery(mux *http.ServeMux, path string) {
	var (
		svgurls, descs []string
		b              *mcbanner.Banner
	)
	for i := mcbanner.PatternLowerThird; i <= mcbanner.PatternLogo; i++ {

		b = mcbanner.NewBanner()
		p, _ := mcbanner.NewPattern(mcbanner.PatternFull, mcbanner.ColorBrightWhite)
		b.AddPattern(p)
		p, _ = mcbanner.NewPattern(i, mcbanner.ColorRed)
		b.AddPattern(p)

		svgxml := b.SVG()

		// Publish the generated SVG as "/img/banner_NNN.svg"
		svgurl := fmt.Sprintf("/img/banner_%d.svg", i)
		mux.HandleFunc(svgurl, func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "image/svg+xml")
			fmt.Fprint(w, svgxml)
		})

		svgurls = append(svgurls, svgurl)
		descs = append(descs, p.PatternString())
	}

	// Generate a Page that includes the svg images
	page := patternGalleryPage("Pattern gallery", svgurls, descs)
	// Publish the generated Page in a way that connects the HTML and CSS
	page.Publish(mux, path, "/css/banner.css", false)
}

func randomBanner(mux *http.ServeMux, path string, r *render.Render) {
	mux.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
		bw := strconv.Itoa(mcbanner.BannerW*zoomFactor) + "px"
		bh := strconv.Itoa(mcbanner.BannerH*zoomFactor) + "px"

		// Reload template (great for development)
		//r = render.New(render.Options{})

		// The recipe

		mcbanner.Seed()

		var howP []*mcbanner.Pattern
		gB, howP = mcbanner.NewRandomBanner()

		// how is a list of pattern.String() based on howP
		how := []string{}
		for _, p := range howP {
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

	// TODO: One url per generated banner. /img/generated/123/123/result.svg
	mux.HandleFunc("/img/random.svg", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "image/svg+xml")
		b, _ := mcbanner.NewRandomBanner()
		fmt.Fprint(w, b.SVG())
	})

}
