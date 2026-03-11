import {
    Box,
    Text,
    Input,
    delegate,
    createCliRenderer, 
    InputRenderableEvents 
} from "@opentui/core"

const renderer = await createCliRenderer()

function LabeledInput(props: { id: string; label: string; placeholder: string }) {
  return delegate(
    { focus: `${props.id}-input` },
    Box(
      { flexDirection: "column", marginBottom: 1 },
      Text({ content: props.label.padEnd(12), fg: "#888888" }),
      Input({
        id: `${props.id}-input`,
        placeholder: props.placeholder,
        backgroundColor: "#222",
        focusedBackgroundColor: "#333",
        textColor: "#FFF",
        cursorColor: "#cfcfcf",
      }),
    ),
  )
}

const searchInput = LabeledInput({
  id: "search",
  label: "",
  placeholder: "Enter search term (artist, album, etc...) and press Enter",
})

const form = Box(
  {
    width: 70,
    borderStyle: "rounded",
    title: "Search Discogs",
    padding: 1,
  },
  searchInput,
)

searchInput.focus()

searchInput.on(InputRenderableEvents.ENTER, (value: string) => {
  console.log("Submitted value:", value)
})

renderer.root.add(form)