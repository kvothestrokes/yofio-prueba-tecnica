package test

import (
	"testing"
	service "yofio-prueba-tecnica/service"
)

func TestCalcularInversion(t *testing.T) {
	tests := []struct {
		inversion int32
		expected  [3]int32
	}{
		// Casos de prueba con combinaciones válidas
		{1000, [3]int32{1, 0, 1}},
		{1500, [3]int32{1, 1, 1}},
		{2000, [3]int32{2, 0, 2}},
		{1700, [3]int32{2, 0, 1}},
		{300, [3]int32{0, 0, 1}},

		// Casos de prueba sin combinaciones válidas
		{250, [3]int32{0, 0, 0}},
		{400, [3]int32{0, 0, 0}},
		{750, [3]int32{0, 0, 0}},
		{1120, [3]int32{0, 0, 0}},
		{1805, [3]int32{0, 0, 0}},
	}

	for _, test := range tests {
		inv_700, inv_500, inv_300 := service.CalcularInversion(test.inversion)
		actual := [3]int32{inv_700, inv_500, inv_300}
		if actual != test.expected {
			t.Errorf("Para la inversion %d, se esperaba %v pero se obtuvo %v", test.inversion, test.expected, actual)
		}
	}
}
