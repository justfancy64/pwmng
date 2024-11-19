package pgen

import (
    "math/rand"
    "strconv"
    "encoding/base64"
    //"fyne.io/fyne/v2"
    "github.com/justfancy64/pwmng/internal/state"

)



func PassGen(s *state.State) []byte{
  src := rand.NewSource(rand.Int63())
  random := rand.New(src)
  data := []byte(strconv.FormatInt(random.Int63(), 36))
  dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
  base64.StdEncoding.Encode(dst, data)
  s.Window.Clipboard().SetContent(string(dst[:15]))

  return dst[:15]




}
