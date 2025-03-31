package main

// A simple example that shows how to render an animated progress bar. In this
// example we bump the progress by 25% every two seconds, animating our
// progress bar to its new target state.

// It's also possible to render a progress bar in a more static fashion without
// transitions. For details on that approach see the progress-static example.

import (
	"fmt"
	"io"
	progress "lighttui/components"
	"lighttui/controllers"
	"lighttui/infra"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

var (
	titleStyle              = lipgloss.NewStyle().MarginLeft(2)
	itemStyle               = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle       = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle         = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle               = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle           = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	brightnessCtlController = controllers.NewBrightnessCtlController()
	nightLightController    = controllers.NewNighLightController()
)

type model struct {
	list list.Model
}

type item struct {
	name       string
	progress   progress.Model
	controller controllers.IController
}

func (i item) FilterValue() string { return i.name }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(item)
	if !ok {
		return
	}

	// Format item string with slider
	progressStr := item.progress.ViewAs(item.controller.GetPercentage())

	str := fmt.Sprintf("%d. %s %s", index+1, item.name, progressStr)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	item, ok := m.list.SelectedItem().(item)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			infra.Save(nightLightController.GetCurrent())
			return m, tea.Quit
		case "l":
			if ok {
				item.controller.IncreasePercentage(0.01)
				cmd := item.progress.SetPercent(item.controller.GetPercentage())
				return m, cmd
			}
		case "h":
			if ok {
				item.controller.DecreasePercentage(0.01)
				cmd := item.progress.SetPercent(item.controller.GetPercentage())
				return m, cmd
			}
		}

	case tea.WindowSizeMsg:
		m.list.SetWidth(min(msg.Width-padding*2-4, maxWidth))
		item.progress.Width = min(msg.Width-padding*2-4, maxWidth)
		return m, nil

	// FrameMsg is sent when the progress bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := item.progress.Update(msg)
		item.progress = progressModel.(progress.Model)
		return m, cmd
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return "\n" + m.list.View()
}

func main() {
	choices := []item{
		{name: "Brightness", progress: progress.New(progress.WithSolidFill("170"),
			progress.WithFillCharacters('█', '█'), progress.WithEmptyColor("238")), controller: brightnessCtlController},
		{name: "Night Light", progress: progress.New(progress.WithSolidFill("170"),
			progress.WithFillCharacters('█', '█'), progress.WithEmptyColor("238")), controller: nightLightController},
	}
	l := list.New([]list.Item{
		choices[0], choices[1],
	}, itemDelegate{}, maxWidth, 14)
	l.Title = "Adjust Your Settings"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{
		list: l,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}
