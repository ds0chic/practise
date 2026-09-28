package practise

import (
	"fmt"
	"io"
	"net/http"
)

func userBodyHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

	fmt.Fprintln(w, "收到数据")
}

func UserBodyDemo() {
	http.HandleFunc("/user", userBodyHandler)
	http.ListenAndServe(":8085", nil)
}
