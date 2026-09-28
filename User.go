package practise

import (
	"fmt"
	"net/http"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	fmt.Println("Path:", r.URL.Path)
	fmt.Println("Useragent:", r.Header.Get("User-Agent"))
	fmt.Fprintln(w, "request received")
}

func UserDemo() {
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8081", nil)
}
