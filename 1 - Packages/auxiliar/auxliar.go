package auxiliar

import "fmt"

/*
We can write private functions (starts with lowercase letters) and calls it on the same package.
btw on external calls of other packages we can only use public functions (starts with uppercase letters)
*/
func Write() {
	fmt.Println("writing from auxiliar package")
	write2()
}
