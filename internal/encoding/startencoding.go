package encoding


import(
  "fmt"
  "github.com/justfancy64/pwmng/internal/state"
  "time"
)

type EncryptedStruct struct {
  CreatedAt  time.Time
  UpdatedAt  time.Time
  Data       []string
}

type Functions struct{
  Encoders       map[string]func(s *state.State,data EncryptedStruct) error
}



func StartEncoding(s *state.State) {

  var funcs Functions
  funcs.Encoders = make(map[string]func(s *state.State,data EncryptedStruct) error)


  funcs.AddFunc("encodePNG", EncodePNG)
  funcs.AddFunc("encodejpg", EncodeJPEG)
  funcs.AddFunc("decodePNG", DecodePNG)
  funcs.AddFunc("decodejpg", DecodeJPEG)
  modetype := s.Mode + s.FileType
  data := EncryptedStruct{
    CreatedAt: time.Now().UTC(),
    UpdatedAt: time.Now().UTC(),
    Data:      s.Data,
  }
  fmt.Println(modetype)
  encoder := funcs.Encoders[modetype]
  err := encoder(s, data)
  if err != nil {
    fmt.Println(err)
  }



}



func (f Functions) AddFunc(ModeFiletype string, function func(s *state.State,data EncryptedStruct) error) {
  f.Encoders[ModeFiletype] = function
}

