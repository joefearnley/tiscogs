package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const BaseURL = "https://api.discogs.com"

type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

type SearchResult struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Resource string `json:"resource_url"`
	URI      string `json:"uri"`
	Year     int    `json:"year,omitempty"`
	Thumb    string `json:"thumb"`
}

type SearchResponse struct {
	Results []SearchResult `json:"results"`
	Pages   int            `json:"pagination,omitempty"`
}

type Release struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Artists     []Artist `json:"artists"`
	Year        int      `json:"year"`
	GenresList  []string `json:"genres"`
	Thumb       string   `json:"thumb"`
	ResourceURL string   `json:"resource_url"`
	URI         string   `json:"uri"`
	Description string   `json:"description,omitempty"`
	Tracklist   []Track  `json:"tracklist,omitempty"`
}

type Artist struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
}

type Track struct {
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Position string `json:"position"`
}

// NewClient creates a new Discogs API client
func NewClient(token string) *Client {
	return &Client{
		baseURL:    BaseURL,
		token:      token,
		httpClient: &http.Client{},
	}
}

// Search searches for releases, artists, or labels on Discogs
func (c *Client) Search(query string, searchType string) ([]SearchResult, error) {
	searchURL := fmt.Sprintf("%s/database/search", c.baseURL)

	fmt.Println("Performing search with query:", query)
	fmt.Println("Search type:", searchType)
	fmt.Println("Using token:", searchURL)

	params := url.Values{}
	params.Add("q", query)
	if searchType != "" {
		params.Add("type", searchType)
	}
	params.Add("token", c.token)

	println("Constructed search URL:", searchURL+"?"+params.Encode())

	req, err := http.NewRequest("GET", searchURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "TiscogsApp/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(body))
	}

	var searchResp SearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		return nil, err
	}

	return searchResp.Results, nil
}

// GetRelease fetches details of a specific release
func (c *Client) GetRelease(releaseID int) (*Release, error) {
	releaseURL := fmt.Sprintf("%s/releases/%d", c.baseURL, releaseID)

	params := url.Values{}
	params.Add("token", c.token)

	req, err := http.NewRequest("GET", releaseURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "TiscogsApp/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(body))
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}

	return &release, nil
}

// GetArtist fetches details of a specific artist
func (c *Client) GetArtist(artistID int) (map[string]interface{}, error) {
	artistURL := fmt.Sprintf("%s/artists/%d", c.baseURL, artistID)

	params := url.Values{}
	params.Add("token", c.token)

	req, err := http.NewRequest("GET", artistURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "TiscogsApp/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode, string(body))
	}

	var artist map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&artist); err != nil {
		return nil, err
	}

	return artist, nil
}
