package other

import (
	"fmt"
	"io"
	"math"
	"os"
	"os/n
exec"
	"strconv"
	"strings"

	"lighttui/brightness"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight = 14
	maxWidth   = 30
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type item struct {
	progress progress.Model
	name     string
	level    int
	max      int
}

func (i item) FilterValue() string { return i.name }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	// Generate slider
	//slider := renderSlider(i.level, 0, i.max, 30)

	// Format item string with slider
	str := fmt.Sprintf("%d. %s %s", index+1, i.name, i.progress.View())

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func renderSlider(current, min, max, width int) string {
	if current > max {
		current = max
	} else if current < min {
		current = min
	}

	// Normalize brightness to a 0-100 scale
	normalized := int(math.Round(float64(current-min) / float64(max-min) * 100))

	// Map 0-100 to the available bar width (fine-grained control)
	totalBlocks := width * 8 // 8 sub-levels per character
	filledBlocks := int(math.Round(float64(normalized) / 100 * float64(totalBlocks)))

	// Full block characters (█) count
	fullChars := filledBlocks / 8

	// Partial block characters (▏▎▍▌▋▊▉█)
	blocks := []string{"", "▏", "▎", "▍", "▌", "▋", "▊", "▉", "█"}
	partialChar := blocks[filledBlocks%8]

	// Styles
	fgStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("170")).Bold(true) // Purple foreground
	bgStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("238"))            // Pale background

	// Build the bar
	filledBar := strings.Repeat("█", fullChars) // Full blocks
	partialBlock := partialChar                 // Single partial block (if any)

	// Apply foreground color to both filled and partial blocks
	filledBar = fgStyle.Render(filledBar)
	if partialBlock != "" {
		partialBlock = fgStyle.Render(partialBlock)
	}

	// Calculate empty width correctly
	emptyWidth := width - fullChars - len(partialBlock)
	emptyBar := bgStyle.Render(strings.Repeat("█", emptyWidth)) // Background-colored empty blocks

	// Concatenate all parts smoothly
	return filledBar + partialBlock + emptyBar
}

type model struct {
	list                 list.Model
	choices              []item
	brightnessController brightness.BrightnessController
	quitting             bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter", "l":
			if err := m.brightnessController.Increase(); err != nil {
				fmt.Println("Error increasing brightness:", err)
				return m, nil
			}

			if item, ok := m.list.SelectedItem().(item); ok {
				item.level = m.brightnessController.GetCurrent()
				item.progress.SetPercent(float64(item.level) / float64(item.max))
				m.choices[m.list.Index()] = item
				m.list.SetItem(m.list.Index(), item)
			}

			return m, nil
		case "h":
			if err := m.brightnessController.Decrease(); err != nil {
				fmt.Println("Error increasing brightness:", err)
				return m, nil
			}

			if item, ok := m.list.SelectedItem().(item); ok {
				item.level = m.brightnessController.GetCurrent()
				m.choices[m.list.Index()] = item
				m.list.SetItem(m.list.Index(), item)
			}

			return m, nil
		}
FrameMsg is sent when the progress bar wants to animate itself
case progress.FrameMsg:
	progressModel, cmd := m.progress.Update(msg)
	m.progress = progressModel.(progress.Model)
	return m, cmd
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return quitTextStyle.Render("Goodbye!")
	}
	return "\n" + m.list.View()
}

// setNightLight updates the temperature and saves it persistently
// func setNightLight(temp int) {
// 	cmd := exec.Command("gammastep", "-O", strconv.Itoa(temp))
// 	cmd.Run()

// 	// Ensure config directory exists
// 	if err := os.MkdirAll(configDir, 0755); err != nil {
// 		fmt.Println("Error creating config directory:", err)
// 		return
// 	}

// 	// Save the temperature value
// 	if err := os.WriteFile(nightLightFile, []byte(strconv.Itoa(temp)), 0644); err != nil {
// 		fmt.Println("Error saving night light temperature:", err)
// 	}
// }

// Ensure the directory exists
func ensureConfigDir() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	configDir := homeDir + "/.config/lighttui"
	if err := os.MkdirAll(configDir, 0755); err != nil {
		fmt.Println("Error creating config directory:", err)
	}
}

// getCurrentNightLight retrieves the last-set temperature from a file
func getCurrentNightLight() int {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return 4500
	}

	nightLightFile := homeDir + "/.config/lighttui/night_light_temp"
	ensureConfigDir()

	if _, err := os.Stat(nightLightFile); os.IsNotExist(err) {
		defaultTemp := 4500
		err := os.WriteFile(nightLightFile, []byte(strconv.Itoa(defaultTemp)), 0644)
		if err != nil {
			fmt.Println("Error writing default night light value:", err)
		}
		return defaultTemp
	}

	// Read the value from file
	data, err := os.ReadFile(nightLightFile)
	if err != nil {
		fmt.Println("Error reading night light value:", err)
		return 4500 // Default fallback
	}

	// Convert string to integer
	temp, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		fmt.Println("Error parsing night light value:", err)
		return 4500 // Default fallback
	}

	return temp
}

func getCurrentBrightness() int {
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

func getMaxBrightness() int {
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

func main() {
	brightnessController := brightness.BrightnessCtl{}
	// Initial choices with levels
	choices := []item{
		{progress: progress.New(progress.WithDefaultGradient()), name: "Brightness", level: brightnessController.GetCurrent(), max: brightnessController.GetMax()},
		// {name: "Night Light", level: getCurrentNightLight(), max: 1000},
	}

	const defaultWidth = 20

	l := list.New([]list.Item{
		choices[0], // choices[1],
	}, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Adjust Your Settings"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{list: l, choices: choices, brightnessController: brightnessController}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}