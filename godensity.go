package godensity

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CTD() float32 {
	return 1
}

func TD() float32 {
	return 1
}

func Density() {

}

func DensitySum() {

}

func DiveIntoDOM(me *goquery.Selection) *Node {
	if me.Children().Length() == 0 {
		re := regexp.MustCompile(`\s+`)
		characters := re.ReplaceAllString(strings.TrimSpace(me.Text()), " ")
		density := float32(len(characters))

		return &Node{
			goqueryNode: me,
			density:     density,
			densitySum:  0,
			images: me.Filter("img").Map(func(_ int, s *goquery.Selection) string {
				src, _ := s.Attr("src")
				return src
			}),
			T:    0,
			text: characters,
			next: nil,
		}
	}

	var densitySum float32
	var maximum float32 = -1
	var next *Node
	var T float32

	me.Children().Each(func(_ int, c *goquery.Selection) {
		child := DiveIntoDOM(c)

		densitySum += child.density
		T += child.T

		if child.density > maximum {
			next = child
			maximum = child.density
		}
	})

	T += float32(me.Children().Length())

	// fmt.Println(len(me.Children().Text()))
	re := regexp.MustCompile(`\s+`)

	var C float32 = float32(len(re.ReplaceAllString(strings.TrimSpace(me.Text()), " ")))
	var LT float32 = float32(me.Find("a").Length())

	var LC float32 = float32(len(re.ReplaceAllString(strings.TrimSpace(me.Find("a").Text()), " ")))
	// var nLC float32 = C - LC
	// var nLC float32 = C - LC
	if T-LT <= 0 {
		T = 1
		LT = 0
	}

	var I float32 = float32(me.Find("img").Length())
	var V float32 = float32(me.Find("video").Length())
	if I == 0 {
		I = 1
	}
	if V == 0 {
		V = 1
	}
	// fmt.Println(class, (C-LC)/(T-LT), I, C, LC, T, LT, re.ReplaceAllString(strings.TrimSpace(me.Text()), " "))
	// fmt.Println(class, (C-LC)/(T-LT), I, C, LC, T, LT)
	var density float32 = (C - LC) / (T - LT)
	// densitySum *= I + V

	// density := (C / T) * float32(math.Log10(float64((C/LC)*(T/LT)))/math.Log10(math.Log(float64(((C/nLC)*LC)+LC+math.E))))

	itsme := &Node{
		goqueryNode: me,
		density:     density,
		densitySum:  densitySum,
		images: me.Find("img").Map(func(_ int, s *goquery.Selection) string {
			src, _ := s.Attr("src")
			return src
		}),
		T:    T,
		text: re.ReplaceAllString(strings.TrimSpace(me.Text()), " "),
		next: next,
	}

	array = append(array, *itsme)

	return itsme
}

var array []Node
