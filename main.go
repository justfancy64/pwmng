package main

import (
 // "github.com/justfancy64/pwmng/internal/window"
  "github.com/justfancy64/pwmng/internal/pgen"
  "github.com/justfancy64/pwmng/internal/encoding"
  "github.com/justfancy64/pwmng/internal/state"
  "github.com/justfancy64/pwmng/internal/window"
  "fmt"


)

func main() {
	var st state.State
	app := window.InputWindow(&st)
	app.Run()
//	st.File = "/home/zach/Downloads/nottree.jpg"
	st.Args = append(st.Args,"quissy quissy")
	pass := pgen.PassGen()  // password []byte 
	fmt.Println(string(pass))
	err := encoding.FileDetector(&st)
	if err != nil {
	fmt.Println(err)
	}
	err = encoding.DecodePNG(&st)
//	encoding.Encode("/home/zach/Downloads/nottree.jpg", pass)


}


