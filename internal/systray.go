package internal

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"log"
)

func Gui() {
	a := app.New()
	w := a.NewWindow("idk how to make systray without this window??")

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("Seadex",
			fyne.NewMenuItem("What to watch today", func() {
				anime := AnilistAnime("Darling in the Franxx").airing
				log.Println(anime)
			}))
		desk.SetSystemTrayMenu(m)
	}
	//w.SetContent(widget.NewLabel("Seadex Test"))
	w.ShowAndRun()
	w.Hide()
}
