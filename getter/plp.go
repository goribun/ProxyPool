package getter

import (
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

//PLP get ip from proxylistplus.com
func PLP() (result []string) {
	pollURL := "https://list.proxylistplus.com/Fresh-HTTP-Proxy-List-1"
	_, body, errs := gorequest.New().Get(pollURL).End()
	if errs != nil {
		log.Println(errs)
		return
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Println(err.Error())
		return
	}
	doc.Find("#page > table.bg > tbody > tr").Each(func(i int, s *goquery.Selection) {
		node := strconv.Itoa(i + 1)
		ss, _ := s.Find("tr:nth-child(" + node + ") > td:nth-child(2)").Html()
		sss, _ := s.Find("tr:nth-child(" + node + ") > td:nth-child(3)").Html()

		result = append(result, ss + ":" + sss)
	})
	if len(result) > 0 {
		result = result[2:]
	}
	log.Println("PLP done.")
	return
}
