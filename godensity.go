package godensity

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ComputeTextDensity calculates the text density of a node
func ComputeTextDensity(node *Node) float32 {
	if node.T == 0 {
		return 0
	}
	return float32(len(node.text)) / node.T
}

// ComputeDensity calculates the overall density of a node
func ComputeDensity(node *Node) float32 {
	textDensity := ComputeTextDensity(node)
	imageDensity := float32(len(node.images)) * 1000 // Assuming each image contributes 1000 to density
	videoDensity := float32(len(node.videos)) * 2000 // Assuming each video contributes 2000 to density
	return textDensity + imageDensity + videoDensity
}

// CalculateDensitySum computes the sum of densities for a node and its children
func CalculateDensitySum(node *Node) float32 {
	if node == nil {
		return 0
	}
	return node.density + node.densitySum
}

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

// IsGIF checks if the provided URL points to a GIF image by examining its content
func IsGIF(src *url.URL, currentURL *url.URL) bool {
	target := src.String()

	if !src.IsAbs() {
		target = currentURL.ResolveReference(src).String()
	}

	res, err := http.Get(target)
	if err != nil {
		log.Printf("Error fetching URL: %v, Source: %v, Target: %v", err, src, target)
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return false
	}

	// Check GIF magic number in first 4 bytes (GIF87a or GIF89a)
	if bytes.HasPrefix(body, []byte{71, 73, 70, 56}) {
		log.Printf("GIF detected! URL: %s", target)
		return true
	}
	return false
}

// diveIntoDOM traverses the DOM structure recursively and calculates density metrics
func diveIntoDOM(selection *goquery.Selection, domain string) *Node {
	// Base case: no children, calculate density for text or images
	if selection.Children().Length() == 0 {
		return createLeafNode(selection, domain)
	}

	// Recursive case: process children and calculate density metrics
	return createInternalNode(selection, domain)
}

// createLeafNode creates a Node for elements without children
func createLeafNode(selection *goquery.Selection, domain string) *Node {
	re := regexp.MustCompile(`\s+`)
	textContent := re.ReplaceAllString(strings.TrimSpace(selection.Text()), " ")
	var density float32

	// Check if it's an image tag
	if goquery.NodeName(selection) == "img" {
		if src, exists := selection.Attr("src"); exists {
			parsedSrc, _ := url.Parse(src)
			parsedDomain, _ := url.Parse(domain)
			if IsGIF(parsedSrc, parsedDomain) {
				density = 1
			}
		} else {
			log.Println("Image tag found with no src attribute")
		}
	}

	return &Node{
		goqueryNode: selection,
		density:     density,
		densitySum:  0,
		images:      getMediaSources(selection, "img"),
		videos:      getMediaSources(selection, "video > source"),
		T:           0,
		text:        textContent,
		next:        nil,
	}
}

// createInternalNode creates a Node for elements with children
func createInternalNode(selection *goquery.Selection, domain string) *Node {
	var (
		densitySum float32
		maxDensity float32 = -1
		nextNode   *Node
		T          float32
	)

	selection.Children().Each(func(_ int, child *goquery.Selection) {
		childNode := diveIntoDOM(child, domain)
		densitySum += childNode.density
		T += childNode.T

		if childNode.density > maxDensity {
			nextNode = childNode
			maxDensity = childNode.density
		}
	})

	T += float32(selection.Children().Length())

	re := regexp.MustCompile(`\s+`)
	textLength := float32(len(re.ReplaceAllString(strings.TrimSpace(selection.Text()), " ")))
	linkTextLength := float32(len(re.ReplaceAllString(strings.TrimSpace(selection.Find("a").Text()), " ")))
	linkCount := float32(selection.Find("a").Length())

	if T-linkCount <= 0 {
		T = 1
		linkCount = 0
	}

	density := (textLength - linkTextLength) / (T - linkCount)

	currentNode := &Node{
		goqueryNode: selection,
		density:     density,
		densitySum:  densitySum,
		videos:      getMediaSources(selection, "video > source"),
		T:           T,
		text:        re.ReplaceAllString(strings.TrimSpace(selection.Text()), " "),
		next:        nextNode,
	}

	// Append to the array for later use
	array = append(array, *currentNode)

	return currentNode
}

// getMediaSources extracts media sources from the given selection
func getMediaSources(selection *goquery.Selection, selector string) []string {
	return selection.Find(selector).Map(func(_ int, s *goquery.Selection) string {
		src, _ := s.Attr("src")
		return src
	})
}

// Global array to store Node structures
var array []Node
