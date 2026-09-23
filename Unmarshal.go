package practise

import (
	"encoding/json"
	"fmt"
)

func UnmarshalDemo() {
	type User struct {
		Name string `json:"your name"`
		Age  int    `json:"your age"`
	}
	userdata := []byte(`{"your name":"Tom","your age":22}`)
	var user User
	err := json.Unmarshal(userdata, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("name:", user.Name)
	fmt.Println("age:", user.Age)
}
