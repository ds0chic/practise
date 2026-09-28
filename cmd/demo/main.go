// 本地看各 server Demo 网页实际输出的统一入口，用法：
//
//	go run ./cmd/demo server	# 跑 ServerDemo，浏览器看 http://localhost:8080/
//	go run ./cmd/demo user		# 跑 UserDemo，浏览器看 http://localhost:8081/user
package main

import (
	"fmt"
	"os"
	"practise"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run ./cmd/demo [server|user]")
		fmt.Println("  server -> practise.ServerDemo() :8080 /")
		fmt.Println("  user   -> practise.UserDemo()   :8081 /user")
		return
	}
	switch os.Args[1] {
	case "server":
		practise.ServerDemo()
	case "user":
		practise.UserDemo()
	default:
		fmt.Println("未知 demo:", os.Args[1])
	}
}
