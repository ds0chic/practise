// 本地看各 server Demo 网页实际输出的统一入口，用法：
//
//	go run ./cmd/demo server	# 跑 ServerDemo，浏览器看 http://localhost:8080/
//	go run ./cmd/demo info		# 跑 RequestInfoDemo，浏览器看 http://localhost:8081/info
package main

import (
	"fmt"
	"os"
	"practise"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run ./cmd/demo [server|info]")
		fmt.Println("  server -> practise.ServerDemo()      :8080 /")
		fmt.Println("  info   -> practise.RequestInfoDemo() :8081 /info")
		return
	}
	switch os.Args[1] {
	case "server":
		practise.ServerDemo()
	case "info":
		practise.RequestInfoDemo()
	default:
		fmt.Println("未知 demo:", os.Args[1])
	}
}
