package godensity

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

// TestDensity tests the density analysis of various web pages
func TestDensity(t *testing.T) {

	urls := []string{
		"https://news.v.daum.net/v/20200216130927758#none",
		// Additional URLs can be added as needed
	}

	for _, url := range urls {
		// Fetch the webpage
		res, err := http.Get(url)
		if err != nil {
			log.Fatalf("Failed to get URL %s: %v", url, err)
		}
		defer res.Body.Close()

		// Parse the HTML document
		doc, err := goquery.NewDocumentFromResponse(res)
		if err != nil {
			log.Fatalf("Failed to parse HTML from URL %s: %v", url, err)
		}

		// Extract and process the body of the page
		body := doc.Find("body")
		Filtering(body)        // Apply filtering logic
		DiveIntoDOM(body, url) // Analyze DOM for density

		// Sort nodes based on densitySum
		sort.Slice(array, func(i, j int) bool {
			return array[i].densitySum > array[j].densitySum
		})

		// Define candidates and threshold
		var candidates []Node
		var threshold float32

		log.Println("------------")
		log.Printf("Analyzing URL: %s", url)

		// Analyze elements and determine the threshold
		for _, element := range array {
			nodeID, _ := element.goqueryNode.Attr("id")
			log.Printf("Node: %s | Density: %.2f | DensitySum: %.2f | ID: %s", goquery.NodeName(element.goqueryNode), element.density, element.densitySum, nodeID)

			if strings.Contains(goquery.NodeName(element.goqueryNode), "body") {
				threshold = element.density
				break
			}
			candidates = append(candidates, element)
		}

		// If no candidates, select a default fallback
		if len(candidates) == 0 {
			candidates = append(candidates, array[1])
		}

		// Display the top 10 elements based on density
		array = array[:10]
		for _, c := range array {
			id, _ := c.goqueryNode.Attr("id")
			class, _ := c.goqueryNode.Attr("class")
			fmt.Printf("Node: %s | Density: %.2f | DensitySum: %.2f | ID: %s | Class: %s\n", goquery.NodeName(c.goqueryNode), c.density, c.densitySum, id, class)
		}

		// Process candidates with density above the threshold
		for _, element := range candidates {
			cursor := &element
			if element.density < threshold {
				continue
			}
			log.Printf("Selected Node: %v", cursor)
			break
		}

		// Clear array for next iteration
		array = array[:0]
	}
}
