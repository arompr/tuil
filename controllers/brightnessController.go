package controllers

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type BrightnessController struct {
	test int
}

func NewBrightnessCtlController() *BrightnessController {
	b := &BrightnessController{}
	return b
}

func (b *BrightnessController) GetCurrent() int {
	cmd := exec.Command("brightnessctl", "get")
	output, err := cmd.Output()
	if err != nil {
		return 50
	}

	value, err := strconv.Atoi(strings.TrimSpace(string(output)))
	if err != nil {
		return 50
	}

	return value
}

func (b *BrightnessController) GetMax() int {
	cmd := exec.Command("brightnessctl", "max")
	output, err := cmd.Output()
	if err != nil {
		return 50
	}

	value, err := strconv.Atoi(strings.TrimSpace(string(output)))
	if err != nil {
		return 50
	}

	return value
}

func (b *BrightnessController) GetPercentage() float64 {
	current := float64(b.GetCurrent())
	max := float64(b.GetMax())

	return current / max
}

func (b *BrightnessController) IncreasePercentage(value float64) {
	exec.Command("brightnessctl", "s", b.formatBrightnessctlArg(value, "+")).Start()
}

func (b *BrightnessController) DecreasePercentage(value float64) {
	exec.Command("brightnessctl", "s", b.formatBrightnessctlArg(value, "-")).Start()
}

// Format as "X%direction" (e.g., "20%+" or "20%-")
func (b *BrightnessController) formatBrightnessctlArg(value float64, direction string) string {
	return fmt.Sprintf("%d%%%s", int(value*100), direction)
}
