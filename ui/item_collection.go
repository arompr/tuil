package ui

import (
	"lighttui/application/usecase"
	"lighttui/ui/progress"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ListModel struct {
	list list.Model
}

func NewListModel(list list.Model) *ListModel {
	return &ListModel{list}
}

func (lic ListModel) isFirstItemFocused() bool {
	return lic.list.Index() == 0
}

func (lic ListModel) isLastItemFocused() bool {
	return lic.list.Index() >= lic.list.Paginator.ItemsOnPage(len(lic.list.VisibleItems())-1)
}

type ListItemCollection struct {
	List []ListItem
}

type ListItem struct {
	name          string
	progress      progress.Model
	increase      usecase.IAdjustableUseCase
	decrease      usecase.IAdjustableUseCase
	getPercentage usecase.IGetAdjustablePercentageUseCase
}

func NewListItemCollection() *ListItemCollection {
	return &ListItemCollection{[]ListItem{}}
}

func (l *ListItemCollection) AddBrightness(increaseBrightnessUseCase *usecase.AdjustUseCase,
	decreaseBrightnessUseCase *usecase.AdjustUseCase,
	getBrightnessPercentageUseCase *usecase.GetPercentageUseCase,
) {
	l.List = append(l.List, ListItem{
		name: "Brightness",
		progress: progress.New(progress.WithSolidFill("170"),
			progress.WithFillCharacters('█', '█'),
			progress.WithEmptyColor("238")),
		increase:      increaseBrightnessUseCase,
		decrease:      decreaseBrightnessUseCase,
		getPercentage: getBrightnessPercentageUseCase,
	})
}

func (l *ListItemCollection) AddNightLight(increaseNightLightUseCase *usecase.AdjustNightlightUseCase,
	decreaseNightLightUseCase *usecase.AdjustNightlightUseCase,
	getNightLightPercentageUseCase *usecase.GetNightlightPercentageUseCase,
) {
	l.List = append(l.List, ListItem{
		name: "NightLight",
		progress: progress.New(progress.WithSolidFill("170"),
			progress.WithFillCharacters('█', '█'),
			progress.WithEmptyColor("238")),
		increase:      increaseNightLightUseCase,
		decrease:      decreaseNightLightUseCase,
		getPercentage: getNightLightPercentageUseCase,
	})
}

func BuildListModel(listItems []ListItem) *ListModel {
	items := make([]list.Item, len(listItems))
	for i, c := range listItems {
		items[i] = c
	}

	listModel := list.New(items, itemDelegate{}, maxWidth, 10)
	listModel.Title = "Adjust Your Settings"
	listModel.SetShowStatusBar(false)
	listModel.SetFilteringEnabled(false)
	listModel.Styles.Title = titleStyle
	listModel.Styles.PaginationStyle = paginationStyle
	listModel.Styles.HelpStyle = helpStyle

	keys := newKeyMap()

	listModel.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			keys.Up,
			keys.Down,
			keys.Increase,
			keys.Decrease,
		}
	}

	listModel.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			keys.Up,
			keys.Down,
			keys.Increase,
			keys.Decrease,
			keys.IncreaseMore,
			keys.DecreaseMore,
		}
	}

	listModel.KeyMap.PrevPage.SetEnabled(false)
	listModel.KeyMap.NextPage.SetEnabled(false)
	listModel.KeyMap.GoToStart.SetEnabled(false)
	listModel.KeyMap.GoToEnd.SetEnabled(false)
	listModel.KeyMap.CursorDown.SetEnabled(false)
	listModel.KeyMap.CursorUp.SetEnabled(false)

	return NewListModel(listModel)
}

func (i ListItem) Increase(percentage float64) tea.Cmd {
	i.increase.Exec(percentage)
	return i.setPercent()
}

func (i ListItem) Decrease(percentage float64) tea.Cmd {
	i.decrease.Exec(percentage)
	return i.setPercent()
}

func (i ListItem) setPercent() tea.Cmd {
	return i.progress.SetPercent(getPercentage(i))
}

func (i ListItem) GetPercentage() (float64, error) {
	return i.getPercentage.Exec()
}

func (i ListItem) FilterValue() string { return i.name }
