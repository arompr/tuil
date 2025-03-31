package controllers

import (
	"fmt"
	"lighttui/infra"
	"os/exec"
	"strconv"
	"strings"
)

const tempFilePath = "~/.local/state/hyprsunset_temp"
const maxTemperature = 1500
const minTemperature = 6500

type NightLightController struct {
	currentTemp int
}

func NewNighLightController() *NightLightController {
	n := &NightLightController{}
	n.init()
	return n
}

func (n *NightLightController) init() {
	tmp, err := infra.ReadOrInitTemperature()
	if err != nil {
		fmt.Print("Something went wrong")
	}
	num, _ := strconv.Atoi(tmp)
	n.currentTemp = num

	// Check for Hyprsunset process
	cmd := exec.Command("pgrep", "-a", "hyprsunset")
	output, err := cmd.Output()
	if err != nil {
		exec.Command("hyprsunset", "-t", tmp).Start()
	}

	// Convert output to string and split lines
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 0 {
		exec.Command("hyprsunset", "-t", tmp).Start()
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
