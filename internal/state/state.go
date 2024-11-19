package state

import(
  "fyne.io/fyne/v2"
)

type State struct {
	Mode string
	Data []string
	File string
	FileType string
	Window   fyne.Window
	App      fyne.App
	Comment  string
	Username string
	Password string
	Contents []Content

}


type Content struct {
  Comment    string
  Username   string
  Password   string
}


func (s *State) Callback(position fyne.Position, uri []fyne.URI) {
  s.File = uri[0].Path()


  
}
