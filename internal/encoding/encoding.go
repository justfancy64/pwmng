package encoding


import (
    "bufio"
    "image/png"
    "image/jpeg"
    "log"
    "os"
    "bytes"
    "io/ioutil"
    "github.com/h2non/filetype"
    "github.com/auyer/steganography"
    "github.com/justfancy64/pwmng/internal/state"
)

func FileDetector(s *state.State) error {
    buf, err := ioutil.ReadFile(s.File)
    if err != nil {
    log.Println(err)
    }
    kind, err := filetype.Match(buf)
    if err != nil {
    log.Println(err)
    }
    if kind == filetype.Unknown {
	log.Println("unknown file type")
    }

    if kind.Extension == "jpeg" {
	s.FileType = kind.Extension

    }

    if kind.Extension == "png" {
	s.FileType = kind.Extension

    }
    return nil

}

type quissy struct {
    femboyscore int
    kittenname string
}
func EncodePNG(s *state.State) error {
    inFile, err := os.Open(s.File)
    if err != nil {
    log.Println(err)
    }


    reader := bufio.NewReader(inFile)
    img, err := png.Decode(reader)
    if err != nil {
	log.Println(err)
    }
    w := new(bytes.Buffer)
    err = steganography.Encode(w, img, []byte(s.Args[0]))
    if err != nil {
	log.Println(err)
    }
    outFile, err := os.Create(s.File)
    if err != nil {
	log.Println(err)
    }

    w.WriteTo(outFile)
    outFile.Close()
    return nil 


}



func EncodeJPG(s *state.State) error {
    inFile, err := os.Open(s.File)
    if err != nil {
    log.Println(err)
    }


    reader := bufio.NewReader(inFile)
    img, err := jpeg.Decode(reader)
    if err != nil {
	log.Println(err)
    }
    w := new(bytes.Buffer)
    err = steganography.Encode(w, img,[]byte(s.Args[0]))
    if err != nil {
	log.Println(err)
    }
    outFile, err := os.Create(s.File)
    if err != nil {
	log.Println(err)
    }

    w.WriteTo(outFile)
    outFile.Close()
    log.Println("encoded successfully")
    return nil 


}
