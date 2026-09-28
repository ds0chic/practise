package practise

import (
	"fmt"
	"net/http"
)

func userMethodHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "获取 user 信息")
	case "POST":
		fmt.Fprintln(w, "创建 user ")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "method not allowed")
	}
}

func UserMethodDemo() {
	http.HandleFunc("/user", userMethodHandler)
	http.ListenAndServe(":8083", nil)
}
