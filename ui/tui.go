package ui

import (
	"fmt"
	"io"
	"lighttui/application/usecase"
	"lighttui/ui/progress"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type model struct {
	list    list.Model
	persist *usecase.PersistUseCase
}

type item struct {
	name          string
	progress      progress.Model
	increase      *usecase.AdjustUseCase
	decrease      *usecase.AdjustUseCase
	getPercentage *usecase.GetPercentageUseCase
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

	progressStr := item.progress.ViewAs(getPercentage(item))

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
			// TODO : Add error handling to persist
			m.persist.Exec()
			return m, tea.Quit
		case "j":
			if m.list.Index() >= m.list.Paginator.ItemsOnPage(len(m.list.VisibleItems())-1) {
				m.list.CursorUp()
			} else {
				m.list.CursorDown()
			}
		case "k":
			if m.list.Index() == 0 {
				m.list.CursorDown()
			} else {
				m.list.CursorUp()
			}
		case "l":
			if ok {
				item.increase.Exec(0.01)
				cmd := item.progress.SetPercent(getPercentage(item))
				return m, cmd
			}
		case "L":
			if ok {
				item.increase.Exec(0.1)
				cmd := item.progress.SetPercent(getPercentage(item))
				return m, cmd
			}
		case "h":
			if ok {
				item.decrease.Exec(0.01)
				cmd := item.progress.SetPercent(getPercentage(item))
				return m, cmd
			}
		case "H":
			if ok {
				item.decrease.Exec(0.1)
				cmd := item.progress.SetPercent(getPercentage(item))
				return m, cmd
			}
		}

	case tea.WindowSizeMsg:
		m.list.SetWidth(min(msg.Width-padding*2-4, maxWidth))
		item.progress.Width = min(msg.Width-padding*2-4, maxWidth)
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func getPercentage(item item) float64 {
	percentage, err := item.getPercentage.Exec()
	if err != nil {
		log.Println(err)
	}
	return percentage
}

func (m model) View() string {
	return "\n" + m.list.View()
}

func NewTUI(increaseNightLightUseCase *usecase.AdjustUseCase,
	decreaseNightLightUseCase *usecase.AdjustUseCase,
	getNightLightPercentageUseCase *usecase.GetPercentageUseCase,
	increaseBrightnessUseCase *usecase.AdjustUseCase,
	decreaseBrightnessUseCase *usecase.AdjustUseCase,
	getBrightnessPercentageUseCase *usecase.GetPercentageUseCase,
	persistNightLightUseCase *usecase.PersistUseCase,
) *tea.Program {
	choices := []item{
		{
			name: "Brightness",
			progress: progress.New(progress.WithSolidFill("170"),
				progress.WithFillCharacters('█', '█'),
				progress.WithEmptyColor("238")),
			increase:      increaseBrightnessUseCase,
			decrease:      decreaseBrightnessUseCase,
			getPercentage: getBrightnessPercentageUseCase,
		},
		{
			name: "NightLight",
			progress: progress.New(progress.WithSolidFill("170"),
				progress.WithFillCharacters('█', '█'),
				progress.WithEmptyColor("238")),
			increase:      increaseNightLightUseCase,
			decrease:      decreaseNightLightUseCase,
			getPercentage: getNightLightPercentageUseCase,
		},
	}
	l := list.New([]list.Item{choices[0], choices[1]}, itemDelegate{}, maxWidth, 10)
	l.Title = "Adjust Your Settings"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	var myKeyMap = struct {
		Up        key.Binding
		Down      key.Binding
		Increase  key.Binding
		Decrease  key.Binding
		Increase2 key.Binding
		Decrease2 key.Binding
	}{
		Up: key.NewBinding(
			key.WithKeys("k"),
			key.WithHelp("k", "Up"),
		),
		Down: key.NewBinding(
			key.WithKeys("j"),
			key.WithHelp("j", "Down"),
		),
		Increase: key.NewBinding(
			key.WithKeys("l"),
			key.WithHelp("l", "Increase"),
		),
		Decrease: key.NewBinding(
			key.WithKeys("h"),
			key.WithHelp("h", "Decrease"),
		),
		Increase2: key.NewBinding(
			key.WithKeys("<S-l>"),
			key.WithHelp("<S-l>", "Increase more"),
		),
		Decrease2: key.NewBinding(
			key.WithKeys("<S-h>"),
			key.WithHelp("<S-h>", "Decrease more"),
		),
	}

	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			myKeyMap.Up,
			myKeyMap.Down,
			myKeyMap.Increase,
			myKeyMap.Decrease,
		}
	}

	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			myKeyMap.Up,
			myKeyMap.Down,
			myKeyMap.Increase,
			myKeyMap.Decrease,
			myKeyMap.Increase2,
			myKeyMap.Decrease2,
		}
	}

	l.KeyMap.PrevPage.SetEnabled(false)
	l.KeyMap.NextPage.SetEnabled(false)
	l.KeyMap.GoToStart.SetEnabled(false)
	l.KeyMap.GoToEnd.SetEnabled(false)
	l.KeyMap.CursorDown.SetEnabled(false)
	l.KeyMap.CursorUp.SetEnabled(false)

	m := model{
		list:    l,
		persist: persistNightLightUseCase,
	}

	return tea.NewProgram(m)
}
