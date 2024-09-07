package godensity

import "github.com/PuerkitoBio/goquery"

// Filtering removes unnecessary HTML elements from the page's body to simplify the DOM analysis.
func Filtering(body *goquery.Selection) {
	// Remove common interactive and irrelevant elements from the DOM
	toRemove := []string{
		"script", "style", "br", "input", "button", "textarea", "label", "form", "iframe",
		"header", "nav", ".comment", "footer", ".footer", "#footer",
	}

	// Iterate over the list and remove each element
	for _, selector := range toRemove {
		body.Find(selector).Remove()
	}
}
