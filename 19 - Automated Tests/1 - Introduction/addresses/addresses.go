package addresses

import (
	"strings"
)

// AddressType check if a address has a valid type and returns it.
func AddressType(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}

	addresOnLowerCase := strings.ToLower(address)
	firstWordOfAddress := strings.Split(addresOnLowerCase, " ")[0]

	addressHaveValidType := false

	for _, typ_e := range validTypes {
		if typ_e == firstWordOfAddress {
			addressHaveValidType = true
		}
	}

	if addressHaveValidType {
		return strings.Title(firstWordOfAddress)
	}

	return "Invalid Type"
}
