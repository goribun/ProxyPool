package getter

import (
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/nladuo/go-phantomjs-fetcher"
)

// XDL get ip from xdaili.cn
func XDL() (result []string) {
	pollURL := "http://www.xdaili.cn/freeproxy.html"

	fetcher, err := phantomjs.NewFetcher(2015, nil)

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
	jsScript := "function() {s=document.documentElement.outerHTML;document.write('<body></body>');document.body.innerText=s;}"
	jsRunAt := phantomjs.RUN_AT_DOC_END
	resp, err := fetcher.GetWithJS(pollURL, jsScript, jsRunAt)
	if err != nil {
		log.Println(err.Error())
		return
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(strings.Replace(strings.Replace(resp.Content, "&lt;", "<", -1), "&gt;", ">", -1)))
	if err != nil {
		log.Println(err.Error())
		return
	}
	doc.Find("#target > tr").Each(func(i int, s *goquery.Selection) {
		node := strconv.Itoa(i + 1)
		ss, _ := s.Find("tr:nth-child(" + node + ") > td:nth-child(1)").Html()
		sss, _ := s.Find("tr:nth-child(" + node + ") > td:nth-child(2)").Html()
		result = append(result, ss+":"+sss)
	})
	log.Println("XDL done.")
	return
}
