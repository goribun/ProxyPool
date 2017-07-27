package getter

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

// IP181 get ip from ip181.com
func IP181() (result []string) {
	pollURL := "http://www.ip181.com/"
	resp, _, errs := gorequest.New().Get(pollURL).End()
	if errs != nil {
		log.Println(errs)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}
	doc.Find("tr.warning").Each(func(i int, s *goquery.Selection) {
		ss := s.Find("td:nth-child(1)").Text()
		sss := s.Find("td:nth-child(2)").Text()

		result = append(result, ss+":"+sss)
	})

	log.Println("IP181 done.")
	return
}
