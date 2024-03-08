package main

import (
	"cyoa"
	"flag"
	"os"

	"github.com/rivo/tview"
)

func main() {
	var app = tview.NewApplication()
	fp := flag.String("file", "gopher.json", "file that contains our story")
	flag.Parse()

	f, err := os.Open(*fp)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonToStory(f)
	if err != nil {
		panic(err)
	}
	if err := app.SetRoot(story.RenderStoryWindow(app, story["intro"]), true).Run(); err != nil {
		panic(err)
	}
}
