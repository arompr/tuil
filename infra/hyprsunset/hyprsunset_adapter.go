package hyprsunset

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"lighttui/domain/adjustable"
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

	if isHyprsunsetRunning() {
		temp, err := getCurrentHyprsunsetTemperature()
		if err != nil {
			return err
		}

		nightlight.ApplyValue(temp)
		if err := h.store.Save(nightlight); err != nil {
			return err
		}
	} else {
		if err := startHyprsunset(nightlight.GetCurrentValue()); err != nil {
			return err
		}
	}

	return nil
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

func execHyprsunsetTemperature(adjustable adjustable.IAdjustable) error {
	return exec.Command("hyprctl", "hyprsunset", "temperature", strconv.Itoa(adjustable.GetCurrentValue())).Start()
}

func getCurrentHyprsunsetTemperature() (int, error) {
	cmd := exec.Command("bash", "-c", "hyprctl hyprsunset temperature 2>/dev/null | grep -oE '[0-9]+'")
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	// Clean up whitespace and convert to int
	valueStr := strings.TrimSpace(string(out))
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, err
	}

	return value, nil
}
