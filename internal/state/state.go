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

}



func (s *State) Callback(position fyne.Position, uri []fyne.URI) {
  fmt.Println(uri[0])
}
