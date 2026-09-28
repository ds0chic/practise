package practise

import (
	"fmt"
	"net/http"
)

func userAuthHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	fmt.Println("User-Agent:", r.Header.Get("User-Agent"))
	fmt.Println("Authorization:", token)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintln(w, `{"status":"ok"}`)
}

func UserAuthDemo() {
	http.HandleFunc("/user", userAuthHandler)
	http.ListenAndServe(":8084", nil)
}
