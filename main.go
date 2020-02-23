package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func DiveInto(me *goquery.Selection) *Node {
	if me.Children().Length() == 0 {
		class, err := me.Attr("class")
		if !err {
			class = fmt.Sprintf("id='%s'", me.AttrOr("id", "none"))
		} else {
			class = fmt.Sprintf("class='%s'", class)
		}

		re := regexp.MustCompile(`\s+`)
		characters := re.ReplaceAllString(strings.TrimSpace(me.Text()), " ")
		density := float32(len(characters))

		return &Node{
			tag:        fmt.Sprintf("<%s %s>", goquery.NodeName(me), class),
			density:    density,
			densitySum: 0,
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
		child := DiveInto(c)

		densitySum += child.density
		T += child.T

		if child.density > maximum {
			next = child
			maximum = child.density
		}
	})

	class, err := me.Attr("class")
	if !err {
		class = fmt.Sprintf("id='%s'", me.AttrOr("id", "none"))
	} else {
		class = fmt.Sprintf("class='%s'", class)
	}

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
		tag:        fmt.Sprintf("<%s %s>", goquery.NodeName(me), class),
		density:    density,
		densitySum: densitySum,
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

type Node struct {
	tag        string
	density    float32
	densitySum float32
	images     []string
	T          float32
	text       string
	next       *Node
}

var array []Node

func main() {
	array = make([]Node, 0)

	// url := "https://news.joins.com/article/23707620?cloc=joongang-home-toptype1basic"
	// url := "https://news.v.daum.net/v/20200216130927758#none"
	// url := "https://www.yna.co.kr/view/AKR20200217067251001?section=politics/national-assembly"
	// url := "https://news.naver.com/main/read.nhn?mode=LSD&mid=shm&sid1=104&oid=001&aid=0011406103"
	// url := "https://gall.dcinside.com/board/lists/?id=baseball_new8"
	// url := "https://blog.bobthedeveloper.io/happy-new-year-from-bob-86b018fd134a"
	// url := "https://namu.wiki/w/설리(배우)"
	// url := "http://www.hani.co.kr/arti/politics/bluehouse/928518.html"
	// url := "https://www.dogdrip.net/246267026"
	// url := "https://m.blog.naver.com/PostView.nhn?blogId=forsun55&logNo=220923292175&proxyReferer=https%3A%2F%2Fwww.google.com%2F"
	// url := "https://www.dogdrip.net/index.php?mid=computer&category=180739712&document_srl=246186396&page=1"
	// url := "http://magazine.hankyung.com/business/apps/news?mode=sub_view&nkey=2011022200795000041"
	// url := "https://gall.dcinside.com/board/view/?id=baseball_new8&no=12410667&exception_mode=recommend&page=1"
	// url := "https://gall.dcinside.com/board/view/?id=hit&no=15657&page=1"
	url := "http://www.donga.com/news/article/all/20200221/99817201/2"

	if res, err := http.Get(url); err != nil {
		log.Fatal(err)
	} else {
		if doc, err := goquery.NewDocumentFromResponse(res); err != nil {
			log.Fatal(err)
		} else {
			body := doc.Find("body")
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

			DiveInto(body)
		}
	}

	sort.Slice(array, func(i, j int) bool {
		return array[i].densitySum > array[j].densitySum
	})

	var candidates []Node
	var threashold float32

	fmt.Println("------------")

	for _, element := range array {
		fmt.Println(element.tag, element.density, element.densitySum)
		if strings.HasPrefix(element.tag, "<body") {
			threashold = element.density
			break
		}

		candidates = append(candidates, element)
	}

	if len(candidates) == 0 {
		candidates = append(candidates, array[1])
	}

	// for _, c := range array {
	// fmt.Println(c.tag, c.density, c.densitySum)
	// }

	// visited := make(map[*Node]bool)

	for _, element := range candidates {
		var cursor *Node = &element
		// var slower *Node = cursor
		if element.density < threashold {
			continue
		}
		fmt.Println(cursor)
		break
		// for cursor.next != nil && cursor.density >= threashold {
		// 	slower = cursor
		// 	cursor = cursor.next
		// }

		// fmt.Println(slower.text)
		// fmt.Println("-------------")
	}
}
