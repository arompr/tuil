package nl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNightLight_Increase(t *testing.T) {
	tests := []struct {
		name                string
		initialTemp         int
		percentage          float64
		expectedTemperature int
	}{
		{
			name:                "Increase by 1%",
			initialTemp:         6500,
			percentage:          0.01,
			expectedTemperature: 6450,
		},
		{
			name:                "Increase by 10%",
			initialTemp:         6500,
			percentage:          0.10,
			expectedTemperature: 6000,
		},
		{
			name:                "Increase without exceeding min",
			initialTemp:         1500,
			percentage:          0.1,
			expectedTemperature: 1500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nl := CreateNewNightlight(tt.initialTemp)
			nl.Increase(tt.percentage)
			assert.Equal(t, tt.expectedTemperature, nl.GetCurrentValue())
		})
	}
}

func TestNightLight_Decrease(t *testing.T) {
	tests := []struct {
		name                string
		initialTemp         int
		percentage          float64
		expectedTemperature int
	}{
		{
			name:                "Decrease by 1%",
			initialTemp:         1500,
			percentage:          0.01,
			expectedTemperature: 1550,
		},
		{
			name:                "Decrease by 10%",
			initialTemp:         1500,
			percentage:          0.10,
			expectedTemperature: 2000,
		},
		{
			name:                "Decrease should not go below min",
			initialTemp:         6480,
			percentage:          0.10,
			expectedTemperature: 6500,
		},
		{
			name:                "Decrease when already at min",
			initialTemp:         6500,
			percentage:          0.10,
			expectedTemperature: 6500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nl := CreateNewNightlight(tt.initialTemp)
			nl.Decrease(tt.percentage)
			assert.Equal(t, tt.expectedTemperature, nl.GetCurrentValue())
		})
	}
}
