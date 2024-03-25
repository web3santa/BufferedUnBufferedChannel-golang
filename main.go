package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	// DecodeJson()
	EncodeJson()
}

func EncodeJson() {
	var mar course
	st, err := json.Marshal(&mar)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(st)

	json.Unmarshal(st, &mar)
	fmt.Printf("%#v\n", mar)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
			"name": "Alice",
			"age": 30,
			"city": "New York"
	  }
	`)
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and vlaue is %v and type of %T\n", k, v, v)
	}
}
