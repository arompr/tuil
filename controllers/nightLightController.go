package controllers

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type NightLightController struct {
	currentTemp    int
	currentProcess int
}

func NewNighLightController() *NightLightController {
	n := &NightLightController{}
	n.currentProcess = 0
	n.CalculateCurrent()
	return n
}

func (n *NightLightController) GetCurrent() int {
	return n.currentTemp
}

func (n *NightLightController) GetMax() int {
	return 1000
}

func (n *NightLightController) GetMin() int {
	return 6500
}

func (n *NightLightController) GetPercentage() float64 {
	// Calculate percentage (invert scale)
	return 1 - (float64(n.GetCurrent()-n.GetMax()) / float64(n.GetMin()-n.GetMax()))
}

func (n *NightLightController) IncreasePercentage(value float64) {
	if n.canIncrease() {
		// Increase by 1% and round to an integer
		newTemp := max(int(float64(n.GetCurrent())*float64(1-(value))), n.GetMax())

		// Apply the new temperature using `hyprsunset -t`
		exec.Command("pkill", "-9", "hyprsunset").Run()
		cmd2 := exec.Command("hyprsunset", "-t", strconv.Itoa(newTemp))
		cmd2.Start()

		n.currentTemp = newTemp
	}
}

// Max value is smaller than Min value (e.g., 1000 is more night light than 6500)
func (n *NightLightController) canIncrease() bool {
	return n.GetCurrent() > n.GetMax()
}

func (n *NightLightController) DecreasePercentage(value float64) {
	if n.canDecrease() {
		// Increase by 1% and round to an integer
		newTemp := min(int(float64(n.GetCurrent())*float64(1+value)), n.GetMin())

		exec.Command("pkill", "-9", "hyprsunset").Run()
		cmd2 := exec.Command("hyprsunset", "-t", strconv.Itoa(newTemp))
		cmd2.Start()

		n.currentTemp = newTemp
	}
}

// Min value is bigger than Max value (e.g., 6500 is less night light than 1000)
func (n *NightLightController) canDecrease() bool {
	return n.GetCurrent() < n.GetMin()
}

func (n *NightLightController) CalculateCurrent() {
	// Run `pgrep -a hyprsunset` to get running instances
	Ttemp := 4500
	cmd := exec.Command("pgrep", "-a", "hyprsunset")
	output, err := cmd.Output()
	if err != nil {
		Ttemp = 4500
	}

	// Convert output to string and split lines
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	// Reverse iterate to check for the last occurrence with `-t <temp>`
	re := regexp.MustCompile(`-t (\d+)`)
	for i := len(lines) - 1; i >= 0; i-- {
		matches := re.FindStringSubmatch(lines[i])
		if len(matches) > 1 {
			temp, err := strconv.Atoi(matches[1])
			if err == nil {
				Ttemp = temp
			}
		}
	}

	// Default to 4500K if no valid temperature was found
	n.currentTemp = Ttemp
}
