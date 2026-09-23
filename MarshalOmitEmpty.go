package practise

import (
	"encoding/json"
	"fmt"
)

func MarshalOmitEmptyDemo() {
	type User struct {
		Name string `json:"your name"`
		Age  int    `json:"your age,omitempty"`
	}
	user := User{
		Name: "Tom",
		Age:  0,
	}
	userdata, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(userdata))
}
