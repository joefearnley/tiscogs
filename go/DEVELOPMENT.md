# Development Guide for Tiscogs

## Project Architecture

Tiscogs follows a clean architecture pattern with clear separation of concerns:

```
tiscogs/
├── main.go                    # Entry point
├── internal/
│   ├── api/                   # Discogs API client layer
│   │   └── client.go          # HTTP client, models, and API methods
│   └── ui/                    # Terminal UI layer
│       ├── app.go             # Main application state machine
│       ├── search.go          # Search view component
│       └── messages.go        # Event messages
└── go.mod, go.sum            # Dependencies
```

## Core Components

### 1. API Client (`internal/api/client.go`)

Handles all communication with the Discogs API:

- **Models**: `SearchResult`, `Release`, `Artist`, `Track`
- **Methods**:
  - `NewClient(token string)` - Create client
  - `Search(query, searchType)` - Search database
  - `GetRelease(releaseID)` - Get release details
  - `GetArtist(artistID)` - Get artist details

### 2. UI Application (`internal/ui/app.go`)

Implements the Bubble Tea application model using the Elm architecture:

- **States**:
  - `SearchState` - User entering search query
  - `ResultsState` - Displaying search results
  - `DetailsState` - Showing details of selected item
  - `LoadingState` - Async operation in progress

- **Main Methods**:
  - `Init()` - Initialize app
  - `Update(msg)` - Handle user input and state changes
  - `View()` - Render UI based on current state

### 3. Search View (`internal/ui/search.go`)

Encapsulates the search input UI component using `bubbles/textinput`.

### 4. Messages (`internal/ui/messages.go`)

Custom event types for async operations:
- `SearchResultsMsg` - Results from API search
- `ReleaseDetailsMsg` - Release details
- `ArtistDetailsMsg` - Artist details

## Extension Points

### Adding New Search Types

To add support for searching by different types (e.g., labels, masters):

1. Modify the search view to allow type selection
2. Update the `handleSearchInput` method in `app.go`
3. The API client already supports different types via the `searchType` parameter

### Adding Release Details View

Currently, the app shows basic result info. To add a detailed view:

1. Create a new component in `internal/ui/details.go`
2. Add `ReleaseDetailsView` struct with `Update()` and `View()` methods
3. Modify `app.go` to fetch and display release details when selected
4. Use the `GetRelease()` method from the API client

Example:
```go
func (a *App) performGetRelease(releaseID int) tea.Cmd {
    return func() tea.Msg {
        release, err := a.apiClient.GetRelease(releaseID)
        return ReleaseDetailsMsg{Release: release, Err: err}
    }
}
```

### Adding Pagination

The Discogs API supports pagination. To implement:

1. Add pagination state to `App` struct: `currentPage int`
2. Modify `SearchResponse` struct to include pagination info
3. Add "Next" and "Previous" navigation in results view
4. Update `Search` method to include page parameter

### Caching Results

To avoid repeated API calls:

1. Add a `results map[string][]SearchResult` to cache
2. Check cache before making API calls
3. Implement cache invalidation strategy

## Testing

Currently, no tests are included. To add tests:

```bash
# Create test files
touch internal/api/client_test.go
touch internal/ui/app_test.go
```

Example test:
```go
func TestSearchSuccess(t *testing.T) {
    client := NewClient("test-token")
    results, err := client.Search("Beatles", "")
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    if len(results) == 0 {
        t.Fatal("Expected results, got none")
    }
}
```

Run tests:
```bash
go test ./...
```

## Building and Deployment

### Build for different platforms

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o tiscogs-linux

# Windows
GOOS=windows GOARCH=amd64 go build -o tiscogs.exe

# macOS Intel
GOOS=darwin GOARCH=amd64 go build -o tiscogs-intel

# macOS ARM
GOOS=darwin GOARCH=arm64 go build -o tiscogs-arm64
```

### Create release

```bash
# Build all platforms
make release

# Creates: dist/tiscogs-{linux,windows,darwin-*}
```

## Code Style

- Follow Go idioms and conventions
- Use `gofmt` for formatting: `go fmt ./...`
- Use meaningful variable names
- Keep functions small and focused
- Document exported functions with comments

## Debugging

### Enable verbose logging

Add to `main.go`:
```go
import "log"

func init() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
}
```

### Debug API responses

Add temporary logging in `client.go`:
```go
body, _ := io.ReadAll(resp.Body)
log.Printf("Response: %s", string(body))
```

## Performance Considerations

1. **API Rate Limiting**: Discogs API has rate limits. Consider:
   - Caching results
   - Implementing backoff strategy
   - Respecting rate limit headers

2. **Memory**: 
   - Large search results are kept in memory
   - Consider pagination for large result sets

3. **Network**:
   - Use timeouts on HTTP requests
   - Handle network errors gracefully

## Dependencies

- **bubbletea** (v0.24.1+): TUI framework
- **bubbles** (v0.18.0+): UI components
- **lipgloss** (v0.9.1+): Terminal styling

Update with: `go get -u`

## Resources

- [Bubbletea Examples](https://github.com/charmbracelet/bubbletea/tree/master/examples)
- [Discogs API Docs](https://www.discogs.com/developers/)
- [Go Best Practices](https://golang.org/doc/effective_go)
- [Elm Architecture](https://guide.elm-lang.org/architecture/)
