package window

import (
  "fmt"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/layout"
  "github.com/justfancy64/pwmng/internal/state"
  "github.com/justfancy64/pwmng/internal/encoding"
  "os"

)

func InputWindow(s *state.State) fyne.App {
  a := app.New()
  w := a.NewWindow("hello")

  hello := widget.NewLabel("Welcome to Pogger Password manager\n Drop in a file or enter the path below")
  headercontainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), hello,layout.NewSpacer())
  inputline := widget.NewEntry()
  inputline.SetPlaceHolder("eg: /home/Downloads/cutepic.jpeg")
  confirm := widget.NewButton("Confirm", func(){
    _,err := os.Open(inputline.Text)
    if err != nil {
    fmt.Println("not a valid address")

    return
    }
    s.File = inputline.Text
    err = encoding.FileDetector(s)
    modetype := s.Mode + s.FileType
    fmt.Println(modetype)
    if err != nil {
      fmt.Println(err)
    }
    w.Hide()
    ModeWindow(s, a)

    })

  content := container.New(layout.NewVBoxLayout(), headercontainer,inputline,confirm) // vertical container input line will be added here
  w.SetContent(content)

  w.SetOnDropped(s.Callback)
  w.Show()

  return a
}





func ModeWindow(s *state.State,a fyne.App) {
	//a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Use encode to if u want to add a password or decode to retrieve one")
	headercontainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), hello,layout.NewSpacer())
	button := container.NewVBox(widget.NewButton("encode", func(){
		  s.Mode = "encode"
		  w.Hide()
		  modetype := s.Mode + s.FileType
		  fmt.Println(modetype)
		  EncodingWindow(s,a)
		  }))
	button2 := container.NewVBox(widget.NewButton("decode", func(){
		  s.Mode = "decode"
		  w.Hide()
		  DecodingWindow(s,a)
		   }))
	content := container.New(layout.NewVBoxLayout(), headercontainer,button,button2)
	w.SetContent(content)


        w.SetOnDropped(s.Callback)

	w.Show()
}






func EncodingWindow(s *state.State,a fyne.App) {
  w := a.NewWindow("Encoding Window")
  title := widget.NewLabel("Enter the data u wish to encode")
  titlecont := container.New(layout.NewHBoxLayout(), layout.NewSpacer(),title,layout.NewSpacer())
  inputComment := widget.NewEntry()
  inputComment.SetPlaceHolder("eg: Username and password for www.examplesite.com")
  inputPassword := widget.NewEntry()
  inputPassword.SetPlaceHolder("Tip: enter multile entries seperated by spaces")
  Confirm := widget.NewButton("Confirm", func(){
    data := inputComment.Text + inputPassword.Text
    s.Data = append(s.Data, data)
    encoding.StartEncoding(s)
     
  })

  exit := widget.NewButton("Exit", func(){
    a.Quit()
  })

  content := container.New(layout.NewVBoxLayout(),titlecont,inputComment,inputPassword,Confirm,exit)


  w.SetContent(content)
  w.Show()
  
}



func DecodingWindow(s *state.State,a fyne.App) {
  w := a.NewWindow("Decoding Window")
  title := widget.NewLabel("Data exracted from image")
  titlecont := container.New(layout.NewHBoxLayout(), layout.NewSpacer(),title,layout.NewSpacer())
  encoding.StartEncoding(s)

  var str string
  //fmt.Println(s.Data)
  for _,line := range s.Data{
    str += line
  }
  data := widget.NewLabel(str)
  
  exit := widget.NewButton("Exit", func(){
    a.Quit()
  })
  content := container.New(layout.NewVBoxLayout(),titlecont,data,exit)
  w.SetContent(content)
  w.Show()
}
