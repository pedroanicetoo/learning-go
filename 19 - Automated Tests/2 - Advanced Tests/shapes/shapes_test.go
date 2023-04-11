package shapes

// TDD - Test Driven Development

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		ret := Rectangle{10, 12}
		expectedArea := float64(120)
		receivedArea := ret.Area()

		if receivedArea != expectedArea {
			t.Fatalf("The received area %f is differend than expected area %f", receivedArea, expectedArea)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circ := Circle{10}

		expectedArea := float64(math.Pi * 100)
		receivedArea := circ.Area()

		if receivedArea != expectedArea {
			t.Fatalf("The received area %f is differend than expected area %f", receivedArea, expectedArea)
		}
	})
}
