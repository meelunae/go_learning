package cyoa

import (
	"encoding/json"
	"io"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (s Story) RenderStoryWindow(tv *tview.Application, a Arc) *tview.Flex {
	// Create a flexbox layout
	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	title := tview.NewTextView().SetText(a.Title).SetTextAlign(tview.AlignCenter).SetTextColor(tcell.ColorYellow)
	flex.AddItem(title, 0, 1, false)

	for _, r := range a.Paragraphs {
		p := tview.NewTextView().SetText(r)
		flex.AddItem(p, 0, 1, false)
	}

	// Selectable Options using List
	list := tview.NewList()
	if len(a.Options) != 0 {
		for i, o := range a.Options {
			list.AddItem(o.Text, o.Chapter, rune(i+1), nil)
		}
	} else {
		list.AddItem("Exit", "leave", 'q', nil)
	}

	list.SetBorder(true).SetTitle("Options").SetTitleAlign(tview.AlignLeft) // Adding a border around the options list
	// Handle keyboard navigation
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyUp:
			if list.HasFocus() {
				currentItem := list.GetCurrentItem()
				if currentItem > 0 {
					list.SetCurrentItem(currentItem - 1)
				}
			}
			return nil
		case tcell.KeyDown:
			if list.HasFocus() {
				currentItem := list.GetCurrentItem()
				if currentItem < list.GetItemCount()-1 {
					list.SetCurrentItem(currentItem + 1)
				}
			}
			return nil
		case tcell.KeyEnter:
			if list.HasFocus() {
				index := list.GetCurrentItem()
				_, selectedChapter := list.GetItemText(index)
				if selectedChapter == "leave" {
					tv.Stop()
				}
				tv.SetRoot(s.RenderStoryWindow(tv, s[selectedChapter]), true)
				return nil
			}
		}
		return event
	})

	flex.AddItem(list, 0, 1, true)

	return flex
}

func JsonToStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Arc

type Arc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
