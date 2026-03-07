package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SearchView struct {
	textInput textinput.Model
}

func NewSearchView() *SearchView {
	ti := textinput.New()
	ti.Placeholder = "Enter artist, album, or label name..."
	ti.Focus()
	ti.CharLimit = 100

	return &SearchView{textInput: ti}
}

func (sv *SearchView) Update(msg tea.KeyMsg) (*SearchView, tea.Cmd) {
	var cmd tea.Cmd
	sv.textInput, cmd = sv.textInput.Update(msg)
	return sv, cmd
}

func (sv *SearchView) View() string {
	title := titleStyle.Render("🎵 Tiscogs - Discogs CLI Explorer")
	subtitle := subtleStyle.Render("A terminal UI for searching the Discogs database")
	
	input := "\n\nWhat would you like to search for?\n"
	input += "(Artists, Albums, Labels, etc.)\n\n"
	input += sv.textInput.View() + "\n"
	input += subtleStyle.Render("Press Enter to search, Ctrl+C to quit")

	return title + "\n" + subtitle + input
}

func (sv *SearchView) GetQuery() string {
	return sv.textInput.Value()
}

func (sv *SearchView) ClearQuery() {
	sv.textInput.SetValue("")
}
