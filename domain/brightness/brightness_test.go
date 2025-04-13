package brightness_test

import (
	"lighttui/domain/brightness"
	"testing"
)

func TestIncreaseBrightness(t *testing.T) {
	tests := []struct {
		name     string
		initial  int
		percent  float64
		expected int
	}{
		{"increase by 20% should be 70", 50, 0.2, 70},
		{"increase past max should not exceed max", 90, 0.5, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := brightness.CreateNewBrightness(tt.initial, 100)
			b.Increase(tt.percent)
			if b.GetCurrentValue() != tt.expected {
				t.Errorf("got %d, want %d", b.GetCurrentValue(), tt.expected)
			}
		})
	}
}

func TestDecreaseBrightness(t *testing.T) {
	tests := []struct {
		name     string
		initial  int
		percent  float64
		expected int
	}{
		{"decrease by 20% should be 30", 50, 0.2, 30},
		{"decrease below min should not exceed min", 10, 0.5, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := brightness.CreateNewBrightness(tt.initial, 100)
			b.Decrease(tt.percent)
			if b.GetCurrentValue() != tt.expected {
				t.Errorf("got %d, want %d", b.GetCurrentValue(), tt.expected)
			}
		})
	}
}

func TestGetPercentage(t *testing.T) {
	b := brightness.CreateNewBrightness(25, 100)
	expected := 0.25
	actual := b.GetPercentage()
	if actual != expected {
		t.Errorf("expected %.2f, got %.2f", expected, actual)
	}
}
