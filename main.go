package main

import (
	"ditto/language"
	"ditto/name"
	"fmt"
)

func main() {
  fmt.Println("Welcome to Ditto project lang randomiser")

  lng := language.SelectRandomLanguage()
  name := name.GenerateName()

  res := fmt.Sprintf("Your selected language is %s, and the name of the project is %s", lng.Name, name)

  fmt.Println(res)
}
