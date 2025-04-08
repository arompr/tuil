package controllers

import (
	"lighttui/domain/nightlight"
	"os/exec"
	"strconv"
	"strings"
)

const (
	maxTemperature = 1500
	minTemperature = 6500
)

type NightLightController struct {
	currentTemp      int
	temperatureStore nightlight.ITemperatureStoreDeprecated
}

func NewNighLightController(
	temperatureStore nightlight.ITemperatureStoreDeprecated,
) *NightLightController {
	n := &NightLightController{
		temperatureStore: temperatureStore,
	}
	n.init()
	return n
}

func (n *NightLightController) init() {
	tmp := n.temperatureStore.GetTemperature()
	n.currentTemp = tmp
	// Check for Hyprsunset process
	cmd := exec.Command("pgrep", "-a", "hyprsunset")
	output, err := cmd.Output()
	if err != nil {
		exec.Command("hyprsunset", "-t", strconv.Itoa(tmp)).Start()
	}

	// Convert output to string and split lines
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 0 {
		exec.Command("hyprsunset", "-t", strconv.Itoa(tmp)).Start()
	}
}

func (n *NightLightController) GetCurrent() int {
	return n.currentTemp
}

func (n *NightLightController) GetPercentage() float64 {
	// Calculate percentage (invert scale)
	return 1 - (float64(n.GetCurrent()-maxTemperature) / float64(minTemperature-maxTemperature))
}

func (n *NightLightController) IncreasePercentage(value float64) {
	if n.canIncrease() {
		// Increase by 1% and round to an integer
		newTemperature := max(int(float64(n.GetCurrent())*float64(1-(value))), maxTemperature)
		n.applyNewTemperature(newTemperature)
	}
}

// Max value is smaller than Min value (e.g., 1500 is more night light than 6500)
func (n *NightLightController) canIncrease() bool {
	return n.GetCurrent() > maxTemperature
}

func (n *NightLightController) DecreasePercentage(value float64) {
	if n.canDecrease() {
		// Increase by 1% and round to an integer
		newTemperature := min(int(float64(n.GetCurrent())*float64(1+value)), minTemperature)
		n.applyNewTemperature(newTemperature)
	}
}

// Min value is bigger than Max value (e.g., 6500 is less night light than 1500)
func (n *NightLightController) canDecrease() bool {
	return n.GetCurrent() < minTemperature
}

func (n *NightLightController) applyNewTemperature(newTemperature int) {
	exec.Command("hyprctl", "hyprsunset", "temperature", strconv.Itoa(newTemperature)).Start()
	n.currentTemp = newTemperature
}
