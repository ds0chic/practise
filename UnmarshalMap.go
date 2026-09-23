package practise

import (
	"encoding/json"
	"fmt"
)

func UnmarshalMapDemo() {
	userdata := []byte(`[{"name":"Tom","age":22},{"name":"Bob","age":25}]`)
	var user []map[string]any
	err := json.Unmarshal(userdata, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
	fmt.Println(user[1]["name"])
}
