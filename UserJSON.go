package practise

import (
	"encoding/json"
	"net/http"
)

func userJSONHandler(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UserJSONDemo() {
	http.HandleFunc("/user", userJSONHandler)
	http.ListenAndServe(":8086", nil)
}
