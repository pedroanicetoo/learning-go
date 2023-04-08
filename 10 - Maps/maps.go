package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"first_name": "Peter",
		"last_name":  "Parker",
	}

	fmt.Println(user) // map[first_name:Peter last_name:Parker]

	user2 := map[string]map[string]string{
		"name": {
			"first": "Peter",
			"last":  "Parker",
		},
		"hero": {
			"name": "Spider Man",
		},
	}

	fmt.Println(user2) // map[hero:map[name:Spider Man] name:map[first:Peter last:Parker]]

	delete(user2, "name")
	fmt.Println(user2) // map[hero:map[name:Spider Man]]

	user2["hero_tutor"] = map[string]string{
		"name": "Tony Start",
	}
	fmt.Println(user2) // map[hero:map[name:Spider Man] hero_tutor:map[name:Tony Start]]
}
