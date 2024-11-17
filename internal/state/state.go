package state

import(
  "fyne.io/fyne/v2"
  "fmt"
)

type State struct {
	Mode string
	Data []string
	File string
	FileType string
	Window   fyne.Window
	App      fyne.App

}



func (s *State) Callback(position fyne.Position, uri []fyne.URI) {
  s.File = uri[0].Path()
  fmt.Println(uri[0].Path())


  
}
