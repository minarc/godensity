package godensity

import "github.com/PuerkitoBio/goquery"

func Filtering(body *goquery.Selection) {
	body.Find("script").Remove()
	body.Find("style").Remove()
	body.Find("br").Remove()
	body.Find("input").Remove()
	body.Find("button").Remove()
	body.Find("textarea").Remove()
	body.Find("label").Remove()
	body.Find("form").Remove()
	body.Find("iframe").Remove()

	body.Find("header").Remove()
	body.Find("nav").Remove()
	body.Find("iframe").Remove()

	body.Find(".comment").Remove()

	body.Find("footer").Remove()
	body.Find(".footer").Remove()
	body.Find("#footer").Remove()
}
