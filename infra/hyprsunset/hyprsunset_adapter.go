package hyprsunset

import (
	"lighttui/domain/adjustable"
	"os/exec"
	"strconv"
	"strings"
)

const maxTemperatureDeprecated = 1500
const minTemperatureDeprecated = 6500

type HyprsunsetAdapter struct {
	store adjustable.IAdjustableStore
}

func NewNighLightAdapter(store adjustable.IAdjustableStore) (*HyprsunsetAdapter, error) {
	n := &HyprsunsetAdapter{store}
	err := n.init()
	return n, err
}

func (n *HyprsunsetAdapter) init() error {
	nightlight, err := n.store.Fetch()
	if err != nil {
		return err
	}

	// Check for Hyprsunset process
	output, err := exec.Command("pgrep", "-a", "hyprsunset").Output()
	if err != nil {
		startHyprsunset(nightlight.GetCurrentValue())
	}

	// Convert output to string and split lines
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 0 {
		startHyprsunset(nightlight.GetCurrentValue())
	}

	return nil
}

func startHyprsunset(temperature int) {
	exec.Command("hyprsunset", "-t", strconv.Itoa(temperature)).Start()
}

func (n *HyprsunsetAdapter) ApplyValue(adjustable adjustable.IAdjustable) error {
	return exec.Command("hyprctl", "hyprsunset", "temperature", strconv.Itoa(adjustable.GetCurrentValue())).Start()
}
