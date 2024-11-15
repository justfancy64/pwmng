package encoding

import(
    "bufio"
    "fmt"
    "os"
    "image/png"
    "encoding/json"
    "image/jpeg"
    "github.com/auyer/steganography"
    "github.com/justfancy64/pwmng/internal/state"
)







func DecodePNG(s *state.State, data EncryptedStruct) error {
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
    var content EncryptedStruct
    err = json.Unmarshal(msg, &content)
    if err != nil {
    return err
  }
    s.Data = content.Data


    fmt.Println(string(msg))

    return nil
}


func DecodeJPEG(s *state.State, data EncryptedStruct) error {
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
    var content EncryptedStruct
    err = json.Unmarshal(msg, &content)
    if err != nil {
    return err
  }
    fmt.Println("idk anymore")
    s.Data = content.Data
    fmt.Println(string(msg))
    return nil

}
