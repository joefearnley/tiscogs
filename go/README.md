# Tiscogs - Discogs CLI Explorer

A beautiful Terminal User Interface (TUI) application written in Go for exploring the Discogs database directly from your terminal.

## Features

- 🎵 Search for artists, albums, and labels
- 📝 View detailed information about releases
- 🎨 Beautiful terminal UI with intuitive navigation
- ⚡ Fast API integration with Discogs
- 🔍 Real-time search results

## Prerequisites

- Go 1.21 or higher
- Discogs API token (free) from https://www.discogs.com/settings/developers

## Installation

1. Set your Discogs API token as an environment variable:
```bash
export DISCOGS_TOKEN=your_token_here
```

2. Build the application:
```bash
go build -o tiscogs
```

## Usage

Run the application:
```bash
./tiscogs
```

### Navigation

- Type your search query (artist, album, or label name)
- Press **Enter** to search
- Use **↑/↓** arrow keys to navigate results
- Press **Enter** to view details
- Press **b** to go back
- Press **q** or **Ctrl+C** to quit

## Project Structure

```
tiscogs/
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
├── go.sum                  # Dependency checksums
├── internal/
│   ├── api/               # Discogs API client
│   │   └── client.go      # HTTP client and API methods
│   └── ui/                # Terminal UI components
│       ├── app.go         # Main app logic
│       ├── search.go      # Search view
│       └── messages.go    # Event messages
└── README.md
```

## Dependencies

- **bubbletea** - TUI framework for Go
- **bubbles** - Components for bubbletea
- **lipgloss** - Terminal styling library

## API Integration

Supports searching the Discogs database for:
- Releases (albums, vinyl, etc.)
- Artists
- Labels

## Getting a Discogs API Token

1. Visit https://www.discogs.com/settings/developers
2. Click "Create an Access Token"
3. Copy your token and set it: `export DISCOGS_TOKEN=your_token_here`

## Development

To run in development mode:
```bash
go run main.go
```

To install dependencies:
```bash
go mod download
```

## License

See the LICENSE file for details.

## Resources

- [Discogs API Documentation](https://www.discogs.com/developers/)
- [Bubbletea Documentation](https://github.com/charmbracelet/bubbletea)
- [Go Documentation](https://golang.org/doc/)
