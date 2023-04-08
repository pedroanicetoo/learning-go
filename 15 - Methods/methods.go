package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Saving the user %s data on database\n", u.name)
}

func (u user) legalAge() bool {
	return u.age >= 18
}

func (u *user) haveBirthday() {
	u.age++
}

func main() {
	user1 := user{"User 1", 12}
	fmt.Println(user1) // {User 1 12}
	user1.save()       // Saving the user User 1 data on database

	user2 := user{"User 2", 30}
	fmt.Println(user2.legalAge()) // true
	user2.haveBirthday()

	fmt.Println(user2.age) // 31

}
