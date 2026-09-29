package practise

import (
	"fmt"
	"io"
	"net/http"
)

func ClientGetDemo() {
	resp, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Status:", resp.Status)
	fmt.Println(string(body))
}
