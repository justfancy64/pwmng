package main

import (
	"github.com/justfancy64/pwmng/internal/state"
	"github.com/justfancy64/pwmng/internal/window"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
)

func main() {

	var s state.State
	s.WinContent = make(map[string]func(*state.State) *fyne.Container)
	s.App = app.New()
	s.Window = s.App.NewWindow("pogger password manager")
	window.WindowParams(&s)
	window.InputWindow(&s)
	s.WinContent["mode"] = window.ModeWindow

	s.Window.Show()

	s.App.Run()


	//s.Window.Show()

}
