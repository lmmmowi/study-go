package io

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name   string
	Age    int
	Gender string
}

func HandleJson() {
	s := encodeJson()
	decodeJson(s)
}

func encodeJson() []byte {
	// marshal
	user := &User{"Alex", 28, "Male"}
	s, _ := json.Marshal(user)
	fmt.Printf("%s\n", s)

	// with encoder
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(user)

	return s
}

func decodeJson(s []byte) {
	// unmarshal
	var u User
	json.Unmarshal(s, &u)
	fmt.Println(u)

	// unmarshal to map
	var f interface{}
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	json.Unmarshal(b, &f)
	m := f.(map[string]interface{})
	for k, v := range m {
		switch v.(type) {
		case string:
			fmt.Printf("%s ==> %s\n", k, v)
		case float64:
			fmt.Printf("%s ==> %f\n", k, v)
		case []interface{}:
			fmt.Printf("%s ==> %v\n", k, v)
		}
	}

	// with decoder
	file, _ := os.Open("assets/user.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	var u2 User
	decoder.Decode(&u2)
	fmt.Println(u2)
}
