package encoding

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/auyer/steganography"
	"github.com/h2non/filetype"
	"github.com/justfancy64/pwmng/internal/state"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
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
		s.FileType = "UnKnown"
	}

	if kind.Extension == "jpg" {
		s.FileType = "JPEG"

	}

	if kind.Extension == "png" || kind.Extension == "PNG" {
		s.FileType = "PNG"

	} else {

		s.FileType = kind.Extension
	}
	return nil

}

type quissy struct {
	femboyscore int
	kittenname  string
}

func EncodePNG(s *state.State, data EncryptedStruct) error {
	inFile, err := os.Open(s.File)
	if err != nil {
		log.Println(err)
	}

	reader := bufio.NewReader(inFile)
	img, err := png.Decode(reader)
	if err != nil {
		log.Println(err)
	}
	sizeofmsg := steganography.GetMessageSizeFromImage(img)
	msg := steganography.Decode(sizeofmsg, img)
	var content EncryptedStruct
	err = json.Unmarshal(msg, &content)
	if err != nil {
		return err
	}
	for _, entry := range content.Contents {
		data.Contents = append(data.Contents, entry)
	}

	w := new(bytes.Buffer)
	bdata, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = steganography.Encode(w, img, bdata)
	if err != nil {
		log.Println(err)
	}
	outFile, err := os.Create(s.File)
	if err != nil {
		log.Println(err)
	}

	_, err = w.WriteTo(outFile)
	if err != nil {
		return err
	}
	outFile.Close()
	return nil

}

func EncodeJPEG(s *state.State, data EncryptedStruct) error {
	inFile, err := os.Open(s.File)
	if err != nil {
		log.Println(err)
	}

	reader := bufio.NewReader(inFile)
	img, err := jpeg.Decode(reader)
	if err != nil {
		log.Println(err)
	}
	sizeofmsg := steganography.GetMessageSizeFromImage(img)
	msg := steganography.Decode(sizeofmsg, img)
	var content EncryptedStruct
	err = json.Unmarshal(msg, &content)
	if err != nil {
		return err
	}
	for _, entry := range content.Contents {
		data.Contents = append(data.Contents, entry)
	}

	w := new(bytes.Buffer)
	bdata, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = steganography.Encode(w, img, bdata)
	if err != nil {
		log.Println(err)
	}
	outFile, err := os.Create(s.File)
	if err != nil {
		log.Println(err)
	}

	_, err = w.WriteTo(outFile)
	if err != nil {
		return err
	}
	outFile.Close()
	log.Println("encoded successfully")
	return nil

}
