package main

import (
	"fmt"

	"github.com/vicanso/umbra/umbra"
)

func main() {
	h := umbra.HTTP{
		URL: "https://www.baidu.com/",
		IP:  "14.215.177.39",
	}
	fmt.Println(h.Check())
}
