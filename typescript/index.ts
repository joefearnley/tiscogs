import { Input, createCliRenderer, InputRenderableEvents } from "@opentui/core"

const renderer = await createCliRenderer()

const input = Input({
  placeholder: "Enter your name...",
  width: 25,
})

input.focus()
renderer.root.add(input)

input.on(InputRenderableEvents.CHANGE, (value: string) => {
  console.log("Value committed:", value)
});

input.on(InputRenderableEvents.ENTER, (value: string) => {
  console.log("Submitted value:", value)
});
