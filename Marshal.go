package practise

import (
	"encoding/json"
	"fmt"
)

func main() {
	type User struct {
		Name string `json:"your name"`
		Age  int    `json:"your age"`
	}
	user := User{
		Name: "Tom",
		Age:  22,
	}
	userdata, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(userdata))
}
