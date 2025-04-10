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
			b := brightness.CreateNewNightLight(tt.initial, 100, 0)
			b.Increase(tt.percent)
			if b.GetCurrentBrightness() != tt.expected {
				t.Errorf("got %d, want %d", b.GetCurrentBrightness(), tt.expected)
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
			b := brightness.CreateNewNightLight(tt.initial, 100, 1)
			b.Decrease(tt.percent)
			if b.GetCurrentBrightness() != tt.expected {
				t.Errorf("got %d, want %d", b.GetCurrentBrightness(), tt.expected)
			}
		})
	}
}

func TestGetPercentage(t *testing.T) {
	b := brightness.CreateNewNightLight(25, 100, 0)
	expected := 0.25
	actual := b.GetPercentage()
	if actual != expected {
		t.Errorf("expected %.2f, got %.2f", expected, actual)
	}
}
