package main

import (
 // "github.com/justfancy64/pwmng/internal/window"
  "github.com/justfancy64/pwmng/internal/pgen"
  "github.com/justfancy64/pwmng/internal/encoding"
  "fmt"


)

func main() {
	//window.CreateWindow()
	pass := pgen.PassGen() 
	fmt.Println(pass)
	encoding.FileDetector( "/home/zach/Downloads/nottree.jpg")
//	encoding.Encode("/home/zach/Downloads/nottree.jpg", pass)

}



