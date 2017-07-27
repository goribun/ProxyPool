package getter

import (
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/nladuo/go-phantomjs-fetcher"
)

// KDL get ip from kuaidaili.com
func KDL() (result []string) {
	pollURL := "http://www.kuaidaili.com/proxylist/"
	//create a fetcher which seems to a httpClient
	fetcher, err := phantomjs.NewFetcher(2016, nil)

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	defer fetcher.ShutDownPhantomJSServer()
	if err != nil {
		log.Println(err.Error())
		return
	}
	//inject the javascript you want to run in the webpage just like in chrome console.
	jsScript := "function() {s=document.documentElement.outerHTML;document.write('<body></body>');document.body.innerText=s;}"
	//run the injected js_script at the end of loading html
	jsRunAt := phantomjs.RUN_AT_DOC_END
	//send httpGet request with injected js

	for i := 1; i <= 10; i++ {
		resp, err := fetcher.GetWithJS(pollURL+strconv.Itoa(i), jsScript, jsRunAt)
		if err != nil {
			log.Println(err.Error())
			return
		}

		//select search results by goquery
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.Content))
		if err != nil {
			log.Println(err.Error())
			return
		}
		doc.Find("#index_free_list > table > tbody > tr").Each(func(i int, s *goquery.Selection) {
			node := strconv.Itoa(i + 1)
			ff, _ := s.Find("tr:nth-child(" + node + ") > td:nth-child(2)").Html()

			result = append(result, ff)
		})
	}
	log.Println("KDL done.")
	return
}
