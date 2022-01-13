package internal

import (
	"fmt"
	"os"

	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const (
	paused State = iota
	playing
	stopped
)

type Ui struct {
	playList      *widgets.List
	infoList      *widgets.List
	volumeGauge   *widgets.Gauge
	progressGauge *widgets.Gauge
	control       *widgets.Paragraph
	songs         []Song
}

func (S State) String() string {

}

func NewUi(songs []Song, path string) *Ui {
	if err := termui.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize ui: %v", err)
		os.Exit(1)
	}
	ui := new(Ui)
	// creating list
	playList := widgets.NewList()
	playList.Block.Title = "Song List"
	playList.TextStyle.Bg = 1
	playList.TextStyle.Fg = 0
	playList.SetRect(0, 5, 25, 12)
	// creating gauge(progress bar) for current playing song

	//creating  image for currently playing song

	//creating paragraph for introduction
	termui.Render(playList)
	return ui
}
