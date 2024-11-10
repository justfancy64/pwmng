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
)

func FileDetector(file string) error {
    buf, err := ioutil.ReadFile(file)
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
    log.Printf("file type: %s.\n", kind.Extension)
    return nil

}

func Encode(file string, message []byte) error {
    inFile, err := os.Open(file)
    if err != nil {
    log.Println(err)
    }


    reader := bufio.NewReader(inFile)
    img, err := png.Decode(reader)
    if err != nil {
	log.Println(err)
    }
    w := new(bytes.Buffer)
    err = steganography.Encode(w, img, message)
    if err != nil {
	log.Println(err)
    }
    outFile, err := os.Create(file)
    if err != nil {
	log.Println(err)
    }

    w.WriteTo(outFile)
    outFile.Close()
    return nil 


}



func EncodeJPG(file string, message []byte) error {
    inFile, err := os.Open(file)
    if err != nil {
    log.Println(err)
    }


    reader := bufio.NewReader(inFile)
    img, err := jpeg.Decode(reader)
    if err != nil {
	log.Println(err)
    }
    w := new(bytes.Buffer)
    err = steganography.Encode(w, img, message)
    if err != nil {
	log.Println(err)
    }
    outFile, err := os.Create(file)
    if err != nil {
	log.Println(err)
    }

    w.WriteTo(outFile)
    outFile.Close()
    log.Println("encoded successfully")
    return nil 


}
