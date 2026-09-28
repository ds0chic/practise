package practise

import (
	"fmt"
	"net/http"
)

func ServerDemo() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})
	http.ListenAndServe(":8080", nil)
}
