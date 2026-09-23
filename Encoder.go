package practise

import (
	"encoding/json"
	"fmt"
	"os"
)

func EncoderDemo() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	user := User{
		Name: "Tom",
		Age:  22,
	}
	encoder := json.NewEncoder(os.Stdout)
	err := encoder.Encode(user)
	if err != nil {
		fmt.Println(err)
		return
	}
}
