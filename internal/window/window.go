package window

import (
  "fmt"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/layout"
  "github.com/justfancy64/pwmng/internal/state"
  "os"

)

func InputWindow(s *state.State) fyne.App {
  a := app.New()
  w := a.NewWindow("hello")

  hello := widget.NewLabel("Welcome to Pogger Password manager\n Drop in a file or enter the path below")
  headercontainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), hello,layout.NewSpacer())
  inputline := widget.NewEntry()
  inputline.SetPlaceHolder("eg: /home/Downloads/cutepic.jpeg")
  confirm := widget.NewButton("Confirm", func(){
    _,err := os.Open(inputline.Text)
    if err != nil {
    fmt.Println("not a valid address")

    return
    }
    s.File = inputline.Text
    w.Hide()
    ModeWindow(s, a)

    })

  content := container.New(layout.NewVBoxLayout(), headercontainer,inputline,confirm) // vertical container input line will be added here
  w.SetContent(content)

  w.SetOnDropped(callback)
  w.Show()

  return a
}





func ModeWindow(s *state.State,a fyne.App) {
	//a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("use encode to if u want to add a password or decode to retrieve one")
	headercontainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), hello,layout.NewSpacer())
	button := container.NewVBox(widget.NewButton("encode", func(){s.Mode = "encode"}))
	button2 := container.NewVBox(widget.NewButton("decode", func(){s.Mode = "decode"}))
	content := container.New(layout.NewVBoxLayout(), headercontainer,button,button2)
	w.SetContent(content)


        w.SetOnDropped(callback)

	w.Show()
}




func callback(position fyne.Position, uri []fyne.URI) {
  fmt.Println(uri[0])

}
