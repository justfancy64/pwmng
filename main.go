package main

import (
  "github.com/justfancy64/pwmng/internal/state"
  "github.com/justfancy64/pwmng/internal/window"


)

func main() {
	var s state.State
	s.App = window.InputWindow(&s)
	s.App.Run()

}


