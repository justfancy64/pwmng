package encoding

import (
	"fmt"
	"github.com/justfancy64/pwmng/internal/state"
	"time"
)

type EncryptedStruct struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Data      []string
	Contents  []state.Content
}

type Functions struct {
	Encoders map[string]func(s *state.State, data EncryptedStruct) error
}

func StartEncoding(s *state.State) {
	FileDetector(s)

	var funcs Functions
	funcs.Encoders = make(map[string]func(s *state.State, data EncryptedStruct) error)

	funcs.AddFunc("encodePNG", EncodePNG)
	funcs.AddFunc("encodejpg", EncodeJPEG)
	funcs.AddFunc("decodePNG", DecodePNG)
	funcs.AddFunc("decodejpg", DecodeJPEG)
	modetype := s.Mode + s.FileType
	cont := state.Content{
		Comment:  s.Comment,
		Username: s.Username,
		Password: s.Password,
	}
	data := EncryptedStruct{
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Data:      s.Data,
		//Contents:  cont,

	}
	data.Contents = append(data.Contents, cont)
	fmt.Println(modetype)
	encoder := funcs.Encoders[modetype]
	err := encoder(s, data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}

func (f Functions) AddFunc(ModeFiletype string, function func(s *state.State, data EncryptedStruct) error) {
	f.Encoders[ModeFiletype] = function
}
