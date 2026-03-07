# Quick Start Guide for Tiscogs

## 1. Get Your Discogs API Token

1. Go to https://www.discogs.com/settings/developers
2. Click "Create an Access Token"
3. Copy the token that appears

## 2. Set Environment Variable

```bash
export DISCOGS_TOKEN=your_token_here
```

Or on macOS, add it to your shell profile:
```bash
echo 'export DISCOGS_TOKEN=your_token_here' >> ~/.zshrc
source ~/.zshrc
```

## 3. Build and Run

```bash
# Navigate to the project directory
cd /Users/joefearnley/projects/tiscogs

# Build the application
go build -o tiscogs

# Run the application
./tiscogs
```

Or use make:
```bash
make run
```

## 4. Using the Application

Once the app starts:

1. **Type a search query** - Enter an artist name, album name, or label you want to search for
   - Example: "The Beatles"
   - Example: "Dark Side of the Moon"
   - Example: "Warp Records"

2. **Press Enter** to search

3. **Navigate results** using arrow keys (↑/↓)

4. **Press Enter** to view details of a selected result

5. **Press 'b'** to go back to the search view

6. **Press 'q'** or **Ctrl+C** to quit the application

## Examples to Try

1. Search for an artist: "Pink Floyd"
2. Search for a label: "Discogs"
3. Search for an album: "Thriller"

## Troubleshooting

### Error: "DISCOGS_TOKEN environment variable not set"
Make sure you've set the environment variable:
```bash
export DISCOGS_TOKEN=your_token_here
```

### API rate limiting
The Discogs API has rate limits. If you get errors, wait a moment before searching again.

### No results found
Try:
- Spelling check
- Use different search terms
- Try artist name or album name separately

## Need Help?

- [Discogs API Docs](https://www.discogs.com/developers/)
- [Bubbletea Examples](https://github.com/charmbracelet/bubbletea/tree/master/examples)
- Check the [README.md](README.md) for more details
