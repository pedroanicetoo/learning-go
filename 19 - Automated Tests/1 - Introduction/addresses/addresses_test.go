package addresses

import "testing"

func TestAddressType(t *testing.T) {
	addressToTest := "Rua ABC"

	expectedAdressType := "Rua"

	receivedAddressType := AddressType(addressToTest)

	if receivedAddressType != expectedAdressType {
		t.Errorf("Received type is different than expected! Expected %s and received %s",
			expectedAdressType,
			receivedAddressType)
	}
}
