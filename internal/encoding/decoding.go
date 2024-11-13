package encoding

import(
    "bufio"
    "fmt"
    "os"
    "image/png"
    "image/jpeg"
    "github.com/auyer/steganography"
    "github.com/justfancy64/pwmng/internal/state"
)







func DecodePNG(s *state.State) error {
    inFile, err := os.Open(s.File)
    if err != nil {
      return err
    }

    reader := bufio.NewReader(inFile)
    img, err := png.Decode(reader)
    if err != nil {
      return err
    }
    sizeofmsg := steganography.GetMessageSizeFromImage(img)
    msg := steganography.Decode(sizeofmsg, img)


    fmt.Println(string(msg))

    return nil
}


func DecodeJPEG(s *state.State) error {
    infile, err := os.Open(s.File)
    if err != nil {
      return err
    }

    reader := bufio.NewReader(infile)
    img, err := jpeg.Decode(reader)
    if err != nil {
      return err
    }
    sizeofmsg := steganography.GetMessageSizeFromImage(img)
    msg := steganography.Decode(sizeofmsg, img)
    
    fmt.Println(string(msg))
    return nil

}
