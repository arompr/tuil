package controllers

import (
	"lighttui/domain/brightness"
	"os/exec"
	"strconv"
	"strings"
)

type BrightnessCtlAdapter struct{}

func NewBrightnessCtlController() *BrightnessCtlAdapter {
	b := &BrightnessCtlAdapter{}
	return b
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

func (b *BrightnessCtlAdapter) ApplyBrightness(brightness *brightness.Brightness) error {
	return exec.Command("brightnessctl", "set", strconv.Itoa(brightness.GetCurrentBrightness())).Start()
}
