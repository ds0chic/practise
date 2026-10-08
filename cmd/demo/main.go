// 本地看各 Demo 实际输出的统一入口，server 类用浏览器看，client 类直接看控制台：
//
//	go run ./cmd/demo server	# 跑 ServerDemo，浏览器看 http://localhost:8080/
//	go run ./cmd/demo user		# 跑 UserDemo，浏览器看 http://localhost:8081/user
//	go run ./cmd/demo userstatus	# 跑 UserStatusDemo，浏览器看 http://localhost:8082/user
//	go run ./cmd/demo usermethod	# 跑 UserMethodDemo，浏览器看 http://localhost:8083/user
//	go run ./cmd/demo userauth	# 跑 UserAuthDemo，浏览器看 http://localhost:8084/user
//	go run ./cmd/demo userbody	# 跑 UserBodyDemo，curl POST 看 http://localhost:8085/user
//	go run ./cmd/demo userjson	# 跑 UserJSONDemo，curl POST JSON 看 http://localhost:8086/user
//	go run ./cmd/demo clientget	# 跑 ClientGetDemo，控制台直接输出
//	go run ./cmd/demo clienttimeout	# 跑 ClientTimeoutDemo，控制台直接输出
//	go run ./cmd/demo ginuser	# 跑 GinUserDemo，浏览器看 http://localhost:8087/user
//	go run ./cmd/demo ginquery	# 跑 GinQueryDemo，浏览器看 http://localhost:8088/user?name=tom&age=22
//	go run ./cmd/demo ginparam	# 跑 GinParamDemo，浏览器看 http://localhost:8089/user/1
//	go run ./cmd/demo ginbind	# 跑 GinBindDemo，curl POST JSON 看 http://localhost:8090/user
//	go run ./cmd/demo ginauth	# 跑 GinAuthDemo，curl 带 token 看 http://localhost:8091/user
//	go run ./cmd/demo gingroup	# 跑 GinGroupDemo，看 http://localhost:8092/api/...
package main

import (
	"fmt"
	"os"
	"practise"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run ./cmd/demo [server|user|userstatus|usermethod|userauth|userbody|userjson|clientget|clienttimeout|ginuser|ginquery|ginparam|ginbind|ginauth|gingroup]")
		fmt.Println("  server     -> practise.ServerDemo()     :8080 /")
		fmt.Println("  user       -> practise.UserDemo()       :8081 /user")
		fmt.Println("  userstatus -> practise.UserStatusDemo() :8082 /user")
		fmt.Println("  usermethod -> practise.UserMethodDemo() :8083 /user")
		fmt.Println("  userauth   -> practise.UserAuthDemo()   :8084 /user")
		fmt.Println("  userbody   -> practise.UserBodyDemo()   :8085 /user")
		fmt.Println("  userjson   -> practise.UserJSONDemo()   :8086 /user")
		fmt.Println("  clientget  -> practise.ClientGetDemo()  GET example.com，看控制台")
		fmt.Println("  clienttimeout -> practise.ClientTimeoutDemo() 带 3s 超时 GET，看控制台")
		fmt.Println("  ginuser -> practise.GinUserDemo() :8087 /user")
		fmt.Println("  ginquery -> practise.GinQueryDemo() :8088 /user?name=&age=")
		fmt.Println("  ginparam -> practise.GinParamDemo() :8089 /user/:id")
		fmt.Println("  ginbind -> practise.GinBindDemo() POST JSON :8090 /user")
		fmt.Println("  ginauth -> practise.GinAuthDemo() token 中间件 :8091 /user")
		fmt.Println("  gingroup -> practise.GinGroupDemo() Group 分组 :8092 /api/...")
		return
	}
	switch os.Args[1] {
	case "server":
		practise.ServerDemo()
	case "user":
		practise.UserDemo()
	case "userstatus":
		practise.UserStatusDemo()
	case "usermethod":
		practise.UserMethodDemo()
	case "userauth":
		practise.UserAuthDemo()
	case "userbody":
		practise.UserBodyDemo()
	case "userjson":
		practise.UserJSONDemo()
	case "clientget":
		practise.ClientGetDemo()
	case "clienttimeout":
		practise.ClientTimeoutDemo()
	case "ginuser":
		practise.GinUserDemo()
	case "ginquery":
		practise.GinQueryDemo()
	case "ginparam":
		practise.GinParamDemo()
	case "ginbind":
		practise.GinBindDemo()
	case "ginauth":
		practise.GinAuthDemo()
	case "gingroup":
		practise.GinGroupDemo()
	default:
		fmt.Println("未知 demo:", os.Args[1])
	}
}
