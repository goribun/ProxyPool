package getter

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

// YDL get ip from youdaili.net
func YDL() (result []string) {
	pollURL := "http://www.youdaili.net/Daili/http/"
	_, body, errs := gorequest.New().Get(pollURL).End()
	if errs != nil {
		log.Println(errs)
		return
	}
	do, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Println(err.Error())
		return
	}

	URL, _ := do.Find("body > div.con.PT20 > div.conl > div.lbtc.l > div.chunlist > ul > li:nth-child(1) > p > a").Attr("href")
	_, content, errs := gorequest.New().Get(URL).End()
	if errs != nil {
		log.Println(errs)
		return
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Println(err.Error())
		return
	}
	doc.Find(".content p").Each(func(_ int, s *goquery.Selection) {
		c := strings.Split(s.Text(), "@")

		result = append(result, c[0])
	})
	log.Println("YDL done.")
	return
}
