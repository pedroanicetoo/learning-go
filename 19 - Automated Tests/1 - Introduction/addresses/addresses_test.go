package addresses_test

// IMPORTANT COMMANDS:
// go test
// go test -v  (verbose)
// go test --cover (coverage)
// go test --coverprofile coverage.txt
// go tool cover --func=coverage.txt (to read the result test file)
// go tool cover --html=coverage.txt (shows a html file with each line covered and not covered by test)

import (
	"testing"
	. "tests-introduction/addresses" // use this notation (.) for auto function reference of this package, its kind of alias.
)

type testScenario struct {
	insertedAddress string
	expectedReturn  string
}

func TestAddressType(t *testing.T) {
	// t.Parallel() - to run parallel tests

	testScenarios := []testScenario{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		// {"Praça das Rosas", "Invalid Type"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		// {"", "Invalid Type"},
	}

	for _, scenario := range testScenarios {
		receivedReturn := AddressType(scenario.insertedAddress)

		if receivedReturn != scenario.expectedReturn {
			t.Errorf("Received type %s is different than expected %s",
				receivedReturn, scenario.expectedReturn)
		}
	}
}

func TestAny(t *testing.T) {
	// t.Parallel()

	if 1 > 2 {
		t.Errorf("Broke test!")
	}
}
