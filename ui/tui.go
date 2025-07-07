package ui

import (
	"fmt"
	"io"
	"lighttui/application/usecase"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	listModel *ListModel
	persist   *usecase.SaveUseCase
}

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(ListItem)
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
	item, ok := m.listModel.list.SelectedItem().(ListItem)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			// TODO : Add error handling to persist
			m.persist.Exec()
			return m, tea.Quit
		case "j":
			if m.listModel.isLastItemFocused() {
				m.listModel.list.CursorUp()
			} else {
				m.listModel.list.CursorDown()
			}
		case "k":
			if m.listModel.isFirstItemFocused() {
				m.listModel.list.CursorDown()
			} else {
				m.listModel.list.CursorUp()
			}
		case "l":
			if ok {
				item.Increase(0.01)
			}
		case "L":
			if ok {
				item.Increase(0.1)
			}
		case "h":
			if ok {
				item.Decrease(0.01)
			}
		case "H":
			if ok {
				item.Decrease(0.1)
			}
		}

	case tea.WindowSizeMsg:
		m.listModel.list.SetWidth(min(msg.Width-padding*2-4, maxWidth))
		item.progress.Width = min(msg.Width-padding*2-4, maxWidth)
		return m, nil
	}

	var cmd tea.Cmd
	m.listModel.list, cmd = m.listModel.list.Update(msg)
	return m, cmd
}

func getPercentage(item ListItem) float64 {
	percentage, err := item.GetPercentage()
	if err != nil {
		log.Println(err)
	}
	return percentage
}

func (m model) View() string {
	return "\n" + m.listModel.list.View()
}

func NewTUI(listModel *ListModel, persistNightLightUseCase *usecase.SaveUseCase,
) *tea.Program {
	m := model{
		listModel: listModel,
		persist:   persistNightLightUseCase,
	}

	return tea.NewProgram(m)
}
