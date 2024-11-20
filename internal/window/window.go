package window

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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
)

var App fyne.App

func WindowParams(s *state.State) {
	var _ myTheme
	s.Window.SetMaster()
	s.Window.SetTitle("Pogger Password Manager")
	s.Window.Resize(fyne.Size{
		Width:  400,
		Height: 300,
	})

	//blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}

	s.Window.CenterOnScreen()
	s.App.Settings().SetTheme(&myTheme{})

}

type myTheme struct {
}

var _ fyne.Theme = (*myTheme)(nil)

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

func InputWindow(s *state.State) fyne.App {
	a := app.New()
	w := a.NewWindow("pogger password manager")
	s.Window = w
	WindowParams(s)

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
		ModeWindow(s, a)

	})
	exit := widget.NewButton("Exit", func() {
		a.Quit()
	})
	content := container.New(layout.NewVBoxLayout(), headercontainer, inputline, confirm, exit) // vertical container input line will be added here
	w.SetContent(content)

	w.Show()

	w.SetOnDropped(s.Callback)
	return a
}

func ModeWindow(s *state.State, a fyne.App) {
	//a := app.New()
	//w := a.NewWindow("Hello")

	hello := widget.NewLabel("Use encode to if u want to add a password or decode to retrieve one")
	headercontainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), hello, layout.NewSpacer())
	button := container.NewVBox(widget.NewButton("encode", func() {
		s.Mode = "encode"
		//w.Hide()
		modetype := s.Mode + s.FileType
		fmt.Println(modetype)
		EncodingWindow(s, a)
	}))
	button2 := container.NewVBox(widget.NewButton("decode", func() {
		s.Mode = "decode"
		//w.Hide()
		DecodingWindow(s, a)
	}))
	exit := widget.NewButton("Exit", func() {
		s.App.Quit()
	})
	content := container.New(layout.NewVBoxLayout(), headercontainer, button, button2, exit)
	s.Window.SetContent(content)

	//w.SetOnDropped(s.Callback)

	//w.Show()
}

func EncodingWindow(s *state.State, a fyne.App) {
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

	})

	exit := widget.NewButton("Exit", func() {
		a.Quit()
	})
	plabel := widget.NewLabel("new password will be copied to your clipboard")
	pgenbtn := widget.NewButton("generate password", func() {
		pass := pgen.PassGen(s)
		plabel = widget.NewLabel(string(pass))

	})
	passwordcont := container.New(layout.NewHBoxLayout(), pgenbtn, plabel)

	content := container.New(layout.NewVBoxLayout(), titlecont, inputComment, inputUsername, inputPassword, passwordcont, Confirm, exit)

	s.Window.SetContent(content)
	//w.Show()

}

func DecodingWindow(s *state.State, a fyne.App) {
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

		dataline := container.New(layout.NewHBoxLayout(), cmnt, userbutton, passbutton)
		content.Add(dataline)

		data = append(data, dataline)
	}
	fmt.Println(len(data))

	exit := widget.NewButton("Exit", func() {
		a.Quit()
	})
	if len(data) == 0 {
		fmt.Println("not data to be found")
		return
	}
	content.Add(exit)
	s.Window.SetContent(content)
	//w.Show()
}
