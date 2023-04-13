package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name string `json:"name"`
	Race string `json:"race"`
	Age  uint   `json:"age"`
}

func main() {
	dog_JSON := `{"name":"Rex","race":"Dálmata","age":3}`

	var d dog

	if erro_r := json.Unmarshal([]byte(dog_JSON), &d); erro_r != nil {
		log.Fatal(erro_r)
	}

	fmt.Println(d) // {Rex Dálmata 3}

	dog2_JSON := `{"name":"Toby","race":"Poodle"}`

	d2 := make(map[string]string)

	if error_r := json.Unmarshal([]byte(dog2_JSON), &d2); error_r != nil {
		log.Fatal(error_r)
	}

	fmt.Println(d2) // map[name:Toby race:Poodle]
}
