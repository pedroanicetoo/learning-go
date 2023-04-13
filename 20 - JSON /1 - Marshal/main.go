package main

import (
	"bytes"
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
	do_g := dog{"Rex", "Dálmata", 3}
	fmt.Println(do_g) // {Rex Dálmata 3}

	do_gtoJSON, erro_r := json.Marshal(do_g)

	if erro_r != nil {
		log.Fatal(erro_r)
	}

	fmt.Println(do_gtoJSON)                  // [123 34 110 97 109 101 34 58 34 82 101 120 34 44 34 114 97 99 101 34 58 34 68 195 161 108 109 97 116 97 34 44 34 97 103 101 34 58 51 125]
	fmt.Println(bytes.NewBuffer(do_gtoJSON)) // {"name":"Rex","race":"Dálmata","age":3}

	do_g2 := map[string]string{
		"name": "Toby",
		"race": "Poodle",
	}

	do_g2toJSON, erro_r := json.Marshal(do_g2)

	if erro_r != nil {
		log.Fatal(erro_r)
	}

	fmt.Println(do_g2toJSON)                  // [123 34 110 97 109 101 34 58 34 84 111 98 121 34 44 34 114 97 99 101 34 58 34 80 111 111 100 108 101 34 125]
	fmt.Println(bytes.NewBuffer(do_g2toJSON)) // {"name":"Toby","race":"Poodle"}
}
