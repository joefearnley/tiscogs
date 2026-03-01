package ui

import "github.com/joefearnley/tiscogs/internal/api"

type SearchResultsMsg struct {
	Results []api.SearchResult
	Err     error
}

type ReleaseDetailsMsg struct {
	Release *api.Release
	Err     error
}

type ArtistDetailsMsg struct {
	Artist map[string]interface{}
	Err    error
}
