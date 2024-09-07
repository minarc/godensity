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

// Placeholder functions for Density-related computations
func ComputeTextDensity() float32 {
	return 1
}

func ComputeDensity() float32 {
	return 1
}

func CalculateDensitySum() float32 {
	return 1
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
		log.Panicf("Error reading response body: %v", err)
	}

	// Check GIF magic number in first 4 bytes (GIF87a or GIF89a)
	if bytes.HasPrefix(body, []byte{71, 73, 70, 56}) {
		log.Printf("GIF detected! URL: %s", target)
		return true
	}
	return false
}

// DiveIntoDOM traverses the DOM structure recursively and calculates density metrics
func DiveIntoDOM(selection *goquery.Selection, domain string) *Node {
	// Base case: no children, calculate density for text or images
	if selection.Children().Length() == 0 {
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
			images: selection.Filter("img").Map(func(_ int, s *goquery.Selection) string {
				src, _ := s.Attr("src")
				return src
			}),
			videos: selection.Filter("video > source").Map(func(_ int, s *goquery.Selection) string {
				src, _ := s.Attr("src")
				return src
			}),
			T:    0,
			text: textContent,
			next: nil,
		}
	}

	// Recursive case: process children and calculate density metrics
	var (
		densitySum float32
		maxDensity float32 = -1
		nextNode   *Node
		T          float32
	)

	selection.Children().Each(func(_ int, child *goquery.Selection) {
		childNode := DiveIntoDOM(child, domain)
		densitySum += childNode.density
		T += childNode.T

		// Track child with maximum density
		if childNode.density > maxDensity {
			nextNode = childNode
			maxDensity = childNode.density
		}
	})

	T += float32(selection.Children().Length())

	// Calculate various metrics for text and links
	re := regexp.MustCompile(`\s+`)
	textLength := float32(len(re.ReplaceAllString(strings.TrimSpace(selection.Text()), " ")))
	linkTextLength := float32(len(re.ReplaceAllString(strings.TrimSpace(selection.Find("a").Text()), " ")))
	linkCount := float32(selection.Find("a").Length())

	if T-linkCount <= 0 {
		T = 1
		linkCount = 0
	}

	// Calculate density: (text - linkText) / (totalText - totalLinkText)
	density := (textLength - linkTextLength) / (T - linkCount)

	// Create Node object for the current element
	currentNode := &Node{
		goqueryNode: selection,
		density:     density,
		densitySum:  densitySum,
		videos: selection.Find("video > source").Map(func(_ int, s *goquery.Selection) string {
			src, _ := s.Attr("src")
			return src
		}),
		T:    T,
		text: re.ReplaceAllString(strings.TrimSpace(selection.Text()), " "),
		next: nextNode,
	}

	// Append to the array for later use
	array = append(array, *currentNode)

	return currentNode
}

// Global array to store Node structures
var array []Node
