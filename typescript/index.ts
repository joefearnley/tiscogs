import {
  Box,
  InputRenderableEvents,
  InputRenderable,
  TextRenderable,
  createCliRenderer,
} from "@opentui/core"

const renderer = await createCliRenderer()

const searchInput = new InputRenderable(renderer, {
  id: "search-input",
  placeholder: "Enter search term (artist, album, etc...) and press Enter",
  backgroundColor: "#222",
  focusedBackgroundColor: "#333",
  textColor: "#FFF",
  cursorColor: "#cfcfcf",
})

const status = new TextRenderable(renderer, {
  id: "search-status",
  content: "Ready to search.",
  fg: "#888888",
})

searchInput.on(InputRenderableEvents.ENTER, (value: string) => {
  status.content = value.trim()
    ? `Search submitted: ${value}`
    : "Enter a search term."
})

const form = Box(
  {
    width: 70,
    borderStyle: "rounded",
    title: "Search Discogs",
    padding: 1,
    flexDirection: "column",
  },
  searchInput,
  status,
)

renderer.root.add(form)
searchInput.focus()
