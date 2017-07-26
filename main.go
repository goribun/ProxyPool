package main

import (
	//"fmt"
	"runtime"
	"log"

	"github.com/goribun/ProxyPool/storage"

	"github.com/goribun/ProxyPool/api"
	"time"
	"sync"
	"github.com/goribun/ProxyPool/getter"
)

func main() {

	conn := storage.NewRedisStorage()
	//fmt.Println(conn)
	//conn.Add("121.68.1.108")
	//
	//bb, _ := conn.GetAll()
	//for xxx, yyy := range bb {
	//	fmt.Println(xxx, yyy, "???")
	//}

	runtime.GOMAXPROCS(runtime.NumCPU())
	ipChan := make(chan string, 2000)
	//	conn := storage.NewStorage()
	//
	//	// 启动HTTP服务
	go func() {
		api.Serve()
	}()

	// 检查Redis内的IP可用性
	go func() {
		storage.CheckProxyInRedis()
	}()

	// 检查channel内的IP，可用则放入Redis
	for i := 0; i < 50; i++ {
		go func() {
			for {
				storage.CheckAndAddProxy(<-ipChan)
			}
		}()
	}

	// Start getters to scraper IP and put it in channel
	for {
		x := conn.Count()
		log.Printf("Chan: %v, IP: %v\n", len(ipChan), x)
		if len(ipChan) < 100 {
			go run(ipChan)
		}
		time.Sleep(100 * time.Minute)
	}
}

func run(ipChan chan<- string) {
	var wg sync.WaitGroup
	funs := []func() []string{
		//getter.Data5u,
		getter.IP66,
		//getter.KDL,
		//getter.GBJ,
		//getter.Xici,
		//getter.XDL,
		//getter.IP181,
		//getter.PLP,
	}
	for _, f := range funs {
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
