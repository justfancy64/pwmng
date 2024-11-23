package window

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/justfancy64/pwmng/internal/encoding"
	"github.com/justfancy64/pwmng/internal/pgen"
	"github.com/justfancy64/pwmng/internal/state"
	"image/color"
	//"fyne.io/fyne/v2/canvas"
	"os"
	"fyne.io/fyne/v2/app"
)


func WindowParams(s *state.State) {
	var _ myTheme
	s.Window.SetMaster()
	s.Window.CenterOnScreen()
	s.Window.SetTitle("Pogger Password Manager")
	s.Window.Resize(fyne.Size{
		Width:  400,
		Height: 300,
	})

	//blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}


}

type myTheme struct {
}


func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	/*
		if name == theme.IconNameHome {
			return fyne.NewStaticResource("myHome", homeBytes)
		}
	*/
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func CreateApp(s *state.State) {
	s.App = app.New()
}

func InputWindow(s *state.State) *fyne.Container{

	

	hello := widget.NewLabel("Welcome to Pogger Password manager\n Drop in a file or enter the path below")
	headercontainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), hello, layout.NewSpacer())
	inputline := widget.NewEntry()
	inputline.SetPlaceHolder("eg: /home/Downloads/cutepic.jpeg")
	confirm := widget.NewButton("Confirm", func() {
		if s.File == "" {
			s.File = inputline.Text
		}
		_, err := os.Open(s.File)
		if err != nil {
			fmt.Println("not a valid address")

			return
		}
		//s.File = inputline.Text
		err = encoding.FileDetector(s)
		modetype := s.Mode + s.FileType
		fmt.Println(modetype)
		if err != nil {
			fmt.Println(err)
		}
		//w.Hide()
		ModeWindow(s)

	})
	exit := widget.NewButton("Exit", func() {
		s.App.Quit()
	})
	content := container.New(layout.NewVBoxLayout(), headercontainer, inputline, confirm, exit) // vertical container input line will be added here
	s.Window.SetContent(content)


	s.Window.SetOnDropped(s.Callback)
	return content
}

func ModeWindow(s *state.State) *fyne.Container {
	//a := app.New()
	//w := a.NewWindow("Hello")

	hello := widget.NewLabel("Use encode to add a password or decode to retrieve image contents")
	headercontainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), hello, layout.NewSpacer())
	button := container.NewVBox(widget.NewButton("encode", func() {
		s.Mode = "encode"
		//w.Hide()
		modetype := s.Mode + s.FileType
		fmt.Println(modetype)
		EncodingWindow(s)
	}))
	button2 := container.NewVBox(widget.NewButton("decode", func() {
		s.Mode = "decode"
		//w.Hide()
		DecodingWindow(s)
	}))
	back := widget.NewButton("Back", func() {
		s.Window.SetContent(InputWindow(s))
	})
	exit := widget.NewButton("Exit", func() {
		s.App.Quit()
	})
	content := container.New(layout.NewVBoxLayout(), headercontainer, button, button2,layout.NewSpacer(),back, exit)
	return content

	//w.SetOnDropped(s.Callback)

	//w.Show()
}

func EncodingWindow(s *state.State) {
	//w := a.NewWindow("Encoding Window")
	title := widget.NewLabel("Enter the data u wish to encode")
	titlecont := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), title, layout.NewSpacer())
	inputComment := widget.NewEntry()
	inputComment.SetPlaceHolder("Enter comment")
	inputUsername := widget.NewEntry()
	inputUsername.SetPlaceHolder("Enter username")
	inputPassword := widget.NewEntry()
	inputPassword.SetPlaceHolder("Enter password")
	Confirm := widget.NewButton("Confirm", func() {
		data := inputComment.Text + " " + inputPassword.Text
		s.Data = append(s.Data, data)
		s.Comment = inputComment.Text
		s.Username = inputUsername.Text
		s.Password = inputPassword.Text
		encoding.StartEncoding(s)
		_ = Completed(s)

	})

	back := widget.NewButton("Back", func(){
		s.Window.SetContent(ModeWindow(s))

	})

	exit := widget.NewButton("Exit", func() {
		s.App.Quit()
	})
	plabel := widget.NewLabel("new password will be copied to your clipboard")
	pgenbtn := widget.NewButton("generate password", func() {
		pass := pgen.PassGen(s)
		plabel = widget.NewLabel(string(pass))

	})
	passwordcont := container.New(layout.NewHBoxLayout(), pgenbtn, plabel)

	content := container.New(layout.NewVBoxLayout(), titlecont, inputComment, inputUsername, inputPassword, passwordcont, layout.NewSpacer(),Confirm, back,exit)

	s.Window.SetContent(content)
	//w.Show()

}

func DecodingWindow(s *state.State) {
	//w := a.NewWindow("Decoding Window")
	title := widget.NewLabel("Data exracted from image")
	titlecont := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), title, layout.NewSpacer())
	encoding.StartEncoding(s)

	var str string
	//fmt.Println(s.Data)
	for _, line := range s.Data {
		str += line + "\n"
	}
	content := container.NewVBox()
	content.Add(titlecont)
	var data []*fyne.Container
	for _, entry := range s.Contents {

		cmnt := widget.NewLabel(entry.Comment)

		userbutton := widget.NewButton("copy username", func() {
			s.Window.Clipboard().SetContent(entry.Username)
		})

		passbutton := widget.NewButton("copy password", func() {
			s.Window.Clipboard().SetContent(entry.Password)
		})

		dataline := container.New(layout.NewHBoxLayout(), cmnt, layout.NewSpacer(),userbutton, passbutton)

		content.Add(dataline)

		data = append(data, dataline)
	}
	fmt.Println(len(data))
	back := widget.NewButton("Back", func(){
		s.Window.SetContent(ModeWindow(s))

	})
	exit := widget.NewButton("Exit", func() {
		s.App.Quit()
	})
	if len(data) == 0 {
		fmt.Println("not data to be found")
		return
	}
	content.Add(layout.NewSpacer())
	content.Add(back)
	content.Add(exit)
	s.Window.SetContent(content)
	//w.Show()
}

func Completed(s *state.State) *fyne.Container {
	text := widget.NewLabel("Data successfuly encrypted")
	row := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text,layout.NewSpacer())
	back := widget.NewButton("Back", func(){
		EncodingWindow(s)
	})
	exit := widget.NewButton("Exit",func(){
		s.App.Quit()
	})
	content := container.New(layout.NewVBoxLayout(),row,layout.NewSpacer(), back,exit)
	s.Window.SetContent(content)
	return content
}
