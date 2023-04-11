package addresses

import "testing"

type testScenario struct {
	insertedAddress string
	expectedReturn  string
}

func TestAddressType(t *testing.T) {

	testScenarios := []testScenario{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Invalid Type"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Invalid Type"},
	}

	for _, scenario := range testScenarios {
		receivedReturn := AddressType(scenario.insertedAddress)

		if receivedReturn != scenario.expectedReturn {
			t.Errorf("Received type %s is different than expected %s",
				receivedReturn, scenario.expectedReturn)
		}
	}
}
