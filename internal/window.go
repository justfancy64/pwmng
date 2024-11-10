package window

import (
  "fmt"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2"

)

func CreateWindow() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))
        w.SetOnDropped(callback)

	w.ShowAndRun()
}

func callback(position fyne.Position, uri []fyne.URI) {
  fmt.Println(uri[0])

}
