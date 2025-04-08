package hyprsunset

import (
	"lighttui/domain/nightlight"
	"os/exec"
	"strconv"
	"strings"
)

const maxTemperatureDeprecated = 1500
const minTemperatureDeprecated = 6500

type HyprsunsetAdapter struct {
	store nightlight.INightLightStore
}

func NewNighLightController(store nightlight.INightLightStore) *HyprsunsetAdapter {
	n := &HyprsunsetAdapter{store}
	n.init()
	return n
}

func (n *HyprsunsetAdapter) init() {
	nightlight := n.store.FetchNightLight()

	// Check for Hyprsunset process
	output, err := exec.Command("pgrep", "-a", "hyprsunset").Output()
	if err != nil {
		startHyprsunset(nightlight.GetCurrentTemperature())
	}

	// Convert output to string and split lines
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 0 {
		startHyprsunset(nightlight.GetCurrentTemperature())
	}
}

func startHyprsunset(temperature int) {
	exec.Command("hyprsunset", "-t", strconv.Itoa(temperature)).Start()
}

func (n *HyprsunsetAdapter) ApplyNightLight(nightlight *nightlight.NightLight) error {
	return exec.Command("hyprctl", "hyprsunset", "temperature", strconv.Itoa(nightlight.GetCurrentTemperature())).Start()
}
