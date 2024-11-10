package pgen

import (
    "math/rand"
    "strconv"
    "encoding/base64"
)



func PassGen() []byte{
  src := rand.NewSource(rand.Int63())
  random := rand.New(src)
  data := []byte(strconv.FormatInt(random.Int63(), 36))
  dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
  base64.StdEncoding.Encode(dst, data)
  return dst[:15]




}
