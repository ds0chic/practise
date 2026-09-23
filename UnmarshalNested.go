package practise

import (
	"encoding/json"
	"fmt"
)

func UnmarshalNestedDemo() {
	type Address struct {
		City   string `json:"city"`
		Street string `json:"street"`
	}
	type User struct {
		Name    string  `json:"name"`
		Age     int     `json:"age"`
		Address Address `json:"address"`
	}

	userdata := []byte(`{
		"name":"Tom",
		"age":22,
		"address":{
			"city":"Shanghai",
			"street":"Nanjing Road"
			}
		}`)
	var user User
	err := json.Unmarshal(userdata, &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
	fmt.Println(user.Address.City)
}
