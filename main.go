package main

import (
 // "github.com/justfancy64/pwmng/internal/window"
  "github.com/justfancy64/pwmng/internal/state"
  "github.com/justfancy64/pwmng/internal/window"


)

func main() {
	var st state.State
	app := window.InputWindow(&st)
	app.Run()

}

