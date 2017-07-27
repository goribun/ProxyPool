package storage

import (
	"log"
	"sync"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/goribun/ProxyPool/util"
)

//检查IP，可用则放入Redis
func CheckAndAddProxy(ip string) {
	if CheckIP(ip) {
		log.Println("Check And Add Proxy IP:(ip)")
		ProxyAdd(ip)
	}
}

// 检查Redis内代理IP是否可用
func CheckProxyInRedis() {

	log.Println("========================Check Proxy InRedis Start===============")

	conn := NewRedisStorage()

	ips, err := conn.GetAll()
	if err != nil {
		log.Println(err.Error())
		return
	}
	var wg sync.WaitGroup
	for _, v := range ips {
		wg.Add(1)

		go func(v string) {
			if !CheckIP(v) {
				ProxyDel(v)
			}
			wg.Done()
		}(v)
	}
	wg.Wait()
	log.Println("========================Check Proxy InRedis Finish===============")
}

// 检查IP是否可用
func CheckIP(ip string) bool {

	start := time.Now()

	pollURL := "http://httpbin.org/get"
	resp, _, errs := gorequest.New().Proxy("http://" + ip).Get(pollURL).Timeout(util.Cfg.Timeout * time.Second).End()

	end := time.Now()
	log.Printf("Check IP:(%v)  Time:(%v)", ip, end.Sub(start))

	if errs != nil {
		return false
	}
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

//随机取出一个IP
func ProxyGet() string {
	conn := NewRedisStorage()
	ip, err := conn.GetOne()
	if err != nil {
		log.Println(err.Error())
	}
	return ip
}

//增加一个IP
func ProxyAdd(ip string) {
	conn := NewRedisStorage()
	err := conn.Add(ip)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Add valid ip:", ip)

}

//删除一个IP
func ProxyDel(ip string) {
	conn := NewRedisStorage()
	if err := conn.Delete(ip); err != nil {
		log.Println(err.Error())
	}
	log.Println("Del invalid ip:", ip)
}
