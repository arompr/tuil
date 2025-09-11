package hyprsunset

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"lighttui/domain/adjustable/nightlight"
)

type HyprsunsetAdapter struct{}

func NewHyprsunsetAdapter() *HyprsunsetAdapter {
	return &HyprsunsetAdapter{}
}

func (adapter *HyprsunsetAdapter) Start(initialValue int) error {
	if !adapter.IsAvailable() {
		return &nightlight.ErrNightlightAdapterUnavailable{Adapter: "HyprsunsetAdapter"}
	}

	if isHyprsunsetRunning() {
		return nil
	}

	return startHyprsunset(initialValue)
}

func (adapter *HyprsunsetAdapter) IsAvailable() bool {
	return isHyprlandRunning()
}

func (adapter *HyprsunsetAdapter) ApplyValue(nightlight *nightlight.Nightlight) error {
	return execHyprsunsetTemperature(nightlight)
}

func isHyprsunsetRunning() bool {
	output, err := exec.Command("pgrep", "-a", "hyprsunset").Output()
	if err != nil {
		// pgrep returns error if no process matched
		return false
	}

	// If output is non-empty, Hyprsunset is running
	return len(strings.TrimSpace(string(output))) > 0
}

func startHyprsunset(temperature int) error {
	return exec.Command("setsid", "hyprsunset", "-t", strconv.Itoa(temperature)).Start()
}

func isHyprlandRunning() bool {
	// 1. Check if HYPRLAND_INSTANCE_SIGNATURE is set
	sig := os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")
	if sig == "" {
		return false // Not running in Hyprland session
	}

	// 2. Try running `hyprctl activewindow` to verify connection
	cmd := exec.Command("hyprctl", "activewindow")
	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}

func execHyprsunsetTemperature(nightlight *nightlight.Nightlight) error {
	return exec.Command("hyprctl", "hyprsunset", "temperature", strconv.Itoa(nightlight.GetCurrentValue())).Start()
}

func (adapter *HyprsunsetAdapter) GetCurrentNightlight() (*nightlight.Nightlight, error) {
	if !adapter.IsAvailable() {
		return nil, &nightlight.ErrNightlightAdapterUnavailable{Adapter: "Hyprsunset"}
	}

	cmd := exec.Command("bash", "-c", "hyprctl hyprsunset temperature 2>/dev/null | grep -oE '[0-9]+'")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Clean up whitespace and convert to int
	valueStr := strings.TrimSpace(string(out))
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return nil, err
	}

	return nightlight.CreateNewNightLight(value), nil
}
