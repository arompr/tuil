package brightnessctl

import (
	"lighttui/domain/adjustable"
	"os/exec"
	"strconv"
	"strings"
)

type BrightnessCtlAdapter struct{}

func NewBrightnessCtlAdapter() *BrightnessCtlAdapter {
	return &BrightnessCtlAdapter{}
}

func (b *BrightnessCtlAdapter) GetCurrentBrightnessValue() (int, error) {
	output, err := exec.Command("brightnessctl", "get").Output()

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(strings.TrimSpace(string(output)))
}

func (b *BrightnessCtlAdapter) GetMaxBrightnessValue() (int, error) {
	output, err := exec.Command("brightnessctl", "max").Output()

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(strings.TrimSpace(string(output)))
}

func (b *BrightnessCtlAdapter) ApplyValue(adjustable adjustable.IAdjustable) error {
	return exec.Command("brightnessctl", "set", strconv.Itoa(adjustable.GetCurrentValue())).Start()
}
