package main

import (
	"runtime"

	"github.com/goribun/ProxyPool/api"
	"github.com/goribun/ProxyPool/storage"
	"github.com/goribun/ProxyPool/tasks"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	//启动抓取和检查任务
	tasks.Schedule()

	// 检查channel内的IP，可用则放入Redis
	for i := 0; i < 50; i++ {
		go func() {
			for {
				storage.CheckAndAddProxy(<-tasks.IpChan)
			}
		}()
	}

	//启动HTTP服务
	api.Serve()
}
