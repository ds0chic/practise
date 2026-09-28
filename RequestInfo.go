package practise

import (
	"fmt"
	"net/http"
)

func requestInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	fmt.Println("Path:", r.URL.Path)
	fmt.Println("Useragent:", r.Header.Get("User-Agent"))
	fmt.Fprintln(w, "request received")
}

func RequestInfoDemo() {
	http.HandleFunc("/info", requestInfoHandler)
	http.ListenAndServe(":8081", nil)
}
