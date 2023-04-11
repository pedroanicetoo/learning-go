package main

import (
	"fmt"
	"introduction-tests/addresses"
)

func main() {
	addressType := addresses.AddressType("Avenida Paulista")
	fmt.Println(addressType)
}
