package main

import (
	"fmt"

	"github.com/vicanso/umbra/umbra"
)

func main() {
	// umbra.CheckDNSServers("store.gf.com.cn", umbra.GetDefaultDNSServers())

	dns := umbra.DNS{
		Server:   "223.5.5.5",
		Hostname: "aslant.site",
	}
	fmt.Println(dns.GetCheckResult("my test"))

	// http := umbra.HTTP{
	// 	URL:         "https://store.gf.com.cn/user/session",
	// 	EnableTrace: true,
	// }
	// http.Check()
	// fmt.Println(http.GetDescription())

	// ping := umbra.Ping{
	// 	IP: "10.2.110.211",
	// }
	// ping.Check()
	// fmt.Println(ping.GetDescription())

	// tcp := umbra.TCP{
	// 	IP:   "10.2.110.211",
	// 	Port: 80,
	// }
	// tcp.Check()
	// fmt.Println(tcp.GetDescription())

	// checkerList := make([]umbra.Checker, 0)
	// baiduIP := "14.215.177.38"

	// dns := &umbra.DNS{
	// 	Server:   "114.114.114.114:53",
	// 	Hostname: "www.baidu.com",
	// }
	// checkerList = append(checkerList, dns)

	// http := &umbra.HTTP{
	// 	URL:         "https://www.baidu.com/",
	// 	IP:          baiduIP,
	// 	EnableTrace: true,
	// }
	// checkerList = append(checkerList, http)

	// ping := &umbra.Ping{
	// 	IP:   baiduIP,
	// 	Type: "ip4",
	// }
	// checkerList = append(checkerList, ping)

	// tcp := &umbra.TCP{
	// 	IP:   baiduIP,
	// 	Port: 443,
	// }
	// checkerList = append(checkerList, tcp)

	// for _, checker := range checkerList {
	// 	fmt.Println(checker.GetDescription())
	// 	healthy, extra, err := checker.Check()
	// 	fmt.Println("healthy:", healthy, ", extra:", extra, ", err:", err)
	// }
}
