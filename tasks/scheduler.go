package tasks

import (
	"time"
	"log"
	"sync"

	"github.com/robfig/cron"
	"github.com/goribun/ProxyPool/storage"
	"github.com/goribun/ProxyPool/util"
	"github.com/goribun/ProxyPool/getter"
)

var IpChan = make(chan string, 500)

func Schedule() {

	c := cron.New()

	//抓取代理IP并放入channel
	c.AddFunc(util.Cfg.Cron.GetProxyCron, func() {
		log.Printf("%v crons running at:%v", "Get proxy", time.Now())
		if len(IpChan) < 100 {
			go run(IpChan)
		}
	})

	// 检查Redis内的IP可用性
	c.AddFunc(util.Cfg.Cron.CheckRedisCron, func() {
		log.Printf("%v crons running at:%v", "Check redis", time.Now())
		go func() {
			storage.CheckProxyInRedis()
		}()
	})

	c.Start()
}

func run(ipChan chan<- string) {
	var wg sync.WaitGroup
	functions := []func() []string{
		getter.Data5u,
		getter.IP66,
		getter.KDL,
		getter.GBJ,
		getter.Xici,
		getter.XDL,
		getter.IP181,
		getter.PLP,
	}
	for _, f := range functions {
		wg.Add(1)
		go func(f func() []string) {
			temp := f()
			for _, v := range temp {
				ipChan <- v
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	log.Println("All getters finished.")
}
