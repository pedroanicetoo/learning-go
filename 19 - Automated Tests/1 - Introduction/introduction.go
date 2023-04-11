package main

import (
	"fmt"
	"tests-introduction/addresses"
)

func main() {
	addressType := addresses.AddressType("Avenida Paulista")
	fmt.Println(addressType)
}
