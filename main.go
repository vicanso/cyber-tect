package main

import (
	"fmt"

	"github.com/vicanso/umbra/umbra"
)

func main() {

	checkerList := make([]umbra.Checker, 0)
	baiduIP := "14.215.177.38"

	dns := &umbra.DNS{
		Server:   "114.114.114.114:53",
		Hostname: "www.baidu.com",
	}
	checkerList = append(checkerList, dns)

	http := &umbra.HTTP{
		URL:         "https://www.baidu.com/",
		IP:          baiduIP,
		EnableTrace: true,
	}
	checkerList = append(checkerList, http)

	ping := &umbra.Ping{
		IP:   baiduIP,
		Type: "ip4",
	}
	checkerList = append(checkerList, ping)

	tcp := &umbra.TCP{
		IP:   baiduIP,
		Port: 443,
	}
	checkerList = append(checkerList, tcp)

	for _, checker := range checkerList {
		fmt.Println(checker.GetDescription())
		healthy, extra, err := checker.Check()
		fmt.Println("healthy:", healthy, ", extra:", extra, ", err:", err)
	}
}
