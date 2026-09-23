package practise

import (
	"encoding/json"
	"fmt"
	"strings"
)

func DecoderDemo() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	reader := strings.NewReader(`{
		"name":"Tom",
		"age":22
	}`)

	decoder := json.NewDecoder(reader)

	var user User

	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
}
