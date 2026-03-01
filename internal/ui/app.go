package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/joefearnley/tiscogs/internal/api"
)

type AppState int

const (
	SearchState AppState = iota
	ResultsState
	DetailsState
	LoadingState
)

type App struct {
	state      AppState
	apiClient  *api.Client
	width      int
	height     int
	searchView *SearchView
	results    []api.SearchResult
	selectedIdx int
	err        error
}

var (
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("170")).
		Padding(1, 0)

	subtleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("246"))

	highlightStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("39")).
		Foreground(lipgloss.Color("230"))
)

// NewApp creates a new App instance
func NewApp(token string) *App {
	return &App{
		state:     SearchState,
		apiClient: api.NewClient(token),
		searchView: NewSearchView(),
	}
}

// Init initializes the app
func (a *App) Init() tea.Cmd {
	return nil
}

// Update handles user input and state updates
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return a, tea.Quit
		}

		switch a.state {
		case SearchState:
			return a.handleSearchInput(msg)
		case ResultsState:
			return a.handleResultsInput(msg)
		case DetailsState:
			return a.handleDetailsInput(msg)
		}

	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		return a, nil

	case SearchResultsMsg:
		a.results = msg.Results
		a.selectedIdx = 0
		a.state = ResultsState
		a.err = msg.Err
		return a, nil

	case ReleaseDetailsMsg:
		a.state = DetailsState
		return a, nil
	}

	return a, nil
}

// View renders the app display
func (a *App) View() string {
	if a.width == 0 {
		return "Initializing..."
	}

	switch a.state {
	case SearchState:
		return a.renderSearchView()
	case ResultsState:
		return a.renderResultsView()
	case DetailsState:
		return a.renderDetailsView()
	case LoadingState:
		return a.renderLoadingView()
	}

	return ""
}

func (a *App) renderSearchView() string {
	return a.searchView.View()
}

func (a *App) renderResultsView() string {
	title := titleStyle.Render("Search Results")
	
	if len(a.results) == 0 {
		return title + "\n\nNo results found. Press 'b' to go back or 'q' to quit."
	}

	results := "\n"
	for i, result := range a.results {
		if i == a.selectedIdx {
			results += highlightStyle.Render("▶ " + result.Title) + "\n"
		} else {
			results += "  " + result.Title + "\n"
		}
	}

	footer := "\n" + subtleStyle.Render("↑/↓: Navigate | Enter: Select | b: Back | q: Quit")
	return title + results + footer
}

func (a *App) renderDetailsView() string {
	if a.selectedIdx >= len(a.results) {
		return "Invalid selection"
	}

	result := a.results[a.selectedIdx]
	title := titleStyle.Render(result.Title)
	details := "\n\nID: " + lipgloss.NewStyle().Foreground(lipgloss.Color("34")).Render(string(rune(result.ID))) + "\n"
	details += "Type: " + result.Type + "\n"
	if result.Year > 0 {
		details += "Year: " + lipgloss.NewStyle().Foreground(lipgloss.Color("34")).Render(string(rune(result.Year))) + "\n"
	}

	footer := "\n" + subtleStyle.Render("b: Back | q: Quit")
	return title + details + footer
}

func (a *App) renderLoadingView() string {
	return titleStyle.Render("Loading...") + "\n\nPlease wait..."
}

func (a *App) handleSearchInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		query := a.searchView.GetQuery()
		if query != "" {
			a.state = LoadingState
			return a, a.performSearch(query, "")
		}
	}

	var cmd tea.Cmd
	a.searchView, cmd = a.searchView.Update(msg)
	return a, cmd
}

func (a *App) handleResultsInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up":
		if a.selectedIdx > 0 {
			a.selectedIdx--
		}
	case "down":
		if a.selectedIdx < len(a.results)-1 {
			a.selectedIdx++
		}
	case "enter":
		a.state = DetailsState
	case "b":
		a.state = SearchState
	}
	return a, nil
}

func (a *App) handleDetailsInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "b":
		a.state = ResultsState
	}
	return a, nil
}

func (a *App) performSearch(query string, searchType string) tea.Cmd {
	return func() tea.Msg {
		results, err := a.apiClient.Search(query, searchType)
		return SearchResultsMsg{Results: results, Err: err}
	}
}
