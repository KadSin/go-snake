package interaction

import (
	"time"

	term "github.com/nsf/termbox-go"
)

const (
	PASS_BY_TTL = iota
	PASS_BY_KEY
)

type Story struct {
	Content       Content
	Background    term.Attribute
	PassMethod    int
	SecondsToLive int
	KeyToPass     term.Key
}

func (story Story) Show() {
	story.Content.Print()

	term.Flush()
	term.Clear(term.ColorDefault, story.Background)

	if story.PassMethod == PASS_BY_KEY {
		sleepUnlessKeyPressed(story.KeyToPass)
	} else {
		sleep(story.SecondsToLive)
	}
}

func sleepUnlessKeyPressed(key term.Key) {
	for {
		var event = term.PollEvent()

		if event.Type == term.EventKey && event.Key == key {
			break
		}
	}
}

func sleep(seconds int) {
	end := time.Now().UnixMilli() + int64(seconds*1000)
	ticker := time.NewTicker(time.Second)

	for t := range ticker.C {
		if t.UnixMilli() > end {
			break
		}
	}
}
