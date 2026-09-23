package practise

import (
	"encoding/json"
	"fmt"
)

func UnmarshalSliceDemo() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	userdata := []byte(`[{"name":"Tom","age":22},{"name":"Bob","age":25}]`)
	var user []User
	err := json.Unmarshal(userdata, &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
	fmt.Println(user[1].Age)
}
