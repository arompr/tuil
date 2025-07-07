package hyprsunset

import (
	"lighttui/domain/adjustable"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	maxTemperatureDeprecated = 1500
	minTemperatureDeprecated = 6500
)

type HyprsunsetAdapter struct {
	store adjustable.IAdjustableStore
}

func NewNighLightAdapter(store adjustable.IAdjustableStore) *HyprsunsetAdapter {
	return &HyprsunsetAdapter{store}
}

func (h *HyprsunsetAdapter) IsAvailable() bool {
	return isHyprlandRunning()
}

func (h *HyprsunsetAdapter) Start() error {
	return h.ensureHyprsunsetRunning()
}

func (h *HyprsunsetAdapter) ApplyValue(adjustable adjustable.IAdjustable) error {
	return execHyprsunsetTemperature(adjustable)
}

func (h *HyprsunsetAdapter) ensureHyprsunsetRunning() error {
	nightlight, err := h.store.Fetch()
	if err != nil {
		return err
	}

	// Check for Hyprsunset process
	output, err := exec.Command("pgrep", "-a", "hyprsunset").Output()
	if err != nil {
		err := startHyprsunset(nightlight.GetCurrentValue())
		if err != nil {
			return err
		}
	}

	// Convert output to string and split lines
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 0 {
		err := startHyprsunset(nightlight.GetCurrentValue())
		if err != nil {
			return err
		}
	}

	return nil
}

func startHyprsunset(temperature int) error {
	return exec.Command("hyprsunset", "-t", strconv.Itoa(temperature)).Start()
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

func execHyprsunsetTemperature(adjustable adjustable.IAdjustable) error {
	return exec.Command("hyprctl", "hyprsunset", "temperature", strconv.Itoa(adjustable.GetCurrentValue())).Start()
}
