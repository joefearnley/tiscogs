package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joefearnley/tiscogs/internal/ui"
)

func main() {
	// Check for Discogs API token
	token := os.Getenv("DISCOGS_TOKEN")
	if token == "" {
		fmt.Println("Error: DISCOGS_TOKEN environment variable not set")
		fmt.Println("Get your token from: https://www.discogs.com/settings/developers")
		os.Exit(1)
	}

	// Initialize the TUI application
	app := ui.NewApp(token)
	p := tea.NewProgram(app, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatalf("Error running program: %v", err)
	}
}
