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

func TestDensity(t *testing.T) {

	urls := []string{
		// "https://gifsf.com/star/4264545",
		"https://www.heye.kr/board/index.html?id=idol&no=57606",

		// "https://news.joins.com/article/23707620?cloc=joongang-home-toptype1basic",
		// "https://news.v.daum.net/v/20200216130927758#none",
		// "https://namu.wiki/w/설리(배우)",
		// "https://www.yna.co.kr/view/AKR20200217067251001?section=politics/national-assembly",
		// "http://www.hani.co.kr/arti/politics/bluehouse/928518.html",
		// "https://gall.dcinside.com/board/lists/?id=baseball_new8",
		// "https://blog.bobthedeveloper.io/happy-new-year-from-bob-86b018fd134a",
		// "http://www.hani.co.kr/arti/politics/bluehouse/928518.html",
		// "https://m.blog.naver.com/PostView.nhn?blogId=forsun55&logNo=220923292175&proxyReferer=https%3A%2F%2Fwww.google.com%2F",
		// "https://www.dogdrip.net/index.php?mid=computer&category=180739712&document_srl=246186396&page=1",
		// "http://www.donga.com/news/article/all/20200221/99817201/2",
		// "https://gall.dcinside.com/board/view/?id=hit&no=15657&page=1",
		// "http://magazine.hankyung.com/business/apps/news?mode=sub_view&nkey=2011022200795000041",
	}

	for _, url := range urls {

		if res, err := http.Get(url); err != nil {
			log.Fatal(err)
		} else {
			if doc, err := goquery.NewDocumentFromResponse(res); err != nil {
				log.Fatal(err)
			} else {
				body := doc.Find("body")
				Filtering(body)
				DiveIntoDOM(body, url)
			}
		}

		sort.Slice(array, func(i, j int) bool {
			return array[i].densitySum > array[j].densitySum
		})

		var candidates []Node
		var threashold float32

		log.Println("------------")
		log.Println(url)

		for _, element := range array {

			class, _ := element.goqueryNode.Attr("id")
			log.Println(goquery.NodeName(element.goqueryNode), element.density, element.densitySum, class)

			if strings.Contains(goquery.NodeName(element.goqueryNode), "body") {
				threashold = element.density
				break
			}

			candidates = append(candidates, element)
		}

		if len(candidates) == 0 {
			candidates = append(candidates, array[1])
		}

		array = array[:10]
		for _, c := range array {
			id, _ := c.goqueryNode.Attr("id")
			class, _ := c.goqueryNode.Attr("class")
			fmt.Println(goquery.NodeName(c.goqueryNode), c.density, c.densitySum, id, class)
		}

		// visited := make(map[*Node]bool)

		// log.Println(threashold)
		// log.Println(candidates)

		for _, element := range candidates {
			var cursor *Node = &element
			// var slower *Node = cursor
			if element.density < threashold {
				continue
			}
			log.Println(cursor)
			break
			// for cursor.next != nil && cursor.density >= threashold {
			// 	slower = cursor
			// 	cursor = cursor.next
			// }

			// fmt.Println(slower.text)
			// fmt.Println("-------------")
		}
		array = array[:0]
	}

}
