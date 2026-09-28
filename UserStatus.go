package practise

import (
	"fmt"
	"net/http"
)

func userStatusHandler(w http.ResponseWriter, r *http.Request) {
	userFound := false

	if !userFound {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "user not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "user found")
}

func UserStatusDemo() {
	http.HandleFunc("/user", userStatusHandler)
	http.ListenAndServe(":8082", nil)
}
