package encoding

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/auyer/steganography"
	"github.com/justfancy64/pwmng/internal/state"
	"image/jpeg"
	"image/png"
	"os"
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
	if len(content.Data) == 0 {
		return fmt.Errorf("not data to be found")
	}
	s.Data = content.Data
	s.Contents = content.Contents

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
	if len(content.Data) == 0 {
		return fmt.Errorf("not data to be found")
	}
	s.Data = content.Data
	s.Contents = content.Contents
	fmt.Println(string(msg))
	return nil

}
